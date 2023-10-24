package main

import (
	"log"
	"net/http"
	"os"
	"swms-sa/internal/routes"
	"time"
)

func main() {

	/* build routing */
	router := http.NewServeMux()
	for _, route := range routes.MakeRoute {
		router.HandleFunc(route.Url, routes.Handler(route.Func, route.Method))
	}

	/* create server */
	server := &http.Server{
		Addr:         ":9095",
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	/* start server */
	log.Println("Server started on port 9095")
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
