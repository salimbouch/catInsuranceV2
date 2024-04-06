/*
 * Cat Insurance API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
	"errors"
)

// CustomerAPIService is a service that implements the logic for the CustomerAPIServicer
// This service should implement the business logic for every endpoint for the CustomerAPI API.
// Include any external packages or services that will be required by this service.
type CustomerAPIService struct {
}

// NewCustomerAPIService creates a default api service
func NewCustomerAPIService() CustomerAPIServicer {
	return &CustomerAPIService{}
}

// CreateCustomer - Create a new customer
func (s *CustomerAPIService) CreateCustomer(ctx context.Context, customerReq CustomerReq) (ImplResponse, error) {
	// TODO - update CreateCustomer with the required logic for this service method.
	// Add api_customer_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, CustomerRes{}) or use other options such as http.Ok ...
	// return Response(201, CustomerRes{}), nil

	// TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	// return Response(400, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateCustomer method not implemented")
}

// DeleteCustomer - Delete a customer
func (s *CustomerAPIService) DeleteCustomer(ctx context.Context, customerId string) (ImplResponse, error) {
	// TODO - update DeleteCustomer with the required logic for this service method.
	// Add api_customer_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	// return Response(200, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteCustomer method not implemented")
}

// GetCustomer - Get customer details
func (s *CustomerAPIService) GetCustomer(ctx context.Context, customerId string) (ImplResponse, error) {
	// TODO - update GetCustomer with the required logic for this service method.
	// Add api_customer_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, CustomerRes{}) or use other options such as http.Ok ...
	// return Response(200, CustomerRes{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetCustomer method not implemented")
}

// GetCustomerContracts - Get customer contracts
func (s *CustomerAPIService) GetCustomerContracts(ctx context.Context, customerId string, page int32, pageSize int32) (ImplResponse, error) {
	// TODO - update GetCustomerContracts with the required logic for this service method.
	// Add api_customer_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []ContractRes{}) or use other options such as http.Ok ...
	// return Response(200, []ContractRes{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetCustomerContracts method not implemented")
}

// GetCustomers - Get all customers
func (s *CustomerAPIService) GetCustomers(ctx context.Context, page int32, pageSize int32) (ImplResponse, error) {
	// TODO - update GetCustomers with the required logic for this service method.
	// Add api_customer_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []CustomerRes{}) or use other options such as http.Ok ...
	// return Response(200, []CustomerRes{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetCustomers method not implemented")
}

// SearchCustomers - Search for customers
func (s *CustomerAPIService) SearchCustomers(ctx context.Context, text string, page int32, pageSize int32) (ImplResponse, error) {
	// TODO - update SearchCustomers with the required logic for this service method.
	// Add api_customer_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []CustomerRes{}) or use other options such as http.Ok ...
	// return Response(200, []CustomerRes{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("SearchCustomers method not implemented")
}

// UpdateCustomer - Update a customer
func (s *CustomerAPIService) UpdateCustomer(ctx context.Context, customerId string, customerReq CustomerReq) (ImplResponse, error) {
	// TODO - update UpdateCustomer with the required logic for this service method.
	// Add api_customer_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	// return Response(200, nil),nil

	// TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	// return Response(400, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateCustomer method not implemented")
}