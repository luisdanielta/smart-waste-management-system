package routes

import (
	"log"
	"net/http"
)

type Route struct {
	Url    string
	Func   []http.HandlerFunc
	Method []string
}

func Handler(funcs []http.HandlerFunc, methods []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/* check method */
		for i := 0; i < len(methods); i++ {
			if r.Method == methods[i] {
				w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
				w.Header().Set("Access-Control-Allow-Methods", methods[i])
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
				log.Printf("url: %s - method: %s", r.URL.Path, r.Method)
				funcs[i](w, r)
				return
			}
		}
		/* method not allowed */
		http.Error(w, "405 Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}
}

var MakeRoute = []Route{ // List of the routes on server
	{Url: "/", Func: []http.HandlerFunc{handleIndex}, Method: []string{"GET"}},

	/* */

}
