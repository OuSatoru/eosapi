package main

import (
	"net/http"
	"github.com/OuSatoru/eosapi/svr"
)

func main() {
	//// client可以共用
	//client := &http.Client{}
	//// jsessionid几乎是唯一标识
	//jsessionid := hoyi.Jsessionid()
	//userName := "09800903"
	//password := "000000"
	//fmt.Println(htmlpick.HasLogin(hoyi.Login(userName, password, client, jsessionid)))
	//b, _ := ioutil.ReadFile("Unread.html")
	//fmt.Println(htmlpick.UnreadListJson(string(b)))
	http.HandleFunc("/login", svr.Login)
	http.ListenAndServe(":2333", nil)
}

