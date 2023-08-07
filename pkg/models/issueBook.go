package models

import (
	"errors"
	"mvc/pkg/types"
)

func IssueBook(data types.IssueBookData) error {
	db, err := Connection()
	if err != nil {
		return err
	}
	query := "SELECT * FROM checkouts WHERE byUser = (?) AND ofBook = (?)"
	res, err := db.Query(query, data.Username, data.Bookid)
	if err != nil {
		return err
	}
	if res.Next() {
		return errors.New("you already own this book")
	} else {
		query := "INSERT INTO checkouts (ofBook, byUser, status) VALUES (?, ?, ?)"
		_, err := db.Exec(query, data.Bookid, data.Username, "pending")
		if err != nil {
			return err
		}
		query = "UPDATE books SET copiesAvailable = copiesAvailable - 1 WHERE bookid = (?)"
		_, err = db.Exec(query, data.Bookid)
		if err != nil {
			return err
		}
		return nil
	}
}
