package location

import (
	"net/http"
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
	panic("unimplemented")
}

func GetLocation(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func EditLocationInfo(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func DeleteLocation(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}
