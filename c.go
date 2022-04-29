package godingtalk

const (
	ScopeOpenid = "openid"
	ScopeOpenIdCropid = "openid corpid"
	LoginUrl = "%s/oauth2/auth?redirect_uri=%s&response_type=code&client_id=%s&scope=%s&state=%s&prompt=consent"
	UserAccessTokenUrl = "%s/v1.0/oauth2/userAccessToken"
	UserInfo4UserAccessTokenUrl = "%s/v1.0/contact/users/me"

)
