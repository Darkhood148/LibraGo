package middleware

import (
	"fmt"
	"mvc/pkg/models"
)

func VerifyAdmin(name string) bool {
	db, err := models.Connection()
	if err != nil {
		fmt.Println("Error Occured")
		return false
	}
	query := "SELECT isAdmin FROM users WHERE username=(?)"
	res, err := db.Query(query, name)
	if err != nil {
		fmt.Println("Error Occured")
		return false
	}
	res.Next()
	var admin bool
	err = res.Scan(&admin)
	if err != nil {
		fmt.Println("Error Occured")
		return false
	}
	if admin {
		return true
	} else {
		return false
	}
}
