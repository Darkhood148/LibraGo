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
		username := middleware.VerifyJWT(w, r)
		deniedRequests, err1 := models.FetchDeniedReqs(username)
		issuedBooks, err2 := models.FetchUserReqs(username)
		pendingRequests, err3 := models.FetchPendingReqs(username)
		if err1 != nil || err2 != nil || err3 != nil {
			renderErrorPage(w, "Error Occured occured while fetching requests")
		} else {
			t := views.RenderPage("profile.html")
			info := types.ProfileInfo{
				Username:    username,
				CheckReqs:   issuedBooks,
				DeniedReqs:  deniedRequests,
				PendingReqs: pendingRequests,
			}
			t.Execute(w, info)
		}
	} else if middleware.TypeOfUser(w, r) == types.Admin {
		username := middleware.VerifyJWT(w, r)
		inventory, err := models.FetchInventory()
		if err != nil {
			renderErrorPage(w, err.Error())
		} else {
			t := views.RenderPage("profileAdmin.html")
			info := types.ProfileAdminInfo{
				Username:  username,
				Inventory: inventory,
			}
			t.Execute(w, info)
		}
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
