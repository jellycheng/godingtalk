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
