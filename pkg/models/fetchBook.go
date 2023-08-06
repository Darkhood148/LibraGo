package models

import (
	"fmt"

	"mvc/pkg/types"
)

func FetchBooks() types.Books {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}
	selectSql := "SELECT * FROM books"
	rows, err := db.Query(selectSql)
	db.Close()

	if err != nil {
		fmt.Printf("error %s querying the database", err)
	}

	var fetchBooks []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.Bookid, &book.Bookname, &book.Author, &book.Quantity)
		if err != nil {
			fmt.Printf("error %s scanning the row", err)
		}
		fetchBooks = append(fetchBooks, book)
	}

	var listBooks types.Books
	listBooks.Books = fetchBooks
	return listBooks

}
