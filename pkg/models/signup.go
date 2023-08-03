package models

import (
	"fmt"
	"mvc/pkg/types"

	"golang.org/x/crypto/bcrypt"
)

func SignUp(data types.SignupData) {
	if (data.Fullname=="" || data.Username=="" || data.Password==""){
		fmt.Println("One or more inputs are null")
		return
	}
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	} else {
		check := "SELECT * FROM users WHERE username=(?)"
		res, err := db.Query(check, data.Username)
		if err != nil {
			fmt.Println("Error Occured")
		} else if res.Next() {
			fmt.Println("User already exists")
		} else {
			if data.Password != data.CPassword {
				fmt.Println("Password and Confirm Passwords do not match")
			} else {
				pswd := []byte(data.Password)
				hashedPassword, err := bcrypt.GenerateFromPassword(pswd, bcrypt.DefaultCost)
				if err != nil {
					panic(err)
				} else {
					insertion := "INSERT INTO users VALUES (?, ?, ?, ?)"
					_, err := db.Exec(insertion, data.Username, data.Fullname, hashedPassword, data.IsAdmin)
					if err != nil {
						fmt.Println("Error Occured")
					} else {
						fmt.Println("Success")
					}
				}
			}
		}
	}
}
