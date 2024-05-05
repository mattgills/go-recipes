package ingredients

import (
	"fmt"
	"html"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ingredients Handler")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
