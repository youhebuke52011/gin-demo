package common

type PublicArg struct {
	AppID int `json:"appid" form:"appid" binding:"required,appId"`
}
