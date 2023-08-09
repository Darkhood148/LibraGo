package models

import (
	"mvc/pkg/types"
	"time"
)

func FetchRequests() (types.CheckRequests, error) {
	db, err := Connection()
	if err != nil {
		return types.CheckRequests{}, err
	}
	defer db.Close()
	selectSql := "SELECT * FROM checkouts WHERE status = \"checkinPending\" OR status = \"pending\""
	rows, err := db.Query(selectSql)

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

	var reqs types.CheckRequests
	reqs.Reqs = fetchReqs
	return reqs, nil

}
