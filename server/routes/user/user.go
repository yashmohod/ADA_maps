package user

import (
	"net/http"
)

func HandelUserRequest(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		GetUser(w, r)
	case http.MethodPost:
		AddUser(w, r)
	case http.MethodDelete:
		DeleteUser(w, r)
	case http.MethodPatch:
		EditUserInfo(w, r)
	default:
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)

	}
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func EditUserInfo(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}
