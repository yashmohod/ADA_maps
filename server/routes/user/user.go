package user

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/yashmohod/server/models"
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

	var resp models.Response
	newHash, _ := HashPassword(r.FormValue("password"))
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	email := r.FormValue("email")

	if firstName == "" || lastName == "" || email == "" {
		resp = models.Response{
			Message: "Some of the data is blank!",
			Payload: "",
		}
	} else {
		_, err := models.GetUser(email)

		if err == nil {
			er := http.StatusConflict
			http.Error(w, "User with email:"+email+" already exists!", er)
			return
		}

		newUser := models.User{
			Id:           uuid.New(),
			FirstName:    firstName,
			LastName:     lastName,
			Email:        email,
			PasswordHash: newHash,
		}
		// fmt.Println(newUser)
		err = models.RegisterUser(newUser)
		if err != nil {
			resp = models.Response{
				Message: "Something went wrong while adding to the database!",
				Payload: "",
			}
		} else {
			resp = models.Response{
				Message: "User added successfully!",
				Payload: "",
			}

		}
	}
	// fmt.Println(resp)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if Authorize(r) {
		curUser, err := models.GetUser(r.FormValue("email"))
		if err != nil {
			er := http.StatusNotFound
			http.Error(w, "User not found!", er)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		tokens := struct {
			Id        string
			Email     string
			FirstName string
			LastName  string
		}{
			Id:        curUser.Id.String(),
			Email:     curUser.Email,
			FirstName: curUser.FirstName,
			LastName:  curUser.LastName,
		}
		json.NewEncoder(w).Encode(tokens)

	} else {
		er := http.StatusUnauthorized
		http.Error(w, "Unauthorized!", er)
		return
	}
}

func EditUserInfo(w http.ResponseWriter, r *http.Request) {
	if Authorize(r) {
		curUser, err := models.GetUser(r.FormValue("email"))
		if err != nil {
			er := http.StatusNotFound
			http.Error(w, "User not found!", er)
			return
		}

		if r.FormValue("newemail") != "" && r.FormValue("newemail") != curUser.Email {
			curUser.Email = r.FormValue("newemail")
		}
		if r.FormValue("firstName") != "" && r.FormValue("firstName") != curUser.FirstName {
			curUser.FirstName = r.FormValue("firstName")
		}
		if r.FormValue("lastName") != "" && r.FormValue("lastName") != curUser.LastName {
			curUser.LastName = r.FormValue("lastName")
		}
		if r.FormValue("password") != "" && !CheckPasswordHash(r.FormValue("password"), curUser.PasswordHash) {
			newHash, err := HashPassword(r.FormValue("password"))
			if err != nil {
				log.Println(err)
				er := http.StatusNotModified
				http.Error(w, "Unable to generate new hash!", er)
				return
			}
			curUser.PasswordHash = newHash
		}
		err = models.UpdateUserInfo(curUser)
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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if Authorize(r) {
		err := models.DeleteUser(r.FormValue("email"))
		if err != nil {
			er := http.StatusNotModified
			http.Error(w, "Unable to delete user!", er)
			return
		} else {
			w.Header().Set("Content-Type", "application/json")

			tokens := struct {
				Message string
			}{
				Message: "User deleted Successfully!",
			}

			json.NewEncoder(w).Encode(tokens)
		}
	} else {
		er := http.StatusUnauthorized
		http.Error(w, "Unauthorized!", er)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	password := r.FormValue("password")

	curUser, err := models.GetUser(email)
	if err != nil {
		er := http.StatusNotFound
		http.Error(w, "User with \" "+email+"\" not found!", er)
		return
	}
	// fmt.Println(curUser)
	passwordMatch := CheckPasswordHash(password, curUser.PasswordHash)

	if passwordMatch {

		sessionToken := GenerateToken(32)
		csrfToken := GenerateToken(32)

		//set session Cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "Session_token",
			Value:    sessionToken,
			Expires:  time.Now().Add(6 * time.Hour),
			HttpOnly: true,
		})

		//set CSRF token in a Cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "csrf_token",
			Value:    csrfToken,
			Expires:  time.Now().Add(6 * time.Hour),
			HttpOnly: false,
		})

		err := models.AddTokens(curUser.Id, sessionToken, csrfToken)
		if err != nil {
			er := http.StatusNotModified
			http.Error(w, "Unable to store tokens!", er)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		tokens := struct {
			Message string
		}{
			Message: "Login Successful!",
		}

		json.NewEncoder(w).Encode(tokens)

	} else {
		er := http.StatusUnauthorized
		http.Error(w, "Wrong email or password!", er)
		return
	}

}
func Logout(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")

	curUser, err := models.GetUser(email)
	if err != nil {
		er := http.StatusNotFound
		http.Error(w, "User with \" "+email+"\" not found!", er)
		return
	}

	if Authorize(r) {
		err = models.AddTokens(curUser.Id, "", "")
		if err != nil {
			er := http.StatusNotModified
			http.Error(w, "Unable to store tokens!", er)
			return
		}

		//set session Cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "Session_token",
			Value:    "",
			Expires:  time.Now().Add(1 * time.Minute),
			HttpOnly: true,
		})

		//set CSRF token in a Cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "csrf_token",
			Value:    "",
			Expires:  time.Now().Add(1 * time.Minute),
			HttpOnly: false,
		})

		w.Header().Set("Content-Type", "application/json")

		tokens := struct {
			Message string
		}{
			Message: "Logged out Successful!",
		}

		json.NewEncoder(w).Encode(tokens)

	} else {
		er := http.StatusUnauthorized
		http.Error(w, "Unauthorized!", er)
		return
	}

}
func Authorize(r *http.Request) bool {
	email := r.FormValue("email")
	curUser, err := models.GetUser(email)
	if err != nil {
		return false
	}

	st, err := r.Cookie("Session_token")
	if err != nil || st.Value == "" || curUser.SessionToken != st.Value {
		return false
	}

	csrf := r.Header.Get("X-CSRF-Token")
	if csrf == "" || curUser.CsrfToken != csrf {
		return false
	}

	return true
}
