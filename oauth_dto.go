package godingtalk

type UserAccessTokenReqDto struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Code         string `json:"code,omitempty"`
	GrantType    string `json:"grantType,omitempty"` // authorization_code 、refresh_token
	RefreshToken string `json:"refreshToken,omitempty"`
}

type UserAccessTokenRespDto struct {
	DingTalkApiErrDto
	ExpireIn     int    `json:"expireIn"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	CorpId       string `json:"corpId"` //所选企业corpId
}


type UserInfoRespDto struct {
	DingTalkApiErrDto
	Nick      string `json:"nick"` // 昵称
	UnionID   string `json:"unionId"`
	OpenID    string `json:"openId"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatarUrl"` // 头像地址
	Mobile    string `json:"mobile"` // 手机号
	StateCode string `json:"stateCode"` // 手机号对应的国家号
}
