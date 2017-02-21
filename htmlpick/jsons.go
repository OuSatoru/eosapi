// TODO: api 格式：

package htmlpick

import (
	"encoding/json"
	"regexp"
	"fmt"
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
	Time       string       `json:"time"`
	Subject    string       `json:"subject"`
	Attachment []attachment `json:"attachment,omitempty"`
	Body       string       `json:"body"`
}

type Err struct {
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
}

func ErrJson(errcode int, errmsg string) string {
	es := Err{ErrCode:errcode, ErrMsg:errmsg}
	if j, err := json.MarshalIndent(es, "", "  "); err != nil {
		panic(err)
	} else {
		return string(j)
	}
}

func UnreadListJson(htm string, userName string) string {
	reg := regexp.MustCompile(`<a href="javascript:return void\(0\);" onclick='executeTask\((\d+?)\)'>\s*(\S+?)\s*</a>`)
	if reg.MatchString(htm) {
		iList := reg.FindAllStringSubmatch(htm, -1)
		mList := make([]MailList, len(iList))
		for k, il := range iList {
			mList[k] = MailList{UserName: userName, TaskId: il[1], MailTitle: il[2]}
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

func MailJson(htm string) string {
	reg1 := regexp.MustCompile(`(发件人:|时&nbsp;&nbsp;&nbsp;&nbsp;间:|主&nbsp;&nbsp;&nbsp;&nbsp;题:)</span></td>\s*?<td colspan="1" *>\s*?<span>(.+?)</span></td>`)
	mail := &MailPage{}
	if reg1.MatchString(htm) {
		//for k, v := range reg1.FindAllStringSubmatch(htm, -1)[0] {
		//	fmt.Println(k, v)
		//}
		ls := reg1.FindAllStringSubmatch(htm, -1)
		mail.From = ls[0][2]
		mail.Time = ls[1][2]
		mail.Subject = ls[2][2]
	} else {
		fmt.Println("aaa")
	}
	reg2 := regexp.MustCompile(`<div id="ididid">\s*?(\S+(\s+\S+)*)\s*?<span></span>`)
	blank := regexp.MustCompile(`\s+`)
	if reg2.MatchString(htm) {
		//for k, v := range reg2.FindAllStringSubmatch(htm, -1)[0] {
		//	fmt.Println(k, v)
		//}
		mail.To = blank.ReplaceAllString(reg2.FindAllStringSubmatch(htm, -1)[0][1], " ")
	} else {
		fmt.Println("bbb")
	}
	reg3 := regexp.MustCompile(`<a href="(/hoyi/comm/attachments/downloadByFilepath\.jsp\?filePath=Z:.+?)" *?>\s+?(\S.+?)\s*?</a>`)
	if reg3.MatchString(htm) {
		for _, v := range reg3.FindAllStringSubmatch(htm, -1) {
			at := attachment{AttachmentName:v[2], AttachmentUrl:v[1]}
			mail.Attachment = append(mail.Attachment, at)
		}
	} else {
		fmt.Println("ccc")
	}
	reg4 := regexp.MustCompile(`<script id="myEditor" style="height:95%;width:100%;" type="html">\s*?(\S.*?)\s*?</script>`)
	if reg4.MatchString(htm) {
		mail.Body = reg4.FindAllStringSubmatch(htm, -1)[0][1]
	}
	if j, err := json.MarshalIndent(mail, "", ""); err != nil {
		panic(err)
	} else {
		//fmt.Println(string(j))
		return string(j)
	}
	fmt.Println(mail)
	return ""
}
