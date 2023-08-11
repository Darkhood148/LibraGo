package models

import (
	"mvc/pkg/types"
)

func FetchInventory() (types.Inventory, error) {
	db, err := Connection()
	if err != nil {
		return types.Inventory{}, err
	}
	defer db.Close()
	selectSql := "SELECT * from books"
	rows, err := db.Query(selectSql)
	if err != nil {
		return types.Inventory{}, err
	}
	var inv []types.BookInventory
	for rows.Next() {
		var inven types.BookInventory
		err := rows.Scan(&inven.Bookid, &inven.Bookname, &inven.Author, &inven.CopiesAvailable)
		if err != nil {
			return types.Inventory{}, err
		}
		query := "SELECT COUNT(*) FROM checkouts WHERE ofBook = (?) AND status != \"pending\""
		var temp int
		err = db.QueryRow(query, inven.Bookid).Scan(&temp)
		if err != nil {
			return types.Inventory{}, err
		}
		inven.TotalCopies = inven.CopiesAvailable + temp
		inv = append(inv, inven)
	}
	var invent types.Inventory
	invent.Books = inv
	return invent, nil
}
