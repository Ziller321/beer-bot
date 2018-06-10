package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
)

var tempDeviceFolder = "/sys/bus/w1/devices/"
var tempFileName = "w1_slave"

// TempSensor is interface for sensor
type TempSensor struct {
	Name        string  `json: name`
	RawValue    int     `json: rawValue`
	Value       float64 `json: value`
	LinkedRelay Relay   `json: relay`
	Target      float64 `json: target`
	Tolerance   float64 `json: tolerance`
}

// FindSensors will find temp sensors from Pi file system
func FindSensors(tempSensors *[]TempSensor) {

	files, err := ioutil.ReadDir(tempDeviceFolder)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		sensorRegexp := regexp.MustCompile("^28.*$")

		if sensorRegexp.MatchString(f.Name()) {
			// register sensor
			*tempSensors = append(*tempSensors, TempSensor{
				f.Name(),
				0,
				0.0,
				Relay{},
				20,
				0.0,
			},
			)
		}
	}
}

// ReadTemp will read temp as int from sensor path
func ReadTemp(sensor *TempSensor) {

	file, err := os.Open(tempDeviceFolder + sensor.Name + "/" + tempFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Regex
	var validTemp = regexp.MustCompile(`t=([0-9]*$)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if validTemp.MatchString(scanner.Text()) {
			m := validTemp.FindStringSubmatch(scanner.Text())
			if l := len(m); l > 1 {
				temp, err := strconv.Atoi(m[1])
				if err != nil {
					log.Fatal(err)
				}

				sensor.RawValue = temp
				sensor.Value = float64(temp) / 1000

				// check if relay needs to change
				if sensor.Value < sensor.Target-sensor.Tolerance {
					sensor.LinkedRelay.On()
				} else if sensor.Value > sensor.Target+sensor.Tolerance {
					sensor.LinkedRelay.Off()
				} else {
					sensor.LinkedRelay.On()
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
