package main

import (
    "fmt"
    "net/http"

    "github.com/abhinandpn/Stock-ED/pkg/loader"  // Import the loader package
    "github.com/abhinandpn/Stock-ED/pkg/route"   // Import the route package
)

func main() {
    // Register asset loader
    loader.AssetLoader()
    fmt.Println("Asset loading complete...")

    // Register routes for HTML files
    route.RegisterRoutes()

    // Register handlers for CSS and JS
    loader.CSSLoader()
    loader.JSLoader()
    fmt.Println("CSS and JS loading complete...")

    // Serve static files (CSS, JS)
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", fs))
    http.Handle("/scripts/", http.StripPrefix("/scripts/", fs))
    fmt.Println("Serving static files accomplished.....")

    // Start HTTP server
    fmt.Println("Server running at port 8080")
    fmt.Println("Check - http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}
