/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package secretutil

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	secretsstorev1 "sigs.k8s.io/secrets-store-csi-driver/apis/v1"

	corev1 "k8s.io/api/core/v1"
)

const (
	certType          = "CERTIFICATE"
	privateKeyType    = "PRIVATE KEY"
	privateKeyTypeRSA = "RSA PRIVATE KEY"
	privateKeyTypeEC  = "EC PRIVATE KEY"
)

const (
	formatJSON      = "json"
	formatPlaintext = "plaintext"
	FormatAuto      = "auto"
)

// getCertPart returns the certificate or the private key part of the cert
func GetCertPart(data []byte, key string) ([]byte, error) {
	if key == corev1.TLSPrivateKeyKey {
		return getPrivateKey(data)
	}
	if key == corev1.TLSCertKey {
		return getCert(data)
	}
	return nil, fmt.Errorf("key '%s' is not supported. Only 'tls.key' and 'tls.crt' are supported", key)
}

// getCert returns the certificate part of a cert
func getCert(data []byte) ([]byte, error) {
	var certs []byte
	for {
		pemBlock, rest := pem.Decode(data)
		if pemBlock == nil {
			break
		}
		if pemBlock.Type == certType {
			block := pem.EncodeToMemory(pemBlock)
			certs = append(certs, block...)
		}
		data = rest
	}
	return certs, nil
}

// getPrivateKey returns the private key part of a cert
func getPrivateKey(data []byte) ([]byte, error) {
	var der, derKey []byte
	privKeyType := privateKeyType

	for {
		pemBlock, rest := pem.Decode(data)
		if pemBlock == nil {
			break
		}
		if pemBlock.Type != certType {
			der = pemBlock.Bytes
		}
		data = rest
	}

	// parses an RSA private key in PKCS #1, ASN.1 DER form
	if key, err := x509.ParsePKCS1PrivateKey(der); err == nil {
		privKeyType = privateKeyTypeRSA
		derKey = x509.MarshalPKCS1PrivateKey(key)
	}
	// parses an unencrypted private key in PKCS #8, ASN.1 DER form
	if key, err := x509.ParsePKCS8PrivateKey(der); err == nil {
		switch key := key.(type) {
		case *rsa.PrivateKey:
			derKey = x509.MarshalPKCS1PrivateKey(key)
			privKeyType = privateKeyTypeRSA
		case *ecdsa.PrivateKey:
			derKey, err = x509.MarshalECPrivateKey(key)
			privKeyType = privateKeyTypeEC
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unknown private key type found while getting key. Only rsa and ecdsa are supported")
		}
	}
	// parses an EC private key in SEC 1, ASN.1 DER form
	if key, err := x509.ParseECPrivateKey(der); err == nil {
		derKey, err = x509.MarshalECPrivateKey(key)
		if err != nil {
			return nil, err
		}
		privKeyType = privateKeyTypeEC
	}
	block := &pem.Block{
		Type:  privKeyType,
		Bytes: derKey,
	}

	return pem.EncodeToMemory(block), nil
}

// GetSecretType returns a k8s secret type.
// Kubernetes doesn't impose any constraints on the type name: https://kubernetes.io/docs/concepts/configuration/secret/#secret-types
// If the secret type is empty, then default is Opaque.
func GetSecretType(secretType, secretName string, syncOptions secretsstorev1.SyncOptions) corev1.SecretType {
	// if secretType == "" {
	// 	return corev1.SecretTypeOpaque
	// }
	// return corev1.SecretType(secretType)

	if secretType != "" {
		return corev1.SecretType(secretType)
	}

	for _, secret := range syncOptions.Secrets {
		if secret.SecretName == secretName && secret.Type == secretType {
			return corev1.SecretType(secret.Type)
		}
	}

	if syncOptions.Type != "" {
		return corev1.SecretType(syncOptions.Type)
	}

	return corev1.SecretTypeOpaque
}

// ValidateSecretObject performs basic validation of the secret provider class
// secret object to check if the mandatory fields - name, type and data are defined
func ValidateSecretObject(secretObj secretsstorev1.SecretObject) error {
	if len(secretObj.SecretName) == 0 {
		return fmt.Errorf("secret name is empty")
	}
	if len(secretObj.Data) == 0 {
		return fmt.Errorf("data is empty")
	}
	return nil
}

