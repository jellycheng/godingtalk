package godingtalk

type DingTalkApiErrDto struct {
	Requestid string `json:"requestid"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}
