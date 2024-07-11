package main

import (
	"net/http"

	"github.com/abhinandpn/Stock-ED/pkg/loader" // Adjust import path as needed
)

func main() {
	// Register handlers for CSS, HTML, JS
	loader.AboutCssLoader()
	loader.HtmlLoader()
	loader.JSLoader()

	// Serve static files (CSS, JS, HTML)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/stylesheets/", fs)
	http.Handle("/scripts/", fs)

	// Start HTTP server
	http.ListenAndServe(":8080", nil)
}
