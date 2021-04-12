package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" //using mySQL driver
)

var (
	//Client :- connection to the database
	Client   *sql.DB
	username = "root"
	password = "mysql"
	host     = "mysqldb:3306"
	schema   = "users"
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, schema)
	var err error
	maxTry := 5
	for maxTry > 0 {
		Client, err = sql.Open("mysql", dataSourceName)
		if err != nil {
			maxTry--
			log.Println("Error While Connecting To The Database....retrying in 5 sec")
			time.Sleep(5 * time.Second)
		} else if err == nil {
			log.Println("Connection To The Database Is Successful....")
			break
		} else if maxTry == 1 {
			log.Fatalln("Unable To Connect To The Database....")
		}
	}
	maxTry = 5
	for maxTry > 5 {
		err = Client.Ping()
		if err != nil && maxTry == 1 {
			log.Fatalln("error while ping to the databse")
		} else if err != nil {
			log.Println("error while ping-ing....trying again in 5 sec")
			time.Sleep(5 * time.Second)
			maxTry -= 1
		} else {
			break
		}
	}
	err = makeTable()
	if err != nil {
		log.Fatalln("error while creating the table")
	}
	log.Println("Connected to the databse successfully")

}
