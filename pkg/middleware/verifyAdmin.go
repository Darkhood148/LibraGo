package middleware

import (
	"fmt"
	"mvc/pkg/models"
)

func VerifyAdmin(name string) bool {
	db, err := models.Connection()
	if err != nil {
		fmt.Println("Error Occured in establishing connection")
		return false
	}
	defer db.Close()
	query := "SELECT isAdmin FROM users WHERE username=(?)"
	res, err := db.Query(query, name)
	if err != nil {
		fmt.Println("Error Occured while querying")
		return false
	}
	res.Next()
	var admin bool
	err = res.Scan(&admin)
	if err != nil {
		fmt.Println("Some Error Occured")
		return false
	}
	if admin {
		return true
	} else {
		return false
	}
}
