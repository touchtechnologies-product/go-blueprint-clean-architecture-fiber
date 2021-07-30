package company

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opentracing/opentracing-go"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/view"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"
)

// Create godoc
// @Tags Companies
// @Summary Create a new company
// @Description A newly created company ID will be given in a Content-Location response header
// @Param input body companyin.CreateInput true "Input"
// @Param X-Authenticated-Userid header string true "User ID"
// @Accept json
// @Produce json
// @Success 201 {object} view.SuccessResp
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 422 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /companies [post]
func (ctrl *Controller) Create(c *fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Context(),
		opentracing.GlobalTracer(),
		"handler.company.Create",
	)
	defer span.Finish()

	input := &companyin.CreateInput{}
	if err := c.BodyParser(input); err != nil {
		return view.MakeErrResp(c, err)
	}

	ID, err := ctrl.service.Create(ctx, input)
	if err != nil {
		return view.MakeErrResp(c, err)
	}

	return view.MakeCreatedResp(c, ID)
}
