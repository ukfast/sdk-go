package ecloud

import "github.com/ukfast/sdk-go/pkg/connection"

// PaginatedVirtualMachine represents a paginated collection of VirtualMachine
type PaginatedVirtualMachine struct {
	*connection.PaginatedBase
	Items []VirtualMachine
}

// NewPaginatedVirtualMachine returns a pointer to an initialized PaginatedVirtualMachine struct
func NewPaginatedVirtualMachine(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []VirtualMachine) *PaginatedVirtualMachine {
	return &PaginatedVirtualMachine{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedTag represents a paginated collection of Tag
type PaginatedTag struct {
	*connection.PaginatedBase
	Items []Tag
}

// NewPaginatedTag returns a pointer to an initialized PaginatedTag struct
func NewPaginatedTag(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Tag) *PaginatedTag {
	return &PaginatedTag{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedSolution represents a paginated collection of Solution
type PaginatedSolution struct {
	*connection.PaginatedBase
	Items []Solution
}

// NewPaginatedSolution returns a pointer to an initialized PaginatedSolution struct
func NewPaginatedSolution(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Solution) *PaginatedSolution {
	return &PaginatedSolution{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedSite represents a paginated collection of Site
type PaginatedSite struct {
	*connection.PaginatedBase
	Items []Site
}

// NewPaginatedSite returns a pointer to an initialized PaginatedSite struct
func NewPaginatedSite(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Site) *PaginatedSite {
	return &PaginatedSite{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedNetwork represents a paginated collection of Network
type PaginatedNetwork struct {
	*connection.PaginatedBase
	Items []Network
}

// NewPaginatedNetwork returns a pointer to an initialized PaginatedNetwork struct
func NewPaginatedNetwork(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Network) *PaginatedNetwork {
	return &PaginatedNetwork{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedHost represents a paginated collection of Host
type PaginatedHost struct {
	*connection.PaginatedBase
	Items []Host
}

// NewPaginatedHost returns a pointer to an initialized PaginatedHost struct
func NewPaginatedHost(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Host) *PaginatedHost {
	return &PaginatedHost{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedDatastore represents a paginated collection of Datastore
type PaginatedDatastore struct {
	*connection.PaginatedBase
	Items []Datastore
}

// NewPaginatedDatastore returns a pointer to an initialized PaginatedDatastore struct
func NewPaginatedDatastore(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Datastore) *PaginatedDatastore {
	return &PaginatedDatastore{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedFirewall represents a paginated collection of Firewall
type PaginatedFirewall struct {
	*connection.PaginatedBase
	Items []Firewall
}

// NewPaginatedFirewall returns a pointer to an initialized PaginatedFirewall struct
func NewPaginatedFirewall(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Firewall) *PaginatedFirewall {
	return &PaginatedFirewall{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedTemplate represents a paginated collection of Template
type PaginatedTemplate struct {
	*connection.PaginatedBase
	Items []Template
}

// NewPaginatedTemplate returns a pointer to an initialized PaginatedTemplate struct
func NewPaginatedTemplate(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Template) *PaginatedTemplate {
	return &PaginatedTemplate{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedPod represents a paginated collection of Pod
type PaginatedPod struct {
	*connection.PaginatedBase
	Items []Pod
}

// NewPaginatedPod returns a pointer to an initialized PaginatedPod struct
func NewPaginatedPod(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Pod) *PaginatedPod {
	return &PaginatedPod{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedAppliance represents a paginated collection of Appliance
type PaginatedAppliance struct {
	*connection.PaginatedBase
	Items []Appliance
}

// NewPaginatedAppliance returns a pointer to an initialized PaginatedAppliance struct
func NewPaginatedAppliance(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Appliance) *PaginatedAppliance {
	return &PaginatedAppliance{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedApplianceParameter represents a paginated collection of ApplianceParameter
type PaginatedApplianceParameter struct {
	*connection.PaginatedBase
	Items []ApplianceParameter
}

// NewPaginatedApplianceParameter returns a pointer to an initialized PaginatedApplianceParameter struct
func NewPaginatedApplianceParameter(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []ApplianceParameter) *PaginatedApplianceParameter {
	return &PaginatedApplianceParameter{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedActiveDirectoryDomain represents a paginated collection of ActiveDirectoryDomain
type PaginatedActiveDirectoryDomain struct {
	*connection.PaginatedBase
	Items []ActiveDirectoryDomain
}

// NewPaginatedActiveDirectoryDomain returns a pointer to an initialized PaginatedActiveDirectoryDomain struct
func NewPaginatedActiveDirectoryDomain(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []ActiveDirectoryDomain) *PaginatedActiveDirectoryDomain {
	return &PaginatedActiveDirectoryDomain{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedVPC represents a paginated collection of VPC
type PaginatedVPC struct {
	*connection.PaginatedBase
	Items []VPC
}

// NewPaginatedVPC returns a pointer to an initialized PaginatedVPC struct
func NewPaginatedVPC(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []VPC) *PaginatedVPC {
	return &PaginatedVPC{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}
