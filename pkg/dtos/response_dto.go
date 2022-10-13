package dtos

type Response[T any] struct {
	Data      *T       `json:"data,omitempty"`
	Meta_data MetaData `json:"meta_data"`
}

type MetaData struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}
