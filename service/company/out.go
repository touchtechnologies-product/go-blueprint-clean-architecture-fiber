package company

import (
	"content-service/domain"
)

type View struct {
	Name string `json:"name"`
}

func companyToView(company *domain.Company) (view *View) {
	return &View{
		Name: company.Name,
	}
}
