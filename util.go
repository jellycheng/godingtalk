package godingtalk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetUrlContent(urlStr string, headers map[string]string) (string, error) {
	req, _ := http.NewRequest("GET", urlStr, nil)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body), nil
}

func PostUrlContnet4json(urlStr string, strJson string, headers map[string]string) (string, error) {
	payload := strings.NewReader(strJson)
	req, _ := http.NewRequest("POST", urlStr, payload)

	req.Header.Add("content-type", "application/json")
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body), nil
}

func DeleteUrlContent(urlStr string, strJson string, headers map[string]string) (string, error) {
	client := &http.Client{}
	payload := strings.NewReader(strJson)
	req, _ := http.NewRequest("DELETE", urlStr, payload)
	req.Header.Add("content-type", "application/json")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body), nil
}

func PutUrlContent(urlStr string, strJson string, headers map[string]string) (string, error) {
	client := &http.Client{}
	payload := strings.NewReader(strJson)
	req, _ := http.NewRequest("PUT", urlStr, payload)
	req.Header.Add("content-type", "application/json")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body), nil
}

func JsonUnmarshal(str string, obj interface{}) error {
	return json.Unmarshal([]byte(str), obj)
}

// 转成json字符串
func ToJson(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

func Debug(msg interface{}) {
	fmt.Println(fmt.Sprintf("%+v", msg))
}

func StringPtr(s string) *string {
	return &s
}

func BoolPtr(b bool) *bool {
	return &b
}
