package dao

import (
	"database/sql"
	"userStory/model"
)

// Create user
func CreateUser(user model.User) error {
	query := `insert into 
			users (name, email, password) 
			values($1, $2, $3)`

	_, err := DB.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

// Check email exist
func CheckEmailExist(user model.User) (bool, error) {
	query := `SELECT id, name, email FROM users where email=$1`
	var usr = model.User{}
	err := DB.QueryRow(query, user.Email).Scan(&usr.ID, &usr.Name, &usr.Email)

	// Email no exist
	if err == sql.ErrNoRows {
		return false, nil
	}
	//db error
	if err != nil {
		return false, err
	}

	return true, nil
}

func LoginUser(credentials model.Login) (model.User, error) {
	query := `Select id, name, email, password 
			from users
			where email=$1`

	var user = model.User{}
	err := DB.QueryRow(query, credentials.Email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return user, err
	}
	if err != nil {
		return user, err
	}

	return user, nil
}
