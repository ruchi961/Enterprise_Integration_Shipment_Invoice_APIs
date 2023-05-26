package main

import (
    	"log"
    	"net/http"
	    "encoding/json"
	    "fmt"
	    "io/ioutil"
      "bytes"
)

type event struct {
PickupLocation string `json:"PickupLocation"`
DeliveryLocation string `json:"DeliveryLocation"`
}


func createEvent(w http.ResponseWriter, r *http.Request) {

  w.Header().Set("Content-Type", "application/json")

  // Change the response depending on the method being requested
  switch r.Method {
    case "GET":
      w.WriteHeader(http.StatusOK)
      w.Write([]byte(`{"message": "GET method requested"}`))
    case "POST":

	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter proper data")
	}
	
	json.Unmarshal(reqBody, &newEvent)
	//events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)
	newEventJSON, err := json.Marshal(newEvent)
	resp, err := http.Post("http://127.0.0.1:5000/", "application/json", bytes.NewBuffer(newEventJSON))
	//json.NewEncoder(w).Encode(newEvent)
  reqBody2, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%+v \n", newEvent)
  fmt.Print(string(reqBody2))
  w.Write([]byte(string(reqBody2)))

  default:
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "Can't find the requested method"}`))
    }



}

func main() {
http.HandleFunc("/shipmentSet",createEvent)
  log.Fatal(http.ListenAndServe(":8080",nil))
}
