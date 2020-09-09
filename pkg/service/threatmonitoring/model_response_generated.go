package threatmonitoring

import "github.com/ukfast/sdk-go/pkg/connection"

// GetAgentSliceResponseBody represents an API response body containing []Agent data
type GetAgentSliceResponseBody struct {
	connection.APIResponseBody
	Data []Agent `json:"data"`
}

// GetAgentResponseBody represents an API response body containing Agent data
type GetAgentResponseBody struct {
	connection.APIResponseBody
	Data Agent `json:"data"`
}

// GetAlertSliceResponseBody represents an API response body containing []Alert data
type GetAlertSliceResponseBody struct {
	connection.APIResponseBody
	Data []Alert `json:"data"`
}

// GetAlertResponseBody represents an API response body containing Alert data
type GetAlertResponseBody struct {
	connection.APIResponseBody
	Data Alert `json:"data"`
}