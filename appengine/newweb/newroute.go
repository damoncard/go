package newweb

import (
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/", land)
}

func land(rw http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("_resource/index.html"))

	if err := t.Execute(rw, "What"); err != nil {
		rw.WriteHeader(500)
		return
	}
}
