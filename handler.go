package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func RecentEntries(write http.ResponseWriter, request *http.Request) {
	write.Header().Set("Content-Type", "application/json")
	json.NewEncoder(write).Encode(AllEntries)
}

func GetEntryByID(write http.ResponseWriter, request *http.Request) {
	write.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, i := range AllEntries {
		if strconv.Itoa(i.ID) == params["id"] {
			json.NewEncoder(write).Encode(i)
			return
		}
	}
	json.NewEncoder(write).Encode(nil)
}

func CreateEntry(write http.ResponseWriter, request *http.Request) {
	write.Header().Set("Content-Type", "application/json")
	var newEntry DetectedEntry
	_ = json.NewDecoder(request.Body).Decode(&newEntry)
	if t, _ := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z"); newEntry.EnterTime == t {
		newEntry.EnterTime = time.Now()
	}
	if newEntry.FilePath == "" {
		json.NewEncoder(write).Encode(nil)
		return
	}
	_, err = db.Exec(fmt.Sprintf("INSERT INTO Entries (Time, FootagePath) VALUES ('%s','%s');", newEntry.EnterTime.Format(time.RFC3339), newEntry.FilePath))
	if err == nil {
		json.NewEncoder(write).Encode(newEntry)
		return
	}
	json.NewEncoder(write).Encode(err)
}

func UpdateEntry(write http.ResponseWriter, request *http.Request) {
	write.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var newEntry DetectedEntry
	_ = json.NewDecoder(request.Body).Decode(&newEntry)
	for n, i := range AllEntries {
		if strconv.Itoa(i.ID) == params["id"] {
			AllEntries[n].EnterTime = newEntry.EnterTime
			json.NewEncoder(write).Encode(AllEntries[n])
			return
		}
	}
	json.NewEncoder(write).Encode(nil)
}

func DeleteEntry(write http.ResponseWriter, request *http.Request) {
	write.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for n, i := range AllEntries {
		if strconv.Itoa(i.ID) == params["id"] {
			json.NewEncoder(write).Encode(i)
			AllEntries = append(AllEntries[:n], AllEntries[n+1:]...)
			return
		}
	}
	json.NewEncoder(write).Encode(nil)
}
