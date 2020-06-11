package securityinsight

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DataConnectorsCheckRequirementsClient is the API spec for Microsoft.SecurityInsights (Azure Security Insights)
// resource provider
type DataConnectorsCheckRequirementsClient struct {
	BaseClient
}

// NewDataConnectorsCheckRequirementsClient creates an instance of the DataConnectorsCheckRequirementsClient client.
func NewDataConnectorsCheckRequirementsClient(subscriptionID string) DataConnectorsCheckRequirementsClient {
	return NewDataConnectorsCheckRequirementsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDataConnectorsCheckRequirementsClientWithBaseURI creates an instance of the DataConnectorsCheckRequirementsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewDataConnectorsCheckRequirementsClientWithBaseURI(baseURI string, subscriptionID string) DataConnectorsCheckRequirementsClient {
	return DataConnectorsCheckRequirementsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Post get requirements state for a data connector type.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// workspaceName - the name of the workspace.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// dataConnectorsCheckRequirements - the parameters for requirements check message
func (client DataConnectorsCheckRequirementsClient) Post(ctx context.Context, resourceGroupName string, workspaceName string, operationalInsightsResourceProvider string, dataConnectorsCheckRequirements BasicDataConnectorsCheckRequirements) (result DataConnectorRequirementsState, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataConnectorsCheckRequirementsClient.Post")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.DataConnectorsCheckRequirementsClient", "Post", err.Error())
	}

	req, err := client.PostPreparer(ctx, resourceGroupName, workspaceName, operationalInsightsResourceProvider, dataConnectorsCheckRequirements)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.DataConnectorsCheckRequirementsClient", "Post", nil, "Failure preparing request")
		return
	}

	resp, err := client.PostSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.DataConnectorsCheckRequirementsClient", "Post", resp, "Failure sending request")
		return
	}

	result, err = client.PostResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.DataConnectorsCheckRequirementsClient", "Post", resp, "Failure responding to request")
	}

	return
}

// PostPreparer prepares the Post request.
func (client DataConnectorsCheckRequirementsClient) PostPreparer(ctx context.Context, resourceGroupName string, workspaceName string, operationalInsightsResourceProvider string, dataConnectorsCheckRequirements BasicDataConnectorsCheckRequirements) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/dataConnectorsCheckRequirements", pathParameters),
		autorest.WithJSON(dataConnectorsCheckRequirements),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PostSender sends the Post request. The method will close the
// http.Response Body if it receives an error.
func (client DataConnectorsCheckRequirementsClient) PostSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// PostResponder handles the response to the Post request. The method always
// closes the http.Response Body.
func (client DataConnectorsCheckRequirementsClient) PostResponder(resp *http.Response) (result DataConnectorRequirementsState, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
