package types

type SignupData struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"pswd"`
	CPassword string `json:"cpswd"`
	IsAdmin bool `json:"isAdmin"`
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"pswd"`
}

type ProfileInfo struct {
	Username string `json:"username"`
}