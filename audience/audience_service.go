package audience

import (
	"github.com/jszwec/csvutil"
	"gorm.io/gorm"
	"io/ioutil"
	"maheswari-boilerplate/domain"
	"os"
)

type User struct {
	FullName string  `csv:"fullName"`
	Email    string  `csv:"email"`
	Detail   *string `csv:"detail,omitempty"`
}

type audienceService struct {
	DB *gorm.DB
}

type Service interface {
	ExtractSave(name string, path string) error
}

func New(db *gorm.DB) Service {
	return &audienceService{
		DB: db,
	}
}

func (a audienceService) ExtractSave(name string, path string) error {
	file, _ := ioutil.ReadFile(path)
	var users []User
	csvutil.Unmarshal(file, &users)
	audience := domain.Audience{
		Name: name,
	}
	audience.Name = name
	a.DB.Save(&audience)
	go func() {
		for _, u := range users {
			a.DB.Save(&domain.User{
				Name:       u.FullName,
				Email:      u.Email,
				AudienceID: audience.ID,
			})
		}
		os.Remove(path)
	}()

	return nil
}
