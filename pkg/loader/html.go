package loader

import (
    "net/http"
)

func HtmlLoader() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/html/home.html")
    })

    http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/html/about.html")
    })

    http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/html/contact.html")
    })
}
