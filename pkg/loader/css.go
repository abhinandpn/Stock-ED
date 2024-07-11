package loader

import (
    "net/http"
)

func AboutCssLoader() {
    http.HandleFunc("/stylesheets/about.css", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/css/about.css")
    })

    http.HandleFunc("/stylesheets/contact.css", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/css/contact.css")
    })

    http.HandleFunc("/stylesheets/home.css", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/css/home.css")
    })
}
