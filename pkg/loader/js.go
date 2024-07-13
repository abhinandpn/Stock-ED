package loader

import (
    "net/http"
)

func JSLoader() {
    http.HandleFunc("/scripts/about.js", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/js/about.js")
    })

    http.HandleFunc("/scripts/contact.js", func(w http.ResponseWriter, r *http.Request) { // Fixed typo
        http.ServeFile(w, r, "./static/js/contact.js")
    })

    http.HandleFunc("/scripts/home.js", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/js/home.js")
    })

    
}
