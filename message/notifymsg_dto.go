package message

import "github.com/jellycheng/godingtalk"

// 发送工作消息通知： https://open.dingtalk.com/document/orgapp-server/asynchronous-sending-of-enterprise-session-messages
// https://open.dingtalk.com/document/orgapp/asynchronous-sending-of-enterprise-session-messages
type SendWorkNotifyReqDto struct {
	Msg        WorkNotifyMsgDto `json:"msg"`                    //必须，消息内容，最长不超过2048个字节
	AgentID    string           `json:"agent_id"`               //必须，发送消息时使用的微应用的AgentID
	ToAllUser  bool             `json:"to_all_user,omitempty"`  //可选，是否发送给企业全部用户，默认false，当设置为false时必须指定userid_list或dept_id_list其中一个参数的值
	DeptIDList string           `json:"dept_id_list,omitempty"` //可选，接收者的部门id列表，最大列表长度20，接收者是部门ID时，包括子部门下的所有用户
	UseridList string           `json:"userid_list,omitempty"`  //可选，接收者的userid列表，最大用户列表长度100
}

// 消息通知类型： https://open.dingtalk.com/document/orgapp-server/message-types-and-data-format
// https://open.dingtalk.com/document/orgapp/message-types-and-data-format
type WorkNotifyMsgDto struct {
	Msgtype    string            `json:"msgtype"`               // 必须，消息类型，文本text、图片image、语音voice、文件file、链接link、oa消息、markdown、action_card卡片消息
	Voice      *VoiceMsgDto      `json:"voice,omitempty"`       //可选，
	Image      *ImageMsgDto      `json:"image,omitempty"`       //可选，
	Oa         *OaMsgDto         `json:"oa,omitempty"`          //可选，
	File       *FileMsgDto       `json:"file,omitempty"`        //可选，
	ActionCard *ActionCardMsgDto `json:"action_card,omitempty"` //可选，
	Link       *LinkMsgDto       `json:"link,omitempty"`        //可选，
	Markdown   *MarkdownMsgDto   `json:"markdown,omitempty"`    //可选，
	Text       *TextMsgDto       `json:"text,omitempty"`        //可选，文本消息
}

// 语音消息
type VoiceMsgDto struct {
	Duration string `json:"duration"`
	MediaID  string `json:"media_id"`
}

// 图片消息
type ImageMsgDto struct {
	MediaID string `json:"media_id"`
}

type OaMsgHeadDto struct {
	Bgcolor string `json:"bgcolor"` //消息头部的背景颜色
	Text    string `json:"text"`    //消息的头部标题
}
type OaMsgStatusBarDto struct {
	StatusValue string `json:"status_value,omitempty"` //状态栏文案
	StatusBg    string `json:"status_bg,omitempty"`    //状态栏背景色
}
type OaMsgBodyFormDto struct {
	Value string `json:"value,omitempty"`
	Key   string `json:"key,omitempty"`
}
type OaMsgBodyRichDto struct {
	Unit string `json:"unit,omitempty"`
	Num  string `json:"num,omitempty"`
}
type OaMsgBodyDto struct {
	FileCount string            `json:"file_count,omitempty"`
	Image     string            `json:"image,omitempty"`
	Form      *OaMsgBodyFormDto `json:"form,omitempty"`
	Author    string            `json:"author,omitempty"`
	Rich      *OaMsgBodyRichDto `json:"rich,omitempty"`
	Title     string            `json:"title,omitempty"`
	Content   string            `json:"content,omitempty"`
}

// OA消息
type OaMsgDto struct {
	Head         OaMsgHeadDto       `json:"head"`                     //消息头部内容，必须
	StatusBar    *OaMsgStatusBarDto `json:"status_bar,omitempty"`     //可选，消息状态栏，只支持接收者的userid列表
	Body         OaMsgBodyDto       `json:"body"`                     //消息体，必须
	MessageURL   string             `json:"message_url"`              //消息点击链接地址，必须
	PcMessageURL string             `json:"pc_message_url,omitempty"` //PC端点击消息时跳转到的地址，可选
}

// 文件消息
type FileMsgDto struct {
	MediaID string `json:"media_id"`
}

type BtnJSONListDto struct {
	Title     string `json:"title"`
	ActionURL string `json:"action_url"`
}

// 卡片消息
type ActionCardMsgDto struct {
	BtnOrientation string          `json:"btn_orientation,omitempty"` //使用独立跳转ActionCard样式时的按钮排列方式，必须与btn_json_list同时设置，0竖排，1横排
	BtnJSONList    *BtnJSONListDto `json:"btn_json_list,omitempty"`   //使用独立跳转ActionCard样式时的按钮列表；必须与btn_orientation同时设置，且长度不超过1000字符
	SingleURL      string          `json:"single_url,omitempty"`      // 消息点击链接地址
	SingleTitle    string          `json:"single_title,omitempty"`    //使用整体跳转ActionCard样式时的标题。必须与single_url同时设置，最长20个字
	Markdown       string          `json:"markdown"`                  // 必须
	Title          string          `json:"title"`                     //可选
}

// 链接消息
type LinkMsgDto struct {
	PicURL     string `json:"picUrl"`
	MessageURL string `json:"messageUrl"`
	Text       string `json:"text"`
	Title      string `json:"title"`
}

// markdown消息
type MarkdownMsgDto struct {
	Text  string `json:"text"`
	Title string `json:"title"`
}

// 文本消息
type TextMsgDto struct {
	Content string `json:"content"`
}

type WorkNotifyRespDto struct {
	godingtalk.DingTalkApiErr2Dto
	TaskID    int64  `json:"task_id"`
	RequestID string `json:"request_id"`
}
