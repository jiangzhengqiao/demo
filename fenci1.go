package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/huichen/sego"
	"wdzj.com/wdzj/logger"
	"wdzj.com/wdzj/utils"
)

func main() {
	content := `808信贷和E速贷本帖最后由p2p 浪子理财 于 20121314-2-11 23:23 编辑安全性，流动性和收益性是资本的三性，对网贷投资做简单的三性分析，抛砖引玉贻笑大方了。
 一，安全性原则。
 安全性是第一性，是流动性和收益性的基础，没有了安全性，其他的都没有意义。
 任何网贷投资，首先要排除非系统性风险。如果本金会由于借款人的逾期而不能收回的，也就是网站不保本的标，坚决不投，如人人贷中的HR标。我没有担保权限，对资本衍生出来的担保权利没有太多的研究，不过，担保对于逾期是要承担垫付责任的，也就是风险完全暴露，那是违背安全性原则的，因此担保，还是应该慎重。在网贷投标中，安全标优先。当然这里的优先是相对的优先，还需要衡量整体收益率水平，做综合考虑。如在红岭，秒标是首选，因为没有风险，快借标次之，然后是担保标和黄牛标，最后是信用标。网站自动投标改革之后，对于小投资者是重大利好，大家可以积累资金，一次投资于快借之中，其他标可以无视开开贷<a href="http://shuju.wdzj.com/">p2p</a>。
`

	// 载入词典
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("/Users/jiangzhengqiao/fenci1.txt")

	// 一、分词
	t := time.Now()
	t1 := time.Now()
	segments := segmenter.Segment([]byte(content))
	words := sego.SegmentsToSlice(segments, true)
	fmt.Println("分词用时:", time.Now().Sub(t1))

	// 二、去重
	t2 := time.Now()
	out, mask, urls := UniqueSlice(words)
	fmt.Println("去重用时:", time.Now().Sub(t2))

	// 三、排序
	t3 := time.Now()
	res := utils.InsertionSort(out)
	fmt.Println("排序用时:", time.Now().Sub(t3))

	// 四、文本替换
	t4 := time.Now()
	for e := res.Front(); nil != e; e = e.Next() {
		val := e.Value.(string)
		content = strings.Replace(content, val, mask[val], 1)

	}

	for key, val := range mask {
		biao := strings.LastIndex(content, val)
		if biao >= 0 {
			start := strings.Count(content[biao:], "<a")
			end := strings.Count(content[biao:], "</a>")

			if start == end {
				content = strings.Replace(content, val, urls[val], 1)
			} else {
				logger.Info("有未关闭 a 标签。", key)
			}
		}
	}

	fmt.Println("文本替换用时:", time.Now().Sub(t4))
	fmt.Println("替换锚文本：", len(out), "个")
	fmt.Println("共耗时:", time.Now().Sub(t))
	fmt.Println("替换后文本:", content)
	fmt.Println("--------------------------------")
	logger.Info("文本长度:", len(content), "替换锚文本：", len(out), "共耗时:", time.Now().Sub(t))
}

func UniqueSlice(slice []string) (ret []string, found map[string]string, urls map[string]string) {
	// t := time.Now()
	found = make(map[string]string)
	urls = make(map[string]string)
	total := 0
	var url string
	for i, val := range slice {
		arrs := strings.Split(val, "\t")
		val, url = arrs[0], arrs[1]
		if _, ok := found[val]; !ok {
			h := md5.New()
			h.Write([]byte(val))
			uuid := hex.EncodeToString(h.Sum(nil))
			found[val] = uuid
			(slice)[total] = strings.Split((slice)[i], "\t")[0]
			urls[uuid] = url
			total++
		}
	}
	ret = (slice)[:total]
	// fmt.Println("去重用时:", time.Now().Sub(t))
	return
}
