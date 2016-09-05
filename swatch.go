package swatch

import (
	"time"
	"fmt"
)

const (
	Biel = "CET"
	BeatsPerSecond = 86.4
)

func Beats(t time.Time) float64 {
	tUTC := t.UTC()

	biel, err := time.LoadLocation(Biel)

	if err != nil {
		panic(err)
	}

	tBiel := tUTC.In(biel)

	secondsPastMidnight := tBiel.Hour() * 3600 + tBiel.Minute() * 60 + tBiel.Second()

	beats := float64(secondsPastMidnight) / BeatsPerSecond

	if beats > 1000.0 {
		beats -= 1000.0
	}

	return beats
}

func Swatch() string {
	return fmt.Sprintf("@%06.2f", Beats(time.Now()))
}
