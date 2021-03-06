package company

import (
	"github.com/gofiber/fiber/v2"
	"net/http"

	"github.com/opentracing/opentracing-go"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/view"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"
)

// Read godoc
// @Tags Companies
// @Summary Get a company by company ID
// @Description Response a company data with a given company ID
// @param company_id path string true "Company ID"
// @Produce  json
// @Success 200 {object} view.SuccessResp{data=out.CompanyView}
// @Success 204 {object} view.SuccessResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /companies/{company_id} [get]
func (ctrl *Controller) Read(c *fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Context(),
		opentracing.GlobalTracer(),
		"handler.company.Read",
	)
	defer span.Finish()

	input := &companyin.ReadInput{CompanyID: c.Params("id")}

	company, err := ctrl.service.Read(ctx, input)
	if err != nil {
		return view.MakeErrResp(c, err)

	}

	return view.MakeSuccessResp(c, http.StatusOK, company)
}
