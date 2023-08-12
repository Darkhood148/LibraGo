package types

import "time"

type SignupData struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"pswd"`
	IsAdmin  User   `json:"isAdmin"`
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"pswd"`
}

type BookInventory struct {
	Book        Book `json:"book_info"`
	TotalCopies int  `json:"total_copies"`
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
	PendingReqs CheckRequests `json:"pending_requests"`
	CheckReqs   CheckRequests `json:"check_requests"`
	DeniedReqs  CheckRequests `json:"denied_requests"`
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
	IssueTime  time.Time `json:"issue_time"`
	Fine       int       `json:"fine"`
}

type CheckRequests struct {
	Reqs []CheckRequest `json:"requests"`
}

type ActionData struct {
	CheckReq CheckRequest `json:"check_request"`
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

type ErrorInfo struct {
	Status     string `json:"status"`
	ErrMessage string `json:"errorMessage"`
}
