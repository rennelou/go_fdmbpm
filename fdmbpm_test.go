package go_fdmbpm

import (
	"testing"

	"github.com/rennelou/go_fdmbpm/types"
	"github.com/rennelou/go_fdmbpm/types/complxtpl"
)

func TestRecurrenceForm(t *testing.T) {
	w := GetWaveguideMock(3)
	d := []complex128{1, 1, 1}
	zIndex := 0

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

func TestAbcArraysSizes(t *testing.T) {
	zIndex := 0

	for i := 5; i < 500; i++ {
		w := GetWaveguideMock(i)
		got := w.Getabcs(zIndex)

		if len(got) != w.NX-2 {
			t.Errorf("iteration %d have wrong dimension result", i)
		}
	}
}

func TestDArraysSizes(t *testing.T) {

	for i := 5; i < 500; i++ {
		var es []complex128
		var ds []complex128
		for j := 0; j < i; j++ {
			es = append(es, 1)
			ds = append(ds, 1)
		}

		got := types.GetD(es, ds)

		if len(got) != i-2 {
			t.Errorf("iteration %d have wrong dimension result", i)
		}
	}
}

func GetWaveguideMock(_nx int) types.Waveguide {
	w := types.NewWaveguide(0, _nx, 0, 1)
	w.SInitialize(2)

	return w
}
