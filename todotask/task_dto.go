package todotask

import "github.com/jellycheng/godingtalk"

// 创建钉钉待办任务： https://open.dingtalk.com/document/orgapp/add-dingtalk-to-do-task
type CreateTodoTaskReqDto struct {
	SourceID           string         `json:"sourceId,omitempty"`       //业务系统侧的唯一标识ID，即业务ID
	Subject            string         `json:"subject"`                  //待办标题，最大长度1024
	CreatorID          string         `json:"creatorId,omitempty"`      //创建者的unionId
	Description        string         `json:"description,omitempty"`    // 待办备注描述，最大长度4096
	DueTime            int            `json:"dueTime,omitempty"`        //截止时间，Unix时间戳，单位毫秒
	ExecutorIds        []string       `json:"executorIds,omitempty"`    //执行者的unionId，最大数量1000
	ParticipantIds     []string       `json:"participantIds,omitempty"` //参与者的unionId，最大数量1000
	DetailURL          *DetailURL     `json:"detailUrl,omitempty"`      //详情页url跳转地址
	IsOnlyShowExecutor bool           `json:"isOnlyShowExecutor"`       // 生成的待办是否仅展示在执行者的待办列表中
	Priority           int            `json:"priority,omitempty"`       //优先级，取值：10：较低、20：普通、30：紧急、40：非常紧急
	NotifyConfigs      *NotifyConfigs `json:"notifyConfigs,omitempty"`  //待办通知配置
}
type DetailURL struct {
	AppURL string `json:"appUrl,omitempty"` //APP端详情页url跳转地址
	PcURL  string `json:"pcUrl,omitempty"`  //PC端详情页url跳转地址
}
type NotifyConfigs struct {
	DingNotify string `json:"dingNotify,omitempty"` //DING通知配置，目前仅支持取值为1，表示应用内DING
}

type CreateTodoTaskRespDto struct {
	godingtalk.DingTalkApiErrDto
	ID                 string        `json:"id"`                 //待办ID
	BizTag             string        `json:"bizTag"`             //接入应用标识
	Subject            string        `json:"subject"`            //待办的标题
	Description        string        `json:"description"`        //待办描述
	DetailURL          DetailURL     `json:"detailUrl"`          //详情页url跳转地址
	CreatorID          string        `json:"creatorId"`          //创建者的unionId
	ExecutorIds        []string      `json:"executorIds"`        //执行者的unionId
	ParticipantIds     []string      `json:"participantIds"`     //参与者的unionId
	ModifierID         string        `json:"modifierId"`         //更新者的unionId
	IsOnlyShowExecutor bool          `json:"isOnlyShowExecutor"` //生成的待办是否仅展示在执行者的待办列表中
	Priority           int           `json:"priority"`           //优先级
	RequestID          string        `json:"requestId"`          //请求ID
	Source             string        `json:"source"`             //业务来源
	SourceID           string        `json:"sourceId"`           // 业务系统侧的唯一标识ID，即业务ID
	CreatedTime        int64         `json:"createdTime"`        //创建时间，Unix时间戳，单位毫秒
	ModifiedTime       int64         `json:"modifiedTime"`       //更新时间，Unix时间戳，单位毫秒
	StartTime          int           `json:"startTime"`          //开始时间，Unix时间戳，单位毫秒
	DueTime            int           `json:"dueTime"`            //截止时间，Unix时间戳，单位毫秒
	FinishTime         int           `json:"finishTime"`         //完成时间，Unix时间戳，单位毫秒
	Done               bool          `json:"done"`               //完成状态
	NotifyConfigs      NotifyConfigs `json:"notifyConfigs"`      //待办通知配置
	TenantID           string        `json:"tenantId"`
	TenantType         string        `json:"tenantType"`
}

type DeleteTodoTaskRespDto struct {
	godingtalk.DingTalkApiErrDto
	RequestID string `json:"requestId"`
	Result    bool   `json:"result"` // 删除结果。true表示删除成功
}

