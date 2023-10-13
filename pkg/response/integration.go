package response

import (
	"net/http"

	"github.com/dedicio/sisgares-integrations-service/internal/dto"
)

type IntegrationResponse struct {
	*dto.IntegrationResponseDTO
}

func (pr *IntegrationResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewIntegrationResponse(product *dto.IntegrationResponseDTO) *IntegrationResponse {
	return &IntegrationResponse{product}
}
