// Code generated by github.com/ukfast/sdk-go/pkg/gen/response. DO NOT EDIT.
package storage

import "github.com/ukfast/sdk-go/pkg/connection"

// GetSolutionSliceResponseBody represents an API response body containing []Solution data
type GetSolutionSliceResponseBody struct {
	connection.APIResponseBody

	Data []Solution `json:"data"`
}

// GetSolutionResponseBody represents an API response body containing Solution data
type GetSolutionResponseBody struct {
	connection.APIResponseBody

	Data Solution `json:"data"`
}

// GetHostSliceResponseBody represents an API response body containing []Host data
type GetHostSliceResponseBody struct {
	connection.APIResponseBody

	Data []Host `json:"data"`
}

// GetHostResponseBody represents an API response body containing Host data
type GetHostResponseBody struct {
	connection.APIResponseBody

	Data Host `json:"data"`
}

// GetHostSetSliceResponseBody represents an API response body containing []HostSet data
type GetHostSetSliceResponseBody struct {
	connection.APIResponseBody

	Data []HostSet `json:"data"`
}

// GetHostSetResponseBody represents an API response body containing HostSet data
type GetHostSetResponseBody struct {
	connection.APIResponseBody

	Data HostSet `json:"data"`
}

// GetVolumeSliceResponseBody represents an API response body containing []Volume data
type GetVolumeSliceResponseBody struct {
	connection.APIResponseBody

	Data []Volume `json:"data"`
}

// GetVolumeResponseBody represents an API response body containing Volume data
type GetVolumeResponseBody struct {
	connection.APIResponseBody

	Data Volume `json:"data"`
}

// GetVolumeSetSliceResponseBody represents an API response body containing []VolumeSet data
type GetVolumeSetSliceResponseBody struct {
	connection.APIResponseBody

	Data []VolumeSet `json:"data"`
}

// GetVolumeSetResponseBody represents an API response body containing VolumeSet data
type GetVolumeSetResponseBody struct {
	connection.APIResponseBody

	Data VolumeSet `json:"data"`
}

// GetIOPSSliceResponseBody represents an API response body containing []IOPS data
type GetIOPSSliceResponseBody struct {
	connection.APIResponseBody

	Data []IOPS `json:"data"`
}

// GetIOPSResponseBody represents an API response body containing IOPS data
type GetIOPSResponseBody struct {
	connection.APIResponseBody

	Data IOPS `json:"data"`
}
