package models

import (
	"fmt"
	"mvc/pkg/types"
)

func Sas(data types.SasData) {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error Occured")
	} else {
		var query string
		if data.Option == "ADD" {
			query = "UPDATE books SET copiesAvailable = copiesAvailable + (?) WHERE bookid = (?)"
		} else if data.Option == "SUBTRACT" {
			query = "UPDATE books SET copiesAvailable = copiesAvailable - (?) WHERE bookid = (?)"
		} else {
			query = "UPDATE books SET copiesAvailable = (?) WHERE bookid = (?)"
		}
		_, err := db.Exec(query, data.Quantity, data.Bookid)
		if err != nil {
			fmt.Println("Error Occured", err)
		} else {
			fmt.Println("Successful")
		}
	}

}
