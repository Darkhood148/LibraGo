package controller

import (
	"fmt"
	"mvc/pkg/middleware"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	uname := middleware.VerifyJWT(w, r)
	if uname != "" {
		t := views.ProfilePage(middleware.VerifyAdmin(uname))
		info := types.ProfileInfo{
			Username:   uname,
			CheckReqs:  models.FetchUserReqs(uname),
			DeniedReqs: models.FetchDeniedReqs(uname),
		}
		fmt.Println("info", info)
		t.Execute(w, info)
	} else {
		w.Write([]byte("Login Please"))
	}
}
