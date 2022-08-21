package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	_"github.com/jackc/pgconn"
	_"github.com/jackc/pgx/v4"
	_"github.com/jackc/pgx/v4/stdlib"
)

const webPort = "8082"
var counts int64 

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication services ...")

	//To connect to DB
	conn := ConnectToDB()
	if conn == nil {
		log.Panic("Cant connect to Postgres!!")
	}

	//Setting up config
	app := Config{
		DB: conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s",webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}


func OpenDB (dsn string) (*sql.DB,error) {
	//? phx is the driver that we imported for postgres database
	//? dsn is the database address 
	db,err := sql.Open("pgx",dsn)
	if err != nil {
		return nil,err
	}
	//? ping is just testing the database whether it is alive or not
	err = db.Ping()
	if err != nil {
		return nil,err
	}
	return db,nil
}

func ConnectToDB() *sql.DB {
	//? declared in the authentication-service in docker compose
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Panic("Couldnt fetch environment variables ...")
		return nil
	}

	//!connecting to database as long as it is ready to accept connections
	for {
		connection,err := OpenDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready ...")
			counts ++
		}else {
			log.Println("Connected to Postgres Successful!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}
		log.Println("Backing off for 4 seconds ...")
		time.Sleep(4 * time.Second)
		continue
	}
}