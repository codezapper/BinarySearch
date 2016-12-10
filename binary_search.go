package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"search"
	"strconv"
	"strings"
)

type Related struct {
	Id   int64
	List []int64
}

func handler(w http.ResponseWriter, r *http.Request) {
	parameter, _ := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
	parameter64 := int64(parameter)
	ret := search.Find_in_sorted_file("test_data.csv", parameter64)

	related := Related{parameter64, ret}

	js, err := json.Marshal(related)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	signal_channel := make(chan os.Signal, 1)
	signal.Notify(signal_channel, os.Interrupt)
	go func() {
		for signal := range signal_channel {
			fmt.Printf("TEST %d\n", signal)
			os.Exit(0)
		}
	}()
	http.HandleFunc("/related/", handler)
	http.ListenAndServe(":8080", nil)
}
