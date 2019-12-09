package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host = "mysql_users_host"
	mysql_users_scheme = "mysql_users_scheme"
)


var (
	Client *sql.DB
	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host = os.Getenv(mysql_users_host)
	scheme = os.Getenv(mysql_users_scheme)
)


func init(){
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, scheme)
	log.Printf("about to connect to %s", dataSourceName)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil{
		panic(err)
	}
	if err = Client.Ping(); err != nil{
		panic(err)
	}
	log.Println("database successfully configured")
}