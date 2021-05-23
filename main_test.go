package main

import (
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
