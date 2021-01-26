package go_fdmbpm

import (
	"testing"

	"github.com/rennelou/go_fdmbpm/types"
	"github.com/rennelou/go_fdmbpm/types/complxtpl"
)

func TestFDMBPM(t *testing.T) {
	w := GetWaveguideMock(0, 5, 0, 1)
	eBoundary := GetOnes(5)
	zIndex := 0
	got := types.FDMBPM(w, eBoundary)

	expected := []complex128{0, 3.5, 5, 3.5, 0}
	if complxtpl.IsEquals(got[zIndex], expected) {
		t.Errorf("got %v, expected %v", got[zIndex], expected)
	}
}

func TestInsuficientSteps(t *testing.T) {
	w := GetWaveguideMock(0, 4, 0, 1)
	eBoundary := []complex128{1, 1, 1}
	zIndex := 0
	got := types.FDMBPM(w, eBoundary)

	if len(got[zIndex]) != 0 {
		t.Errorf("got should be empty")
	}
}

func TestAbcArraysSizes(t *testing.T) {
	zIndex := 0

	for i := 5; i < 500; i++ {
		w := GetWaveguideMock(0, i, 0, 1)
		got := w.Getabcs(zIndex)

		if len(got) != w.NX-2 {
			t.Errorf("iteration %d have wrong array size result", i)
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
			t.Errorf("iteration %d have wrong array size result", i)
		}
	}
}

func GetWaveguideMock(_dx int, _nx int, _dz int, _nz int) types.Waveguide {
	w := types.NewWaveguide(_dx, _nx, _dz, _nz)
	SInitialize(w, 2)
	QInitialize(w, 3)
	QBoundaryInitialize(w, 1)

	return w
}

func GetOnes(n int) []complex128 {
	result := make([]complex128, n)
	for i := 0; i < n; i++ {
		result[i] = 1
	}

	return result
}

// SInitialize ...
func SInitialize(w types.Waveguide, n float64) {
	for i := 0; i < w.NZ; i++ {
		for j := 0; j < w.NX; j++ {
			w.S[i][j] = complex(n, 0)
		}
	}
}

// QInitialize ...
func QInitialize(w types.Waveguide, n float64) {
	for i := 0; i < w.NZ; i++ {
		for j := 0; j < w.NX; j++ {
			w.Q[i][j] = complex(n, 0)
		}
	}
}

// QBoundaryInitialize ...
func QBoundaryInitialize(w types.Waveguide, n float64) {
	for i := 0; i < w.NX; i++ {
		w.QBoundary[i] = complex(n, 0)
	}
}
