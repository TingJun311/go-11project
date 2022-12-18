package connection

import (
	"database/sql"
	"log"
	"mvc_app/local_package/config"

	_ "github.com/go-sql-driver/mysql"
)

func DatabaseConnect() (*sql.DB) {
	conn, err := sql.Open(
		config.DBServer, 
		config.DBName + ":" + config.DBPassword + "@" + config.DBConnection + "(" + config.DBPort + ")/" + config.DBTableName)
	if err != nil {
		log.Println(err.Error())
	}
	
	return conn
}