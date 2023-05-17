package models

import (
	"errors"
)

type User struct {
	ID        int    `json:"user_id" db:"user_id"`
	Name      string `json:"login" db:"user_name"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password" binding:"required"`
}

func (u *User) Register() error {
	err := db.QueryRow("INSERT INTO users (user_name, first_name, last_name, email, password) VALUES ($1, $2, $3, $4, $5) RETURNING user_id",
		u.Name, u.FirstName, u.LastName, u.Email, u.Password).Scan(&u.ID)
	if err != nil {
		switch err.Error() {
		case "pq: duplicate key value violates unique constraint \"users_email_key\"":
			return errors.New("email already exists")
		}
	}

	u.Password = ""

	return nil
}

func (u *User) Login() error {
	return errors.New("not implemented")
}

func (u *User) GetUser() error {
	err := db.Get(u, "SELECT * FROM users WHERE user_id=$1", u.ID)
	if err != nil {
		switch err.Error() {
		case "sql: no rows in result set":
			return errors.New("user not found")
		}
	}

	u.Password = ""

	return nil
}

func (u *User) GetFriends() ([]User, error) {
	var Users []User
	err := db.Select(&Users, "SELECT u.first_name, u.last_name FROM users_users JOIN users u on users_users.user1_id = u.user_id JOIN users u2 on users_users.user2_id = u2.user_id WHERE (user1_id=$1 OR user2_id=$1) AND u.user_id!=$1", u.ID)
	if err != nil {
		switch err.Error() {
		case "sql: no rows in result set":
			return nil, errors.New("user not found")
		}
	}

	return Users, nil
}
