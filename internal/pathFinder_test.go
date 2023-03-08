package pathFinder_test

/**
 * @author Mohamed-Aly Bou-Hanane
 * © 2023
 */

import (
	"testing"

	"github.com/stretchr/testify/assert"

	pathFinder "github.com/reactivejson/flaypath/internal"
)

func TestTrack_OneFlight(t *testing.T) {
	fl := []pathFinder.Flight{{Source: "SFO", Destination: "LAX"}}

	tr := pathFinder.NewTracker(fl)

	assert.Equal(t, fl[0], tr.PathFinder())
}

func TestTrack_MultipleFlights(t *testing.T) {
	fl := []pathFinder.Flight{
		{Source: "SFO", Destination: "LAX"},
		{Source: "LAX", Destination: "JFK"},
	}

	tr := pathFinder.NewTracker(fl)

	f := pathFinder.Flight{Source: "SFO", Destination: "JFK"}
	assert.Equal(t, f, tr.PathFinder())
}

func TestDiff(t *testing.T) {

	a := []string{"apple", "banana", "cherry"}
	b := []string{"apple", "cherry"}
	diff := pathFinder.DiffSlice(a, b)

	assert.Equal(t, "banana", diff)
}

func TestDiff_NoDiff(t *testing.T) {

	a := []string{"apple", "banana", "cherry"}
	b := []string{"apple", "banana", "cherry"}
	diff := pathFinder.DiffSlice(a, b)

	assert.Empty(t, diff)
}

func TestDiff_EmptyB(t *testing.T) {

	a := []string{"apple", "banana", "cherry"}
	b := []string{}
	diff := pathFinder.DiffSlice(a, b)

	assert.Equal(t, "apple", diff)
}

func TestDiff_EmptyA(t *testing.T) {

	a := []string{}
	b := []string{"apple", "banana", "cherry"}
	diff := pathFinder.DiffSlice(a, b)

	assert.Empty(t, diff)
}
