package base

type Page struct {
	Current int         `json:"current" form:"current"`
	Size    int         `json:"size" form:"size"`
	Total   int64       `json:"total" form:"total"`
	Records interface{} `json:"records" form:"records"`
}
