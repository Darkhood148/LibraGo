package models

import (
	"mvc/pkg/types"
)

func FetchUserReqs(data string) (types.CheckRequests, error) {
	db, err := Connection()
	if err != nil {
		return types.CheckRequests{}, err
	}
	selectSql := "SELECT * FROM checkouts WHERE byUser = (?) AND status = \"issued\""
	rows, err := db.Query(selectSql, data)
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

	var listReqs types.CheckRequests
	listReqs.Reqs = fetchReqs
	return listReqs, nil

}
