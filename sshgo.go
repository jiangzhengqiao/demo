package main

import (
	"bufio"
	"bytes"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"strconv"
)

// 自定义类型shuju_server_resources，列的首字母对应自定义类型，必须首字母大写
type ShuJuServerResources struct {
	Key, Port, Resource_type     int
	Ip, Username, Password, Info string
}

// 全局变量
var (
	DRIVER   = "mysql"
	USERNAME = "root"
	PASSWORD = "8.3.root"
	IP       = "192.168.11.140"
	PORT     = "3306"
	DBNAME   = "ProductData"
)

var (
	err error
	db  gorm.DB
)

// 主方法
func main() {
	// 初始化数据库连接
	db, err = db_handle()
	check(err, "连接数据库错误")

	// 查询数据库，返回数组
	ShuJuServerResource := find_resources()
	log.Println(len(*ShuJuServerResource))
	if len(*ShuJuServerResource) == 0 {
		log.Println("查询数据库无法获取配置")
		return
	}

	// 输出序号、IP、描述，并且封装成map
	fmt.Println("序号", "\t", "IP", "\t", "描述")
	// num := len(ShuJuServerResources)
	ips := make(map[string]ShuJuServerResources, 0)
	for _, v := range *ShuJuServerResource {
		if v.Key == 0 {
			return
		}
		fmt.Println(v.Key, "\t", v.Ip, "\t", v.Info)
		ips[strconv.Itoa(v.Key)] = v
	}

	// for i := 0; i < num; i++ {
	// 	serverR := arrs[i]
	// 	fmt.Println(serverR.key, "\t", serverR.ip, "\t", serverR.info)
	// 	ips[strconv.Itoa(serverR.key)] = serverR
	// }

	// 等待用户输入
	fmt.Print("请如果要编号：")
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()

	command := string(data)

	// 获取详细信息，连接
	serverResurce := ips[command]
	ssh_connect(serverResurce.Username, serverResurce.Password, serverResurce.Ip+":"+strconv.Itoa(serverResurce.Port))
}

// 查找资源
func find_resources() *[]ShuJuServerResources {
	// 查询IP列表
	var ShuJuServerResource []ShuJuServerResources
	db.Select("`key`, `ip`, `port`, `username`, `password`, `info`, `resource_type`").Table("ProductData.shuju_server_resources").Where("resource_type = 1").Find(&ShuJuServerResource)
	return &ShuJuServerResource
}

// 检查是否有错误
func check(err error, msg string) {
	if err != nil {
		log.Fatalf("%s error: %v", msg, err)
	}
}

// 连接数据库
func db_handle() (gorm.DB, error) {
	var buffer bytes.Buffer
	buffer.WriteString(USERNAME)
	buffer.WriteString(":")
	buffer.WriteString(PASSWORD)
	buffer.WriteString("@tcp(")
	buffer.WriteString(IP)
	buffer.WriteString(":")
	buffer.WriteString(PORT)
	buffer.WriteString(")/")
	buffer.WriteString(DBNAME)
	buffer.WriteString("?charset=utf8")
	db, err := gorm.Open(DRIVER, buffer.String())
	db.DB()
	db.DB().Ping()
	db.DB().SetMaxIdleConns(2)
	db.DB().SetMaxOpenConns(10)
	return db, err
}

// ssh 连接ip_port的格式例如：192.168.1.1:22
func ssh_connect(username, password, ip_port string) {
	client, err := ssh.Dial("tcp", ip_port, &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{ssh.Password(password)},
	})
	check(err, "dial")

	session, err := client.NewSession()
	check(err, "new session")
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	err = session.RequestPty("xterm", 25, 100, modes)
	check(err, "request pty")

	err = session.Shell()
	check(err, "start shell")

	err = session.Wait()
	check(err, "return")
}
