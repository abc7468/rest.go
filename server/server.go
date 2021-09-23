package server

import (
	"fmt"
	"log"
	"net/http"

	user "github.com/abc7468/rest.go/user"
)

func home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello World")
}

func query(rw http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("query")
	if q == "" {
		q = "nothing"
	}
	fmt.Fprintf(rw, "Hi %s", q)
}

func Start(port int) {

	mux := http.NewServeMux()
	mux.Handle("/user", &user.UserHandler{})
	mux.HandleFunc("/", home)
	mux.HandleFunc("/query", query)
	fmt.Printf("Listening on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))

}
