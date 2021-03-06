package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetDHCPs retrieves a list of dhcps
func (s *Service) GetDHCPs(parameters connection.APIRequestParameters) ([]DHCP, error) {
	var dhcps []DHCP

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetDHCPsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, dhcp := range response.(*PaginatedDHCP).Items {
			dhcps = append(dhcps, dhcp)
		}
	}

	return dhcps, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetDHCPsPaginated retrieves a paginated list of dhcps
func (s *Service) GetDHCPsPaginated(parameters connection.APIRequestParameters) (*PaginatedDHCP, error) {
	body, err := s.getDHCPsPaginatedResponseBody(parameters)

	return NewPaginatedDHCP(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetDHCPsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getDHCPsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetDHCPSliceResponseBody, error) {
	body := &GetDHCPSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/dhcps", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetDHCP retrieves a single dhcp by id
func (s *Service) GetDHCP(dhcpID string) (DHCP, error) {
	body, err := s.getDHCPResponseBody(dhcpID)

	return body.Data, err
}

func (s *Service) getDHCPResponseBody(dhcpID string) (*GetDHCPResponseBody, error) {
	body := &GetDHCPResponseBody{}

	if dhcpID == "" {
		return body, fmt.Errorf("invalid dhcp id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/dhcps/%s", dhcpID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DHCPNotFoundError{ID: dhcpID}
		}

		return nil
	})
}
