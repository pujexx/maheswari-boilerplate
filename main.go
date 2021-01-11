package main

import (
	"github.com/gorilla/mux"
	muxlogrus "github.com/pytimer/mux-logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"maheswari-boilerplate/audience"
	"maheswari-boilerplate/domain"
	"maheswari-boilerplate/lib"
	"maheswari-boilerplate/template"
	"net/http"
	"time"
)

func main() {
	db, err := gorm.Open(sqlite.Open("db"), &gorm.Config{})
	if err != nil {
		panic("error connection database")
	}
	//audience
	au := audience.NewAudience(db)
	t := template.NewHandler(db)

	db.AutoMigrate(
		&domain.Template{},
		&domain.User{},
		&domain.Audience{},
		&domain.Template{},
	)

	r := mux.NewRouter()

	v1 := r.PathPrefix("/v1").Subrouter()
	r.Use(muxlogrus.NewLogger().Middleware)
	v1.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		lib.BaseResponse(domain.Response{Data: domain.Template{}}, w, r)
	}).Methods("GET")

	v1.HandleFunc("/audience", au.CreateAudience).Methods("POST")
	v1.HandleFunc("/audience", au.List).Methods("GET")
	v1.HandleFunc("/audience-upload", au.UploadAudience).Methods("POST")

	v1.HandleFunc("/template", t.Store).Methods("POST")
	v1.HandleFunc("/template", t.List).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
