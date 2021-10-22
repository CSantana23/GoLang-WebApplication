package handlers

import (
	"myApp/pkg/render"
	"net/http"
)

// Home functions that handle web request need to handle two parameters a responseWriter and request
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
