package models

import (
	"rest-api/db"
	"rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() (*User, error) {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return nil, err
	}

	userID, err := result.LastInsertId()
	u.ID = userID
	u.Password = hashedPassword

	return u, err
}
