package types

import "time"

type SignupData struct {
	Fullname  string `json:"fullname"`
	Username  string `json:"username"`
	Password  string `json:"pswd"`
	CPassword string `json:"cpswd"`
	IsAdmin   bool   `json:"isAdmin"`
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"pswd"`
}

type ProfileInfo struct {
	Username   string        `json:"username"`
	CheckReqs  CheckRequests `json:"checkreq"`
	DeniedReqs CheckRequests `json:"deniedreq"`
}

type Book struct {
	Bookid   int    `json:"bookid"`
	Bookname string `json:"bookname"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

type Books struct {
	Books []Book `json:"books"`
}

type SasData struct {
	Bookid   int    `json:"bookid"`
	Option   string `json:"options"`
	Quantity int    `json:"quanity"`
}

type IssueBookData struct {
	Bookid   int    `json:"bookid"`
	Username string `json:"username"`
}

type CheckRequest struct {
	Checkoutid int       `json:"checkoutid"`
	OfBook     int       `json:"bookid"`
	ByUser     string    `json:"username"`
	Status     string    `json:"status"`
	IssueTime  time.Time `json:"issueTime"`
}

type CheckRequests struct {
	Reqs []CheckRequest `json:"reqs"`
}

type ActionData struct {
	CheckReq CheckRequest `json:"checkrequest"`
	Action   string       `json:"action"`
}

type User string

const (
	Unverified User = "Unverified"
	Client     User = "Client"
	Admin      User = "Admin"
)

type Response string

const ( //Some standard error messages
	NotAdmin    Response = "You need to be an admin to access this"
	NotLoggedIn Response = "You need to be logged in to access this"
	Success     Response = "Successful"
)
