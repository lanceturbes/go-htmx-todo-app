package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		// Handle 404 ("/" is also used as a catch-all route, so you need to check the URL specifically)
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			Render(w, notFound())
			return
		}

		w.WriteHeader(http.StatusOK)
		Render(w, homePage())
	})

	// Serve all files in the static folder under the "/static" url path
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("GET /static", fileServer)

	// Start the server, and log when it shuts down
	fmt.Println("Listening at http://localhost:4200 ...")
	err := http.ListenAndServe("localhost:4200", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Writes a templ component to an http response body
func Render(w http.ResponseWriter, t templ.Component) {
	t.Render(context.Background(), w)
}
