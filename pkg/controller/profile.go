package controller

import (
	"mvc/pkg/middleware"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) != types.Unverified {
		uname := middleware.VerifyJWT(w, r)
		denreq, err1 := models.FetchDeniedReqs(uname)
		usereq, err2 := models.FetchUserReqs(uname)
		if err1 != nil || err2 != nil {
			w.Write([]byte("Error Occured occured while fetching requests"))
		} else {
			t := views.ProfilePage(middleware.VerifyAdmin(uname))
			info := types.ProfileInfo{
				Username:   uname,
				CheckReqs:  usereq,
				DeniedReqs: denreq,
			}
			t.Execute(w, info)
		}
	} else {
		w.Write([]byte(types.NotLoggedIn))
	}
}
