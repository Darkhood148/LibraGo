package controller

import (
	"mvc/pkg/middleware"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
)

func GetAdminAccess(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == types.Client {
		renderAdReqPage(w, "0", "")
	} else {
		renderInvalidPage(w, string(types.NotClient))
	}
}

func GetAdminAccessPost(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == types.Client {
		pswd := r.FormValue("pswd")
		uname := middleware.VerifyJWT(w, r)

		res, err := models.RequestAdminAccess(uname, pswd)
		if err != nil {
			renderAdReqPage(w, "2", err.Error())
		} else {
			if res {
				http.Redirect(w, r, "/profile", http.StatusSeeOther)
			} else {
				renderAdReqPage(w, "2", "Incorrect Password")
			}
		}

	} else {
		renderInvalidPage(w, string(types.NotClient))
	}
}

func renderAdReqPage(w http.ResponseWriter, status string, errMess string) {
	t := views.RenderPage("requestAdmin.html")
	info := types.ErrorInfo{
		Status:     status,
		ErrMessage: errMess,
	}
	t.Execute(w, info)
}
