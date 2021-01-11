package user

import (
	"maheswari-boilerplate/lib"
	"net/http"
)

func CreateAudience(w http.ResponseWriter, r *http.Request) {
	lib.BaseResponse("data", w, r)
}
