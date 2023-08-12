package controller

import (
	"mvc/pkg/middleware"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
	"strconv"
)

func IssueBook(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == types.Client {
		renderIssuePage(w, "0", "")
	} else {
		renderInvalidPage(w, string(types.NotClient))
	}
}

func IssueBookPost(w http.ResponseWriter, r *http.Request) {
	if middleware.TypeOfUser(w, r) == types.Client {
		val, err := strconv.Atoi(r.FormValue("bookid"))
		if err != nil {
			renderIssuePage(w, "2", err.Error())
		} else {
			data := types.IssueBookData{
				Bookid:   val,
				Username: middleware.VerifyJWT(w, r),
			}
			err := models.IssueBook(data)
			if err != nil {
				renderIssuePage(w, "2", err.Error())
			} else {
				renderIssuePage(w, "1", "")
			}
		}
	} else {
		renderInvalidPage(w, string(types.NotClient))
	}
}

func renderIssuePage(w http.ResponseWriter, status string, errMess string) {
	t := views.RenderPage("issueBook.html")
	info := types.ErrorInfo{
		Status:     status,
		ErrMessage: errMess,
	}
	t.Execute(w, info)
}
