package godingtalk

import (
	"fmt"
	"github.com/jellycheng/gosupport"
)

// 通讯录管理封装： 组织架构（部门、部门下员工）、角色

// 获取部门列表：只支持获取当前部门的下一级部门基础信息，不支持获取当前部门下所有层级子部门
func GetDepartmentList(accessToken string, deptId int, language string) DepartmentListRespDto {
	lang := "zh_CN"
	if language != "" {
		lang = language
	}
	reqDto := DepartmentReqDto{Language:lang}
	if deptId>0 {
		reqDto.DeptID = deptId
	}
	urlStr := fmt.Sprintf(DepartmentUrl, DingOapiDomain, accessToken)
	resp, _ := PostUrlContnet4json(urlStr, gosupport.ToJson(reqDto),
					map[string]string{"content-type": "application/json",})
	ret := DepartmentListRespDto{}
	_ = JsonUnmarshal(resp, &ret)
	return ret
}

func GetDepartmentListTree(accessToken string, deptId int, language string) ([]DepartmentSimpleTree, DingTalkApiErr3Dto) {
	apiErr := DingTalkApiErr3Dto{}
	ret := []DepartmentSimpleTree{}
	curDept := GetDepartmentList(accessToken, deptId, language)
	if curDept.Errcode == 0 {
		for _,v := range curDept.Result {
			tmp := DepartmentSimpleTree{
						AutoAddUser:v.AutoAddUser,
						CreateDeptGroup:v.CreateDeptGroup,
						DeptID:v.DeptID,
						Name:v.Name,
						ParentID:v.ParentID,
					}
			subDept,_ := GetDepartmentListTree(accessToken, v.DeptID, language)
			tmp.SubDeptList = subDept
			ret = append(ret, tmp)
		}
	} else {
		apiErr.Errcode = curDept.Errcode
		apiErr.Errmsg = curDept.Errmsg
	}

	return ret,apiErr
}


// 获取指定部门的userid列表
func GetDepartmentUseridList(accessToken string, deptId int) DeptUseridListRespDto {
	urlStr := fmt.Sprintf(DepartmentUseridListUrl, DingOapiDomain, accessToken)
	reqDto := struct {
		DeptId int `json:"dept_id"`
	}{DeptId:deptId}

	resp, _ := PostUrlContnet4json(urlStr, gosupport.ToJson(reqDto),
		map[string]string{"content-type": "application/json",})
	ret := DeptUseridListRespDto{}
	_ = JsonUnmarshal(resp, &ret)
	return ret
}


// 获取指定用户的详细信息
func GetUserInfo(accessToken string, userid string, lang string) UserProfileRespDto {
	urlStr := fmt.Sprintf(UserInfoUrl, DingOapiDomain, accessToken)
	if lang == "" {
		lang = "zh_CN"
	}
	reqDto := struct {
		Language string `json:"language"`
		Userid string `json:"userid"`
	}{
		Language:lang,
		Userid:userid,
	}
	resp, _ := PostUrlContnet4json(urlStr, gosupport.ToJson(reqDto),
		map[string]string{"content-type": "application/json",})

	ret := UserProfileRespDto{}
	_ = JsonUnmarshal(resp, &ret)
	return ret
}
