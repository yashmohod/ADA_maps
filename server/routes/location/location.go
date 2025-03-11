package location

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/yashmohod/server/models"
	"github.com/yashmohod/server/routes/user"
)

func HandelLocationRequest(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		GetLocation(w, r)
	case http.MethodPost:
		AddLocation(w, r)
	case http.MethodDelete:
		DeleteLocation(w, r)
	case http.MethodPatch:
		EditLocationInfo(w, r)
	default:
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)

	}
}

func AddLocation(w http.ResponseWriter, r *http.Request) {
	if user.Authorize(r) {
		curUser, err := models.GetUser(r.FormValue("email"))
		if err != nil {
			er := http.StatusNotFound
			http.Error(w, "User not found!", er)
			return
		}
		var resp models.Response

		name := r.FormValue("name")
		latitude, err := strconv.ParseFloat(r.FormValue("latitude"), 64)
		longitude, er := strconv.ParseFloat(r.FormValue("longitude"), 64)

		if err != nil || er != nil {
			log.Println(err, er)
			er := http.StatusConflict
			http.Error(w, "Latitude or longitude contain forbidden chars!", er)
			return
		}

		if name == "" || latitude == 0 || longitude == 0 {
			resp = models.Response{
				Message: "Some of the data is blank!",
				Payload: "",
			}
		} else {

			newLocation := models.Location{
				Id:        uuid.New(),
				Name:      name,
				Latitude:  latitude,
				Longitude: longitude,
				EntryBy:   curUser.Id,
			}
			// fmt.Println(newUser)
			err := models.AddLocation(newLocation)
			if err != nil {
				log.Println(err)
				er := http.StatusConflict
				http.Error(w, "Unable to add location. Try again!", er)
				return
			}
			resp = models.Response{
				Message: "Location added successfully!",
				Payload: "",
			}

		}
		// fmt.Println(resp)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	} else {
		er := http.StatusUnauthorized
		http.Error(w, "Unauthorized!", er)
		return
	}

}

func GetLocation(w http.ResponseWriter, r *http.Request) {
	if user.Authorize(r) {
		locations, err := models.GetLocations()
		if err != nil {
			log.Println(err)
			er := http.StatusConflict
			http.Error(w, "Unable to get locations!", er)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(locations)
	} else {
		er := http.StatusUnauthorized
		http.Error(w, "Unauthorized!", er)
		return
	}
}

func EditLocationInfo(w http.ResponseWriter, r *http.Request) {
	if user.Authorize(r) {
		curLocation, err := models.GetLocation(r.FormValue("locationId"))
		if err != nil {
			er := http.StatusNotFound
			http.Error(w, "User not found!", er)
			return
		}
		latitude, err := strconv.ParseFloat(r.FormValue("latitude"), 64)
		longitude, er := strconv.ParseFloat(r.FormValue("longitude"), 64)
		if err != nil || er != nil {
			log.Println(err, er)
			er := http.StatusConflict
			http.Error(w, "Latitude or longitude contain forbidden chars!", er)
			return
		}

		if r.FormValue("name") != "" && r.FormValue("name") != curLocation.Name {
			curLocation.Name = r.FormValue("name")
		}
		if r.FormValue("latitude") != "" && latitude != curLocation.Latitude {
			curLocation.Latitude = latitude
		}
		if r.FormValue("longitude") != "" && longitude != curLocation.Longitude {
			curLocation.Longitude = longitude
		}

		err = models.UpdateLocation(curLocation)
		if err != nil {
			log.Println(err)
			er := http.StatusNotModified
			http.Error(w, "Unable to update user info!", er)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		tokens := models.Response{
			Message: "User info updated!",
			Payload: "",
		}
		json.NewEncoder(w).Encode(tokens)

	} else {
		er := http.StatusUnauthorized
		http.Error(w, "Unauthorized!", er)
		return
	}
}

func DeleteLocation(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}
