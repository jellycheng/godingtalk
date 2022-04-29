package main

import (
	"fmt"
	"github.com/jellycheng/gosupport"
	"github.com/jellycheng/gosupport/env"
	"net/http"
)

var globalEnv = gosupport.NewGlobalEnvSingleton()
func init()  {
	err := env.LoadEnv2DataManage("./.env")
	if err!=nil {
		fmt.Println(err.Error())
	}
}

func main()  {
	if globalEnv.GetString("IS_VERIFY") == "1" {
		step01()
	} else {
		dingTalkLogin()
	}

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(""))
	})

	http.ListenAndServe(":" + globalEnv.GetString("APP_PORT"), nil)
}

func step01()  {

}

func dingTalkLogin()  {

}
