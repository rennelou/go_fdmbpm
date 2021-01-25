package go_fdmbpm

import (
	"testing"

	"github.com/rennelou/go_fdmbpm/types"
	"github.com/rennelou/go_fdmbpm/types/complxtpl"
)

func TestRecurrenceForm(t *testing.T) {
	nx := 3
	w := types.NewWaveguide(0, nx, 0, 1)
	d := []complex128{1, 1, 1}
	zIndex := 0
	w.SInitialize(2)

	expected := []complxtpl.ComplexTupla{
		{Alpha: complex(0.5, 0), Beta: complex(0.5, 0)},
		{Alpha: complex(1/1.5, 0), Beta: complex(1, 0)},
		{Alpha: complex(0, 0), Beta: complex(2/(2-(1/1.5)), 0)},
	}

	got := types.GetAlphasBetas(w.Getabcs(zIndex), d)

	if len(got) < 1 {
		t.Errorf("got is empty")
	}

	for i := 0; i < len(got); i++ {
		if !got[i].IsEquals(expected[i]) {
			t.Errorf("got %v; expected %v", got, expected)
		}
	}
}
