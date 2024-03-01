package godingtalk

import (
	"fmt"
	"github.com/jellycheng/gosupport"
	"net/url"
)

// GetLoginUrl 第1步： 拼接钉钉登录地址: https://open.dingtalk.com/document/isvapp-server/obtain-identity-credentials
func GetLoginUrl(redirectUri, clientId, scope, state string) string {
	if scope == "" {
		scope = ScopeOpenid
	}
	sUrl := fmt.Sprintf(LoginUrl, DingLoginDomain, url.QueryEscape(redirectUri), clientId, url.QueryEscape(scope), url.QueryEscape(state))
	return sUrl
}

// 第2步： 通过用户授权成功后的authCode获取访问令牌
func GetUserAccessToken4authCode(authCode string, dingCfg DingAgentConfig) UserAccessTokenRespDto {
	urlStr := fmt.Sprintf(UserAccessTokenUrl, DingApiDomain)
	reqDto := UserAccessTokenReqDto{
		ClientID:     dingCfg.AppKey,
		ClientSecret: dingCfg.AppSecret,
		Code:         authCode,
		GrantType:    "authorization_code",
	}
	jsonStr := gosupport.ToJson(reqDto)
	resp, _ := PostUrlContnet4json(urlStr, jsonStr, map[string]string{})
	ret := UserAccessTokenRespDto{}
	_ = JsonUnmarshal(resp, &ret)
	return ret
}

// 第3步： 获取用户信息：https://open.dingtalk.com/document/isvapp-server/dingtalk-retrieve-user-information
func GetUserInfo4UserAccessToken(userAccessToken string) UserInfoRespDto {
	urlStr := fmt.Sprintf(UserInfo4UserAccessTokenUrl, DingApiDomain)
	resp, _ := GetUrlContent(urlStr, map[string]string{
		"x-acs-dingtalk-access-token": userAccessToken,
		"content-type":                "application/json",
	},
	)
	ret := UserInfoRespDto{}
	_ = JsonUnmarshal(resp, &ret)
	return ret
}

// 刷新访问令牌
func RefreshToken(refreshToken string, dingCfg DingAgentConfig) UserAccessTokenRespDto {
	urlStr := fmt.Sprintf(UserAccessTokenUrl, DingApiDomain)
	reqDto := UserAccessTokenReqDto{
		ClientID:     dingCfg.AppKey,
		ClientSecret: dingCfg.AppSecret,
		GrantType:    "refresh_token",
		RefreshToken: refreshToken,
	}
	jsonStr := gosupport.ToJson(reqDto)
	resp, _ := PostUrlContnet4json(urlStr, jsonStr, map[string]string{})
	ret := UserAccessTokenRespDto{}
	_ = JsonUnmarshal(resp, &ret)
	return ret
}
