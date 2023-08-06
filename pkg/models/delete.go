package models

import "fmt"

func DeleteBook(data int) {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error Occured")
	} else {
		query := "DELETE FROM books WHERE bookid = (?)"
		_, err := db.Exec(query, data)
		if err != nil {
			fmt.Println("Error Occured")
		} else {
			fmt.Println("Success")
		}
	}
}
