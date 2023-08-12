package models

import (
	"database/sql"
	"mvc/pkg/types"

	_ "github.com/go-sql-driver/mysql"
)

func FetchBooks(*sql.DB) (types.Books, error) {
	db, err := Connection()
	if err != nil {
		return types.Books{}, err
	}
	defer db.Close()
	selectSql := "SELECT * FROM books WHERE copiesAvailable > 0"
	rows, err := db.Query(selectSql)
	db.Close()
	if err != nil {
		return types.Books{}, err
	}
	var fetchBooks []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.Bookid, &book.Bookname, &book.Author, &book.Quantity)
		if err != nil {
			return types.Books{}, err
		}
		fetchBooks = append(fetchBooks, book)
	}
	var listBooks types.Books
	listBooks.Books = fetchBooks
	return listBooks, nil
}
