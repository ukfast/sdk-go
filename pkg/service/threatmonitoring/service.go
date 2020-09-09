package threatmonitoring

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// ThreatMonitoringService is an interface for managing Threat Monitoring
type ThreatMonitoringService interface {
	GetAgents(parameters connection.APIRequestParameters) ([]Agent, error)
	GetAgentsPaginated(parameters connection.APIRequestParameters) (*PaginatedAgent, error)
	GetAgent(agentID string) (Agent, error)

	GetAlerts(parameters connection.APIRequestParameters) ([]Alert, error)
	GetAlertsPaginated(parameters connection.APIRequestParameters) (*PaginatedAlert, error)
	GetAlert(alertID string) (Alert, error)
}

// Service implements ThreatMonitoringService for managing the Threat Monitoring service
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of ThreatMonitoringService
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}