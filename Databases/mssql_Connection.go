package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

var server string = ""
var user string = ""
var pass string = ""
var database string = ""

func getDBConnection(addr string, unam string, pwrd string, db string) *sql.DB {
	var connString string = fmt.Sprintf("server=%s;user id=%s;password=%s;port=1433;database=%s", addr, unam, pwrd, db)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Failed to open connection: ", err.Error())
	}
	ctx := context.Background()
	err = conn.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	return conn
}

// func closeDBconnection(conn *sql.DB){
// 	conn.Close()
// }

func main() {
	var dbConn *sql.DB = getDBConnection(server, user, pass, database)
	dbConn.Close()
}
