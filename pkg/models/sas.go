package models

import (
	"errors"
	"mvc/pkg/types"
)

func Sas(data types.SasData) error {
	db, err := Connection()
	if err != nil {
		return err
	}
	var query string
	if data.Option == "ADD" {
		query = "UPDATE books SET copiesAvailable = copiesAvailable + (?) WHERE bookid = (?)"
	} else if data.Option == "SUBTRACT" {
		query := "SELECT copiesAvailable FROM books WHERE bookid = (?)"
		res, err := db.Query(query, data.Bookid)
		if err != nil {
			return err
		}
		var temp int
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
	if err != nil {
		return err
	}
	return nil
}
