package test

import (
	"net/http"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/view"
)

type ItemStruct struct {
	Title string
	Body  string
}

func (suite *PackageTestSuite) TestMakeSuccessResp() {
	st := struct {
		Title string `validate:"required"`
		Body  string `validate:"required"`
	}{
		Title: "",
		Body:  "",
	}
	view.MakeSuccessResp(suite.ctx, http.StatusOK, st)
}

func (suite *PackageTestSuite) TestMakePaginatorResp() {
	items := []ItemStruct{
		{
			Title: "Test",
			Body:  "Test",
		},
	}
	view.MakePaginatorResp(suite.ctx, 1, items)
}

func (suite *PackageTestSuite) TestMakePaginatorRespNoContent() {
	var items []ItemStruct
	view.MakePaginatorResp(suite.ctx, 0, items)
}

func (suite *PackageTestSuite) TestMakeCreatedResp() {
	newID := "new ID"
	view.MakeCreatedResp(suite.ctx, newID)
}
