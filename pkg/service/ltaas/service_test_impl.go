package ltaas

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetTests retrieves a list of tests
func (s *Service) GetTests(parameters connection.APIRequestParameters) ([]Test, error) {
	var sites []Test

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetTestsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, site := range response.(*PaginatedTest).Items {
			sites = append(sites, site)
		}
	}

	return sites, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetTestsPaginated retrieves a paginated list of tests
func (s *Service) GetTestsPaginated(parameters connection.APIRequestParameters) (*PaginatedTest, error) {
	body, err := s.getTestsPaginatedResponseBody(parameters)

	return NewPaginatedTest(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetTestsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getTestsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetTestsResponseBody, error) {
	body := &GetTestsResponseBody{}

	response, err := s.connection.Get("/ltaas/v1/tests", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetTest retrieves a single test by id
func (s *Service) GetTest(testID string) (Test, error) {
	body, err := s.getTestResponseBody(testID)

	return body.Data, err
}

func (s *Service) getTestResponseBody(testID string) (*GetTestResponseBody, error) {
	body := &GetTestResponseBody{}

	if testID == "" {
		return body, fmt.Errorf("invalid test id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ltaas/v1/tests/%s", testID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TestNotFoundError{ID: testID}
		}

		return nil
	})
}
