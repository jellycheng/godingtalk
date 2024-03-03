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
    go get -u github.com/jellycheng/godingtalk@master
    
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


##  钉钉文档
```
钉钉开放平台： https://open.dingtalk.com
创建h5微应用： https://open.dingtalk.com/document/orgapp/create-an-h5-micro-application


```

