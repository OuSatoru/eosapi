package svr

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/OuSatoru/eosapi/hoyi"
	"github.com/OuSatoru/eosapi/htmlpick"
	"github.com/OuSatoru/eosapi/vals"
)

func Login(w http.ResponseWriter, r *http.Request) {
	vals.Jsessionid = hoyi.Jsessionid()
	if r.Method == "POST" {
		r.ParseForm()
		vals.UserName = r.Form["username"][0]
		vals.Password = r.Form["password"][0]
		if htmlpick.HasLogin(hoyi.LoginEos(vals.UserName, vals.Password, vals.Client, vals.Jsessionid)) {
			fmt.Fprint(w, "OK")
		} else {
			fmt.Fprint(w, "Login Failed")
		}
	} else {
		t, _ := template.ParseFiles("post.html")
		t.Execute(w, nil)
	}
}

//每次到其他页都会读首页，只为读取总邮件数
func TodoList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	htm := hoyi.MailUnread(vals.UserName, vals.Password, vals.Client, vals.Jsessionid)
	page := 1
	if r.Form.Get("page") != "" && r.Form.Get("page") != "1" {
		page, _ = strconv.Atoi(r.Form.Get("page"))
	}
	pageLen := 10
	if r.Form.Get("len") != "" && r.Form.Get("len") != "10" {
		pageLen, _ = strconv.Atoi(r.Form.Get("len"))
	}
	if page == 1 && pageLen == 10 {
		fmt.Fprint(w, htmlpick.UnreadListJson(htm))
	} else {
		pageCount := htmlpick.UnreadCount(htm)
		htm2 := hoyi.MailUnreadPage(vals.UserName, vals.Password, vals.Client, vals.Jsessionid, page, pageLen, pageCount)
		fmt.Fprint(w, htmlpick.UnreadListJson(htm2))
	}
}

func ExecTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
}
