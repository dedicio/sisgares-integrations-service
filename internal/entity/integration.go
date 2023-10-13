package entity

import "github.com/dedicio/sisgares-integrations-service/pkg/utils"

type Integration struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	CompanyID        string `json:"company_id"`
	PlatformID       string `json:"platform_id"`
	PlatformUsername string `json:"platform_username"`
	PlatformToken    string `json:"platform_token"`
	Active           bool   `json:"active"`
}

type IntegrationRepository interface {
	// FindByCompanyIDAndPlatformID(companyID string, platformID string) (*Integration, error)
	// FindAllByCompanyID(companyID string) ([]*Integration, error)
	// FindAllByCompanyIDAndActive(companyID string, active bool) ([]*Integration, error)
	// FindAllByCompanyIDAndPlatformID(companyID string, platformID string) ([]*Integration, error)
	Update(integration *Integration) error
	Create(integration *Integration) error
	Delete(id string) error
}

func NewIntegration(
	name string,
	companyID string,
	platformID string,
	platformUsername string,
	platformToken string,
	active bool,
) *Integration {
	id := utils.NewUUID()
	return &Integration{
		ID:               id,
		Name:             name,
		CompanyID:        companyID,
		PlatformID:       platformID,
		PlatformUsername: platformUsername,
		PlatformToken:    platformToken,
		Active:           active,
	}
}
