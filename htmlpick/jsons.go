// TODO: api 格式：

package htmlpick

import (
	"encoding/json"
	"regexp"

	"github.com/OuSatoru/eosapi/vals"
)

type MailList struct {
	UserName  string `json:"user_name"`
	TaskId    string `json:"task_id"`
	MailTitle string `json:"mail_title"`
}

type attachment struct {
	AttachmentName string `json:"attachment_name"`
	AttachmentUrl  string `json:"attachment_url"`
}

type MailPage struct {
	From       string       `json:"from"`
	To         string       `json:"to"`
	Attachment []attachment `json:"attachment,omitempty"`
	Body       string       `json:"body"`
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
