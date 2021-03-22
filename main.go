package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/nehal1992/Go-Clean-Architecture/httpframework"
	"github.com/nehal1992/Go-Clean-Architecture/repository"
	"github.com/nehal1992/Go-Clean-Architecture/repository/postrepository"
)

func main() {
	var muxFramwork httpframework.HttpFramework = httpframework.NewMuxHttp()
	fmt.Println("uncle bob clean Architecture main func")
	//connect to mysql db
	conn, dbclient := ConnectMysql()
	defer conn.Close()
	//connect with mongo client
	// dbclient := ConnectMongoDB()

	m := httpframework.NewHttpRoutesEngine(dbclient)
	muxFramwork.HandelRoutes(m)
}

func ConnectMysql() (db *sql.DB, mysql repository.PostRepo) {
	db, err := repository.ConnectMysql()
	if err != nil {
		log.Fatalf("error happen : %f", err)
	}
	mysql = postrepository.NewMysqlPost(db)
	return
}

func ConnectMongoDB() (mongo repository.PostRepo) {
	db, err := repository.ConnectMongoClient(context.Background(), "localhost", "27017")
	if err != nil {
		log.Fatalf("error happen : %f", err)
	}
	mongo = postrepository.NewMongoPost(db)
	return
}
