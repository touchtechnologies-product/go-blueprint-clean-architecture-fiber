package company

import (
	"github.com/gofiber/fiber/v2"
	"net/http"

	"github.com/opentracing/opentracing-go"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/view"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"
)

// Delete godoc
// @Tags Companies
// @Summary Soft delete a company
// @Description Mark a company with a given company ID as deleted but keep its content
// @param company_id path string true "Company ID"
// @Param X-Authenticated-Userid header string true "User ID"
// @Produce  json
// @Success 200 {object} view.SuccessResp
// @Success 204 {object} view.SuccessResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 422 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /companies/{company_id} [delete]
func (ctrl *Controller) Delete(c *fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Context(),
		opentracing.GlobalTracer(),
		"handler.company.Delete",
	)
	defer span.Finish()

	input := &companyin.DeleteInput{
		ID: c.Params("id"),
	}

	err := ctrl.service.Delete(ctx, input)
	if err != nil {
		return view.MakeErrResp(c, err)
	}

	return view.MakeSuccessResp(c, http.StatusOK, nil)
}
