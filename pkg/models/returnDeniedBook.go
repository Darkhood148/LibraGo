package models

func ReturnDeniedBook(data string) error {
	db, err := Connection()
	if err != nil {
		return err
	}
	defer db.Close()
	query := "UPDATE checkouts SET status = \"checkinPending\" WHERE checkoutID = (?)"
	_, err = db.Exec(query, data)
	if err != nil {
		return err
	}
	return nil
}
