package main

//borcaar/lorawan
//NOTE: NEEDS GCC IN ORDER TO COMPILE

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type DetectedEntry struct {
	ID        int
	EnterTime time.Time
	FilePath  string
}

const dbpath string = "SecuritySystem.db"

var db, err = sql.Open("sqlite3", dbpath)

// DUMMY DATA
var AllEntries = []DetectedEntry{
	{0, time.Date(2024, time.January, 1, 15, 37, 24, 45, time.UTC), ""},
	{1, time.Date(2024, time.January, 2, 10, 23, 55, 79, time.UTC), ""},
	{2, time.Date(2024, time.March, 5, 1, 0, 30, 42, time.UTC), ""},
	{3, time.Date(2024, time.July, 23, 16, 12, 5, 31, time.UTC), ""},
	{4, time.Date(2024, time.August, 15, 23, 59, 59, 99, time.UTC), ""}}

func main() {
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	api := mux.NewRouter().StrictSlash(true)
	api.HandleFunc("/RecentEntries", RecentEntries).Methods("GET")
	api.HandleFunc("/Entries/{id}", GetEntryByID).Methods("GET")
	api.HandleFunc("/CreateEntry", CreateEntry).Methods("POST")
	api.HandleFunc("/Entries/{id}", UpdateEntry).Methods("PUT")
	api.HandleFunc("/Entries/{id}", DeleteEntry).Methods("DELETE")
	//initialization notification
	println("API launced at", time.Now().GoString())
	log.Fatal(http.ListenAndServe(":10000", api))
}
