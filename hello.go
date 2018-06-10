package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Slice to hold sensor data
var tempSensors []TempSensor

func webServer() {
	// Start web server
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}

func embedServer() {

	// Relay
	SetRelays()

	// Find sensors
	FindSensors(&tempSensors)

	// Try linking
	LinkRelay(tempSensors[0], 0)
	LinkRelay(tempSensors[1], 1)

	// Run loop
	count := time.Tick(3 * time.Second)
	for range count {

		for index := range tempSensors {

			sensor := &tempSensors[index]
			ReadTemp(sensor)

			fmt.Println("Sensor: " + sensor.Name + " v:" + strconv.FormatFloat(sensor.Value, 'f', 2, 64) + " t:" + strconv.FormatFloat(sensor.Target, 'f', 2, 64))

		}
	}
}

func main() {
	fmt.Println("Hello from goBerry")

	go webServer()
	embedServer()
}

func hello(w http.ResponseWriter, r *http.Request) {

	testJSON, err := json.Marshal(&tempSensors)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(testJSON)

}
