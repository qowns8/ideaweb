package middlewares

import (
	"net/http"
	"log"
	"golang.org/x/oauth2"
	"io/ioutil"
	"encoding/json"
	"github.com/qowns8/ideaweb/utils"
	"github.com/qowns8/ideaweb/models"
	"fmt"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func AuthCheck() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			token := &oauth2.Token{AccessToken:r.Header.Get("access-key")}
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
			authUser.Name, authUser.Email,"\n access token :", token.AccessToken,
			)

			var us models.User
			us.Access_key = token.AccessToken
			check := us.Get()
			if check {
				f(w, r)
			} else {
				fmt.Fprintf(w, "token access fail")
			}
		}
	}
}

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			defer func() {
			}()
			log.Println("visit : ", r.URL.Path)
			f(w, r)
		}
	}
}

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func BasicLogic(f http.HandlerFunc) http.HandlerFunc {
	return Chain(f, AuthCheck(), Logging())
}