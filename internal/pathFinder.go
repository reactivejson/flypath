package pathFinder

// Tracker keeps track of flights and airports
type Tracker struct {
	flights          []Flight // slice of Flight objects
	startingAirports []string // slice of starting airports
	endingAirports   []string // slice of ending airports
}

// NewTracker creates a new Tracker with the given flights
func NewTracker(flights []Flight) *Tracker {
	// Create a new Tracker with the given flights
	tracker := &Tracker{
		flights: flights,
	}

	// Initialize the startingAirports and endingAirports slices with the same length as the flights slice
	startingAirports := make([]string, len(tracker.flights))
	endingAirports := make([]string, len(tracker.flights))
	tracker.startingAirports = startingAirports
	tracker.endingAirports = endingAirports

	return tracker
}

// PathFinder returns a Flight object representing the complete flight path from the starting airport to the ending airport
func (tracker *Tracker) PathFinder() Flight {
	// If there is only one flight, return that flight
	if len(tracker.flights) == 1 {
		return tracker.flights[0]
	}

	// Update the startingAirports and endingAirports slices
	tracker.setAirports()

	// Find the starting and ending airports that are not in both slices
	sourceAirport := DiffSlice(tracker.startingAirports, tracker.endingAirports)
	destinationAirport := DiffSlice(tracker.endingAirports, tracker.startingAirports)

	// Return a Flight object with the starting and ending airports
	return Flight{
		Source:      sourceAirport,
		Destination: destinationAirport,
	}
}

// setAirports sets the startingAirports and endingAirports slices based on the Source and Destination fields of the Flight objects in the flights slice
func (tracker *Tracker) setAirports() {
	for i, flight := range tracker.flights {
		tracker.startingAirports[i] = flight.Source
		tracker.endingAirports[i] = flight.Destination
	}
}

// DiffSlice returns the first string in the first slice that does not appear in the second slice
func DiffSlice(firstSlice, secondSlice []string) string {
	// Create a map of the strings in the second slice
	secondSliceMap := make(map[string]bool, len(secondSlice))
	for _, value := range secondSlice {
		secondSliceMap[value] = true
	}

	// Iterate through the strings in the first slice to find the first string that does not appear in the second slice
	for _, value := range firstSlice {
		if !secondSliceMap[value] {
			return value
		}
	}

	// If no string is found, return an empty string
	return ""
}
