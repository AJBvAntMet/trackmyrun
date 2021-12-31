package main

import (
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	const shortForm = "2006-Jan-02"
	dateGot, _ := time.Parse(shortForm, "2013-Feb-03")
	dateWant := "2013-Feb-03"

	cases := []struct {
		Description string
		Run         Run
		DistWant    string
		TimeWant    RunTime
		PaceWant    string
		DateWant    string
	}{
		{"0Km Run", Run{dateGot, 0, RunTime{0, 4, 1}}, "0km", RunTime{0, 4, 1}, "Invalid Distance", dateWant},
		{"-10Km Run", Run{dateGot, -10, RunTime{0, 4, 1}}, "-10km", RunTime{0, 4, 1}, "Invalid Distance", dateWant},
		{"1Km Run", Run{dateGot, 1, RunTime{0, 4, 1}}, "1km", RunTime{0, 4, 1}, "4.02", dateWant},
		{"5Km Run", Run{dateGot, 5, RunTime{0, 34, 1}}, "5km", RunTime{0, 34, 1}, "6.80", dateWant},
		{"5Km Run", Run{dateGot, 5.42, RunTime{0, 34, 52}}, "5.42km", RunTime{0, 34, 52}, "6.43", dateWant},
		{"10Km Run", Run{dateGot, 10, RunTime{1, 4, 1}}, "10km", RunTime{1, 4, 1}, "6.40", dateWant},
		{"100Km Run", Run{dateGot, 100, RunTime{36, 4, 1}}, "100km", RunTime{36, 4, 1}, "21.64", dateWant},
	}

	for _, test := range cases {
		t.Run(test.Description+"Date", func(t *testing.T) {
			got := test.Run.GetRunDate()
			if got != test.DateWant {
				t.Errorf("got %q want %v", got, test.DateWant)
			}
		})
		t.Run(test.Description+"Dist", func(t *testing.T) {
			got := test.Run.GetRunDistanceKm()
			if got != test.DistWant {
				t.Errorf("got %q want %q", got, test.DistWant)
			}
		})
		t.Run(test.Description+"Time", func(t *testing.T) {
			got := test.Run.GetRunTime()
			if got != test.TimeWant {
				t.Errorf("got %v want %v", got, test.TimeWant)
			}
		})
		t.Run(test.Description+"Pace", func(t *testing.T) {
			got := test.Run.GetRunPace()
			if got != test.PaceWant {
				t.Errorf("got %q want %q", got, test.PaceWant)
			}
		})
	}
}
