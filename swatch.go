// Package swatch offers utility functions for formatting time as Swatch Internet Beats.
package swatch

import (
	"fmt"
	"log"
	"time"
)

// Version is semver.
var Version = "0.0.2"

// Swatch Internet Time is relative to the Biel, Switzerland timezone, at a ratio of 1000 beats per 24 hour day.
const (
	Biel           = "CET"
	BeatsPerSecond = 86.4
)

// Beats formats a given time of day as an amount of Swatch Beats.
func Beats(t time.Time) float64 {
	tUTC := t.UTC()

	biel, err := time.LoadLocation(Biel)

	if err != nil {
		log.Panic(err)
	}

	tBiel := tUTC.In(biel)

	secondsPastMidnight := tBiel.Hour()*3600 + tBiel.Minute()*60 + tBiel.Second()

	beats := float64(secondsPastMidnight) / BeatsPerSecond

	if beats > 1000.0 {
		beats -= 1000.0
	}

	return beats
}

// Swatch formats the current time in Internet Time.
func Swatch() string {
	return fmt.Sprintf("@%06.2f", Beats(time.Now()))
}
