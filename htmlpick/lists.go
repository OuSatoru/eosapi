package htmlpick

import (
	"regexp"
	"strconv"
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

func UnreadCount(htm string) int {
	reg := regexp.MustCompile(`共\s*?(\d+)\s*?条记录`)
	if reg.MatchString(htm) {
		ls := reg.FindAllStringSubmatch(htm, -1)
		count, err := strconv.Atoi(ls[0][1])
		if err != nil {
			panic(err)
		}
		return count
	} else {
		return 0
	}
}

func UnreadCountChan(htm string, c chan int) {
	c <- UnreadCount(htm)
}