// GetSecretData gets the object contents from the pods target path and returns a
// map that will be populated in the Kubernetes secret data field
func GetSecretData(secretObjData []*secretsstorev1.SecretObjectData, secretType corev1.SecretType, files map[string]string, format, jsonPath string) (map[string][]byte, error) {
	datamap := make(map[string][]byte)
	for _, data := range secretObjData {
		objectName := strings.TrimSpace(data.ObjectName)
		dataKey := strings.TrimSpace(data.Key)

		if len(objectName) == 0 {
			return datamap, fmt.Errorf("object name in secretObjects.data is empty")
		}
		if len(dataKey) == 0 {
			return datamap, fmt.Errorf("key in secretObjects.data is empty")
		}

		file, ok := files[objectName]
		if !ok {
			return datamap, fmt.Errorf("file matching objectName %s not found in the pod", objectName)
		}
		content, err := os.ReadFile(file)
		if err != nil {
			return datamap, fmt.Errorf("failed to read file %s, err: %w", objectName, err)
		}

		// TODO (manedurphy) Take auto-detection into consideration
		switch format {
		case formatJSON:
			var (
				jsonContent map[string]interface{}
				valBytes    []byte
				err         error
			)

			if err = json.Unmarshal(content, &jsonContent); err == nil {
				if jsonPath != "" {
					var (
						jsonPathSplit []string
						valid         bool
					)
					jsonPathSplit = strings.Split(jsonPath, ".")[1:]
					for _, path := range jsonPathSplit {
						if jsonContent, valid = jsonContent[path].(map[string]interface{}); !valid {
							return datamap, fmt.Errorf("invalid json path: %s", jsonPath)
						}
					}
				}

				// extract key-value pairs from the json object file content
				for key, val := range jsonContent {
					switch val := val.(type) {
					case string:
						valBytes = []byte(val)
					default:
						// we can marshal non-string types to get unquoted values as well as handle nested objects
						if valBytes, err = json.Marshal(val); err != nil {
							return datamap, fmt.Errorf("failed to marshal value %v, err: %w", val, err)
						}
					}
					datamap[key] = valBytes
				}
				continue
			}
			return datamap, fmt.Errorf("failed to unmarshal JSON file contents %s, err: %w", file, err)
		default:
			// set the contents of the mounted secret as the value
			datamap[dataKey] = content
			if secretType == corev1.SecretTypeTLS {
				c, err := GetCertPart(content, dataKey)
				if err != nil {
					return datamap, fmt.Errorf("failed to get cert data from file %s, err: %w", file, err)
				}
				datamap[dataKey] = c
			}
		}
	}
	return datamap, nil
}

// GetSHAFromSecret gets SHA for the secret data
func GetSHAFromSecret(data map[string][]byte) (string, error) {
	var values []string
	for k, v := range data {
		values = append(values, k+"="+string(v[:]))
	}
	// sort the values to always obtain a deterministic SHA for
	// same content in different order
	sort.Strings(values)
	return generateSHA(strings.Join(values, ";"))
}

// generateSHA generates SHA from string
func generateSHA(data string) (string, error) {
	hasher := sha256.New()
	_, err := io.WriteString(hasher, data)
	if err != nil {
		return "", err
	}
	sha := hasher.Sum(nil)
	return fmt.Sprintf("%x", sha), nil
}

// TODO (manedurphy) Add description, write unit tests
func GetSecretFormat(secretName string, syncOptions secretsstorev1.SyncOptions) (string, error) {
	var format string

	// top level value for secret format
	if syncOptions.Format != "" {
		format = syncOptions.Format
	}
	for _, secret := range syncOptions.Secrets {
		if secret.SecretName == secretName {
			if secret.Format != "" {
				// secret format is set on a specific secret
				format = secret.Format
			}
		}
	}
	if format == "" {
		format = formatPlaintext
	}

	// TODO (manedurphy) Add support for YAML
	if format != formatPlaintext && format != formatJSON {
		return format, fmt.Errorf("unsupported secret format: %s", format)
	}

	return format, nil
}

// TODO (manedurphy) Add description, write unit tests
func GetJsonPath(secretName string, syncOptions secretsstorev1.SyncOptions) string {
	var jsonPath string

	// top level value for json path
	if syncOptions.JsonPath != "" {
		jsonPath = syncOptions.JsonPath
	}
	for _, secret := range syncOptions.Secrets {
		if secret.SecretName == secretName {
			if secret.JsonPath != "" {
				// json path is set on a specific secret
				jsonPath = secret.JsonPath
			}
		}
	}

	return jsonPath
}
