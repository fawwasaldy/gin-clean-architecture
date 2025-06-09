package pagination

type (
	Request struct {
		Search  string `form:"search"`
		Page    int    `form:"page"`
		PerPage int    `form:"per_page"`
	}
)

func (p *Request) GetOffset() int {
	return (p.Page - 1) * p.PerPage
}

func (p *Request) GetLimit() int {
	return p.PerPage
}

func (p *Request) GetPage() int {
	return p.Page
}

func (p *Request) Default() {
	if p.Page == 0 {
		p.Page = 1
	}

	if p.PerPage == 0 {
		p.PerPage = 10
	}
}
