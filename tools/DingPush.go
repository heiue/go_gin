package tools

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type SimpleRequest struct {
	Msgtype string `json:"msgtype"`
	Text SimpleRequestContent `json:"text"`
	MarkDown SimpleMarkDown `json:"markdown"`
}

type SimpleRequestContent struct {
	Content string `json:"content"`
}

// markdown
type SimpleMarkDown struct {
	Title string `json:"title"`
	Text string `json:"text"`
}

var access_token = "a84a49cee61270e8dbebe5ec6f7cd4b1160e02f3685bfdfee9536ea8cb16b812"
var secert = "SEC54866c8f6f0f0ff0a8c2f04156337f751eea8e848a8c7c0246146b99ec8a8abc"
var api = "https://oapi.dingtalk.com/robot/send?"

func Sign(secret string) (int64, string) {
	timesp := time.Now().UnixNano() / 1e6
	string_sign := fmt.Sprintf("%d\n%s", timesp, secret)
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(string_sign))
	sum := hash.Sum(nil)
	encode_sign := base64.StdEncoding.EncodeToString(sum)
	urlencode_sign := url.QueryEscape(encode_sign)
	return timesp, urlencode_sign
}

func Send()  {
	times, str := Sign(secert)

	res_api := fmt.Sprintf("%saccess_token=%s&timestamp=%d&sign=%s", api, access_token, times, str)

	// 设置请求头 test
	/*requestBody := SimpleRequest{
		Text: SimpleRequestContent{
			Content: "simple request",
		},
		Msgtype: "text",
	}*/
	requestBody := SimpleRequest{
		Msgtype: "markdown",
		MarkDown: SimpleMarkDown{
			Title: "天气",
			Text: "#### 杭州天气 @150XXXXXXXX \n > 9度，西北风1级，空气良89，相对温度73%\\n> ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\\n> ###### 10点20分发布 [天气](https://www.dingalk.com) \\n",
		},
	}
	reqBodyBox, _ := json.Marshal(requestBody)
	body := string(reqBodyBox)
	fmt.Println(body)


	request, _ := http.NewRequest("POST", res_api, strings.NewReader(body))

	// 设置库端口
	client := &http.Client{}

	// 请求头添加内容
	request.Header.Set("Content-Type", "application/json")

	// 发送请求
	response, _ := client.Do(request)
	fmt.Println("response: ", response)

	// 关闭 读取 reader
	defer response.Body.Close()

	// 读取内容
	all, _ := ioutil.ReadAll(response.Body)
	fmt.Println("all: ", string(all))
}