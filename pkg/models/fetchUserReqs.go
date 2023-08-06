package models

import (
	"fmt"

	"mvc/pkg/types"
)

func FetchUserReqs(data string) types.CheckRequests {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}
	selectSql := "SELECT * FROM checkouts WHERE byUser = (?) AND status = \"issued\""
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
		fmt.Println("req", req)
		fetchReqs = append(fetchReqs, req)
	}

	fmt.Println("fetchReqs", fetchReqs)
	var listReqs types.CheckRequests
	listReqs.Reqs = fetchReqs
	fmt.Println("listReqs", listReqs)
	return listReqs

}
