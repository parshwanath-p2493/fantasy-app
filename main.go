package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"fantasy-backend/config"
	"fantasy-backend/database"
	"fantasy-backend/routes"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()
	database.Connect()

	router := mux.NewRouter()
	routes.AuthRoutes(router)

	port := os.Getenv("PORT")
	fmt.Println("🚀 Server running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
