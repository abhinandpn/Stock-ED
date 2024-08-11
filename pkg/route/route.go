package route

import (
	"net/http"
)

// RegisterRoutes sets up the root paths for the HTML files.
func RegisterRoutes() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
}

// homeHandler serves the home.html file.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/html/home.html")
}

// aboutHandler serves the about.html file.
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/html/about.html")
}

// contactHandler serves the contact.html file.
func contactHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/html/contact.html")
}
