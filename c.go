package godingtalk

const (
	ScopeOpenid                 = "openid"
	ScopeOpenIdCropid           = "openid corpid"
	LoginUrl                    = "%s/oauth2/auth?redirect_uri=%s&response_type=code&client_id=%s&scope=%s&state=%s&prompt=consent"
	UserAccessTokenUrl          = "%s/v1.0/oauth2/userAccessToken"
	UserInfo4UserAccessTokenUrl = "%s/v1.0/contact/users/me"
	// 企业内部应用获取access_token
	GetAccessTokenUrl = "%s/gettoken?appkey=%s&appsecret=%s"
	// 部门列表
	DepartmentUrl = "%s/topapi/v2/department/listsub?access_token=%s"
	// 获取指定部门的userid列表
	DepartmentUseridListUrl = "%s/topapi/user/listid?access_token=%s"
	// 获取指定用户的详细信息
	UserInfoUrl = "%s/topapi/v2/user/get?access_token=%s"

	// 发送工作通知
	WorkNotifyMsgUrl = "%s/topapi/message/corpconversation/asyncsend_v2?access_token=%s"
	// 撤回工作消息
	RecallWorkNotifyMsgUrl = "%s/topapi/message/corpconversation/recall?access_token=%s"

	// 根据手机号获取企业账号用户的userId
	GetUseridByMobileUrl = "%s/topapi/v2/user/getbymobile?access_token=%s"

	// 创建钉钉待办任务 ?operatorId=
	CreateTodoTaskUrl = "%s/v1.0/todo/users/%s/tasks%s"
	// 删除钉钉待办任务 ?operatorId=
	DeleteTodoTaskUrl = "%s/v1.0/todo/users/%s/tasks/%s%s"
	// 查询企业下用户待办列表
	GetTodoTaskUrl = "%s/v1.0/todo/users/%s/org/tasks/query"
	// 更新待办任务状态 /v1.0/todo/users/{unionId}/tasks/{taskId}/executorStatus?operatorId=
	UpdateTodoTaskStatusUrl = "%s/v1.0/todo/users/%s/tasks/%s/executorStatus%s"
	// 更新待办任务信息 /v1.0/todo/users/{unionId}/tasks/{taskId}?operatorId=
	UpdateTodoTaskInfoUrl = "%s/v1.0/todo/users/%s/tasks/%s%s"
)
