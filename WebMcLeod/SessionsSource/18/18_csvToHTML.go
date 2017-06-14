package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Window is a struct that contains OHLC data for an asset time series, indexed by date
type Window struct {
	Asset string
	Date  time.Time
	Open  float64
	High  float64
	Low   float64
	Close float64
}

func main() {

	http.HandleFunc("/", dataService)

	port := ":8070"
	fmt.Println("Listening and Serving at port", port)
	http.ListenAndServe(port, nil)
}

// dataService function __________________________________________________

func dataService(w http.ResponseWriter, r *http.Request) {

	// csv parsing
	windows := csvObtainer("data.csv")

	// template parsing
	templat, err := template.ParseFiles("template1.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// template execution
	err = templat.Execute(w, windows)
	if err != nil {
		log.Fatalln(err)
	}
}

// csvObtainer function __________________________________________________

func csvObtainer(filepath string) []Window {

	src, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	reader := csv.NewReader(src)
	rows, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	windows := make([]Window, 0, len(rows))

	for i, row := range rows {
		if i == 0 {
			continue
		}

		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 32)
		high, _ := strconv.ParseFloat(row[2], 32)
		low, _ := strconv.ParseFloat(row[3], 32)
		close, _ := strconv.ParseFloat(row[4], 32)

		windows = append(windows, Window{
			Date:  date,
			Open:  open,
			High:  high,
			Low:   low,
			Close: close,
		})
	}

	return windows
}
