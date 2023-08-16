package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vseriousv/blockchainkeys"
)

var Db *sql.DB

func main() {
	var wg sync.WaitGroup
	numRoutines := 10 // 可以根据你的硬件配置调整协程数量

	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func() {
			loop()
			wg.Done()
		}()
	}

	wg.Wait()
}

// 连接数据库
func init() {
	var err error
	Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/clipbd?charset=utf8")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	Db.SetMaxOpenConns(20)
	Db.SetMaxIdleConns(20)
}

// 循环生成地址和私钥 并插入数据库函数
func loop() {
	var num = 0 // 生成地址和私钥的数量
	for i := 0; i < 10000; i++ {
		bc, err := blockchainkeys.NewBlockchain(blockchainkeys.Tron)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		privateKey, _, address, err := bc.GenerateKeyPair()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// fmt.Println("Private Key:", privateKey)
		// fmt.Println("Address:", address)
		save(address, privateKey)
		num++
		fmt.Println("已生成:", num)
	}
}

// 保存地址和私钥到数据库函数 接受参数为地址和私钥
func save(address, privateKey string) {
	stmt, err := Db.Prepare("insert into address_map(address, private) values(?, ?)")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(address, privateKey)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
