package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Person struct {
	Nama string `json:"nama"`
	Umur int    `json:"umur"`
}

type Respond struct {
	Msg        string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func HandlePerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}
	var p Person
	d := json.NewDecoder(r.Body)
	if err := d.Decode(&p); err != nil {
		log.Println(err)
	}
	var rsp Respond
	rsp.Msg = "Hello, " + p.Nama + ". Usia kamu; " + strconv.Itoa(p.Umur) + " Tahun."
	rsp.StatusCode = http.StatusOK
	e := json.NewEncoder(w)
	e.Encode(rsp)
}

func HandleHello(w http.ResponseWriter, r *http.Request) {
	var rsp Respond
	rsp.Msg = "Hello World"
	rsp.StatusCode = http.StatusOK
	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w)
	e.Encode(rsp)
}

func main() {
	http.Handle("/", http.HandlerFunc(HandleHello))
	http.Handle("/person", http.HandlerFunc(HandlePerson))
	http.ListenAndServe("0.0.0.0:8080", http.DefaultServeMux)
}
