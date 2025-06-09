package pagination

type (
	Response struct {
		Page    int   `json:"page"`
		PerPage int   `json:"per_page"`
		MaxPage int64 `json:"max_page"`
		Count   int64 `json:"count"`
	}

	ResponseWithData struct {
		Data []any `json:"data"`
		Response
	}
)
