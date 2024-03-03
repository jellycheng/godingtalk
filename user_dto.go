package godingtalk

type GetUseridByMobileReqDto struct {
	Mobile                        string `json:"mobile"`                           //手机号
	SupportExclusiveAccountSearch bool   `json:"support_exclusive_account_search"` //false、true
}

type GetUseridByMobileRespDto struct {
	DingTalkApiErr2Dto
	Result    UseridByMobileResultDto `json:"result"`
	RequestID string                  `json:"request_id"`
}

type UseridByMobileResultDto struct {
	ExclusiveAccountUseridList []interface{} `json:"exclusive_account_userid_list"`
	Userid                     string        `json:"userid"`
}
