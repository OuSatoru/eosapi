package htmlpick

import "strings"

func HasLogin(htm string) bool {
	return !strings.Contains(htm, "密码输入错误")
}
