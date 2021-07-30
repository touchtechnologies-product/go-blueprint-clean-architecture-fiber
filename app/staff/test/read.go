package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/out"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestRead() {
	input := staffin.MakeTestReadInput()
	req, _, err := makeReadReq(input)
	suite.NoError(err)

	suite.service.On("Read", mock.Anything, input).Return(&out.StaffView{}, nil)
	res, _ := suite.router.Test(req, -1)

	suite.Equal(http.StatusOK, res.StatusCode)
}

func makeReadReq(input *staffin.ReadInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	req, err = http.NewRequest("GET", fmt.Sprintf("/staffs/%s", input.StaffID), nil)
	if err != nil {
		return nil, nil, err
	}
	return req, httptest.NewRecorder(), nil
}
