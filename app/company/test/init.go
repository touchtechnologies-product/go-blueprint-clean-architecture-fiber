package test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/company"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/config"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/mocks"

	"github.com/stretchr/testify/suite"
)

type PackageTestSuite struct {
	suite.Suite
	router  *fiber.App
	ctx     *fiber.Ctx
	conf    *config.Config
	ctrl    *company.Controller
	service *mocks.Service
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.conf = config.Get()
	suite.service = &mocks.Service{}
	suite.ctrl = company.New(suite.service)
	suite.router = fiber.New()

	suite.router.Get("/companies", suite.ctrl.List)
	suite.router.Post("/companies", suite.ctrl.Create)
	suite.router.Get("/companies/:id", suite.ctrl.Read)
	suite.router.Put("/companies/:id", suite.ctrl.Update)
	suite.router.Delete("/companies/:id", suite.ctrl.Delete)
}
