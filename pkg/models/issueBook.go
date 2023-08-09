package models

import (
	"errors"
	"mvc/pkg/types"
	"time"
)

func IssueBook(data types.IssueBookData) error {
	db, err := Connection()
	if err != nil {
		return err
	}
	defer db.Close()
	query := "SELECT copiesAvailable FROM books WHERE bookid = (?)"
	res, err := db.Query(query, data.Bookid)
	if err != nil {
		return err
	}
	res.Next()
	var temp int
	res.Scan(&temp)
	if temp > 0 {
		query := "SELECT * FROM checkouts WHERE byUser = (?) AND ofBook = (?)"
		res, err := db.Query(query, data.Username, data.Bookid)
		if err != nil {
			return err
		}
		if res.Next() {
			return errors.New("you already own this book")
		} else {
			query := "INSERT INTO checkouts (ofBook, byUser, status, issueTime) VALUES (?, ?, ?, ?)"
			_, err := db.Exec(query, data.Bookid, data.Username, "pending", time.Now())
			if err != nil {
				return err
			}
			return nil
		}
	} else {
		return errors.New("no copies are available")
	}
}
