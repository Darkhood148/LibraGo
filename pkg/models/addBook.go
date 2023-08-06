package models

import (
	"fmt"
	"mvc/pkg/types"
)

func AddBook(data types.Book) {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error Occured")
	} else {
		query := "SELECT * FROM books WHERE name=(?) AND author=(?)"
		res, err := db.Query(query, data.Bookname, data.Author)
		if err != nil {
			fmt.Println("Error Occured")
		} else {
			if res.Next() {
				fmt.Println("Book already exists")
			} else {
				query := "INSERT INTO books (name, author, copiesAvailable) VALUES (?, ?, ?)"
				_, err := db.Exec(query, data.Bookname, data.Author, data.Quantity)
				if err != nil {
					fmt.Println("Error Occured")
				} else {
					fmt.Println("Successful")
				}
			}
		}
	}

}
