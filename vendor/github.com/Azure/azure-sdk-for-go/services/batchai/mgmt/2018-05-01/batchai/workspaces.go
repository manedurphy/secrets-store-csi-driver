package batchai

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// WorkspacesClient is the the Azure BatchAI Management API.
type WorkspacesClient struct {
	BaseClient
}

// NewWorkspacesClient creates an instance of the WorkspacesClient client.
func NewWorkspacesClient(subscriptionID string) WorkspacesClient {
	return NewWorkspacesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkspacesClientWithBaseURI creates an instance of the WorkspacesClient client.
func NewWorkspacesClientWithBaseURI(baseURI string, subscriptionID string) WorkspacesClient {
	return WorkspacesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a Workspace.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// workspaceName - the name of the workspace. Workspace names can only contain a combination of alphanumeric
// characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
// parameters - workspace creation parameters.
func (client WorkspacesClient) Create(ctx context.Context, resourceGroupName string, workspaceName string, parameters WorkspaceCreateParameters) (result WorkspacesCreateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[-\w_]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Location", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batchai.WorkspacesClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, workspaceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.WorkspacesClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.WorkspacesClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client WorkspacesClient) CreatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, parameters WorkspaceCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspacesClient) CreateSender(req *http.Request) (future WorkspacesCreateFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client WorkspacesClient) CreateResponder(resp *http.Response) (result Workspace, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a Workspace.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// workspaceName - the name of the workspace. Workspace names can only contain a combination of alphanumeric
// characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
func (client WorkspacesClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string) (result WorkspacesDeleteFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[-\w_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batchai.WorkspacesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.WorkspacesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.WorkspacesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client WorkspacesClient) DeletePreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspacesClient) DeleteSender(req *http.Request) (future WorkspacesDeleteFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client WorkspacesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about a Workspace.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// workspaceName - the name of the workspace. Workspace names can only contain a combination of alphanumeric
// characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
func (client WorkspacesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string) (result Workspace, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[-\w_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batchai.WorkspacesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.WorkspacesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.WorkspacesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.WorkspacesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkspacesClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspacesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkspacesClient) GetResponder(resp *http.Response) (result Workspace, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of Workspaces associated with the given subscription.
// Parameters:
// maxResults - the maximum number of items to return in the response. A maximum of 1000 files can be returned.
func (client WorkspacesClient) List(ctx context.Context, maxResults *int32) (result WorkspaceListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: maxResults,
			Constraints: []validation.Constraint{{Target: "maxResults", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "maxResults", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "maxResults", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("batchai.WorkspacesClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, maxResults)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.WorkspacesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.wlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.WorkspacesClient", "List", resp, "Failure sending request")
		return
	}

	result.wlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.WorkspacesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client WorkspacesClient) ListPreparer(ctx context.Context, maxResults *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if maxResults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxResults)
	} else {
		queryParameters["maxresults"] = autorest.Encode("query", 1000)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.BatchAI/workspaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspacesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WorkspacesClient) ListResponder(resp *http.Response) (result WorkspaceListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client WorkspacesClient) listNextResults(lastResults WorkspaceListResult) (result WorkspaceListResult, err error) {
	req, err := lastResults.workspaceListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batchai.WorkspacesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batchai.WorkspacesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.WorkspacesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client WorkspacesClient) ListComplete(ctx context.Context, maxResults *int32) (result WorkspaceListResultIterator, err error) {
	result.page, err = client.List(ctx, maxResults)
	return
}

// ListByResourceGroup gets a list of Workspaces within the specified resource group.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// maxResults - the maximum number of items to return in the response. A maximum of 1000 files can be returned.
func (client WorkspacesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, maxResults *int32) (result WorkspaceListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: maxResults,
			Constraints: []validation.Constraint{{Target: "maxResults", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "maxResults", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "maxResults", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("batchai.WorkspacesClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, maxResults)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.WorkspacesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.wlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.WorkspacesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.wlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.WorkspacesClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client WorkspacesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, maxResults *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if maxResults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxResults)
	} else {
		queryParameters["maxresults"] = autorest.Encode("query", 1000)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspacesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client WorkspacesClient) ListByResourceGroupResponder(resp *http.Response) (result WorkspaceListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client WorkspacesClient) listByResourceGroupNextResults(lastResults WorkspaceListResult) (result WorkspaceListResult, err error) {
	req, err := lastResults.workspaceListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batchai.WorkspacesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batchai.WorkspacesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.WorkspacesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client WorkspacesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, maxResults *int32) (result WorkspaceListResultIterator, err error) {
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, maxResults)
	return
}

// Update updates properties of a Workspace.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// workspaceName - the name of the workspace. Workspace names can only contain a combination of alphanumeric
// characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
// parameters - additional parameters for workspace update.
func (client WorkspacesClient) Update(ctx context.Context, resourceGroupName string, workspaceName string, parameters WorkspaceUpdateParameters) (result Workspace, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[-\w_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batchai.WorkspacesClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, workspaceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.WorkspacesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.WorkspacesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.WorkspacesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client WorkspacesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, parameters WorkspaceUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspacesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client WorkspacesClient) UpdateResponder(resp *http.Response) (result Workspace, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
