package models

type X struct {
	Text string `json:"text"`
}

type LoginRequest struct {
	Username string `json:"username" form:"username" xml:"username" query:"username"`
	Password string `json:"password" form:"username" xml:"password" query:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
