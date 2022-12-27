package main

import (
	"encoding/json"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Data struct {
	Status struct {
		Water int `json:"water"`
		Wind  int `json:"wind"`
	} `json:"status"`
}

var templates = template.Must(template.ParseFiles("templates/index.html"))

func renderTemplate(w http.ResponseWriter, nav string, data map[string]interface{}) {
	err := templates.ExecuteTemplate(w, nav, data)
	// errs := templates.Execute(w, nav, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func weatherChecker(weather *Data) {
	rand.Seed(time.Now().UnixNano())
	for {
		weather.Status.Water = rand.Intn(100-1) + 1
		weather.Status.Wind = rand.Intn(100-1) + 1

		jsonData, err := json.Marshal(weather)
		if err != nil {
			log.Fatal("error:", err.Error())
		}

		err = os.WriteFile("data.json", jsonData, 0o644)
		if err != nil {
			log.Fatal("error: ", err.Error())
		}

		time.Sleep(15 * time.Second)
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fileData, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatal("error: ", err.Error())
	}

	var statusData Data

	err = json.Unmarshal(fileData, &statusData)
	if err != nil {
		log.Fatal("error: ", err.Error())
	}

	var (
		waterStatus string
		windStatus  string
	)

	switch {
	case statusData.Status.Water <= 5:
		waterStatus = "Aman"
	case statusData.Status.Water >= 6 && statusData.Status.Water <= 8:
		waterStatus = "Siaga"
	case statusData.Status.Water > 8:
		waterStatus = "Bahaya"
	default:
		waterStatus = "Tidak Ditemukan"
	}

	switch {
	case statusData.Status.Wind <= 6:
		windStatus = "Aman"
	case statusData.Status.Wind >= 7 && statusData.Status.Wind <= 15:
		windStatus = "Siaga"
	case statusData.Status.Wind > 15:
		windStatus = "Bahaya"
	default:
		windStatus = "Tidak Ditemukan"
	}

	data := map[string]interface{}{
		"waterStatus": waterStatus,
		"windStatus":  windStatus,
		"waterValue":  statusData.Status.Water,
		"windValue":   statusData.Status.Wind,
	}

	renderTemplate(w, "index", data)
}

func main() {
	var data Data
	go weatherChecker(&data)
	http.HandleFunc("/", statusHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
