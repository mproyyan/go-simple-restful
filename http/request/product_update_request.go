package request

type ProductUpdateRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
