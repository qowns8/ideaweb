package router

import (
	"net/http"
	"github.com/qowns8/ideaweb/models"
	"fmt"
	"encoding/json"
)
func token_ok (w http.ResponseWriter, r *http.Request) {
	access_key := r.FormValue("access_token")
	user := models.User{
		Access_key : access_key,
	}
	result := user.Get()
	if result {
		fmt.Fprintf(w, "success" + user.Email)
	} else {
		fmt.Fprintf(w, "fail")
	}
}

func getUserInfo (w http.ResponseWriter, r *http.Request) {
	access_key := r.Header.Get("access_token")
	user := models.User{
		Access_key: access_key,
	}
	result := user.Get()
	if result {
		r, _ := json.Marshal(user)
		w.Write(r)
	} else {
		fmt.Fprintf(w, "fail")
	}
}
