package template

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"maheswari-boilerplate/domain"
	"maheswari-boilerplate/lib"
	"net/http"
	"strconv"
)

type templateHandler struct {
	DB *gorm.DB
}

func NewHandler(DB *gorm.DB) templateHandler {
	return templateHandler{
		DB,
	}
}

func (t templateHandler) Store(w http.ResponseWriter, r *http.Request) {
	var template domain.Template
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&template); err != nil {
		logrus.Println("decoder", err)
		lib.BaseResponse(err, w, r)
		return
	}

	if status, err := lib.ValidateStruct(&template); !status {
		lib.BaseResponse(err, w, r)
		return
	}

	t.DB.Save(&template)
	lib.BaseResponse("success", w, r)
	return
}

func (t templateHandler) Delete(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (t templateHandler) Update(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (t templateHandler) List(w http.ResponseWriter, r *http.Request) {
	var templates []domain.Template

	page, _ := strconv.ParseInt(r.FormValue("page"), 0, 32)
	logrus.Println(page)
	paging := lib.Paginate(&templates, &lib.Param{
		DB:      t.DB,
		Page:    int(page),
		PerPage: 20,
		OrderBy: 0,
	})

	lib.BaseResponse(paging, w, r)
}

func (t templateHandler) Detail(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
