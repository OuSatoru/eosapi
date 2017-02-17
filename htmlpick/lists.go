package htmlpick

import (
	"regexp"
)

// {"114514": "第一条", "23333": "第二条"}
func UnreadListMap(htm string) map[string]string {
	m := make(map[string]string)
	reg := regexp.MustCompile(`<a href="javascript:return void\(0\);" onclick='executeTask\((\d+?)\)'>\s*(\S+?)\s*</a>`)
	if reg.MatchString(htm) {
		for _, il := range reg.FindAllStringSubmatch(htm, -1) {
			m[il[1]] = il[2]
		}
		return m
	} else {
		return nil
	}
}
