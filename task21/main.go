package main

import "fmt"

// MetricSystemEnjoyer describes a guy that
// tells distance in kilometers
type MetricSystemEnjoyer interface {
	TellDistanceInKilometers() float32
}

// MetricSystemGuy represents a person
// who calculates distance in metric system
type MetricSystemGuy struct {
	distance float32
}

// TellDistanceInKilometers method represents
// MS Enjoyer telling distance in kilometers
func (msg MetricSystemGuy) TellDistanceInKilometers() float32 {
	return msg.distance
}

// ImperialSystemFan represents a person
// who calculates distance in imperial system
type ImperialSystemFan struct {
	distance float32
}

// TellDistanceInMiles method represents
// IS fan telling distance in miles
func (isf ImperialSystemFan) TellDistanceInMiles() float32 {
	return isf.distance
}

// Traveler is a client type of MetricSystemGuy type
type Traveler struct {
	// ...
}

// Traveler wants to ask distance in kilometers
func (t Traveler) AskDistanceInKilometers(guy MetricSystemEnjoyer) {
	fmt.Printf("That guy told me I have %.1f kilometers left to go\n", guy.TellDistanceInKilometers())
}

// Adapter allows using ImperialSystemFan as MetricSystemEnjoyer
type Adapter struct {
	fan *ImperialSystemFan
}

// TellDistanceInKilometers makes Adapter to implement
// MetricSystemEnjoyer
func (a Adapter) TellDistanceInKilometers() float32 {
	return a.fan.distance * 1.6
}

func main() {
	john := &ImperialSystemFan{1} // John tells distance in miles
	pier := &MetricSystemGuy{1}   // Pier tells distance in kilometers
	adaptedJohn := &Adapter{john} // adapted John tells distance in kilometers

	simone := new(Traveler) // Simone is a traveller that understands only metric system

	simone.AskDistanceInKilometers(pier)
	simone.AskDistanceInKilometers(adaptedJohn)
}
