package usecase

import (
	"github.com/dedicio/sisgares-integrations-service/internal/dto"
	"github.com/dedicio/sisgares-integrations-service/internal/entity"
)

type UpdateIntegrationUseCase struct {
	Repository entity.IntegrationRepository
}

func NewUpdateIntegrationUseCase(IntegrationRepository entity.IntegrationRepository) *UpdateIntegrationUseCase {
	return &UpdateIntegrationUseCase{
		Repository: IntegrationRepository,
	}
}

func (uc *UpdateIntegrationUseCase) Execute(integration *dto.IntegrationDTO) (*dto.IntegrationResponseDTO, error) {
	newIntegration := entity.NewIntegration(
		integration.Name,
		integration.CompanyID,
		integration.PlatformID,
		integration.PlatformUsername,
		integration.PlatformToken,
		integration.Active,
	)

	err := uc.Repository.Update(newIntegration)
	if err != nil {
		return nil, err
	}

	output := &dto.IntegrationResponseDTO{
		ID:               newIntegration.ID,
		Name:             newIntegration.Name,
		PlatformID:       newIntegration.PlatformID,
		PlatformUsername: newIntegration.PlatformUsername,
		Active:           newIntegration.Active,
	}

	return output, nil
}
