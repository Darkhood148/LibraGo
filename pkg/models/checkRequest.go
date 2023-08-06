package models

import (
	"fmt"
	"mvc/pkg/types"
)

func CheckRequest(data int, status string) {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error Occured")
	} else {
		query := "SELECT * FROM checkouts WHERE checkoutid = (?)"
		res, err := db.Query(query, data)
		if err != nil {
			fmt.Println("Error Occured")
		} else {
			if res.Next() {
				var resData types.CheckRequest
				err := res.Scan(&resData.Checkoutid, &resData.OfBook, &resData.ByUser, &resData.Status)
				fmt.Println(resData)
				if err != nil {
					fmt.Println("Error Occured")
				} else {
					if resData.Status == "pending" {
						if status == "approve" {
							query := "UPDATE checkouts SET status = \"issued\" WHERE checkoutid = (?)"
							_, err := db.Exec(query, resData.Checkoutid)
							if err != nil {
								fmt.Println("Error Occured")
							} else {
								fmt.Println("SUCCESS")
							}
						} else {
							query := "DELETE FROM checkouts WHERE checkoutid = (?)"
							_, err := db.Exec(query, resData.Checkoutid)
							if err != nil {
								fmt.Println("Error Occured")
							} else {
								query := "UPDATE books SET copiesAvailable = copiesAvailable + 1 WHERE bookid = (?)"
								_, err := db.Exec(query, resData.OfBook)
								if err != nil {
									fmt.Println("Error Occured")
								} else {
									fmt.Println("Success")
								}
							}
						}
					} else {
						if status == "approve" {
							query := "DELETE FROM checkouts WHERE checkoutid = (?)"
							_, err := db.Exec(query, resData.Checkoutid)
							if err != nil {
								fmt.Println("Error Occured")
							} else {
								query := "UPDATE books SET copiesAvailable = copiesAvailable + 1 WHERE bookid = (?)"
								_, err := db.Exec(query, resData.OfBook)
								if err != nil {
									fmt.Println("Error Occured")
								} else {
									fmt.Println("Success")
								}
							}
						} else {
							query := "UPDATE checkouts SET status = \"checkinDenied\" WHERE checkoutid = (?)"
							_, err := db.Exec(query, resData.Checkoutid)
							if err != nil {
								fmt.Println("Error Occured")
							} else {
								fmt.Println("SUCCESS")
							}
						}
					}
				}
			}
		}
	}
}
