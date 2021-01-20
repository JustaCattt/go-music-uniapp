package model

type UserInfo struct {
	Username	string	`json:"username"`
	Telephone	string	`json:"telephone"`
}

//User转UserInfo
func ToUserInfo(user User) UserInfo {
	return UserInfo{
		Username: user.Username,
		Telephone: user.Telephone,
	}
}
