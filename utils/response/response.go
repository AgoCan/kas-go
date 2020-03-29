package response

// 回调使用

import (
	"fmt"
	"time"

	"bytes"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

// 定义常用ContentType类型
const (
	ContentTypeJson = "application/json"
	ContentTypeText = "text/xml"
)

// Data 回调必须使用下面的格式
type Data struct {
	Code  int                    `json:"code"`
	Data  map[string]interface{} `json:"data"`
	Error string                 `json:"error"`
}

// SendMessage 主动发送请求
func SendMessage(url string, data Data, contentType string) (string, error) {
	client := &http.Client{Timeout: 5 * time.Second}

	// json 序列化数据
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// 发送请求
	response, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}

	// 函数结束后关闭请求
	defer func(response *http.Response) {
		err = response.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(response)

	// 返回 返回值
	result, err := ioutil.ReadAll(response.Body)
	return string(result), err
}
