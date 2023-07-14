package global

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
)

// 获取标识代码,如果没有就用时间戳
func GetCode(code string) string {
	if code == "" {
		return fmt.Sprint(time.Now().Unix())
	} else {
		return code
	}
}

// 忽略了错误信息,没有测试前不要用这个函数
func ToInt(num string, default_ int) int {
	number, err := strconv.Atoi(num)
	if err != nil {
		return default_
	}
	return number
}

// 发送文件
func SendFile(path string, data []byte, name string) bool {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	part, err := writer.CreateFormFile(name, name)
	if err != nil {
		return false
	}
	part.Write(data)
	if err = writer.Close(); err != nil {
		return false
	}
	req, err := http.NewRequest(http.MethodPost, path, buf)
	if err != nil {
		return false
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return true
}
