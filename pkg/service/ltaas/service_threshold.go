package ltaas

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetThresholds retrieves a list of thresholds
func (s *Service) GetThresholds(parameters connection.APIRequestParameters) ([]Threshold, error) {
	var thresholds []Threshold

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetThresholdsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, threshold := range response.(*PaginatedThreshold).Items {
			thresholds = append(thresholds, threshold)
		}
	}

	return thresholds, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetThresholdsPaginated retrieves a paginated list of thresholds
func (s *Service) GetThresholdsPaginated(parameters connection.APIRequestParameters) (*PaginatedThreshold, error) {
	body, err := s.getThresholdsPaginatedResponseBody(parameters)

	return NewPaginatedThreshold(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetThresholdsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getThresholdsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetThresholdSliceResponseBody, error) {
	body := &GetThresholdSliceResponseBody{}

	response, err := s.connection.Get("/ltaas/v1/thresholds", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetThreshold retrieves a single threshold by id
func (s *Service) GetThreshold(thresholdID string) (Threshold, error) {
	body, err := s.getThresholdResponseBody(thresholdID)

	return body.Data, err
}

func (s *Service) getThresholdResponseBody(thresholdID string) (*GetThresholdResponseBody, error) {
	body := &GetThresholdResponseBody{}

	if thresholdID == "" {
		return body, fmt.Errorf("invalid threshold id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ltaas/v1/thresholds/%s", thresholdID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ThresholdNotFoundError{ID: thresholdID}
		}

		return nil
	})
}
