package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yashmohod/server/models"
	"github.com/yashmohod/server/routes/location"
	"github.com/yashmohod/server/routes/user"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Got a hit!")
		w.Header().Set("Content-Type", "application/json")
		resp := models.Response{
			Message: "Reach maps api!",
			Payload: "",
		}
		json.NewEncoder(w).Encode(resp)
	})

	mux.HandleFunc("/user", user.HandelUserRequest)
	mux.HandleFunc("/location", location.HandelLocationRequest)

	fmt.Println(http.ListenAndServe("localhost:8080", mux))

}
