package utils

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"github.com/gorilla/sessions"
)

var (
	key = []byte("api-key")
	Store = sessions.NewCookieStore(key)
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

const (
	CallBackURL = "http://Testserver-env.miwxirpfwi.ap-northeast-2.elasticbeanstalk.com:5000/auth/callback"

	UserInfoAPIEndpoint = "https://www.googleapis.com/oauth2/v3/userinfo"
	ScopeEmail          = "https://www.googleapis.com/auth/userinfo.email"
	ScopeProfile        = "https://www.googleapis.com/auth/userinfo.profile"
)

var OAuthConf *oauth2.Config

func init() {
	OAuthConf = &oauth2.Config{
		ClientID:     "712337399460-oqb7nfc6d4rsbfres8mr1jh9m4l5iv5b.apps.googleusercontent.com",
		ClientSecret: "PLcmyqjSkZ5xNGV_CvLD_uoM",
		RedirectURL:  CallBackURL,
		Scopes:       []string{ScopeEmail, ScopeProfile},
		Endpoint:     google.Endpoint,
	}
}

func GetLoginURL(state string) string {
	return OAuthConf.AuthCodeURL(state)
}

func RandToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}