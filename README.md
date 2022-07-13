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

## 调用示例
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
	urlStr := godingtalk.GetLoginUrl(redirectUri, cfg.AppKey,scope,state)
	fmt.Println(urlStr)

}


```


## 发送工作消息示例
```
curl -X POST \
  'https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2?access_token=a3a63fc592083b2fb74bcc5f195015dd' \
  -H 'content-type: application/json' \
  -d '{"msg":{"msgtype":"text","text":{"content":"文本消息"}},"agent_id":"1605203268","userid_list":"manager4056"}
'

```
