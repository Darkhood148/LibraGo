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
	defer db.Close()
	var updateQuery string
	if data.Option == "ADD" {
		updateQuery = "UPDATE books SET copiesAvailable = copiesAvailable + (?) WHERE bookID = (?)"
	} else if data.Option == "SUBTRACT" {
		selectQuery := "SELECT copiesAvailable FROM books WHERE bookID = (?)"
		res, err := db.Query(selectQuery, data.Bookid)
		if err != nil {
			return err
		}
		var copiesAvailable int
		res.Next()
		res.Scan(&copiesAvailable)
		if copiesAvailable > data.Quantity {
			updateQuery = "UPDATE books SET copiesAvailable = copiesAvailable - (?) WHERE bookID = (?)"
		} else {
			return errors.New("cannot remove more copies than exist")
		}
	} else {
		updateQuery = "UPDATE books SET copiesAvailable = (?) WHERE bookID = (?)"
	}
	_, err = db.Exec(updateQuery, data.Quantity, data.Bookid)
	if err != nil {
		return err
	}
	return nil
}
