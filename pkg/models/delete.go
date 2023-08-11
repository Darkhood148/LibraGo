package models

import "errors"

func DeleteBook(data int) error {
	db, err := Connection()
	if err != nil {
		return err
	} else {
		defer db.Close()
		query := "SELECT * FROM books WHERE ofBook = (?)"
		res, err := db.Query(query, data)
		if err != nil {
			return err
		} else if res.Next() {
			return errors.New("book has been lent out and hence cannot be removed")
		}
		query = "DELETE FROM books WHERE bookid = (?)"
		_, err = db.Exec(query, data)
		if err != nil {
			return err
		}
		return nil
	}
}
