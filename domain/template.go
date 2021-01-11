package domain

type Template struct {
	BaseDomain
	Name string `json:"name" validate:"required"`
	Html string `json:"html" validate:"required"`
}
