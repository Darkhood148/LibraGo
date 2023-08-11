package controller

import (
	"mvc/pkg/middleware"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == types.Client {
		uname := middleware.VerifyJWT(w, r)
		denreq, err1 := models.FetchDeniedReqs(uname)
		usereq, err2 := models.FetchUserReqs(uname)
		penreq, err3 := models.FetchPendingReqs(uname)
		if err1 != nil || err2 != nil || err3 != nil {
			w.Write([]byte("Error Occured occured while fetching requests"))
		} else {
			t := views.ProfilePage(false)
			info := types.ProfileInfo{
				Username:    uname,
				CheckReqs:   usereq,
				DeniedReqs:  denreq,
				PendingReqs: penreq,
			}
			t.Execute(w, info)
		}
	} else if middleware.TypeOfUser(w, r) == types.Admin {
		uname := middleware.VerifyJWT(w, r)
		inventory, err := models.FetchInventory()
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			t := views.ProfilePage(true)
			info := types.ProfileAdminInfo{
				Username:  uname,
				Inventory: inventory,
			}
			t.Execute(w, info)
		}
	} else {
		w.Write([]byte(types.NotLoggedIn))
	}
}
