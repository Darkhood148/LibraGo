package models

import (
	"mvc/pkg/types"
	"time"
)

func FetchUserReqs(data string) (types.CheckRequests, error) {
	db, err := Connection()
	if err != nil {
		return types.CheckRequests{}, err
	}
	defer db.Close()
	selectSql := "SELECT * FROM checkouts WHERE byUser = (?) AND status = \"issued\""
	rows, err := db.Query(selectSql, data)
	db.Close()

	if err != nil {
		return types.CheckRequests{}, err
	}

	var fetchReqs []types.CheckRequest
	for rows.Next() {
		var req types.CheckRequest
		var temp string
		err := rows.Scan(&req.Checkoutid, &req.OfBook, &req.ByUser, &req.Status, &temp)
		if err != nil {
			return types.CheckRequests{}, err
		}
		req.IssueTime, err = time.Parse("2006-01-02 15:04:05", temp)
		if err != nil {
			return types.CheckRequests{}, err
		}
		fetchReqs = append(fetchReqs, req)
	}

	var listReqs types.CheckRequests
	listReqs.Reqs = fetchReqs
	return listReqs, nil

}
