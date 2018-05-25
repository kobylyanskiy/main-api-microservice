package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	pb "github.com/kobylyanskiy/dgraph-api/dgraph"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn
var err error
var DgraphServiceConnection pb.DgraphServiceClient
var ctx context.Context
var cancel context.CancelFunc

const (
	address     = "localhost:50051"
	defaultName = "world"
)

type Agent struct {
	Codename   string   `json:"codename"`
	Age        int      `json:"age"`
	Operations []string `json:"operations"`
}

type operation struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func get_agent(w http.ResponseWriter, r *http.Request) {

	var operations []pb.Operation

	if err := r.ParseForm(); err != nil {
		log.Println("Error parsing form")
	}
	get_operations, err := strconv.ParseBool(r.Form.Get("get_operations"))
	if err != nil {
		log.Println("Error parsing bool from string")
	}

	// Get operations from dgraph only if we need
	if get_operations {
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		get_operations_result, err := DgraphServiceConnection.GetOperations(ctx, &pb.Agent{Codename: "AGENT"})
		if err != nil {
			log.Fatalf("Could not get operations: %v", err)
		}
		log.Println("Result: %s", get_operations_result)
		// TODO remove it with the real data
		operations = append(operations, pb.Operation{Codename: "Codename"})
	}

	// Call Elastic API
	vars := mux.Vars(r)
	codename := vars["codename"]
	log.Println(codename)
	log.Println(operations)
	// resp, _ := http.Post("http://35.195.7.131/agents")
	// res1D := &agent{
	// Codename:   codename,
	// Age:        27,
	// Operations: []string{"apple", "peach", "pear"},
	// }
	// res1B, _ := json.Marshal(res1D)
	// w.Write(res1B)
	log.Println("[GET] /agents - OK")
}

func add_agent(w http.ResponseWriter, r *http.Request) {
	var agent Agent
	var error_buffer bytes.Buffer

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&agent)
	if err != nil {
		log.Fatalf("Couldn't decode Agent")
	}

	// Call Elastic API

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	dgraph_result, err := DgraphServiceConnection.AddAgent(ctx, &pb.Agent{
		Codename: agent.Codename,
	})
	if err != nil || dgraph_result.Result != true {
		log.Println("Cannot insert agent to the graph database: %v", err)
		error_buffer.WriteString("Cannot insert agent to the graph database")
	}

	result := pb.Result{
		Result:       true,
		ErrorMessage: error_buffer.String(),
	}
	res1B, _ := json.Marshal(result)
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
	// Set up a connection to the server.
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	DgraphServiceConnection = pb.NewDgraphServiceClient(conn)
	defer cancel()

	r := mux.NewRouter()
	r.HandleFunc("/agents/{codename}", update_agent).Methods("POST")
	r.HandleFunc("/agents/{codename}", get_agent).Methods("GET")
	r.HandleFunc("/agents", add_agent).Methods("POST")
	r.HandleFunc("/agents", get_agents).Methods("GET")

	log.Println("Start listening on 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}

//