// 查询待办任务列表入参： https://open.dingtalk.com/document/orgapp/query-the-to-do-list-of-enterprise-users
type GetTodoTaskReqDto struct {
	NextToken *string `json:"nextToken,omitempty"`
	IsDone    *bool   `json:"isDone,omitempty"`
}

// 查看待办任务列表出参
type GetTodoTaskRespDto struct {
	RequestID  string      `json:"requestId"`
	TodoCards  []TodoCards `json:"todoCards"`  //待办卡片列表
	TotalCount int         `json:"totalCount"` // 总记录数
	NextToken  string      `json:"nextToken"`  //下一次请求的token
}

// 组织信息
type OrgInfo struct {
	CorpID string `json:"corpId"` //企业id
	Name   string `json:"name"`   // 企业名称
}
type OriginalSource struct {
	SourceTitle string `json:"sourceTitle"`
}
type TodoCardView struct {
	ActionType          string        `json:"actionType"`
	CardType            string        `json:"cardType"`
	CircleELType        string        `json:"circleELType"`
	ContentType         string        `json:"contentType"`
	Icon                string        `json:"icon"`
	TodoCardContentList []interface{} `json:"todoCardContentList"`
	TodoCardTitle       string        `json:"todoCardTitle"` // 待办标题
}
type TodoCards struct {
	TaskID         string         `json:"taskId"`  //待办ID，任务ID
	Subject        string         `json:"subject"` //待办标题
	BizTag         string         `json:"bizTag"`
	Category       string         `json:"category"`
	CreatedTime    int64          `json:"createdTime"`  //创建时间
	ModifiedTime   int64          `json:"modifiedTime"` //更新时间
	CreatorID      string         `json:"creatorId"`    //创建者ID
	DetailURL      DetailURL      `json:"detailUrl"`    //详情页链接
	IsDone         bool           `json:"isDone"`       //待办完成状态
	OrgInfo        OrgInfo        `json:"orgInfo"`      // 组织信息
	OriginalSource OriginalSource `json:"originalSource"`
	Priority       int            `json:"priority"` //优先级
	TodoCardView   TodoCardView   `json:"todoCardView"`
	TodoStatus     string         `json:"todoStatus"`
	SourceID       string         `json:"sourceId"`
}

// 更新待办任务状态: https://open.dingtalk.com/document/orgapp/update-dingtalk-to-do-status
type UpdateTodoTaskStatusReqDto struct {
	ExecutorStatusList []ExecutorStatusList `json:"executorStatusList"`
}
type ExecutorStatusList struct {
	ID     string `json:"id"`     // 执行者的unionId
	IsDone bool   `json:"isDone"` // 更新结果。true表示成功
}

type UpdateTodoTaskStatusRespDto struct {
	godingtalk.DingTalkApiErrDto
	RequestID string `json:"requestId"`
	Result    bool   `json:"result"`
}

// 更新待办任务信息： https://open.dingtalk.com/document/orgapp/updates-dingtalk-to-do-tasks
type UpdateTodoTaskInfoReqDto struct {
	Subject        string   `json:"subject,omitempty"`        //待办标题，最大长度1024
	Description    string   `json:"description,omitempty"`    //待办描述，最大长度4096
	DueTime        int      `json:"dueTime,omitempty"`        //截止时间，Unix时间戳，单位毫秒
	Done           bool     `json:"done,omitempty"`           //完成状态,true完成，false未完成
	ExecutorIds    []string `json:"executorIds,omitempty"`    //执行者的unionId列表，最大数量1000
	ParticipantIds []string `json:"participantIds,omitempty"` //参与者的unionId列表，最大数量1000
}

type UpdateTodoTaskInfoRespDto struct {
	godingtalk.DingTalkApiErrDto
	RequestID string `json:"requestId"`
	Result    bool   `json:"result"`
}
