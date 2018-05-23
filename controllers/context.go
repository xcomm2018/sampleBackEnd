package controllers

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Context struct {
	DB *sql.DB
}

func DBInitial(*sql.DB) *sql.DB {

	var err error

	// db, err := sql.Open("mysql", "root:root@/ApotikDb")
	// db, err := sql.Open("mysql", "id5879709_xcommapotikdb:xcommApotikDb@tcp(xcomm2018.000webhostapp.com)/id5879709_apotikdb")
	db, err := sql.Open("mysql", "sql12239446:69Ey22hwA4@tcp(sql12.freemysqlhosting.net:3306)/sql12239446")
	if err != nil {
		fmt.Println("errorKoneksi")
		log.Fatal(err)

	}
	//fmt.Println(db)
	return db

}

// func NewContext() *Context {
// 	// untuk buat session
// 	// perintah untuk buat session
// 	//......
// 	return nil
// }
