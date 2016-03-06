package controllers

import (
	"encoding/json"
	"log"
	"memorymaps-backend/config"
	"memorymaps-backend/db/postgres"
	"net/http"
	"strconv"
	"time"
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
	ID           string `json:"id"`
	CreationTime string `json:"creationtime"`
}

// TextMemoriesResponse is used to send back a reply of visitors array
type TextMemoriesResponse struct {
	TextMemories TextMemories `json:"memories"`
}

// TextMemories is an array of memories
type TextMemories []TextMemoryResponse

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

// GetAll : Used to get all events from DB
func (t TextMemory) GetAll(w http.ResponseWriter, r *http.Request) {

	var (
		id           string
		text         string
		lat          float64
		long         float64
		creationtime time.Time
	)
	// Used for per user connection to DB
	dbconn, err := db.GetDBConn(config.DBName)
	defer dbconn.Close()

	if err != nil {
		log.Println(err)
		ThrowInternalErrAndExit(w)
		return
	}

	rows, err := dbconn.Query("SELECT * FROM TextMemory")
	if err != nil {
		log.Println(err)
		ThrowInternalErrAndExit(w)
		return
	}

	var textMemArr TextMemories

	for rows.Next() {

		// Scan all the values for the given row
		rows.Scan(&id, &text, &lat, &long, &creationtime)

		// Convert lat long to string
		latStr := strconv.FormatFloat(lat, 'E', -1, 64)
		longStr := strconv.FormatFloat(long, 'E', -1, 64)
		// Create a visitor object
		TextMemoryResp := TextMemoryResponse{
			TextMemory{text, latStr, longStr},
			id,
			creationtime.Format(time.RFC3339)}

		textMemArr = append(textMemArr, TextMemoryResp)

	}

	memories := TextMemoriesResponse{
		TextMemories: textMemArr,
	}

	jsonResponse, err := json.Marshal(memories)
	if err != nil {

		log.Println(err)
		ThrowInternalErrAndExit(w)
		return
	}

	// Append the data to response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}
