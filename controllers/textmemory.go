package controllers

import (
	"encoding/json"
	"log"
	"memorymaps-backend/config"
	"memorymaps-backend/db/postgres"
	"net/http"
	"strconv"
)

// TextMemory : Used to refer to text memories
type TextMemory struct {
	Text      string `json:"text"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// TextMemoryResponse : Used to send individual text memories
type TextMemoryResponse struct {
	TextMemory
	ID string `json:"id"`
}

// Create : Create Text Memory
func (t TextMemory) Create(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	// Expand the json attached in post request
	err := decoder.Decode(&t)
	if err != nil {
		log.Println(err)
		ThrowInternalErrAndExit(w)
		return
	}

	// Used for per user connection to DB
	dbconn, err := db.GetDBConn(config.DBName)
	defer dbconn.Close()

	if err != nil {
		log.Println(err)
		ThrowInternalErrAndExit(w)
		return
	}

	// Insert into DB
	stmt, _ := dbconn.Prepare(`INSERT INTO TextMemory(TextMem, Latitude, Longitude) VALUES($1,$2,$3);`)

	lat, err := strconv.ParseFloat(t.Latitude, 64)
	if err != nil {
		log.Println(err)
		ThrowInternalErrAndExit(w)
		return
	}

	long, err := strconv.ParseFloat(t.Longitude, 64)
	if err != nil {
		log.Println(err)
		ThrowInternalErrAndExit(w)
		return
	}

	_, execerr := stmt.Exec(t.Text, lat, long)

	if execerr != nil {
		// If execution err occurs then throw error
		log.Println(err)
		ThrowInternalErrAndExit(w)
		return
	}

	// If no error then give a success response
	RespondSuccessAndExit(w, "Text Memory Added Successfully")
}
