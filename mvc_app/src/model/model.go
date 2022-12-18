package model

import (
	"database/sql"
	"fmt"
	"mvc_app/local_package/connection"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func PrintfModel() {
	con := connection.DatabaseConnect()
	defer con.Close()

	go insert(con)
	go printf()
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
}

func printf() {
	fmt.Println("HERE")
}

func insert(con *sql.DB) {
	insert, err := con.Query("INSERT INTO `view_list` (`id`, `content`) VALUES (6, 'something')")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer insert.Close()
	fmt.Println("OOO")
}
 