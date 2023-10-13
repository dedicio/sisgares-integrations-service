package routes

import (
	"github.com/dedicio/sisgares-integrations-service/internal/controllers"
	"github.com/dedicio/sisgares-integrations-service/internal/entity"
	"github.com/go-chi/chi/v5"
)

type IntegrationRoutes struct {
	Controller controllers.IntegrationController
}

func NewIntegrationRoutes(repository entity.IntegrationRepository) *IntegrationRoutes {
	return &IntegrationRoutes{
		Controller: *controllers.NewIntegrationController(repository),
	}
}

func (ir IntegrationRoutes) Routes() chi.Router {
	router := chi.NewRouter()

	router.Route("/", func(router chi.Router) {
		router.Post("/", ir.Controller.Create)

		router.Route("/{id}", func(router chi.Router) {
			// router.Get("/", ir.Controller.FindById)
			router.Put("/status/{status}", ir.Controller.Update)
			router.Delete("/items/{itemId}", ir.Controller.Delete)
		})
	})

	return router
}
