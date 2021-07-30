package domain

type Paginator struct {
	Total   int
	CurPage int
	PerPage int
	HasMore bool
}

type PageOption struct {
	Page    int      `form:"page" query:"page" validate:"min=0"`
	PerPage int      `form:"per_page" query:"per_page"  validate:"min=0"`
	Filters []string `form:"filters" query:"filters"`
	Sorts   []string `form:"sorts" query:"sorts"`
}

type SetOpParam struct {
	Filters      []string
	SetFieldName string
	Item         interface{}
}
