package go_fdmbpm

import (
	"math"
	"testing"

	"github.com/Arafatk/glot"
	"github.com/rennelou/go_fdmbpm/types"
	"github.com/rennelou/go_fdmbpm/types/cmplxfp"
)

func TestPlot(t *testing.T) {
	w := types.NewSlabWaveguide(10, 2, 10, 5, 1/1550.0, 3.4757, 1, 0.2, complex(1, 0), complex(1, 0))

	eBoundary := GetOnes(int(math.Round(10 / 2)))
	got := w.FDMBPM(eBoundary)

	plot, _ := glot.NewPlot(2, false, false)
	plot.AddPointGroup("eletric filed", "point", got)

	plot.SavePlot("fdmbpm.png")
}

func TestFDMBPM(t *testing.T) {
	w := types.NewSlabWaveguide(10, 2, 10, 5, 1/1550.0, 3.4757, 1, 0.2, complex(1, 0), complex(1, 0))

	eBoundary := GetOnes(int(math.Round(10 / 2)))
	got := w.FDMBPM(eBoundary)

	expected := []complex128{0, 3.5, 5, 3.5, 0}
	if !cmplxfp.IsEquals(got[1], expected) {
		t.Errorf("got %v, expected %v", got[1], expected)
	}
}

func TestInsuficientSteps(t *testing.T) {
	w := GetWaveguideMock(12, 4, 2.0, 1.0, 1/1550.0, 3.4757, 1, 0.2, complex(0, 0), complex(0, 0))
	eBoundary := GetOnes(4)
	got := w.FDMBPM(eBoundary)

	if len(got[1]) != 0 {
		t.Errorf("got should be empty")
	}
}

func TestAbcArraysSizes(t *testing.T) {
	zIndex := 0

	for i := 1; i < 10; i++ {
		w := GetWaveguideMock(100, float64(i), 2.0, 1.0, 1550.0, 3.4757, 1, 0.2, complex(0, 0), complex(0, 0))
		got := w.Getabcs(zIndex)

		if len(got) != w.XSteps-2 {
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

func GetWaveguideMock(_dx float64, _xSteps float64, _dz float64, _zSteps float64, k float64,
	n float64, n0 float64, alpha float64, _kLeft complex128, _kRight complex128) types.SlabWaveguide {

	w := types.NewSlabWaveguide(_dx, _xSteps, _dz, _zSteps, k, n, n0, alpha, _kLeft, _kRight)

	return w
}

func GetOnes(n int) []complex128 {
	result := make([]complex128, n)
	for i := 0; i < n; i++ {
		result[i] = 1
	}

	return result
}
