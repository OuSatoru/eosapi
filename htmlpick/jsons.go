// api返回的json格式，还没想好= =

package htmlpick

import (
	"encoding/json"
	"regexp"
	"github.com/OuSatoru/eosapi/vals"
)

type MailList struct {
	UserName string `json:"user_name"`
	TaskId string `json:"task_id"`
	MailTitle string `json:"mail_title"`
}

func UnreadListJson(htm string) string {
	reg := regexp.MustCompile(`<a href="javascript:return void\(0\);" onclick='executeTask\((\d+?)\)'>\s*(\S+?)\s*</a>`)
	if reg.MatchString(htm) {
		iList := reg.FindAllStringSubmatch(htm, -1)
		mList := make([]MailList, len(iList))
		for k, il := range iList {
			mList[k] = MailList{UserName: vals.UserName, TaskId: il[1], MailTitle: il[2]}
		}
		if j, err := json.MarshalIndent(mList, "", "  "); err != nil {
			panic(err)
		} else {
			return string(j)
		}
	} else {
		return ""
	}
}