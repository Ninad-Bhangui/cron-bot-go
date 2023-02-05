package main

import (
	"testing"
    "time"
)

func TestSum(t *testing.T) {
	t.Run("difference of days between times", func(t *testing.T) {
		timeFormat := "2006-01-02"
		source, _ := time.Parse(timeFormat, "2022-02-05")
		dest, _ := time.Parse(timeFormat, "2022-02-07")

		got := timeDiff(source, dest)
		want := 2

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
