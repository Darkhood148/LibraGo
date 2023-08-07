package models

func ReturnDeniedBook(data string) error {
	db, err := Connection()
	if err != nil {
		return err
	}
	query := "UPDATE checkouts SET status = \"checkinPending\" WHERE checkoutid = (?)"
	_, err = db.Exec(query, data)
	if err != nil {
		return err
	}
	return nil
}
