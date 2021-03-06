package ltaas

import "github.com/ukfast/sdk-go/pkg/connection"

// GetDomainSliceResponseBody represents an API response body containing []Domain data
type GetDomainSliceResponseBody struct {
	connection.APIResponseBody
	Data []Domain `json:"data"`
}

// GetDomainResponseBody represents an API response body containing Domain data
type GetDomainResponseBody struct {
	connection.APIResponseBody
	Data Domain `json:"data"`
}

// GetTestSliceResponseBody represents an API response body containing []Test data
type GetTestSliceResponseBody struct {
	connection.APIResponseBody
	Data []Test `json:"data"`
}

// GetTestResponseBody represents an API response body containing Test data
type GetTestResponseBody struct {
	connection.APIResponseBody
	Data Test `json:"data"`
}

// GetJobSliceResponseBody represents an API response body containing []Job data
type GetJobSliceResponseBody struct {
	connection.APIResponseBody
	Data []Job `json:"data"`
}

// GetJobResponseBody represents an API response body containing Job data
type GetJobResponseBody struct {
	connection.APIResponseBody
	Data Job `json:"data"`
}

// GetJobResultsSliceResponseBody represents an API response body containing []JobResults data
type GetJobResultsSliceResponseBody struct {
	connection.APIResponseBody
	Data []JobResults `json:"data"`
}

// GetJobResultsResponseBody represents an API response body containing JobResults data
type GetJobResultsResponseBody struct {
	connection.APIResponseBody
	Data JobResults `json:"data"`
}

// GetJobSettingsSliceResponseBody represents an API response body containing []JobSettings data
type GetJobSettingsSliceResponseBody struct {
	connection.APIResponseBody
	Data []JobSettings `json:"data"`
}

// GetJobSettingsResponseBody represents an API response body containing JobSettings data
type GetJobSettingsResponseBody struct {
	connection.APIResponseBody
	Data JobSettings `json:"data"`
}

// GetThresholdSliceResponseBody represents an API response body containing []Threshold data
type GetThresholdSliceResponseBody struct {
	connection.APIResponseBody
	Data []Threshold `json:"data"`
}

// GetThresholdResponseBody represents an API response body containing Threshold data
type GetThresholdResponseBody struct {
	connection.APIResponseBody
	Data Threshold `json:"data"`
}

// GetScenarioSliceResponseBody represents an API response body containing []Scenario data
type GetScenarioSliceResponseBody struct {
	connection.APIResponseBody
	Data []Scenario `json:"data"`
}

// GetScenarioResponseBody represents an API response body containing Scenario data
type GetScenarioResponseBody struct {
	connection.APIResponseBody
	Data Scenario `json:"data"`
}

// GetAgreementSliceResponseBody represents an API response body containing []Agreement data
type GetAgreementSliceResponseBody struct {
	connection.APIResponseBody
	Data []Agreement `json:"data"`
}

// GetAgreementResponseBody represents an API response body containing Agreement data
type GetAgreementResponseBody struct {
	connection.APIResponseBody
	Data Agreement `json:"data"`
}

// GetAccountSliceResponseBody represents an API response body containing []Account data
type GetAccountSliceResponseBody struct {
	connection.APIResponseBody
	Data []Account `json:"data"`
}

// GetAccountResponseBody represents an API response body containing Account data
type GetAccountResponseBody struct {
	connection.APIResponseBody
	Data Account `json:"data"`
}
