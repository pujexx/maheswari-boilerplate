package domain

type User struct {
	BaseDomain
	Name       string `json:"name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Detail     string `json:"Detail" validate:"omitempty"`
	AudienceID string
	//Audience *Audience `gorm:"foreignKey:AudienceID;references:Id"`

}
