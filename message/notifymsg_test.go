package message

import (
	"fmt"
	"github.com/jellycheng/godingtalk"
	"github.com/jellycheng/gosupport"
	"github.com/jellycheng/gosupport/env"
	"testing"
)

// go test -run=TestSendWorkNotifyReqDto
func TestSendWorkNotifyReqDto(t *testing.T) {
	err := env.LoadEnv2DataManage("../example/.env")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	globalEnv := gosupport.NewGlobalEnvSingleton()

	req := SendWorkNotifyReqDto{}
	req.AgentID = globalEnv.GetString("DingTalk_1_AgentId")
	req.UseridList = globalEnv.GetString("UseridListDemo")
	req.Msg = WorkNotifyMsgDto{
		Msgtype: "text",
		Text: &TextMsgDto{
			Content: "文本消息",
		}}
	resp := godingtalk.ToJson(req)
	fmt.Println(resp)

	req2 := SendWorkNotifyReqDto{}
	req2.AgentID = globalEnv.GetString("DingTalk_1_AgentId")
	req2.UseridList = globalEnv.GetString("UseridListDemo")
	markdownText := `
标题
# 一级标题
## 二级标题
### 三级标题
#### 四级标题
##### 五级标题
###### 六级标题
 
引用
> 引用文案.
 
文字加粗、斜体
**bold**
*italic*
 
链接
[this is a link，跳转baidu](http://www.baidu.com)
 
图片
![](http://name.com/pic.jpg)
 
无序列表
- item1
- item2
 
有序列表
1. item1
2. item2
`
	req2.Msg = WorkNotifyMsgDto{
		Msgtype: "markdown",
		Markdown: &MarkdownMsgDto{
			Title: "消息标题01",
			Text:  markdownText,
		}}
	resp2 := godingtalk.ToJson(req2)
	fmt.Println(resp2)

}

// go test -run=TestSendWorkMsg
func TestSendWorkMsg(t *testing.T) {
	err := env.LoadEnv2DataManage("../example/.env")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	globalEnv := gosupport.NewGlobalEnvSingleton()

	req := SendWorkNotifyReqDto{}
	req.AgentID = globalEnv.GetString("DingTalk_1_AgentId")
	req.UseridList = globalEnv.GetString("UseridListDemo")
	markdownText := `
标题
# 一级标题
## 二级标题
### 三级标题
#### 四级标题
##### 五级标题
###### 六级标题
 
引用
> 引用文案.

[跳转baidu](http://www.baidu.com)
 
图片
![](http://name.com/pic.jpg)

`
	req.Msg = WorkNotifyMsgDto{
		Msgtype: "markdown",
		Markdown: &MarkdownMsgDto{
			Title: "消息标题01",
			Text:  markdownText,
		}}
	resp := SendWorkMsg(req, globalEnv.GetString("AccessTokenDemo"))
	godingtalk.Debug(resp)
}
