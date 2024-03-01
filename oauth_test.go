package godingtalk

import (
	"fmt"
	"github.com/jellycheng/gosupport"
	"github.com/jellycheng/gosupport/env"
	"testing"
)

// 第1步：go test -run=TestGetLoginUrl
func TestGetLoginUrl(t *testing.T) {
	err := env.LoadEnv2DataManage("./example/.env")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	globalEnv := gosupport.NewGlobalEnvSingleton()
	cfg := DingAgentConfig{
		AppKey:    globalEnv.GetString("DingTalk_1_AppKey"),
		AppSecret: globalEnv.GetString("DingTalk_1_AppSecret"),
	}
	redirectUri := globalEnv.GetString("DingTalk_1_RedirectUri")
	clientId := cfg.AppKey
	scope := ScopeOpenIdCropid
	state := ""
	res := GetLoginUrl(redirectUri, clientId, scope, state)
	Debug(res)
}

// 第2步：go test -run=TestGetUserAccessToken4authCode
func TestGetUserAccessToken4authCode(t *testing.T) {
	err := env.LoadEnv2DataManage("./example/.env")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	globalEnv := gosupport.NewGlobalEnvSingleton()
	authcode := globalEnv.GetString("AuthCodeDemo") // 这个其实是从url中获取的authCode参数值
	cfg := DingAgentConfig{
		AppKey:    globalEnv.GetString("DingTalk_1_AppKey"),
		AppSecret: globalEnv.GetString("DingTalk_1_AppSecret"),
	}
	res := GetUserAccessToken4authCode(authcode, cfg)
	Debug(gosupport.ToJson(res))

}

// 第3步： go test -run=TestGetUserInfo4UserAccessToken
func TestGetUserInfo4UserAccessToken(t *testing.T) {
	err := env.LoadEnv2DataManage("./example/.env")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	globalEnv := gosupport.NewGlobalEnvSingleton()
	userAccessToken := globalEnv.GetString("userAccessTokenDemo")

	res := GetUserInfo4UserAccessToken(userAccessToken)
	Debug(gosupport.ToJson(res))
}

// 刷新访问令牌： go test -run=TestRefreshToken
func TestRefreshToken(t *testing.T) {
	err := env.LoadEnv2DataManage("./example/.env")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	globalEnv := gosupport.NewGlobalEnvSingleton()
	cfg := DingAgentConfig{
		AppKey:    globalEnv.GetString("DingTalk_1_AppKey"),
		AppSecret: globalEnv.GetString("DingTalk_1_AppSecret"),
	}

	rtoken := globalEnv.GetString("refreshTokenDemo")

	res := RefreshToken(rtoken, cfg)
	Debug(gosupport.ToJson(res))
}
