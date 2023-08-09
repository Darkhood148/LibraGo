package models

import (
	"errors"
	"mvc/pkg/types"
)

func AddBook(data types.Book) error {
	db, err := Connection()
	if err != nil {
		return err
	}
	defer db.Close()
	query := "SELECT * FROM books WHERE name=(?) AND author=(?)"
	res, err := db.Query(query, data.Bookname, data.Author)
	if err != nil {
		return err
	}
	if res.Next() {
		return errors.New("book already exists")
	} else {
		query := "INSERT INTO books (name, author, copiesAvailable) VALUES (?, ?, ?)"
		_, err := db.Exec(query, data.Bookname, data.Author, data.Quantity)
		if err != nil {
			return err
		} else {
			return nil
		}
	}

}
