package flights_test

import (
	"github.com/reactivejson/flaypath/.idea/internal/flights"
	"reflect"
	"testing"
)

func TestBuildGraph(t *testing.T) {
	tests := []struct {
		flights [][]string
		want    flights.Graph
	}{
		{
			flights: [][]string{
				{"SFO", "EWR"},
			},
			want: flights.Graph{
				"SFO": {"EWR"},
				"EWR": {},
			},
		},
		{
			flights: [][]string{
				{"ATL", "EWR"},
				{"SFO", "ATL"},
			},
			want: flights.Graph{
				"ATL": {"EWR"},
				"SFO": {"ATL"},
				"EWR": {},
			},
		},
		{
			flights: [][]string{
				{"IND", "EWR"},
				{"SFO", "ATL"},
				{"GSO", "IND"},
				{"ATL", "GSO"},
			},
			want: flights.Graph{
				"IND": {"EWR"},
				"SFO": {"ATL"},
				"GSO": {"IND"},
				"ATL": {"GSO"},
				"EWR": {},
			},
		},
	}
	for _, tt := range tests {
		got := flights.BuildGraph(tt.flights)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("BuildGraph(%v) = %v; want %v", tt.flights, got, tt.want)
		}
	}
}

func TestFindPath(t *testing.T) {
	tests := []struct {
		flights [][]string
		want    []string
	}{
		{
			flights: [][]string{
				{"SFO", "EWR"},
			},
			want: []string{"SFO", "EWR"},
		},
		{
			flights: [][]string{
				{"ATL", "EWR"},
				{"SFO", "ATL"},
			},
			want: []string{"SFO", "EWR"},
		},
		{
			flights: [][]string{
				{"IND", "EWR"},
				{"SFO", "ATL"},
				{"GSO", "IND"},
				{"ATL", "GSO"},
			},
			want: []string{"SFO", "EWR"},
		},
	}
	for _, tt := range tests {
		got := flights.FindPath(tt.flights)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FindPath(%v) = %v; want %v", tt.flights, got, tt.want)
		}
	}
}
