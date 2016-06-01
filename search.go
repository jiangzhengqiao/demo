package main

import (
	"bufio"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"
)

var (
	// searcher是协程安全的
	searcher = engine.Engine{}
)

func main() {
	path := "/Users/jiangzhengqiao/dictionary.txt"

	// 初始化词库
	searcher.Init(types.EngineInitOptions{
		SegmenterDictionaries: path})
	defer searcher.Close()

	dictFile, err := os.Open("/Users/jiangzhengqiao/fenci1.txt")
	defer dictFile.Close()
	if err != nil {
		log.Fatalf("无法载入文档索引。 \"%s\" \n", dictFile)
	}

	// 将文档加入索引
	reader := bufio.NewReader(dictFile)
	var i uint64 = 0
	var text string
	var pos string
	names, urls := []string{}, []string{}
	dictionary := make(map[string]string)
	for {
		line, _ := reader.ReadString('\n')
		arrs := strings.Split(line, "\t")
		size := len(arrs)
		if size == 2 {
			text = arrs[0]
			pos = strings.Replace(arrs[1], "\n", "", -1)
		} else if size == 1 {
			break
		}

		searcher.IndexDocument(i, types.DocumentIndexData{Content: text})
		names = append(names, text)
		urls = append(urls, pos)
		dictionary[text] = pos
		i++
	}

	// 等待索引刷新完毕
	searcher.FlushIndex()

	// 查询
	t := time.Now()
	results := make(map[string]string)
	keywords := [5]string{"P2P", "p2p", "人人贷", "123", "金融"}
	if len(keywords) > 0 {
		for _, key := range keywords {
			for _, document := range searcher.Search(types.SearchRequest{Text: key}).Docs {
				if len(results) == 20 {
					goto OVER_FOR
				}
				key := names[document.DocId]
				val := urls[document.DocId]
				results[key] = val
			}
		}
	}
OVER_FOR:
	// 随机取得结果
	length := len(results)
	if length < 20 {
		length = 20 - length
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		var i int = 1
		for {
			if len(results) > 20 {
				break
			}

			num := r.Intn(len(names))
			key := names[num]
			if _, ok := results[key]; !ok {
				val := urls[num]
				results[key] = val
			}
			i++
		}
	}

	// 输出结果
	result, _ := json.Marshal(results)
	log.Println("查询用时:", time.Now().Sub(t), "返回结果：", string(result))
}
