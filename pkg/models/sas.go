package models

import (
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
		query = "UPDATE books SET copiesAvailable = copiesAvailable - (?) WHERE bookid = (?)"
	} else {
		query = "UPDATE books SET copiesAvailable = (?) WHERE bookid = (?)"
	}
	_, err = db.Exec(query, data.Quantity, data.Bookid)
	if err != nil {
		return err
	}
	return nil
}
