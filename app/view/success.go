package view

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

const (
	okStatus       = "OK"
	xContentLength = "X-Content-Length"
	location       = "Content-Location"
)

type SuccessResp struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
} // @Name SuccessResponse

func MakeSuccessResp(c *fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(SuccessResp{
		Status: okStatus,
		Code:   status,
		Data:   data,
	})
}

func MakePaginatorResp(c *fiber.Ctx, total int, items interface{}) error {
	status := http.StatusOK
	if total < 1 {
		status = http.StatusNoContent
	}
	c.Set(xContentLength, strconv.Itoa(total))
	return MakeSuccessResp(c, status, items)
}

func MakeCreatedResp(c *fiber.Ctx, ID string) error {
	c.Set(location, ID)
	return MakeSuccessResp(c, http.StatusCreated, nil)
}
