package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/yinheli/qqwry"
	"net/http"
	"runtime"
	"strconv"

	"wdzj.com/bbs/dao"
	"wdzj.com/bbs/logger"
	"wdzj.com/bbs/supper"
)

var (
	q   *qqwry.QQwry
	err error
)

const (
	pageSize = 1000
)

func init() {
	q = qqwry.NewQQwry("qqwry.dat")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	r := mux.NewRouter()
	r.HandleFunc("/ip-user-nums/{p:[0-9]+}", GetIpUserNums)
	r.HandleFunc("/ip-users/{ip}", GetUsersByIp)
	// r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8888", nil)

}

func GetIpUserNums(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p := vars["p"]
	page, _ := strconv.Atoi(p)
	page = (page - 1) * pageSize
	ipusernums := dao.GetIpUserNum(strconv.Itoa(page), strconv.Itoa(pageSize))

	nums := make([]supper.IpUserNum, 0, len(ipusernums))
	if len(ipusernums) > 0 {
		for _, val := range ipusernums {
			q.Find(val.Ip)
			val.Addr = q.Country
			val.Url = "http://localhost:8888/ip-users/" + val.Ip
			nums = append(nums, val)
		}
	}

	if b, err := json.Marshal(nums); err == nil {
		w.Write(b)
	}

	logger.Info("转换完成。")
}

func GetUsersByIp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ip := vars["ip"]
	members := dao.GetUsersByIp(ip)
	if b, err := json.Marshal(members); err == nil {
		w.Write(b)
	}

	logger.Info("查询完成。")
}
