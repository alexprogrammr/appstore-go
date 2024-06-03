package appstore

type Paging struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
}

type PagingLinks struct {
	Self  string `json:"self"`
	Next  string `json:"next"`
	First string `json:"first"`
}
