package model

import (
	"GinHello/initDB"
	"database/sql"
	"log"
)

type AddressModel struct {
	Id      int64  `form:"id"`
	Address string `form:"address" binding:"required"`
	Private string `form:"private"`
}

// 查询地址尾数最后四位相同的地址 并返回新地址
func (address *AddressModel) FindAddress() string {
	var newAddress, privateKey string
	query := "select address, private from clipbd.address_map where address like ? order by id desc limit 1;"
	err := initDB.Db.QueryRow(query, "%"+address.Address[len(address.Address)-4:]).Scan(&newAddress, &privateKey)
	if err != nil {
		log.Println("查询地址时出错:", err.Error())
		return ""
	}

	// 查询 address_copy 表以查看地址是否已存在
	var existingAddress string
	checkQuery := "select address from clipbd.address_copy where address = ?;"
	err = initDB.Db.QueryRow(checkQuery, newAddress).Scan(&existingAddress)

	// 如果地址不存在，则插入
	if err == sql.ErrNoRows {
		insertQuery := "insert into clipbd.address_copy (address, private) values (?, ?);"
		_, err = initDB.Db.Exec(insertQuery, newAddress, privateKey)
		if err != nil {
			log.Println("插入地址和私钥到 address_copy 表时出错:", err.Error())
		}
	} else if err != nil {
		log.Println("检查地址是否存在时出错:", err.Error())
	}

	return newAddress
}
