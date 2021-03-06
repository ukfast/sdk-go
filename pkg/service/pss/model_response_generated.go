package pss

import "github.com/ukfast/sdk-go/pkg/connection"

// GetRequestSliceResponseBody represents an API response body containing []Request data
type GetRequestSliceResponseBody struct {
	connection.APIResponseBody
	Data []Request `json:"data"`
}

// GetRequestResponseBody represents an API response body containing Request data
type GetRequestResponseBody struct {
	connection.APIResponseBody
	Data Request `json:"data"`
}

// GetReplySliceResponseBody represents an API response body containing []Reply data
type GetReplySliceResponseBody struct {
	connection.APIResponseBody
	Data []Reply `json:"data"`
}

// GetReplyResponseBody represents an API response body containing Reply data
type GetReplyResponseBody struct {
	connection.APIResponseBody
	Data Reply `json:"data"`
}
