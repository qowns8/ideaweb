package router

import (
	"net/http"
	"github.com/qowns8/ideaweb/middlewares"
)

type Router struct {
	handler * http.Handler
}

func (ro * Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	switch path {
	case "/" :
		RenderMainView(w, r)
	case "/auth" :
		RenderAuthView(w, r)
	case "/auth/callback" :
		Authenticate(w, r)
	case "/test" :
		test := middlewares.BasicLogic(
			func(writer http.ResponseWriter, request *http.Request) {
				testPrint(writer, request)
			})
		test(w, r)
	case "/user/info" :
		getUserInfo(w, r)
	}
/*
	usinfo := middlewares.BasicLogic(
		func(writer http.ResponseWriter, request *http.Request) {
			getUserInfo(writer, request)
		})
*/
}
