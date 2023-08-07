package models

import (
	"errors"
	"mvc/pkg/types"

	"golang.org/x/crypto/bcrypt"
)

func SignUp(data types.SignupData) error {
	if data.Fullname == "" || data.Username == "" || data.Password == "" {
		return errors.New("one or more inputs are null")
	}
	db, err := Connection()
	if err != nil {
		return err
	}
	check := "SELECT * FROM users WHERE username=(?)"
	res, err := db.Query(check, data.Username)
	if err != nil {
		return err
	} else if res.Next() {
		return errors.New("username already exists")
	} else {
		if data.Password != data.CPassword {
			return errors.New("passwords and confirmed passwords do not match")
		} else {
			pswd := []byte(data.Password)
			hashedPassword, err := bcrypt.GenerateFromPassword(pswd, bcrypt.DefaultCost)
			if err != nil {
				return err
			} else {
				insertion := "INSERT INTO users VALUES (?, ?, ?, ?)"
				_, err := db.Exec(insertion, data.Username, data.Fullname, hashedPassword, data.IsAdmin)
				if err != nil {
					return err
				}
				return nil
			}
		}
	}
}
