package app

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/company"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/staff"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/docs"
	companyService "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company"
	staffService "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff"
)

type App struct {
	staff   *staff.Controller
	company *company.Controller
}

func New(staffService staffService.Service, companyService companyService.Service) *App {
	return &App{
		staff:   staff.New(staffService),
		company: company.New(companyService),
	}
}

func (app *App) RegisterRoute(router *fiber.App) *App {
	docs.SwaggerInfo.Title = "Touch Tech API"
	docs.SwaggerInfo.Description = "API Spec Demo."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	apiRoutes := router.Group(docs.SwaggerInfo.BasePath)
	apiRoutes.Get("/companies", app.company.List)
	apiRoutes.Post("/companies", app.company.Create)
	apiRoutes.Get("/companies/:id", app.company.Read)
	apiRoutes.Put("/companies/:id", app.company.Update)
	apiRoutes.Delete("/companies/:id", app.company.Delete)

	apiRoutes.Get("/staffs", app.staff.Update)
	apiRoutes.Post("/staffs", app.staff.Create)
	apiRoutes.Get("/staffs/:id", app.staff.Read)
	apiRoutes.Put("/staffs/:id", app.staff.Update)
	apiRoutes.Delete("/staffs/:id", app.staff.Delete)
	router.Get("/swagger/*", swagger.Handler)

	return app
}
