// Code generated by github.com/ukfast/sdk-go/pkg/gen/model_paginated. DO NOT EDIT.
package threatmonitoring

import "github.com/ukfast/sdk-go/pkg/connection"

// PaginatedAgent represents a paginated collection of Agent
type PaginatedAgent struct {
	*connection.PaginatedBase

	Items []Agent
}

// NewPaginatedAgent returns a pointer to an initialized PaginatedAgent struct
func NewPaginatedAgent(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Agent) *PaginatedAgent {
	return &PaginatedAgent{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}
