package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"product-api/cmd/data"
	"time"

	_ "github.com/lib/pq"
)

type Config struct {
	DB     *sql.DB
	Models data.Models
}

const webPort = "4000"

var counts int64

func main() {
	log.Println("Starting api on port", webPort)
	conn := connectToDB()
	if conn == nil {
		log.Fatal("Cannot connect to db")
	}

	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	// ping db
	err = db.Ping()
	if err != nil {
		_ = db.Close()
		return nil, err
	}
	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Cannot connect to db, retrying...")
			counts++
		} else {
			log.Println("Connected to db")
			return connection
		}
		if counts > 10 {
			log.Println(err)
			return nil
		}
		log.Println("Sleeping for 2 seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}
