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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// CustomerAPIController binds http requests to an api service and writes the service results to the http response
type CustomerAPIController struct {
	service CustomerAPIServicer
	errorHandler ErrorHandler
}

// CustomerAPIOption for how the controller is set up.
type CustomerAPIOption func(*CustomerAPIController)

// WithCustomerAPIErrorHandler inject ErrorHandler into controller
func WithCustomerAPIErrorHandler(h ErrorHandler) CustomerAPIOption {
	return func(c *CustomerAPIController) {
		c.errorHandler = h
	}
}

// NewCustomerAPIController creates a default api controller
func NewCustomerAPIController(s CustomerAPIServicer, opts ...CustomerAPIOption) Router {
	controller := &CustomerAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the CustomerAPIController
func (c *CustomerAPIController) Routes() Routes {
	return Routes{
		"CreateCustomer": Route{
			strings.ToUpper("Post"),
			"/v1/customers",
			c.CreateCustomer,
		},
		"DeleteCustomer": Route{
			strings.ToUpper("Delete"),
			"/v1/customers/{customerId}",
			c.DeleteCustomer,
		},
		"GetCustomer": Route{
			strings.ToUpper("Get"),
			"/v1/customers/{customerId}",
			c.GetCustomer,
		},
		"GetCustomerContracts": Route{
			strings.ToUpper("Get"),
			"/v1/customers/{customerId}/contracts",
			c.GetCustomerContracts,
		},
		"GetCustomers": Route{
			strings.ToUpper("Get"),
			"/v1/customers",
			c.GetCustomers,
		},
		"SearchCustomers": Route{
			strings.ToUpper("Get"),
			"/v1/customers/search",
			c.SearchCustomers,
		},
		"UpdateCustomer": Route{
			strings.ToUpper("Patch"),
			"/v1/customers/{customerId}",
			c.UpdateCustomer,
		},
	}
}

// CreateCustomer - Create a new customer
func (c *CustomerAPIController) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	customerReqParam := CustomerReq{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&customerReqParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCustomerReqRequired(customerReqParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertCustomerReqConstraints(customerReqParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateCustomer(r.Context(), customerReqParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteCustomer - Delete a customer
func (c *CustomerAPIController) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerIdParam := params["customerId"]
	if customerIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"customerId"}, nil)
		return
	}
	result, err := c.service.DeleteCustomer(r.Context(), customerIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetCustomer - Get customer details
func (c *CustomerAPIController) GetCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerIdParam := params["customerId"]
	if customerIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"customerId"}, nil)
		return
	}
	result, err := c.service.GetCustomer(r.Context(), customerIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetCustomerContracts - Get customer contracts
func (c *CustomerAPIController) GetCustomerContracts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	customerIdParam := params["customerId"]
	if customerIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"customerId"}, nil)
		return
	}
	var pageParam int32
	if query.Has("page") {
		param, err := parseNumericParameter[int32](
			query.Get("page"),
			WithParse[int32](parseInt32),
			WithMinimum[int32](1),
			WithMaximum[int32](100),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		pageParam = param
	} else {
	}
	var pageSizeParam int32
	if query.Has("pageSize") {
		param, err := parseNumericParameter[int32](
			query.Get("pageSize"),
			WithParse[int32](parseInt32),
			WithMinimum[int32](1),
			WithMaximum[int32](100),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		pageSizeParam = param
	} else {
	}
	result, err := c.service.GetCustomerContracts(r.Context(), customerIdParam, pageParam, pageSizeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetCustomers - Get all customers
func (c *CustomerAPIController) GetCustomers(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var pageParam int32
	if query.Has("page") {
		param, err := parseNumericParameter[int32](
			query.Get("page"),
			WithParse[int32](parseInt32),
			WithMinimum[int32](1),
			WithMaximum[int32](100),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		pageParam = param
	} else {
	}
	var pageSizeParam int32
	if query.Has("pageSize") {
		param, err := parseNumericParameter[int32](
			query.Get("pageSize"),
			WithParse[int32](parseInt32),
			WithMinimum[int32](1),
			WithMaximum[int32](100),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		pageSizeParam = param
	} else {
	}
	result, err := c.service.GetCustomers(r.Context(), pageParam, pageSizeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// SearchCustomers - Search for customers
func (c *CustomerAPIController) SearchCustomers(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var textParam string
	if query.Has("text") {
		param := query.Get("text")

		textParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "text"}, nil)
		return
	}
	var pageParam int32
	if query.Has("page") {
		param, err := parseNumericParameter[int32](
			query.Get("page"),
			WithParse[int32](parseInt32),
			WithMinimum[int32](1),
			WithMaximum[int32](100),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		pageParam = param
	} else {
	}
	var pageSizeParam int32
	if query.Has("pageSize") {
		param, err := parseNumericParameter[int32](
			query.Get("pageSize"),
			WithParse[int32](parseInt32),
			WithMinimum[int32](1),
			WithMaximum[int32](100),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		pageSizeParam = param
	} else {
	}
	result, err := c.service.SearchCustomers(r.Context(), textParam, pageParam, pageSizeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateCustomer - Update a customer
func (c *CustomerAPIController) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerIdParam := params["customerId"]
	if customerIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"customerId"}, nil)
		return
	}
	customerReqParam := CustomerReq{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&customerReqParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCustomerReqRequired(customerReqParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertCustomerReqConstraints(customerReqParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateCustomer(r.Context(), customerIdParam, customerReqParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
