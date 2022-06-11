package types

type Error struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Status int    `json:"status"`
}

type Errors []Error

type ErrorsResponse struct {
	Errors Errors `json:"errors"`
}
