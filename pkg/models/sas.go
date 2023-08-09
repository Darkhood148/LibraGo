package models

import (
	"errors"
	"fmt"
	"mvc/pkg/types"
)

func Sas(data types.SasData) error {
	db, err := Connection()
	if err != nil {
		return err
	}
	defer db.Close()
	var query string
	if data.Option == "ADD" {
		query = "UPDATE books SET copiesAvailable = copiesAvailable + (?) WHERE bookid = (?)"
	} else if data.Option == "SUBTRACT" {
		query2 := "SELECT copiesAvailable FROM books WHERE bookid = (?)"
		res, err := db.Query(query2, data.Bookid)
		if err != nil {
			return err
		}
		var temp int
		res.Next()
		res.Scan(&temp)
		if temp > data.Quantity {
			query = "UPDATE books SET copiesAvailable = copiesAvailable - (?) WHERE bookid = (?)"
		} else {
			return errors.New("cannot remove more copies than exist")
		}
	} else {
		query = "UPDATE books SET copiesAvailable = (?) WHERE bookid = (?)"
	}
	_, err = db.Exec(query, data.Quantity, data.Bookid)
	fmt.Println(query)
	if err != nil {
		return err
	}
	return nil
}
