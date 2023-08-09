package models

import (
	"mvc/pkg/types"
)

func FetchRequests() (types.CheckRequests, error) {
	db, err := Connection()
	if err != nil {
		return types.CheckRequests{}, err
	}
	defer db.Close()
	selectSql := "SELECT * FROM checkouts WHERE status = \"checkinPending\" OR status = \"pending\""
	rows, err := db.Query(selectSql)
	db.Close()

	if err != nil {
		return types.CheckRequests{}, err
	}

	var fetchReqs []types.CheckRequest
	for rows.Next() {
		var req types.CheckRequest
		err := rows.Scan(&req.Checkoutid, &req.OfBook, &req.ByUser, &req.Status)
		if err != nil {
			return types.CheckRequests{}, err
		}
		fetchReqs = append(fetchReqs, req)
	}

	var reqs types.CheckRequests
	reqs.Reqs = fetchReqs
	return reqs, nil

}
