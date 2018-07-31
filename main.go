package main

import (
	"github.com/qowns8/ideaweb/router"
	"log"
	"net/http"
)

func main() {

	var route = router.Router{}

	println("serve start")
	log.Fatal(http.ListenAndServe(":5000", &route ))
}