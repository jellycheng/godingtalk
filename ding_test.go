package godingtalk

import (
	"fmt"
	"github.com/jellycheng/gosupport"
	"github.com/jellycheng/gosupport/env"
	"testing"
)

// go test -run=TestGetAccessToken
func TestGetAccessToken(t *testing.T) {
	err := env.LoadEnv2DataManage("./example/.env")
	if err!=nil {
		fmt.Println(err.Error())
		return
	}
	globalEnv := gosupport.NewGlobalEnvSingleton()
	cfg := DingAgentConfig{
		AppKey:globalEnv.GetString("DingTalk_1_AppKey"),
		AppSecret:globalEnv.GetString("DingTalk_1_AppSecret"),
	}
	at := GetAccessToken(cfg)
	Debug(at)

}

