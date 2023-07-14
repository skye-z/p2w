package global

import (
	"fmt"
	"time"
)

func GetCode(code string) string {
	if code == "" {
		return fmt.Sprint(time.Now().Unix())
	} else {
		return code
	}
}
