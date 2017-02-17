package svr

import (
	"net/http"
	"github.com/OuSatoru/eosapi/hoyi"
	"fmt"
	"github.com/OuSatoru/eosapi/htmlpick"
)

var Client = &http.Client{}
var Jsessionid = hoyi.Jsessionid()

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		userName := r.Form["username"][0]
		password := r.Form["password"][0]
		if htmlpick.HasLogin(hoyi.LoginEos(userName, password, Client, Jsessionid)) {
			fmt.Fprint(w, "OK")
		} else {
			fmt.Fprint(w, "Login Failed")
		}
	} else {
		fmt.Fprint(w, "请登录")
	}
}
