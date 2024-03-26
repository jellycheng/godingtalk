package message

import (
	"fmt"
	"github.com/jellycheng/godingtalk"
)

// 发送工作通知
func SendWorkMsg(w SendWorkNotifyReqDto, at string) WorkNotifyRespDto {
	ret := WorkNotifyRespDto{}
	jsonStr := godingtalk.ToJson(w)
	urlStr := fmt.Sprintf(godingtalk.WorkNotifyMsgUrl, godingtalk.DingOapiDomain, at)
	resp, _ := godingtalk.PostUrlContnet4json(urlStr, jsonStr, map[string]string{})
	_ = godingtalk.JsonUnmarshal(resp, &ret)
	return ret
}

// 撤回工作消息
func RecallWorkMsg(w RecallWorkNotifyReqDto, at string) RecallWorkNotifyRespDto {
	ret := RecallWorkNotifyRespDto{}
	jsonStr := godingtalk.ToJson(w)
	urlStr := fmt.Sprintf(godingtalk.RecallWorkNotifyMsgUrl, godingtalk.DingOapiDomain, at)
	resp, _ := godingtalk.PostUrlContnet4json(urlStr, jsonStr, map[string]string{})
	_ = godingtalk.JsonUnmarshal(resp, &ret)
	return ret
}
