package main

import (
	// "bytes"
	"github.com/fedesog/webdriver"
	"log"
	"strconv"
)

const (
	CONN string = "="
	END  string = ";"
)

var (
	chromeDriver *webdriver.ChromeDriver
)

func init() {
	chromeDriver = webdriver.NewChromeDriver("/Users/jiangzhengqiao/go/src/demo/chromedriver")
}

func main() {
	// // 一、获取cookie
	// url := "http://weixin.sogou.com/weixin?query=陆金所&type=1&ie=utf8&page=1"
	// chromeDirverPath := "/Users/jiangzhengqiao/go/src/demo/chromedriver"
	// cookie := GetCookie(url, chromeDirverPath)

	// // 二、
	// log.Println(cookie)

	err := chromeDriver.Start()
	if err != nil {
		log.Println(err)
	}

	page := getPage("http://weixin.sogou.com/weixin?query=陆金所&type=1&ie=utf8&page=1")
	if page > 1 {
		for i := 1; i <= page; i++ {

		}
	}

	chromeDriver.Stop()

}

// func GetCookie(url, chromeDirverPath string) string {
// 	chromeDriver := webdriver.NewChromeDriver(chromeDirverPath)
// 	err := chromeDriver.Start()
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	session, err := chromeDriver.NewSession(webdriver.Capabilities{}, webdriver.Capabilities{})
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	err = session.Url(url)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	page := getPage(session)
// 	if page > 1 {
// 		getContent(session, page)
// 	}

// 	results, _ := session.FindElement(webdriver.XPath, "//*[@id='main']/div/div[2]/div")
// 	divs, _ := results.FindElements(webdriver.ClassName, "wx-rb")
// 	if len(divs) > 0 {
// 		for _, div := range divs {
// 			url, _ := div.GetAttribute("href")
// 			log.Println("微信地址 : \t" + url)

// 			val, _ := div.FindElement(webdriver.ClassName, "txt-box")
// 			h3, _ := val.FindElement(webdriver.TagName, "h3")
// 			h3html, _ := h3.Text()
// 			log.Println("微信名称 : \t" + h3html)

// 			h4, _ := val.FindElement(webdriver.TagName, "h4")
// 			h4, _ = h4.FindElement(webdriver.Name, "em_weixinhao")
// 			h4html, _ := h4.Text()
// 			log.Println("微信号 : \t" + h4html)

// 			divs, _ := val.FindElements(webdriver.ClassName, "s-p3")
// 			sptxt, _ := divs[0].FindElement(webdriver.ClassName, "sp-txt")
// 			sphtml, _ := sptxt.Text()
// 			log.Println("功能介绍 : \t" + sphtml)

// 			ico, _ := div.FindElement(webdriver.ClassName, "pos-ico")
// 			box, _ := ico.FindElement(webdriver.ClassName, "pos-box")
// 			img, _ := box.FindElement(webdriver.TagName, "img")
// 			imgurl, _ := img.GetAttribute("src")
// 			log.Println("二维码地址 : \t" + imgurl)
// 			log.Println("-------------------------------------------------------")
// 		}
// 	}

// 	// arrs, _ := session.FindElements(webdriver.ClassName, "txt-box")
// 	// if len(arrs) > 0 {
// 	// 	for _, val := range arrs {
// 	// 		h3, _ := val.FindElement(webdriver.TagName, "h3")
// 	// 		h3html, _ := h3.Text()

// 	// 		h4, _ := val.FindElement(webdriver.TagName, "h4")
// 	// 		h4, _ = h4.FindElement(webdriver.Name, "em_weixinhao")
// 	// 		h4html, _ := h4.Text()

// 	// 		divs, _ := val.FindElements(webdriver.ClassName, "s-p3")

// 	// 		sptxt, _ := divs[0].FindElement(webdriver.ClassName, "sp-txt")
// 	// 		sphtml, _ := sptxt.Text()

// 	// 		log.Println("title : \t" + h3html)
// 	// 		log.Println("微信号 : \t" + h4html)
// 	// 		log.Println("功能介绍 : \t" + sphtml)
// 	// 		// log.Println("二维码地址 : \t" + imghtml)
// 	// 		log.Println("-------------------------------------------------------")
// 	// 	}
// 	// }

// 	cookies, err := session.GetCookies()
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	var cookie *bytes.Buffer
// 	if len(cookies) > 0 {
// 		cookie = bytes.NewBufferString("")
// 		for _, val := range cookies {
// 			cookie.WriteString(val.Name)
// 			cookie.WriteString(CONN)
// 			cookie.WriteString(val.Value)
// 			cookie.WriteString(END)
// 		}
// 	}

// 	session.Delete()
// 	chromeDriver.Stop()
// 	return cookie.String()
// }

//
func getPage(url string) int {
	session, err := chromeDriver.NewSession(webdriver.Capabilities{}, webdriver.Capabilities{})
	if err != nil {
		log.Println(err)
	}

	err = session.Url(url)

	if err != nil {
		log.Println(err)
	}

	p, _ := session.FindElement(webdriver.ID, "pagebar_container")
	mun, _ := p.FindElement(webdriver.ClassName, "mun")
	scdnum, _ := mun.FindElement(webdriver.ID, "scd_num")
	pageText, _ := scdnum.Text()
	pageCount, _ := strconv.Atoi(pageText)

	var page int = 1
	if pageCount > 10 {
		page = pageCount/10 + 1
	}
	session.Delete()

	return page
}
