package mysql

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func getMySqlConnString() string {
	username := "root"
	password := "password"
	host := "localhost"
	port := "3306"
	db := "parkinglot"

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		username, password,
		host, port, db)
}

func GetMySqlConn() (dbConn *gorm.DB, err error) {
	connStr := getMySqlConnString()
	if dbConn, err = gorm.Open("mysql", connStr); err != nil {
		log.Fatal(err)
	}
	return
}
