package appstore

type getResponse[T any] struct {
	Data  T     `json:"data"`
	Links links `json:"links"`
}

type listResponse[T any] struct {
	Data  []T   `json:"data"`
	Links links `json:"links"`
	Meta  meta  `json:"meta"`
}

type createRequest[T any] struct {
	Data struct {
		Type string `json:"type"`
		Attr T      `json:"attributes"`
		Rel  any    `json:"relationships"`
	} `json:"data"`
}

type relation struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func newCreateRequest[T any](attr T, typ string, rel relation) createRequest[T] {
	r := createRequest[T]{}
	r.Data.Type = typ
	r.Data.Attr = attr
	r.Data.Rel = map[string]any{
		rel.Type[:len(rel.Type)-1]: map[string]any{
			"data": rel,
		},
	}

	return r
}

type createResponse[T any] struct {
	Data  T     `json:"data"`
	Links links `json:"links"`
}

type meta struct {
	Paging paging `json:"paging"`
}

type paging struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
}

type links struct {
	Self  string `json:"self"`
	Next  string `json:"next"`
	First string `json:"first"`
}
