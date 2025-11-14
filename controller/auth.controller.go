package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"userStory/dao"
	"userStory/model"
	"userStory/utils"
)

// Create user
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var usr = model.User{}

	if err := json.NewDecoder(r.Body).Decode(&usr); err != nil {
		http.Error(w, "Invalid user", http.StatusBadRequest)
	}

	if usr.Name == "" || usr.Email == "" {
		http.Error(w, "Name and Email is required", http.StatusNotFound)
		return
	}

	isExist, err := dao.CheckEmailExist(usr)
	if isExist {
		http.Error(w, "Email already Exist", http.StatusConflict)
		return
	}
	if err != nil {
		log.Println(err)
	}
	hashedPassword, err := utils.Hash(usr.Password)
	if err != nil {
		log.Println(err)
	}

	usr.Password = string(hashedPassword)

	err = dao.CreateUser(usr)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User Created Successfully"})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var credentials = model.Login{}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid User", http.StatusBadRequest)
	}

	if credentials.Email == "" || credentials.Password == "" {
		http.Error(w, "Email and Password required", http.StatusNotAcceptable)
	}

	user, err := dao.LoginUser(credentials)
	if err != nil {
		http.Error(w, "user not exist", http.StatusUnauthorized)
		log.Println(err)
		return
	}

	err = utils.VerifyHash(user.Password, credentials.Password)
	if err != nil {
		http.Error(w, "Wrong password", http.StatusUnauthorized)
		return
	}

	type userWithoutPass struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	var loggedInUser = userWithoutPass{ID: user.ID, Name: user.Name, Email: user.Email}
	json.NewEncoder(w).Encode(loggedInUser)
}
