package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	username   string = "" // 用户名
	password   string = ""    // 密码
	tradePwd   string = ""         // 支付密码
	investPart string = ""              // 投资金额
)

var token string
var borrowId string // 专享标ID
var rate string     // 转出利率

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU() - 1)

	go GetNewBorrow()
	time.Sleep(5 * time.Second)

	token = login()
	// fmt.Println("token:", token)

	for {
		result := robBorrow()
		if result == "未登录" {
			token = login()
		} else if result == "" {
			log.Println("估计抢到了。")
		} else if result != "债权购买异常:该债权状态不是投标中,暂时不能购买!" {
			log.Println(result)
		}
	}
}

type Login struct {
	Auth_key      string `json:"auth_key"`
	UserId        string `json:"userId"`
	LlpayBankList string `json:"llpayBankList"`
}

// 登陆，返回Authorization信息
func login() string {
	body := bytes.NewBuffer([]byte(`{"username":"` + username + `", "password":"` + password + `"}`))
	res, err := http.Post("http://api.touzhijia.com/login", "application/json;charset=utf-8", body)
	if err != nil {
		return ""
	}

	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return ""
	}

	var login Login
	if err := json.Unmarshal(result, &login); err != nil {
		// fmt.Println("auth_key:", login.Auth_key)
		return `Basic ` + base64.StdEncoding.EncodeToString([]byte(login.Auth_key+":"))
	}
	return ""
}

// 投标
func robBorrow() string {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://api.touzhijia.com/borrows/buy", strings.NewReader(`{"tradePwd":"`+tradePwd+`","borrowid":"`+borrowId+`","investPart":"`+investPart+`"}`))
	if err != nil {
		// handle error
	}

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Accept", "text/plain")
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	var rest map[string]interface{}
	json.Unmarshal(body, &rest)
	return rest["message"].(string)
}

type BorrowList struct {
	List []Borrow `json:"list"`
}

type Borrow struct {
	Id   string `json:"id"`
	Tag  string `json:"tag"`
	Rate string `json:"rate"`
}

// 获取新标,频率1分钟
func GetNewBorrow() {
	for {
		client := &http.Client{}

		req, err := http.NewRequest("POST", "http://api.touzhijia.com/borrows/v3?page=1&per_page=20", strings.NewReader(""))
		if err != nil {
			// handle error
		}

		req.Header.Set("Cache-Control", "no-cache")
		req.Header.Set("Accept", "text/plain")
		req.Header.Set("Authorization", token)
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}

		var borLst BorrowList
		json.Unmarshal(body, &borLst)

		for _, val := range borLst.List {
			if val.Tag == "专享债权" {
				if borrowId != val.Id {
					log.Printf("成功更新标ID，原ID：%s \t 新ID：%s \n", borrowId, val.Id)
					borrowId = val.Id

					rate, _ := strconv.ParseFloat(val.Rate, 64)
					principal, _ := strconv.ParseFloat(investPart, 64)
					z1 := (rate/12/100*principal - 8) * 12 / principal * 100
					z2 := (rate/12/100*principal - 10) * 12 / principal * 100
					z3 := (rate/12/100*principal - 12) * 12 / principal * 100
					log.Printf("赚 8 块 利率：%f", z1)
					log.Printf("赚 10 块 利率：%f", z2)
					log.Printf("赚 12 块 利率：%f", z3)
				} else {
					log.Printf("%s 标ID不需要更新。\n", borrowId)
				}
			}
		}
		time.Sleep(60 * time.Second)
	}
}
