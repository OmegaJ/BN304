package main

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	newEntry.ID = AllEntries[len(AllEntries)-1].ID + 1
	AllEntries = append(AllEntries, newEntry)
	json.NewEncoder(write).Encode(newEntry)
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
