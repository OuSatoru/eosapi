package htmlpick

import (
	"regexp"
	"fmt"
)

// {"第一条": "114514", "第二条": "23333"}
func UnreadList(htm string) {
	reg := regexp.MustCompile(`<a href="javascript:return void\(0\);" onclick='executeTask\((\d+?)\)'>\s*(\S+?)\s*</a>`)
	if reg.MatchString(htm) {
		fmt.Println(reg.FindAllStringSubmatch(htm, -1)[0][1])
	} else {
		fmt.Println("None.")
	}
}
