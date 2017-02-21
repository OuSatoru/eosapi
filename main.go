package main

import (
	"io/ioutil"
	"fmt"
	"github.com/OuSatoru/eosapi/htmlpick"
)

func main() {
	//// client可以共用
	//client := &http.Client{}
	//// jsessionid几乎是唯一标识
	//jsessionid := hoyi.Jsessionid()
	//userName := "09800903"
	//password := "000000"
	//fmt.Println(htmlpick.HasLogin(hoyi.Login(userName, password, client, jsessionid)))

	b, _ := ioutil.ReadFile("Task.html")
	fmt.Println(htmlpick.MailJson(string(b)))

	//vals.OpenDb()
	//http.HandleFunc("/login", svr.Login)
	//http.HandleFunc("/todolist", svr.TodoList)
	//http.ListenAndServe(":2333", nil)
}
