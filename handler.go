package todo

import (
	"log"
	"net/http"
	"time"
)

func todoIndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte("This is todo index handler"))
	} else if r.Method == http.MethodPost {
		todoSaveHandler(w, r)
	} else {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
	}
}

func todoSaveHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body")
	// layout, string to be parsed
	t, err := time.Parse(time.RFC3339, r.FormValue("time"))
	if err != nil {
		log.Print(err)
		w.Write([]byte("time cant be parsed. please use RFC3339 format."))
		return
	}

	p := &Todo{Time: t, Body: body}
	jsoned, err := p.save()
	if err != nil {
		w.WriteHeader(400)
	} else {
		w.WriteHeader(200)
		w.Write(jsoned)
	}
}
