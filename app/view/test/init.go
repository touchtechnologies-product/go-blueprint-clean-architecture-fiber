package test

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
	"github.com/valyala/fasthttp"
)

type PackageTestSuite struct {
	suite.Suite
	router   *fiber.App
	ctx      *fiber.Ctx
	validate *validator.Validate
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.router = fiber.New()
	suite.ctx = suite.router.AcquireCtx(&fasthttp.RequestCtx{})
	suite.validate = validator.New()
}
