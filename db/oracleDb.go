package main

import (
	"database/sql"
	_ "github.com/mattn/go-oci8"
	"log"
)

func main() {
	sqlconn := "SWGD/easipass@192.168.131.52:1521/dev12c"
	db, err := sql.Open("oci8", sqlconn)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
}
