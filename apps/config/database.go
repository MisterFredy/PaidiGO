package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

/*
 DB Driver
 1.host = "it's host your DB"

		  -localhost if your host in your pc
		  -192.168.1.3 if you have network connection Your IP.

 2.port = "port on your database"

		   -3306  Default MYSQL
		   -5432 Default Postgress
		   -27017  Default MongoDB

 3.user = username your database
 4.password = password your database
 5.dbname = your database name

 6.adapter = list adapter

			-mysql
			-mongo
			-postgre
			-sqlite
*/

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = ""
	dbname   = "ez-mags"
	adapter  = "mysql"
)

/*
ConnectionDB
* connection DB function
*/
func ConnectionDB() {
	if adapter == "mongo" {
		hostmongo := fmt.Sprintf("host=%s:port=%d ", host, port)
		info := &mgo.DialInfo{
			Addrs:    []string{hostmongo},
			Timeout:  60 * time.Second,
			Database: dbname,
			Username: user,
			Password: password,
		}
		err := info
		if err != nil {
			log.Fatalf("error on mongoDB", err)
		}
	} else if (adapter == "mysql") || (adapter == "postgre") || (adapter == "sqlite") {
		info := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

		db, err := sql.Open(adapter, info)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			panic(err)
		}
		fmt.Println("Successfully connected! MYSQL")
	}
}
