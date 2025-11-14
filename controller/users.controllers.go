package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"userStory/model"
)

var users = []model.User{
	{ID: 1, Name: "Hasib", Email: "md.hasibuzzaman001@gmail.com"},
	{ID: 2, Name: "Itu", Email: "iffath.com"},
}

// GET: All users
func GetAlluser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}

// GET: User by Id
func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if user.ID == id {
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	http.Error(w, "User Not Found", http.StatusNotFound)
}

// DELETE: user by id
func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
	}

	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "User deleted"})
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)

}
