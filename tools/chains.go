package tools

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"github.com/astaxie/beego/httplib"
)

func JsonStream() {
	const jsonStream = `
{"appid":"1212","eventtime":"1479781226209","spreadurl":"kkkk","appkey":"36bbf2fbaeb64d7b9623ec4ee3cde575","devicetype":"iPhone8,2","adnetname":"\u667a\u6167\u63a8-\u96f7\u9706IOS-\u5929\u5929\u5feb\u62a5\u5927\u56fe","osversion":"10.1.1","idfa":"idfa","clicktime":"1479781180851","ip":"192.1.168.192"}
{"appid":"2323","eventtime":"1479781341934","spreadurl":"kkkk","appkey":"36bbf2fbaeb64d7b9623ec4ee3cde575","devicetype":"iPhone7,1","adnetname":"\u667a\u6c47\u63a8-\u96f7\u9706\u4e4b\u6218-\u817e\u8baf\u5927\u56fe","osversion":"10.0.1","idfa":"idfa","clicktime":"1479780829406","ip":"192.1.168.192"}
`
	JsonHandle(jsonStream)
}

func httpGet(url string) {
	req := httplib.Get(url)
	str, err := req.String()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}

func JsonStreamFile() {
	dat, err := ioutil.ReadFile("/Users/sino/Downloads/3.txt")
	if err != nil {
		log.Fatal(err)
	}

	JsonHandle(string(dat))
}

func JsonHandle(jsonStream string) {
	type Message struct {
		Appid     string `json:"appid"`
		Spreadurl string `json:"spreadurl"`
		Adnetname string `json:"adnetname"`
		Idfa      string `json:"idfa"`
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		switch m.Adnetname {
		case "智汇推-腾讯新闻客户端小图":
			m.Adnetname = "智汇推-雷霆IOS-腾讯小图"
		case "智汇推-雷霆之战-腾讯大图":
			m.Adnetname = "智慧推-雷霆IOS-腾讯大图"
		case "智慧推-雷霆IOS-天天快报大图":
			m.Adnetname = "智汇推-雷霆IOS-快报大图"
		case "智汇推-雷霆IOS-天天快报小图":
			m.Adnetname = "智汇推-雷霆IOS-天天快报小图"
		}
		fmt.Println(m)

		//		url := fmt.Sprintf("xxx?appid=%s&idfa=%s&adnetname=%s&spreadurl=%s", m.Appid, m.Idfa, m.Adnetname, m.Spreadurl)

	}
}