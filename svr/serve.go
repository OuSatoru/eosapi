package svr

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/OuSatoru/eosapi/hoyi"
	"github.com/OuSatoru/eosapi/htmlpick"
	"github.com/OuSatoru/eosapi/vals"
	"github.com/OuSatoru/eosapi/db"
)

func Login(w http.ResponseWriter, r *http.Request) {
	jsessionid := hoyi.Jsessionid()
	if r.Method == "POST" {
		r.ParseForm()
		userName := r.Form["username"][0]
		password := r.Form["password"][0]
		if htmlpick.HasLogin(hoyi.LoginEos(userName, password, vals.Client, jsessionid)) {
			db.InsUser(vals.Db, userName, password)
			db.UpdJsession(vals.Db, userName, jsessionid)
			fmt.Fprint(w, "OK")
		} else {
			fmt.Fprint(w, "Login Failed")
		}
	} else {
		t, _ := template.ParseFiles("post.html")
		t.Execute(w, nil)
	}
}

func TodoList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userName := r.Form.Get("username")
	if userName == "" {
		fmt.Fprint(w, htmlpick.ErrJson(1, "无用户名"))
		return
	}
	jsessionid := db.SelJsession(vals.Db, userName)
	password := db.SelPwd(vals.Db, userName)
	c := make(chan int)
	go htmlpick.UnreadCountChan(hoyi.MailUnread(userName, password, vals.Client, jsessionid), c)
	var page int
	var pageLen int
	var err error
	if ps := r.Form.Get("page"); ps != "" {
		page, err = strconv.Atoi(ps)
		if err != nil {
			panic(err)
		}
	} else {
		page = 1
	}
	if pls := r.Form.Get("len"); pls != "" {
		pageLen, err = strconv.Atoi(pls)
		if err != nil {
			panic(err)
		}
	} else {
		pageLen = 10
	}
	htm := hoyi.MailUnreadPage(userName, password, vals.Client, jsessionid, page, pageLen, <-c)
	fmt.Fprint(w, htmlpick.UnreadListJson(htm, userName))

}

func ExecTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
}
