package controller

import (
	"mvc/pkg/middleware"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == types.Unverified {
		renderSignupPage(w, "0", "")
	} else {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}
}

func SignupPost(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == types.Unverified {
		if r.FormValue("password") == r.FormValue("confirmPassword") {
			data := types.SignupData{
				Fullname: r.FormValue("fullName"),
				Username: r.FormValue("username"),
				Password: r.FormValue("password"),
				IsAdmin:  types.Client,
			}
			err := models.SignUp(data)
			if err != nil {
				renderSignupPage(w, "2", err.Error())
			} else {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
			}
		} else {
			renderSignupPage(w, "2", "Password and Confirm-Password do not match")
		}
	} else {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}
}

func renderSignupPage(w http.ResponseWriter, status string, errMess string) {
	t := views.RenderPage("signup.html")
	info := types.ErrorInfo{
		Status:     status,
		ErrMessage: errMess,
	}
	t.Execute(w, info)
}
