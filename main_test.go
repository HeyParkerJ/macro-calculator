package main

import (
	"os"
	"testing"
)

func TestRoundTwoSigFigs(t *testing.T) {
	t.Run("Rounds to two sig figs", func(t *testing.T) {
		var got = roundToTwoSigFigs(20.119)
		var want = 20.12

		if want != got {
			t.Errorf("it broke %f, %f", want, got)
		}
	})
}

func TestConvertArgsToVars(t *testing.T) {
	t.Run("Converts args to the vars we need", func(t *testing.T) {
		oldArgs := os.Args
		defer func() { os.Args = oldArgs }()
		os.Args = []string{"cmd", "192", "70", "30", "3"}
		var got1, got2, got3, got4 = convertArgsToVars()
		var want1, want2, want3, want4 float64 = 192, 70, 30, 3
		if got1 != want1 {
			t.Error("Failed Test, 1", got1, want1)
		}
		if got2 != want2 {
			t.Error("Failed Test, 2", got2, want2)
		}
		if got3 != want3 {
			t.Error("Failed Test, 3", got3, want3)
		}
		if got4 != want4 {
			t.Error("Failed Test, 4", got4, want4)
		}
	})
}
