package main

import (
	"fmt"
	"net/http"

	"github.com/abhinandpn/Stock-ED/pkg/loader" // Adjust import path as needed
)

func main() {
	// Register asset loader
	loader.AssetLoader()
	fmt.Println("asset loading complete...")

	// Register handlers for CSS, HTML, JS
	loader.CSSLoader()
	loader.HtmlLoader()
	loader.JSLoader()
	fmt.Println("css,html,js loding complete...")

	// Serve static files (CSS, JS, HTML)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/stylesheets/", fs)
	http.Handle("/scripts/", fs)
    fmt.Println("server static file accomplished.....")

	// Start HTTP server
	fmt.Println("server runnign at port 8080")
    fmt.Println("check - http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
