package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type agent struct {
	Codename   string
	Age        int
	Operations []string
}

type operation struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func get_agent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	codename := vars["codename"]
	resp, _ := http.Get("http://mongo-api:5000/agents")
	fmt.Println(resp)
	res1D := &agent{
		Codename:   codename,
		Age:        27,
		Operations: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	w.Write(res1B)
	log.Println("[GET] /agents - OK")
}
func add_agent(w http.ResponseWriter, r *http.Request) {
	res1D := &agent{
		Codename:   "Agent007",
		Age:        27,
		Operations: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	w.Write(res1B)
	log.Println("[POST] /agents - OK")
}
func update_agent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var buffer bytes.Buffer

	codename := vars["codename"]

	buffer.WriteString("updating agent ")
	buffer.WriteString(codename)

	w.Write(buffer.Bytes())
	log.Println("[POST] /agents/{codename} - OK")
}

func get_agents(w http.ResponseWriter, r *http.Request) {
	res1B, _ := json.Marshal("Get all agents")
	w.Write(res1B)
	log.Println("[GET] /agents/{codename} - OK")
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/agents/{codename}", update_agent).Methods("POST")
	r.HandleFunc("/agents/{codename}", get_agent).Methods("GET")
	r.HandleFunc("/agents", add_agent).Methods("POST")
	r.HandleFunc("/agents", get_agents).Methods("GET")

	log.Println("Start listening on 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
