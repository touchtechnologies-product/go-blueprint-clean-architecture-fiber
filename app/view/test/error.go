package test

import (
	"errors"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/view"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

type ValidatorTestStruct struct {
	Title string `validate:"required" json:"title"`
	Body  string `validate:"required"`
}

func (suite *PackageTestSuite) TestMakeErrResp() {
	err := util.ConvertInputToDomainErr(errors.New("test"))
	view.MakeErrResp(suite.ctx, err)
}

func (suite *PackageTestSuite) TestMakeConvertInputToDomainErrResp() {
	err := util.ConvertInputToDomainErr(errors.New("test"))
	view.MakeErrResp(suite.ctx, err)
}

func (suite *PackageTestSuite) TestMakeRepoListErrResp() {
	err := util.RepoListErr(errors.New("test"))
	view.MakeErrResp(suite.ctx, err)
	//suite.Equal(http.StatusNoContent, suite.ctx.Writer.Status())
}

func (suite *PackageTestSuite) TestMakeRepoReadErrResp() {
	err := util.RepoReadErr(errors.New("test"))
	view.MakeErrResp(suite.ctx, err)
}

func (suite *PackageTestSuite) TestMakeRepoCreateErrResp() {
	err := util.RepoCreateErr(errors.New("test"))
	view.MakeErrResp(suite.ctx, err)
}

func (suite *PackageTestSuite) TestMakeUnknownErrResp() {
	err := errors.New("some error message")
	view.MakeErrResp(suite.ctx, err)
}

func (suite *PackageTestSuite) TestMakeValidationCreateErrResp() {
	st := &ValidatorTestStruct{
		Title: "",
		Body:  "",
	}
	vErrs := suite.validate.Struct(st)

	err := util.ValidationCreateErr(vErrs)
	view.MakeErrResp(suite.ctx, err)
}

func (suite *PackageTestSuite) TestMakeValidationUpdateErrResp() {
	st := &ValidatorTestStruct{
		Title: "",
		Body:  "",
	}
	vErrs := suite.validate.Struct(st)

	err := util.ValidationUpdateErr(vErrs)
	view.MakeErrResp(suite.ctx, err)
}
