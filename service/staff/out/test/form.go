package test

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/out"
)

func (suite *PackageTestSuite) TestStaffToView() {
	given := domain.MakeTestStaff()

	actual := out.StaffToView(given)

	suite.Equal(given.ID, actual.ID)
	suite.Equal(given.Name, actual.Name)
}
