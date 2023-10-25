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

	/* auto migrate */
	// pkg.ConnDB.GetConn()

	// db.AutoMigrate(&models.ControllerC{})
	// db.AutoMigrate(&models.Container{})

	/* create server */
	PORT := os.Getenv("PORT_SERVER")
	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	/* start server */
	log.Println("Server started, port:", PORT)
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
