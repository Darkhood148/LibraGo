package models

func DeleteBook(data int) error {
	db, err := Connection()
	if err != nil {
		return err
	} else {
		query := "DELETE FROM books WHERE bookid = (?)"
		_, err := db.Exec(query, data)
		if err != nil {
			return err
		}
		return nil
	}
}
