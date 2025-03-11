package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/yashmohod/server/middleware"
	"github.com/yashmohod/server/models"
	"github.com/yashmohod/server/routes/location"
	"github.com/yashmohod/server/routes/user"
)

func main() {
	// establishing connection to database
	models.ConnectDatabase()
	// makes sure to terminate connection to database after main ends
	defer models.DisconnectDatabase()

	// api server init
	mux := http.NewServeMux()

	// healthCheck
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Got a hit!")
		w.Header().Set("Content-Type", "application/json")
		resp := models.Response{
			Message: "Reach maps api!",
			Payload: "",
		}
		json.NewEncoder(w).Encode(resp)
	})

	// routes
	mux.HandleFunc("/user", user.HandelUserRequest)
	mux.HandleFunc("POST /login", user.Login)
	mux.HandleFunc("POST /logout", user.Logout)
	mux.HandleFunc("/location", location.HandelLocationRequest)

	// starting api server
	log.Println("Server starting...")
	log.Print("Listening on :8000...")
	server := http.Server{
		Addr:    ":8000",
		Handler: middleware.Logging(mux),
	}
	err := server.ListenAndServe()
	log.Fatal(err)

}
