package apimanagement

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
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// ProductGroupsClient is the use these REST APIs for performing operations on
// entities like API, Product, and Subscription associated with your Azure API
// Management deployment.
type ProductGroupsClient struct {
	ManagementClient
}

// NewProductGroupsClient creates an instance of the ProductGroupsClient
// client.
func NewProductGroupsClient(subscriptionID string) ProductGroupsClient {
	return NewProductGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProductGroupsClientWithBaseURI creates an instance of the
// ProductGroupsClient client.
func NewProductGroupsClientWithBaseURI(baseURI string, subscriptionID string) ProductGroupsClient {
	return ProductGroupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Add adds the association between the specified developer group with the
// specified product.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. productID is product identifier. Must be
// unique in the current API Management service instance. groupID is group
// identifier. Must be unique in the current API Management service instance.
func (client ProductGroupsClient) Add(resourceGroupName string, serviceName string, productID string, groupID string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: productID,
			Constraints: []validation.Constraint{{Target: "productID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "productID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "productID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: groupID,
			Constraints: []validation.Constraint{{Target: "groupID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "groupID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "groupID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.ProductGroupsClient", "Add")
	}

	req, err := client.AddPreparer(resourceGroupName, serviceName, productID, groupID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.ProductGroupsClient", "Add", nil, "Failure preparing request")
	}

	resp, err := client.AddSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "apimanagement.ProductGroupsClient", "Add", resp, "Failure sending request")
	}

	result, err = client.AddResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ProductGroupsClient", "Add", resp, "Failure responding to request")
	}

	return
}

// AddPreparer prepares the Add request.
func (client ProductGroupsClient) AddPreparer(resourceGroupName string, serviceName string, productID string, groupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId":           autorest.Encode("path", groupID),
		"productId":         autorest.Encode("path", productID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/groups/{groupId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// AddSender sends the Add request. The method will close the
// http.Response Body if it receives an error.
func (client ProductGroupsClient) AddSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// AddResponder handles the response to the Add request. The method always
// closes the http.Response Body.
func (client ProductGroupsClient) AddResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByProduct lists the collection of developer groups associated with the
// specified product.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. productID is product identifier. Must be
// unique in the current API Management service instance. filter is | Field
// | Supported operators    | Supported functions                         |
// |-------------|------------------------|---------------------------------------------|
// | id          | ge, le, eq, ne, gt, lt | substringof, contains, startswith,
// endswith |
// | name        | ge, le, eq, ne, gt, lt | substringof, contains, startswith,
// endswith |
// | description | ge, le, eq, ne, gt, lt | substringof, contains, startswith,
// endswith |
// | type        | eq, ne                 | N/A
// | top is number of records to return. skip is number of records to skip.
func (client ProductGroupsClient) ListByProduct(resourceGroupName string, serviceName string, productID string, filter string, top *int32, skip *int32) (result GroupCollection, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: productID,
			Constraints: []validation.Constraint{{Target: "productID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "productID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "productID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.ProductGroupsClient", "ListByProduct")
	}

	req, err := client.ListByProductPreparer(resourceGroupName, serviceName, productID, filter, top, skip)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.ProductGroupsClient", "ListByProduct", nil, "Failure preparing request")
	}

	resp, err := client.ListByProductSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.ProductGroupsClient", "ListByProduct", resp, "Failure sending request")
	}

	result, err = client.ListByProductResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ProductGroupsClient", "ListByProduct", resp, "Failure responding to request")
	}

	return
}

// ListByProductPreparer prepares the ListByProduct request.
func (client ProductGroupsClient) ListByProductPreparer(resourceGroupName string, serviceName string, productID string, filter string, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"productId":         autorest.Encode("path", productID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/groups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByProductSender sends the ListByProduct request. The method will close the
// http.Response Body if it receives an error.
func (client ProductGroupsClient) ListByProductSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByProductResponder handles the response to the ListByProduct request. The method always
// closes the http.Response Body.
func (client ProductGroupsClient) ListByProductResponder(resp *http.Response) (result GroupCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByProductNextResults retrieves the next set of results, if any.
func (client ProductGroupsClient) ListByProductNextResults(lastResults GroupCollection) (result GroupCollection, err error) {
	req, err := lastResults.GroupCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.ProductGroupsClient", "ListByProduct", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByProductSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.ProductGroupsClient", "ListByProduct", resp, "Failure sending next results request")
	}

	result, err = client.ListByProductResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ProductGroupsClient", "ListByProduct", resp, "Failure responding to next results request")
	}

	return
}

// Remove deletes the association between the specified group and product.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. productID is product identifier. Must be
// unique in the current API Management service instance. groupID is group
// identifier. Must be unique in the current API Management service instance.
func (client ProductGroupsClient) Remove(resourceGroupName string, serviceName string, productID string, groupID string) (result ErrorBodyContract, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: productID,
			Constraints: []validation.Constraint{{Target: "productID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "productID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "productID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: groupID,
			Constraints: []validation.Constraint{{Target: "groupID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "groupID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "groupID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.ProductGroupsClient", "Remove")
	}

	req, err := client.RemovePreparer(resourceGroupName, serviceName, productID, groupID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.ProductGroupsClient", "Remove", nil, "Failure preparing request")
	}

	resp, err := client.RemoveSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.ProductGroupsClient", "Remove", resp, "Failure sending request")
	}

	result, err = client.RemoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.ProductGroupsClient", "Remove", resp, "Failure responding to request")
	}

	return
}

// RemovePreparer prepares the Remove request.
func (client ProductGroupsClient) RemovePreparer(resourceGroupName string, serviceName string, productID string, groupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId":           autorest.Encode("path", groupID),
		"productId":         autorest.Encode("path", productID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/groups/{groupId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// RemoveSender sends the Remove request. The method will close the
// http.Response Body if it receives an error.
func (client ProductGroupsClient) RemoveSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// RemoveResponder handles the response to the Remove request. The method always
// closes the http.Response Body.
func (client ProductGroupsClient) RemoveResponder(resp *http.Response) (result ErrorBodyContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusBadRequest),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
