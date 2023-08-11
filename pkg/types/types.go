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

type BookInventory struct {
	Bookid          int    `json:"bookid"`
	Bookname        string `json:"bookname"`
	Author          string `json:"string"`
	TotalCopies     int    `json:"totcopies"`
	CopiesAvailable int    `json:"copiesAvailable"`
}

type Inventory struct {
	Books []BookInventory `json:"books"`
}

type ProfileAdminInfo struct {
	Username  string    `json:"username"`
	Inventory Inventory `json:"inventory"`
}

type ProfileInfo struct {
	Username    string        `json:"username"`
	PendingReqs CheckRequests `json:"pendingreqs"`
	CheckReqs   CheckRequests `json:"checkreq"`
	DeniedReqs  CheckRequests `json:"deniedreq"`
	Fine        int           `json:"fine"`
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
	Fine       int       `json:"fine"`
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
	NotClient   Response = "You need to be a client to access this"
	Success     Response = "Successful"
)
