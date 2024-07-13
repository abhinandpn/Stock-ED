// asset.go
package loader

import (
	"net/http"
)

// AssetLoader serves static assets like images
func AssetLoader() {

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/asset"))))
}
