package usecase

import (
	"github.com/dedicio/sisgares-integrations-service/internal/dto"
	"github.com/dedicio/sisgares-integrations-service/internal/entity"
	"github.com/dedicio/sisgares-integrations-service/pkg/crm"
)

type IntegrateOnOrderCreatedUseCase struct {
	Repository entity.IntegrationRepository
}

func NewIntegrateOnOrderCreatedUseCase(
	repository entity.IntegrationRepository,
) *IntegrateOnOrderCreatedUseCase {
	return &IntegrateOnOrderCreatedUseCase{
		Repository: repository,
	}
}

func (uc *IntegrateOnOrderCreatedUseCase) Execute(order *dto.OrderMessagingDTO) error {
	integrations, err := uc.Repository.FindAllByCompanyIDAndActive(order.CompanyID)
	if err != nil {
		return err
	}

	for _, integration := range integrations {
		if integration.Platform == "crm" {
			err = crm.NewCRMExample(crm.Config{
				Username: integration.PlatformUsername,
				Token:    integration.PlatformToken,
			}).IntegrateOrderCreated(*order)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
