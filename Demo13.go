package main

import (
	"fmt"
	"regexp"
	// "strings"
)

func main() {
	var str string = `顶也～～～











[img]static/image/common/sigline.gif[/img]
[size=2][url=http://www.60kv.com]http://www.60kv.com[/url] 快播电影 [url=http://www.92kmv.com]http://www.92kmv.com[/url] 看美女图 [url=http://www.44hb.com]http://www.44hb.com[/url] 百度影音[/size]`
	regexp, _ := regexp.Compile("\\[[^\\]]+\\]")
	str = regexp.ReplaceAllString(str, "")
	fmt.Println(str)
	// str = strings.TrimSpace(str) //去空格
	// //将HTML标签全转换成小写
	// re, _ := regexp.Compile("\\[[\\S\\s]+?\\]")
	// str = re.ReplaceAllStringFunc(str, strings.ToLower)
}
