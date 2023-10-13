package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dedicio/sisgares-integrations-service/internal/dto"
	"github.com/dedicio/sisgares-integrations-service/internal/entity"
	usecase "github.com/dedicio/sisgares-integrations-service/internal/usecase"
	httpResponsePkg "github.com/dedicio/sisgares-integrations-service/pkg/response"
	"github.com/go-chi/render"
)

type IntegrationController struct {
	Repository entity.IntegrationRepository
	// Publisher  entity.IntegrationPublisher
}

func NewIntegrationController(
	integrationRepository entity.IntegrationRepository,
	// integrationPublisher entity.IntegrationPublisher,
) *IntegrationController {
	return &IntegrationController{
		Repository: integrationRepository,
		// Publisher:  integrationPublisher,
	}
}

func (ctrl *IntegrationController) Create(w http.ResponseWriter, r *http.Request) {
	companyID := r.Header.Get("X-Company-ID")
	payload := json.NewDecoder(r.Body)
	var integration *dto.IntegrationDTO
	err := payload.Decode(&integration)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	integration.CompanyID = companyID
	output, err := usecase.NewCreateIntegrationUseCase(ctrl.Repository).Execute(integration)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	// go usecase.NewPublishCreateIntegrationUseCase(ctrl.Publisher).Execute(integrationSaved)

	render.Render(w, r, httpResponsePkg.NewIntegrationResponse(output))
}

func (ctrl *IntegrationController) Update(w http.ResponseWriter, r *http.Request) {
	integrationID := r.Header.Get("X-Integration-ID")
	payload := json.NewDecoder(r.Body)
	var integration *dto.IntegrationDTO
	err := payload.Decode(&integration)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	integration.ID = integrationID
	output, err := usecase.NewUpdateIntegrationUseCase(ctrl.Repository).Execute(integration)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, httpResponsePkg.NewIntegrationResponse(output))
}

func (ctrl *IntegrationController) Delete(w http.ResponseWriter, r *http.Request) {
	integrationID := r.Header.Get("X-Integration-ID")
	err := usecase.NewDeleteIntegrationUseCase(ctrl.Repository).Execute(integrationID)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, httpResponsePkg.NewIntegrationResponse(nil))
}
