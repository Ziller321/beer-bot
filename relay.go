package main

import (
	"fmt"
	"log"

	rpio "github.com/stianeikeland/go-rpio"
)

// Relay is struct for single relay
type Relay struct {
	pin  int
	isOn bool
}

var relays [2]Relay

// SetRelays will add controlable relays to array
func SetRelays() {
	fmt.Println("Set relays")

	relays[0] = Relay{
		2,
		false,
	}

	relays[1] = Relay{
		3,
		false,
	}

	// Initial state, all off

	for _, r := range relays {
		fmt.Println("initial turn off ")
		r.On()
	}

}

// LinkRelay will link relay to thermal sensor
func LinkRelay(sensor TempSensor, relayIndex int) {
	sensor.LinkedRelay = relays[relayIndex]
}

// On will turn relay on/high
func (r *Relay) On() {
	err := rpio.Open()
	defer rpio.Close()

	if err != nil {
		log.Fatal(err)
	}
	pin := rpio.Pin(r.pin)
	pin.Output()

	r.isOn = true
	pin.Low()

}

// Off will turn relay off/low
func (r *Relay) Off() {
	err := rpio.Open()
	defer rpio.Close()

	if err != nil {
		log.Fatal(err)
	}
	pin := rpio.Pin(r.pin)
	pin.Output()

	r.isOn = false
	pin.High()
}
