package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}
*/

func hello(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Printf("v T of err: %v, %T\n", err, err)
		fmt.Printf("v T of err.Error(): %v, %T\n", err.Error(), err.Error())
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func headers(w http.ResponseWriter, req *http.Request) {

	fmt.Printf("host: %s\n", req.Host)
	for name, headers := range req.Header {
		//fmt.Printf("name: %v %T\n", name, name)
		//fmt.Printf("headers: %v %T\n", headers, headers)
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
			//fmt.Fprintf(os.Stderr, "%v: %v\n", name, h)
		}
	}
}

func ybikeDataHandler(w http.ResponseWriter, req *http.Request) {
	var fieldCount = make(map[string]int)
	fieldCount["host_name"] = 0
	fieldCount["battery_no"] = 0
	fieldCount["battery_id"] = 0
	fieldCount["battery_life"] = 0
	fieldCount["charging_cycle"] = 0
	fieldCount["battery_cap"] = 0
	fieldCount["remain_cap"] = 0
	fieldCount["battery_status"] = 0
	fieldCount["charging_state"] = 0
	switch req.Method {
	case http.MethodGet:
		//fmt.Fprintf(w, "hello Get\n")
	case http.MethodPost:
		//fmt.Fprintf(w, "hello Post\n")
		// Process Request
		if err := req.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//fmt.Println("Post with PostForm: ", req.PostForm)
		//fmt.Println("Post with Data: ", req.Form) // url here!
		//fmt.Println("Post with Bosy: ", req.Body)
		for key, value := range req.Form {
			fmt.Printf("%s = %s\n", key, value[0])
			fieldCount[key] += 1
		}
		for key, value := range fieldCount {
			fmt.Printf("count of %s = %d\n", key, value)
		}

		// Process Response
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["return_code"] = "200"
		resp["return_msg"] = "Status_OK"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal, Err: %s\n", err)
		}
		w.Write(jsonResp)
	case http.MethodPut:
		//fmt.Fprintf(w, "hello Put\n")
	case http.MethodDelete:
		//fmt.Fprintf(w, "hello Delete\n")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func ybikeEventHandler(w http.ResponseWriter, req *http.Request) {
	var fieldCount = make(map[string]int)
	fieldCount["host_name"] = 0
	fieldCount["battery_no"] = 0
	fieldCount["battery_id"] = 0
	fieldCount["battery_event"] = 0
	fieldCount["time_stamp"] = 0
	switch req.Method {
	case http.MethodGet:
		//fmt.Fprintf(w, "hello Get\n")
	case http.MethodPost:
		//fmt.Fprintf(w, "hello Post\n")
		// Process Request
		if err := req.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//fmt.Println("Post with PostForm: ", req.PostForm)
		//fmt.Println("Post with Event: ", req.Form) // here!
		//fmt.Println("Post with Bosy: ", req.Body)
		for key, value := range req.Form {
			fmt.Printf("%s = %s\n", key, value[0])
			fieldCount[key] += 1
		}
		for key, value := range fieldCount {
			fmt.Printf("count of %s = %d\n", key, value)
		}

		// Process Response
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["return_code"] = "200"
		resp["return_msg"] = "Status_OK"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal, Err: %s\n", err)
		}
		w.Write(jsonResp)
	case http.MethodPut:
		//fmt.Fprintf(w, "hello Put\n")
	case http.MethodDelete:
		//fmt.Fprintf(w, "hello Delete\n")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/api/batteryData", ybikeDataHandler)
	http.HandleFunc("/api/batteryEvent", ybikeEventHandler)
	http.ListenAndServe(":8080", nil)
}
