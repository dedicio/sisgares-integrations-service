package usecase

import (
	"github.com/dedicio/sisgares-integrations-service/internal/entity"
)

type DeleteIntegrationUseCase struct {
	Repository entity.IntegrationRepository
}

func NewDeleteIntegrationUseCase(IntegrationRepository entity.IntegrationRepository) *DeleteIntegrationUseCase {
	return &DeleteIntegrationUseCase{
		Repository: IntegrationRepository,
	}
}

func (uc *DeleteIntegrationUseCase) Execute(id string) error {
	err := uc.Repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
