package models

import (
	"fmt"

	"mvc/pkg/types"
)

func FetchRequests() types.CheckRequests {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}
	selectSql := "SELECT * FROM checkouts WHERE status = \"checkinPending\" OR status = \"pending\""
	rows, err := db.Query(selectSql)
	db.Close()

	if err != nil {
		fmt.Printf("error %s querying the database", err)
	}

	var fetchReqs []types.CheckRequest
	for rows.Next() {
		fmt.Println("Works")
		var req types.CheckRequest
		err := rows.Scan(&req.Checkoutid, &req.OfBook, &req.ByUser, &req.Status)
		if err != nil {
			fmt.Printf("error %s scanning the row", err)
		}
		fetchReqs = append(fetchReqs, req)
	}

	var reqs types.CheckRequests
	reqs.Reqs = fetchReqs
	fmt.Println(reqs)
	return reqs

}
