package hoyi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const (
	mainpage    string = "http://32.185.24.71/hoyi/com.hoyi.desktop.login.flow"
	list4todo   string = "http://32.185.24.71/hoyi/oa/workcenter/com.hoyi.oa.workcenter.workitem_list4todo.flow"
	executeTask string = "http://32.185.24.71/hoyi/oa/workcenter/com.hoyi.workflow.client.pageflow.executeTask.flow?_eosFlowAction=action0&workitemID=" //in executeTask(%d)
)

func Jsessionid() string {
	resp, err := http.Get(mainpage)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	setCookie := resp.Header.Get("Set-Cookie")
	reg := regexp.MustCompile(`JSESSIONID=(.+);`)
	if reg.MatchString(setCookie) {
		return reg.FindAllStringSubmatch(setCookie, -1)[0][1]
	} else {
		return ""
	}
}

func LoginEos(userName string, password string, client *http.Client, jsessionid string) string {
	req, err := http.NewRequest("POST", mainpage, strings.NewReader(fmt.Sprintf("_eosFlowAction=login&username=%s&password=%s&companyId=1", userName, password)))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", fmt.Sprintf("JSESSIONID=%s; hoyi_username=%s; hoyi_password=%s; hoyi_companyid=1; hoyi_version=", jsessionid, userName, password))
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}

func getWithCookie(userName string, password string, client *http.Client, jsessionid string, url string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Cookie", fmt.Sprintf("hoyi_username=%s; hoyi_password=%s; hoyi_companyid=1; hoyi_version=; JSESSIONID=%s", userName, password, jsessionid))
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}

func DesktopWorkList(userName string, password string, client *http.Client, jsessionid string) string {
	//request body 不必要
	req, err := http.NewRequest("POST", "http://32.185.24.71/hoyi/desktop/desktop/com.hoyi.workflow.client.pageflow.Worklist4desktop.flow?_eosFlowAction=queryWithPage", strings.NewReader("competePara=ALL&bizPara=ALL&workitemID=&showType=&pageCond%2Fbegin=0&pageCond%2Flength=6&pageCond%2FisCount=true&pageCond%2Fcount=3"))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//JSESSIONID 必要 会变 与用户名无关
	req.Header.Set("Cookie", fmt.Sprintf("hoyi_username=%s; hoyi_password=%s; hoyi_companyid=1; hoyi_version=; JSESSIONID=%s; hoyi_last_visit_time=20170217095659", userName, password, jsessionid))
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}

func MailUnread(userName string, password string, client *http.Client, jsessionid string) string {
	return getWithCookie(userName, password, client, jsessionid, list4todo)
}

func ExecTaskPage(userName string, password string, client *http.Client, jsessionid string, taskId string) string {
	return getWithCookie(userName, password, client, jsessionid, executeTask+taskId)
}

func MailUnreadPage(userName string, password string, client *http.Client, jsessionid string, page int, pageLen int, pageCount int) string {
	req, err := http.NewRequest("POST", list4todo, strings.NewReader(fmt.Sprintf("_eosFlowAction=queryWithPage&criteria%%2F_entity=com.eos.workflow.data.WFWorkItem&criteria%%2F_orderby%%5B1%%5D%%2F_property=startTime&criteria%%2F_orderby%%5B1%%5D%%2F_sort=desc&page%%2Fbegin=%d&page%%2Flength=%d&page%%2Fcount=%d&page%%2FisCount=true", beginCk(page, pageLen, pageCount), pageLenCk(page, pageLen, pageCount), pageCount)))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", fmt.Sprintf("hoyi_username=%s; hoyi_password=%s; hoyi_companyid=1; hoyi_version=; JSESSIONID=%s; hoyi_last_visit_time=20170217095659", userName, password, jsessionid))
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}

func beginCk(page int, pageLen int, pageCount int) int {
	if page > pageCount/pageLen {
		return (pageCount / pageLen) * pageLen
	} else {
		return (page - 1) * pageLen
	}
}

func pageLenCk(page int, pageLen int, pageCount int) int {
	if page > pageCount/pageLen {
		return pageCount % pageLen
	} else {
		return pageLen
	}
}
