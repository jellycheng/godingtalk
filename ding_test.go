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
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	globalEnv := gosupport.NewGlobalEnvSingleton()
	cfg := DingAgentConfig{
		AppKey:    globalEnv.GetString("DingTalk_1_AppKey"),
		AppSecret: globalEnv.GetString("DingTalk_1_AppSecret"),
	}
	at := GetAccessToken(cfg)
	Debug(at)

}

// go test -run=GetDepartmentListTree
func TestGetDepartmentListTree(t *testing.T) {
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
	at := GetAccessToken(cfg)
	atVal := at.AccessToken
	deptId := 1 // 部门ID
	if dtl, err := GetDepartmentListTree(atVal, deptId, "zh_CN"); err.Errcode == 0 {
		for _, v := range dtl {
			fmt.Println(fmt.Sprintf("部门ID：%d，部门：%s，上级部门ID：%d,子部门信息：%+v", v.DeptID, v.Name, v.ParentID, v.SubDeptList))
		}
	} else {
		fmt.Println(err.Errmsg)
	}

}

// go test -run=TestGetDepartmentUseridList
func TestGetDepartmentUseridList(t *testing.T) {
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
	at := GetAccessToken(cfg)
	atVal := at.AccessToken
	deptId := 1 // 部门ID
	// 获取部门的userid列表: https://open-dev.dingtalk.com/apiExplorer#/?devType=org&api=dingtalk.oapi.user.listid
	ret := GetDepartmentUseridList(atVal, deptId)
	fmt.Println(ret.Result.UseridList)
}
