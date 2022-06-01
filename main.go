package main

import (
	"fmt"
	"net/http"
)

type router struct {
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	fmt.Fprint(w, "Wecome to home site")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	fmt.Fprint(w, "Any question please contact me "+
		"<a href=\"mailto:luan.nguyen@netpower.no\">luan.nguyen@netpower.no</a>")
}

func (router router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact/":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	var router router
	fmt.Println("The server listen on: 3000")
	http.ListenAndServe(":3000", router)
}
