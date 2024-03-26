package todotask

import (
	"fmt"
	"github.com/jellycheng/godingtalk"
)

// 创建待办任务
func CreateTodoTask(w CreateTodoTaskReqDto, at string, unionId string, operatorId string) CreateTodoTaskRespDto {
	ret := CreateTodoTaskRespDto{}
	jsonStr := godingtalk.ToJson(w)
	operatorIdParam := ""
	if operatorId != "" {
		operatorIdParam = fmt.Sprintf("?operatorId=%s", operatorId)
	}
	urlStr := fmt.Sprintf(godingtalk.CreateTodoTaskUrl, godingtalk.DingApiDomain, unionId, operatorIdParam)
	resp, _ := godingtalk.PostUrlContnet4json(urlStr, jsonStr, map[string]string{
		"x-acs-dingtalk-access-token": at,
	})
	_ = godingtalk.JsonUnmarshal(resp, &ret)
	return ret
}

// 删除待办任务： https://open.dingtalk.com/document/orgapp/delete-dingtalk-to-do-tasks
func DeleteTodoTask(at string, unionId string, operatorId string, taskId string) DeleteTodoTaskRespDto {
	ret := DeleteTodoTaskRespDto{}
	operatorIdParam := ""
	if operatorId != "" {
		operatorIdParam = fmt.Sprintf("?operatorId=%s", operatorId)
	}
	urlStr := fmt.Sprintf(godingtalk.DeleteTodoTaskUrl, godingtalk.DingApiDomain, unionId, taskId, operatorIdParam)
	resp, _ := godingtalk.DeleteUrlContent(urlStr, "", map[string]string{
		"x-acs-dingtalk-access-token": at,
	})
	_ = godingtalk.JsonUnmarshal(resp, &ret)
	return ret
}

// 查询企业下用户待办列表
func GetTodoTask(w GetTodoTaskReqDto, at string, unionId string) GetTodoTaskRespDto {
	ret := GetTodoTaskRespDto{}
	resp := GetTodoTaskToStr(w, at, unionId)
	_ = godingtalk.JsonUnmarshal(resp, &ret)
	return ret
}

func GetTodoTaskToStr(w GetTodoTaskReqDto, at string, unionId string) string {
	jsonStr := godingtalk.ToJson(w)
	urlStr := fmt.Sprintf(godingtalk.GetTodoTaskUrl, godingtalk.DingApiDomain, unionId)
	resp, _ := godingtalk.PostUrlContnet4json(urlStr, jsonStr, map[string]string{
		"x-acs-dingtalk-access-token": at,
	})
	return resp
}

// 更新待办任务状态
func UpdateTodoTaskStatus(at string, unionId string, operatorId string, taskId string, done bool) UpdateTodoTaskStatusRespDto {
	ret := UpdateTodoTaskStatusRespDto{}
	operatorIdParam := ""
	if operatorId != "" {
		operatorIdParam = fmt.Sprintf("?operatorId=%s", operatorId)
	}
	urlStr := fmt.Sprintf(godingtalk.UpdateTodoTaskStatusUrl, godingtalk.DingApiDomain, unionId, taskId, operatorIdParam)
	reqDto := UpdateTodoTaskStatusReqDto{
		ExecutorStatusList: []ExecutorStatusList{
			{
				ID:     unionId,
				IsDone: done,
			},
		},
	}
	jsonStr := godingtalk.ToJson(reqDto)
	resp, _ := godingtalk.PutUrlContent(urlStr, jsonStr, map[string]string{
		"x-acs-dingtalk-access-token": at,
	})
	_ = godingtalk.JsonUnmarshal(resp, &ret)
	return ret
}

// 更新待办信息
func UpdateTodoTaskInfo(w UpdateTodoTaskInfoReqDto, at string, unionId string, operatorId string, taskId string) UpdateTodoTaskInfoRespDto {
	ret := UpdateTodoTaskInfoRespDto{}
	operatorIdParam := ""
	if operatorId != "" {
		operatorIdParam = fmt.Sprintf("?operatorId=%s", operatorId)
	}
	urlStr := fmt.Sprintf(godingtalk.UpdateTodoTaskInfoUrl, godingtalk.DingApiDomain, unionId, taskId, operatorIdParam)
	jsonStr := godingtalk.ToJson(w)
	resp, _ := godingtalk.PutUrlContent(urlStr, jsonStr, map[string]string{
		"x-acs-dingtalk-access-token": at,
	})
	_ = godingtalk.JsonUnmarshal(resp, &ret)
	return ret
}
