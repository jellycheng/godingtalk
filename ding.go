package godingtalk

import "fmt"

// 企业内部应用获取access_token
func GetAccessToken(cfg DingAgentConfig) AccessTokenRespDto {
	urlStr := fmt.Sprintf(GetAccessTokenUrl, DingOapiDomain, cfg.AppKey, cfg.AppSecret)
	resp, _ := GetUrlContent(urlStr, map[string]string{
		"content-type": "application/json",
	},
	)

	ret := AccessTokenRespDto{}
	_ = JsonUnmarshal(resp, &ret)
	return ret
}
