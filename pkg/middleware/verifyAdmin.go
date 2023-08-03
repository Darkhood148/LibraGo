package middleware

import (
	"fmt"
	"mvc/pkg/models"
)

func VerifyAdmin(name string) bool {
	db, err := models.Connection()
	if err != nil {
		fmt.Println("Error Occured")
		fmt.Println("Err1")
		return false
	}
	query := "SELECT isAdmin FROM users WHERE username=(?)"
	res, err := db.Query(query, name)
	if err != nil {
		fmt.Println("Error Occured")
		fmt.Println("Err2")
		return false
	}
	res.Next()
	var admin bool
	err = res.Scan(&admin)
	if err != nil {
		fmt.Println("Error Occured")
		fmt.Println("Err3")
		return false
	}
	if admin {
		return true
	} else {
		return false
	}
}
