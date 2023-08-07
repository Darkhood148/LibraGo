package models

func ReturnBook(data string) error {
	db, err := Connection()
	if err != nil {
		return err
	} else {
		query := "UPDATE checkouts SET status = \"checkinPending\" WHERE checkoutid = (?)"
		_, err := db.Exec(query, data)
		if err != nil {
			return err
		}
		return nil
	}
}
