package test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/staff"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/config"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/mocks"
)

type PackageTestSuite struct {
	suite.Suite
	router  *fiber.App
	ctx     *fiber.Ctx
	conf    *config.Config
	ctrl    *staff.Controller
	service *mocks.Service
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.conf = config.Get()
	suite.service = &mocks.Service{}
	suite.ctrl = staff.New(suite.service)
	suite.router = fiber.New()

	suite.router.Get("/staffs", suite.ctrl.List)
	suite.router.Post("/staffs", suite.ctrl.Create)
	suite.router.Get("/staffs/:id", suite.ctrl.Read)
	suite.router.Put("/staffs/:id", suite.ctrl.Update)
	suite.router.Delete("/staffs/:id", suite.ctrl.Delete)
}
