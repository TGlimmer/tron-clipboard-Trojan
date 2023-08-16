package initDB

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/clipbd?charset=utf8")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	createTables()
}

func createTables() {
	createTable("address_map")
	createTable("address_copy")
}

func createTable(tableName string) {
	query := `
		CREATE TABLE IF NOT EXISTS ` + tableName + ` (
			id INT(11) NOT NULL AUTO_INCREMENT,
			address VARCHAR(255) NOT NULL,
			private VARCHAR(255),
			PRIMARY KEY (id)
		) CHARSET=utf8;
	`

	_, err := Db.Exec(query)
	if err != nil {
		log.Panicln("Failed to create table", tableName, ":", err.Error())
	}
}
