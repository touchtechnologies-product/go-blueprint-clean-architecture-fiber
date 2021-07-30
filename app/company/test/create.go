package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestCreate() {
	input := companyin.MakeTestCreateInput()
	req, resp, err := suite.makeCreateReq(input)
	suite.NoError(err)

	newID := ""
	suite.service.On("Create", mock.Anything, input).Return(newID, nil)
	res, _ := suite.router.Test(req, -1)

	suite.Equal(http.StatusCreated, res.StatusCode)
	suite.Equal(newID, resp.Header().Get("Content-Location"))
}

func (suite *PackageTestSuite) makeCreateReq(input *companyin.CreateInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return nil, nil, err
	}

	req, err = http.NewRequest("POST", "/companies", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) TestCreateInvalidJSON() {
	req, _, err := suite.makeCreateReqInvalidJSON()
	suite.NoError(err)

	res, _ := suite.router.Test(req, -1)

	suite.Equal(http.StatusInternalServerError, res.StatusCode)
}

func (suite *PackageTestSuite) makeCreateReqInvalidJSON() (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes := []byte("{{{}}}")

	req, err = http.NewRequest("POST", "/companies", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) TestCreateServiceError() {
	input := companyin.MakeTestCreateInputErr()
	req, _, err := suite.makeCreateReq(input)
	suite.NoError(err)

	givenError := errors.New("some error message")
	suite.service.On("Create", mock.Anything, input).Once().Return("", givenError)
	res, _ := suite.router.Test(req, -1)
	suite.Equal(http.StatusInternalServerError, res.StatusCode)
}
