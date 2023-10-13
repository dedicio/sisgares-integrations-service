package dto

type IntegrationDTO struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	CompanyID        string `json:"company_id"`
	PlatformID       string `json:"platform_id"`
	PlatformUsername string `json:"platform_username"`
	PlatformToken    string `json:"platform_token"`
	Active           bool   `json:"active"`
}

type IntegrationResponseDTO struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	PlatformID       string `json:"platform_id"`
	PlatformUsername string `json:"platform_username"`
	Active           bool   `json:"active"`
}
