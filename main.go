package main

import (
	"database/sql"
	"errors"
	"fmt"
	movie3 "github.com/vageeshabr/cinema-service/internal/http/movie"
	movie2 "github.com/vageeshabr/cinema-service/internal/services/movie"
	"github.com/vageeshabr/cinema-service/internal/stores/movie"
	"log"
	"net/http"
)

func getDB() (*sql.DB, error) {
	//connections string user:pass@proto(host:port)/database
	conn := "root:@tcp(localhost:3306)/test"

	db, err := sql.Open("mysql", conn)

	if err != nil {
		log.Println("Failed to connect to db, err:", err)
		return nil, errors.New(fmt.Sprintln("Failed to connect to db, err:", err))
	}

	if err := db.Ping(); err != nil {
		log.Println("Failed to connect to db, err:", err)
		return nil, errors.New(fmt.Sprintln("Failed to connect to db, err:", err))
	}

	return db, nil
}

func main() {
	db, err := getDB()
	if err != nil {
		panic(err)
	}

	movieStore := movie.New(db)
	movieSvc := movie2.New(movieStore)
	movieHTTP := movie3.New(movieSvc)

	http.HandleFunc("/movies", movieHTTP.Index)

	http.ListenAndServe(":3000", nil)
}
