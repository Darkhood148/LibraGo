package models

import "fmt"

func ReturnDeniedBook(data string) {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error Occured")
	} else {
		query := "UPDATE checkouts SET status = \"checkinPending\" WHERE checkoutid = (?)"
		_, err := db.Exec(query, data)
		if err != nil {
			fmt.Println("Error Occured")
		} else {
			fmt.Println("Success")
		}
	}
}
