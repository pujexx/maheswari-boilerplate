package domain

type Audience struct {
	BaseDomain
	Name  string `json:"name" validate:"required"`
	Users []User `json:"-"`
}
