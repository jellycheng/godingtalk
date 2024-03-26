# godingtalk
```
封装钉钉开放的接口

```

## Install下载依赖
```
go get -u github.com/jellycheng/godingtalk
或者
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/jellycheng/godingtalk

直接获取master分支代码：
    go get -u github.com/jellycheng/godingtalk@main
    
```

## Documentation
[https://pkg.go.dev/github.com/jellycheng/godingtalk](https://pkg.go.dev/github.com/jellycheng/godingtalk)

## 拼接钉钉登录地址
```
package main

import (
	"fmt"
	"github.com/jellycheng/godingtalk"
)

func main() {
	cfg := godingtalk.DingAgentConfig{
		AppKey:"your appkey",
		AppSecret:"your appsecret",
	}
	// 授权成功后跳转地址
	redirectUri := "http://www.xxx.com/callback/dingdingtalk"
	scope := "openid corpid"
	state := ""
	// 拼接钉钉登录地址
	urlStr := godingtalk.GetLoginUrl(redirectUri, cfg.AppKey,scope,state)
	fmt.Println(urlStr)

}


```

## 企业内部应用获取access_token
```
package main

import (
	"fmt"
	"github.com/jellycheng/godingtalk"
)

func main() {
	cfg := godingtalk.DingAgentConfig{
		AppKey:    "钉钉应用AppKey",
		AppSecret: "钉钉应用AppSecret",
	}
	at := godingtalk.GetAccessToken(cfg)
	fmt.Println(at.AccessToken)
}

```


## 发送工作通知消息示例
```
发送文本消息：
    curl -X POST \
      'https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2?access_token=a3a63fc592083b2fb74bcc5f195015dd' \
      -H 'content-type: application/json' \
      -d '{"msg":{"msgtype":"text","text":{"content":"文本消息"}},"agent_id":"1605203268","userid_list":"manager4056"}
    '
    代码示例如下：
    package main
    
    import (
    	"fmt"
    	"github.com/jellycheng/godingtalk/message"
    )
    
    func main() {
    
    	req := message.SendWorkNotifyReqDto{}
    	req.AgentID = "1605203268"     //agentid
    	req.UseridList = "manager4056" // 钉钉用户ID，多个逗号分隔
    
    	req.Msg = message.WorkNotifyMsgDto{
    		Msgtype: "text",
    		Text: &message.TextMsgDto{
    			Content: "文本消息内容，如：请尽快提交本周周报",
    		}}
    	accessToken := "1cd01dc304d434e988afd680561712fb" //调用接口：企业内部应用获取access_token
    	resp := message.SendWorkMsg(req, accessToken)
    	fmt.Println(resp)
    }


发送链接消息：
第1步：企业内部应用获取access_token
第2步：
    package main
    
    import (
    	"fmt"
    	"github.com/jellycheng/godingtalk/message"
    )
    
    func main() {
    
    	req := message.SendWorkNotifyReqDto{}
    	req.AgentID = "1605203268"     //agentid
    	req.UseridList = "manager4056" // 钉钉用户ID，多个逗号分隔
    
    	req.Msg = message.WorkNotifyMsgDto{
    		Msgtype: "link",
    		Link: &message.LinkMsgDto{
    			PicURL:     "@lADOADmaWMzazQKA", // 图片资源ID
    			MessageURL: "http://www.baidu.com", //链接地址
    			Title:      "要货单审批", //链接标题名
    			Text:       "张三提交的要货单", //链接描述
    		}}
    	accessToken := "1cd01dc304d434e988afd680561712fb" //调用接口：企业内部应用获取access_token
    	resp := message.SendWorkMsg(req, accessToken)
    	fmt.Println(resp)
    }

```

## 发送markdown格式的工作工作消息
```
package main

import (
	"fmt"
	"github.com/jellycheng/godingtalk/message"
)

func main() {

	req := message.SendWorkNotifyReqDto{}
	req.AgentID = "1605203268"     //agentid
	req.UseridList = "manager4056" // 钉钉用户ID，多个逗号分隔
	markdownText := `
# 张三提交的要货单
[去审批](http://www.baidu.com)

`
	req.Msg = message.WorkNotifyMsgDto{
		Msgtype: "markdown",
		Markdown: &message.MarkdownMsgDto{
			Title: "张三提交的要货单", //标题
			Text:  markdownText, // markdown内容
		}}
	accessToken := "1cd01dc304d434e988afd680561712fb" //调用接口：企业内部应用获取access_token
	resp := message.SendWorkMsg(req, accessToken)
	fmt.Println(resp)
}


```

## 根据手机号获取企业账号用户的userId
```
package main

import (
	"fmt"
	"github.com/jellycheng/godingtalk"
)

func main() {
	req := godingtalk.GetUseridByMobileReqDto{
		Mobile:                        "手机号",
		SupportExclusiveAccountSearch: false,
	}
	accessToken := "1cd01dc304d434e988afd680561712fb" //调用接口：企业内部应用获取access_token
	res := godingtalk.GetUseridByMobile(req, accessToken)
	fmt.Println(res.Result.Userid)

}

```

## 查询钉钉用户详情
```
package main

import (
	"fmt"
	"github.com/jellycheng/godingtalk"
)

func main() {
	accessToken := "68e9ee0cac0035a295923b0bd7dbef76"                          //调用接口：企业内部应用获取access_token
	userid := "18024719687686"                                                 //钉钉用户ID
	userprofileRespDto := godingtalk.GetUserInfo(accessToken, userid, "zh_CN") // 查询钉钉用户信息，包括 unionid
	if userprofileRespDto.Errcode == 0 {
		fmt.Println(fmt.Sprintf("用户详情： %+v", userprofileRespDto.Result))
		fmt.Println("Unionid: ", userprofileRespDto.Result.Unionid)
	} else {
		fmt.Println("错误：", userprofileRespDto.Errmsg)
	}

}

```

## 创建钉钉待办任务
```
package main

import (
	"fmt"
	"github.com/jellycheng/godingtalk/todotask"
)

func main() {
	at := "cb5673b689e132bc93ff6eede12fbe2a" // 访问凭证
	unionId := "G5x1lVRm5TSw2JToFbkOmAiEiE"  // 钉钉用户unionId
	operatorId := unionId
	detailUrl := new(todotask.DetailURL)
	detailUrl.PcURL = "https://www.php.net"    // PC端详情页url跳转地址
	detailUrl.AppURL = "https://www.baidu.com" //APP端详情页url跳转地址

	reqDto := todotask.CreateTodoTaskReqDto{
		Subject:            "钉钉待办标题369",
		CreatorID:          unionId, //创建者的unionId
		Description:        "钉钉代办描述，你有xxx提交的审批单审批",
		ExecutorIds:        []string{unionId},
		IsOnlyShowExecutor: true,
		DetailURL:          detailUrl,
	}
	respDto := todotask.CreateTodoTask(reqDto, at, unionId, operatorId)
	fmt.Println(fmt.Sprintf("%+v", respDto))

}

```

## 删除钉钉待办任务
```
package main

import (
	"fmt"
	"github.com/jellycheng/godingtalk/todotask"
)

func main() {
	at := "5bbb653cff473e279360c7c4755d6ecf" // 访问凭证
	unionId := "G5x1lVRm5TSw2JToFbkOmAiEiE"  // 钉钉用户unionId
	operatorId := unionId
	taskId := "taska684c561500c8674e4f08a57c298f6f7" //任务ID

	respDto := todotask.DeleteTodoTask(at, unionId, operatorId, taskId)
	fmt.Println(fmt.Sprintf("%+v", respDto))

}

```

## 更新钉钉待办任务完成
```
package main

import (
	"fmt"
	"github.com/jellycheng/godingtalk/todotask"
)

func main() {
	at := "5bbb653cff473e279360c7c4755d6ecf" // 访问凭证
	unionId := "G5x1lVRm5TSw2JToFbkOmAiEiE"  // 钉钉用户unionId
	operatorId := unionId
	taskId := "taska731540ae5c7cba052df63a374336912" //任务ID
	status := true // 更新任务完成
	respDto := todotask.UpdateTodoTaskStatus(at, unionId,operatorId,taskId, status)
	fmt.Println(fmt.Sprintf("%+v", respDto))

}

```

##  钉钉文档
```
钉钉开放平台： https://open.dingtalk.com
创建h5微应用： https://open.dingtalk.com/document/orgapp/create-an-h5-micro-application


```

