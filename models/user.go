package models

import (
	"errors"
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

func (u User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("invalid credentials")
	}

	isPasswordCorrect := utils.IsPasswordCorrect(retrievedPassword, u.Password)
	if !isPasswordCorrect {
		return errors.New("invalid credentials")
	}

	return nil
}
