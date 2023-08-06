package models

import (
	"fmt"

	"mvc/pkg/types"
)

func FetchDeniedReqs(data string) types.CheckRequests {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}
	selectSql := "SELECT * FROM checkouts WHERE byUser = (?) AND status = \"checkinDenied\""
	rows, err := db.Query(selectSql, data)
	db.Close()

	if err != nil {
		fmt.Printf("error %s querying the database", err)
	}

	var fetchReqs []types.CheckRequest
	for rows.Next() {
		var req types.CheckRequest
		err := rows.Scan(&req.Checkoutid, &req.OfBook, &req.ByUser, &req.Status)
		if err != nil {
			fmt.Printf("error %s scanning the row", err)
		}
		fetchReqs = append(fetchReqs, req)
	}

	var listReqs types.CheckRequests
	listReqs.Reqs = fetchReqs
	return listReqs

}
