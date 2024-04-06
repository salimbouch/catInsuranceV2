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

// EmployeeAPIController binds http requests to an api service and writes the service results to the http response
type EmployeeAPIController struct {
	service EmployeeAPIServicer
	errorHandler ErrorHandler
}

// EmployeeAPIOption for how the controller is set up.
type EmployeeAPIOption func(*EmployeeAPIController)

// WithEmployeeAPIErrorHandler inject ErrorHandler into controller
func WithEmployeeAPIErrorHandler(h ErrorHandler) EmployeeAPIOption {
	return func(c *EmployeeAPIController) {
		c.errorHandler = h
	}
}

// NewEmployeeAPIController creates a default api controller
func NewEmployeeAPIController(s EmployeeAPIServicer, opts ...EmployeeAPIOption) Router {
	controller := &EmployeeAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the EmployeeAPIController
func (c *EmployeeAPIController) Routes() Routes {
	return Routes{
		"CreateEmployee": Route{
			strings.ToUpper("Post"),
			"/v1/employees",
			c.CreateEmployee,
		},
		"GetEmployee": Route{
			strings.ToUpper("Get"),
			"/v1/employees/{employeeId}",
			c.GetEmployee,
		},
		"UpdateEmployee": Route{
			strings.ToUpper("Patch"),
			"/v1/employees",
			c.UpdateEmployee,
		},
	}
}

// CreateEmployee - Create a new employee
func (c *EmployeeAPIController) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	employeeReqParam := EmployeeReq{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&employeeReqParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertEmployeeReqRequired(employeeReqParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertEmployeeReqConstraints(employeeReqParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateEmployee(r.Context(), employeeReqParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetEmployee - Get employee details
func (c *EmployeeAPIController) GetEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	employeeIdParam := params["employeeId"]
	if employeeIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"employeeId"}, nil)
		return
	}
	result, err := c.service.GetEmployee(r.Context(), employeeIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateEmployee - Update an employee
func (c *EmployeeAPIController) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	employeeReqParam := EmployeeReq{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&employeeReqParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertEmployeeReqRequired(employeeReqParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertEmployeeReqConstraints(employeeReqParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateEmployee(r.Context(), employeeReqParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}