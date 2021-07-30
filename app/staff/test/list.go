package test

import (
	"net/http"
	"net/http/httptest"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/out"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestList() {
	req, _, err := makeListReq()
	suite.NoError(err)

	opt := &domain.PageOption{
		Page:    1,
		PerPage: 10,
		Sorts:   []string{"createdAt:desc"},
	}

	suite.service.On("List", mock.Anything, opt).Return(0, []*out.StaffView{}, nil)
	res, _ := suite.router.Test(req, -1)

	suite.Equal(http.StatusNoContent, res.StatusCode)
}

func makeListReq() (req *http.Request, w *httptest.ResponseRecorder, err error) {
	req, err = http.NewRequest("GET", "/staffs?page=1&per_page=10", nil)
	if err != nil {
		return nil, nil, err
	}
	return req, httptest.NewRecorder(), nil
}
