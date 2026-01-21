package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type SearchResponse struct {
	Result      *Employee `json:"result"`
	Method      string    `json:"method"`
	TimeNano    int64     `json:"time_ns"`
	TimeMicro   float64   `json:"time_us"`
}


func searchHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	method := r.URL.Query().Get("method")
	n, _ := strconv.Atoi(r.URL.Query().Get("n"))

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", 500)
		return
	}
	defer db.Close()

	data, err := getEmployees(db, n)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	start := time.Now()

	var result *Employee
	if method == "iterative" {
		result = sequentialIterative(data, name)
	} else {
		result = sequentialRecursive(data, name, 0)
	}

	duration := time.Since(start)

	response := SearchResponse{
		Result:    result,
		Method:    method,
		TimeNano:  duration.Nanoseconds(),
		TimeMicro: float64(duration.Nanoseconds()) / 1000.0,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// serve frontend
	http.Handle("/", http.FileServer(http.Dir("../frontend")))

	// API endpoint
	http.HandleFunc("/search", searchHandler)

	http.ListenAndServe(":8080", nil)
}
