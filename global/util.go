package global

import (
	"fmt"
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
func ToInt(num string) int {
	number, _ := strconv.Atoi(num)
	return number
}

// 发送文件
func SendFile(path string, data []byte, code string) {

}
