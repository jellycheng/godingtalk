package godingtalk

type DingTalkApiErrDto struct {
	Requestid string `json:"requestid"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}

type DingTalkApiErr2Dto struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type RequestIdDto struct {
	RequestId string `json:"request_id"`
}

type AccessTokenRespDto struct {
	DingTalkApiErr2Dto
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type DingTalkApiErr3Dto struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"request_id"`
}
