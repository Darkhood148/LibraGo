package models

import (
	"errors"
	"mvc/pkg/types"
)

func CheckRequest(data int, status string) error {
	db, err := Connection()
	if err != nil {
		return err
	}
	query := "SELECT * FROM checkouts WHERE checkoutID = (?)"
	res, err := db.Query(query, data)
	if err != nil {
		return err
	}
	defer db.Close()
	if res.Next() {
		var resData types.CheckRequest
		var temp string
		err := res.Scan(&resData.Checkoutid, &resData.OfBook, &resData.ByUser, &resData.Status, &temp)
		if err != nil {
			return err
		}
		if resData.Status == "pending" {
			if status == "approve" {
				query := "UPDATE checkouts SET status = \"issued\" WHERE checkoutID = (?)"
				_, err := db.Exec(query, resData.Checkoutid)
				if err != nil {
					return err
				}
				query = "UPDATE books SET copiesAvailable = copiesAvailable - 1 WHERE bookID = (?)"
				_, err = db.Exec(query, resData.OfBook)
				if err != nil {
					return err
				}
				query = "SELECT copiesAvailable FROM books WHERE bookid = (?)"
				res, err = db.Query(query, resData.OfBook)
				if err != nil {
					return err
				}
				res.Next()
				var ca int
				res.Scan(&ca)
				if ca <= 0 {
					query = "DELETE FROM checkouts WHERE ofBook = (?) AND status = \"pending\""
					_, err = db.Exec(query, resData.OfBook)
					if err != nil {
						return err
					}
				}
				return nil
			} else {
				query := "DELETE FROM checkouts WHERE checkoutID = (?)"
				_, err := db.Exec(query, resData.Checkoutid)
				if err != nil {
					return err
				}
				return nil
			}
		} else {
			if status == "approve" {
				query := "DELETE FROM checkouts WHERE checkoutID = (?)"
				_, err := db.Exec(query, resData.Checkoutid)
				if err != nil {
					return err
				}
				query = "UPDATE books SET copiesAvailable = copiesAvailable + 1 WHERE bookID = (?)"
				_, err = db.Exec(query, resData.OfBook)
				if err != nil {
					return err
				}
				return nil
			} else {
				query := "UPDATE checkouts SET status = \"checkinDenied\" WHERE checkoutID = (?)"
				_, err := db.Exec(query, resData.Checkoutid)
				if err != nil {
					return err
				}
				return nil
			}
		}
	} else {
		return errors.New("Checkout-ID is not valid")
	}
}
