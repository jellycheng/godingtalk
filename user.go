package godingtalk

import (
	"fmt"
)

// 根据手机号获取企业账号用户的userId：https://open.dingtalk.com/document/orgapp/query-users-by-phone-number
func GetUseridByMobile(w GetUseridByMobileReqDto, at string) GetUseridByMobileRespDto {
	ret := GetUseridByMobileRespDto{}
	jsonStr := ToJson(w)
	urlStr := fmt.Sprintf(GetUseridByMobileUrl, DingOapiDomain, at)
	resp, _ := PostUrlContnet4json(urlStr, jsonStr, map[string]string{})
	_ = JsonUnmarshal(resp, &ret)
	return ret
}
