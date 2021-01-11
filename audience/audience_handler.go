package audience

import (
	"encoding/json"
	"fmt"
	"github.com/adelowo/filer"
	"github.com/adelowo/filer/validator"
	validator2 "github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/thanhpk/randstr"
	"gorm.io/gorm"
	"io"
	"maheswari-boilerplate/domain"
	"maheswari-boilerplate/lib"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

type audienceHandler struct {
	DB *gorm.DB
}

func NewAudience(DB *gorm.DB) audienceHandler {
	return audienceHandler{
		DB: DB,
	}
}
func (a audienceHandler) Store(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (a audienceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (a audienceHandler) Update(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (a audienceHandler) List(w http.ResponseWriter, r *http.Request) {
	var audiences []domain.Audience

	page, _ := strconv.ParseInt(r.FormValue("page"), 0, 32)
	logrus.Println(page)
	paging := lib.Paginate(&audiences, &lib.Param{
		DB:      a.DB,
		Page:    int(page),
		PerPage: 20,
		OrderBy: 0,
	})

	lib.BaseResponse(paging, w, r)
}

func (a audienceHandler) Detail(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (a audienceHandler) CreateAudience(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var audience domain.Audience
	if err := decoder.Decode(&audience); err != nil {
		logrus.Println("decoder", err)
		lib.BaseResponse(err, w, r)
		return
	}

	if status, errors := lib.ValidateStruct(audience); !status {
		lib.BaseResponse(errors, w, r)
		return
	}

	logrus.Println(audience)
	lib.BaseResponse("Success", w, r)
	return
}

func (a audienceHandler) UploadAudience(w http.ResponseWriter, r *http.Request) {
	max, _ := filer.LengthInBytes("20MB")
	min, _ := filer.LengthInBytes("2KB")

	val := validator.NewSizeValidator(max, min)
	//mime := validator.NewMimeTypeValidator([]string{ "text/csv"})
	extenision := validator.NewExtensionValidator([]string{"csv"})

	file, handler, err := r.FormFile("file")
	if err != nil {
		lib.BaseResponse(err, w, r)
		return
	}
	nameAudience := r.PostFormValue("name")
	ext := path.Ext(handler.Filename)[1:]

	name := randstr.String(40) + "." + ext
	year, month, day := time.Now().Date()
	randomName := fmt.Sprintf("%v_%v_%v_%v", year, int(month), day, name)
	tempVar := fmt.Sprintf("temp/%v", randomName)
	dst, err := os.Create(tempVar)
	defer dst.Close()
	if err != nil {
		logrus.Println(err)
		lib.BaseResponse(err, w, r)
		return
	}

	if _, err := io.Copy(dst, file); err != nil {
		logrus.Println(err)
		lib.BaseResponse(err, w, r)
		return
	}

	file2, _ := os.Open(tempVar)
	defer file2.Close()

	validate := validator.NewChainedValidator(val, extenision)
	isValid, errFile := validate.Validate(file2)
	if !isValid && errFile != nil {
		go func() {
			os.Remove(file2.Name())
		}()
		lib.BaseResponse([]lib.ValidateError{{
			Field: "file",
			Error: errFile.Error(),
		}}, w, r)
		return
	}
	validate1 := validator2.New()
	if errs := validate1.Var(nameAudience, "required"); errs != nil {
		os.Remove(file2.Name())
		lib.BaseResponse(errs, w, r)
		return
	}

	service := New(a.DB)
	service.ExtractSave(nameAudience, file2.Name())
	lib.BaseResponse("Success", w, r)
	return

}
