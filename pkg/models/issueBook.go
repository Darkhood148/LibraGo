package models

import (
	"fmt"
	"mvc/pkg/types"
)

func IssueBook(data types.IssueBookData) {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error Occured")
	} else {
		query := "SELECT * FROM checkouts WHERE byUser = (?) AND ofBook = (?)"
		res, err := db.Query(query, data.Username, data.Bookid)
		if err != nil {
			fmt.Println("Error Occured 1", err)
		} else {
			if res.Next() {
				fmt.Println("User has already this book")
			} else {
				query := "INSERT INTO checkouts (ofBook, byUser, status) VALUES (?, ?, ?)"
				_, err := db.Exec(query, data.Bookid, data.Username, "pending")
				if err != nil {
					fmt.Println("Error Occured 2", err)
				} else {
					query := "UPDATE books SET copiesAvailable = copiesAvailable - 1 WHERE bookid = (?)"
					_, err := db.Exec(query, data.Bookid)
					if err != nil {
						fmt.Println("Error Occured 3", err)
					} else {
						fmt.Println("Success")
					}
				}
			}
		}
	}
}
