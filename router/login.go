package router

import (
	"net/http"
	"log"
	"golang.org/x/oauth2"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/sessions"
	"html/template"
	"github.com/qowns8/ideaweb/utils"
	"github.com/qowns8/ideaweb/models"
)

func testPrint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world \n access_key" + r.Header.Get("access_key")))
}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl, _ := template.ParseFiles(name)
	tmpl.Execute(w, data)
}

func RenderMainView(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "main.html", nil)
}

func RenderAuthView(w http.ResponseWriter, r *http.Request) {
	session, _ := utils.Store.Get(r, "session")
	session.Options = &sessions.Options{
		Path:   "/auth",
		MaxAge: 300,
	}
	state := utils.RandToken()
	session.Values["state"] = state
	session.Save(r, w)
	RenderTemplate(w, "auth.html", utils.GetLoginURL(state))
}

func Authenticate(w http.ResponseWriter, r *http.Request) {

	session, _ := utils.Store.Get(r, "session")
	state := session.Values["state"]

	delete(session.Values, "state")
	session.Save(r, w)

	if state != r.FormValue("state") {
		http.Error(w, "Invalid session state", http.StatusUnauthorized)
		return
	}

	log.Println("state : ", state)

	token, err := utils.OAuthConf.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client := utils.OAuthConf.Client(oauth2.NoContext, token)
	userInfoResp, err := client.Get(utils.UserInfoAPIEndpoint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer userInfoResp.Body.Close()
	userInfo, err := ioutil.ReadAll(userInfoResp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var authUser utils.User
	json.Unmarshal(userInfo, &authUser)

	log.Println(
		authUser.Name, authUser.Email,"\n acess token :", token.AccessToken,
	)

	models.GoogleAccessUpdate(token.AccessToken, authUser.Email)

	http.Redirect(w, r, "/", http.StatusFound)
}
