package staff

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
)

type View struct {
	Name      string `json:"name"`
	CompanyID string `json:"companyId" validator:"required"`
	Tel       string `json:"tel"`
	CreatedAt int64  `bson:"createdAt"`
	UpdatedAt int64  `bson:"updatedAt"`
}

func staffToView(staff *domain.Staff) (view *View) {
	return &View{
		Name: staff.Name,
	}
}