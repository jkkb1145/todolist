package model

type User struct {
	UserID   int
	UserName string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
}

type Userregistermodel struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	Confirmpassword string `json:"confirmpassword"`
	Nickname        string `json:"nickname"`
}

type UserLogIn struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type ModelForCP struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ModelForCI struct {
	NickName string `json:"nickname"`
}
