package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var Router = mux.NewRouter()

// handle "/user" routes
var UserPath = Router.PathPrefix("/user").Subrouter()

func CreateServer() {
	http.Handle("/", Router)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	// Add CORS Policies
	serverHandler := cors.AllowAll().Handler(Router)

	fmt.Printf("The Server is UP in Localhost:%v \n", PORT)
	// Turns up server and listen connections
	log.Fatal(http.ListenAndServe(":"+PORT, serverHandler))
}
