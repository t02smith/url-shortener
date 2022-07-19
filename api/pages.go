package api

import (
	"html/template"
	"log"
	"net/http"

	"t02smith.com/url-shortener/db"
)

func setupTemplates() *template.Template {
	var templates *template.Template
	templates = template.Must(templates.ParseGlob("./static/html/*.html"))
	return templates
}

var t *template.Template = setupTemplates()

type IndexForm struct {
	URL string
}

func Index(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		log.Println("Directing to Index page")
		t.ExecuteTemplate(w, "index.html", nil)
		return
	case "POST":
		r.ParseForm()
		var Form IndexForm = IndexForm{db.RequestURL(db.Database, r.Form["url"][0], r.Form["request"][0]).New_link}
		t.ExecuteTemplate(w, "index.html", Form)
		return
	}

}
