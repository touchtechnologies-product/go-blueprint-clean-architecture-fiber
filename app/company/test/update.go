package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestUpdate() {
	input := companyin.MakeTestUpdateInput()
	req, _, err := suite.makeUpdateReq(input)
	suite.NoError(err)

	suite.service.On("Update", mock.Anything, input).Return(nil)
	res, _ := suite.router.Test(req, -1)

	suite.Equal(http.StatusOK, res.StatusCode)
}

func (suite *PackageTestSuite) makeUpdateReq(input *companyin.UpdateInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return nil, nil, err
	}

	req, err = http.NewRequest("PUT", fmt.Sprintf("/companies/%s", input.ID), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	return req, httptest.NewRecorder(), nil
}
