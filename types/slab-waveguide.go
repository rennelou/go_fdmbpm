package types

import (
	"math"
	"math/cmplx"

	"github.com/rennelou/go_fdmbpm/types/cmplxfp"
)

const MINIMALSTEP = 5

// SlabWaveguide ...
type SlabWaveguide struct {
	XSteps int
	ZSteps int

	xDelta float64

	kRight complex128
	kLeft  complex128

	S [][]complex128
	Q [][]complex128
}

// ABC ...
type ABC struct {
	A complex128
	B complex128
	C complex128
}

// NewSlabWaveguide ...
func NewSlabWaveguide(_DX float64, _XDelta float64, _DZ float64, _ZDelta float64,
	k float64, n float64, n0 float64, alpha float64, _kLeft complex128, _kRight complex128) SlabWaveguide {

	xSteps := int(math.Round(_DX / _XDelta))
	zSteps := int(math.Round((_DZ / _ZDelta)))

	_s := make([][]complex128, zSteps)
	_q := make([][]complex128, zSteps)

	for i := 0; i < zSteps; i++ {
		_s[i] = make([]complex128, xSteps)
		_q[i] = make([]complex128, xSteps)

		guidingSpace := complex(math.Sqrt(k)*math.Sqrt(_XDelta)*(math.Sqrt(n)-math.Sqrt(n0)), 0)
		freeSpace := complex(0, 4*k*n0*math.Sqrt(_XDelta)/_ZDelta)
		loss := complex(0, 2*k*n0*math.Sqrt((_XDelta)*math.Sqrt(alpha)))

		for j := 0; j < xSteps; j++ {
			_s[i][j] = complex(2, 0) - guidingSpace + freeSpace + loss  // okamoto 7.98
			_q[i][j] = complex(-2, 0) + guidingSpace + freeSpace - loss // okamoto 7.99
		}
	}

	return SlabWaveguide{
		XSteps: xSteps,
		ZSteps: zSteps,
		xDelta: _XDelta,
		kRight: _kRight,
		kLeft:  _kLeft,
		S:      _s,
		Q:      _q,
	}
}

// FDMBPM ...
func (w SlabWaveguide) FDMBPM(eInput []complex128) [][]complex128 {

	result := make([][]complex128, w.ZSteps)
	result[0] = eInput

	ds := GetD(eInput, w.Q[0])

	for i := 1; i < w.ZSteps; i++ {
		abcs := w.Getabcs(i)
		es := w.insertBoundaryValues(i, GetRecurrenceForm(GetAlphasBetas(abcs, ds)))
		ds = GetD(es, w.Q[i])

		result[i] = es
	}

	return result
}

func (w SlabWaveguide) insertBoundaryValues(z int, es []complex128) []complex128 {
	// okamoto 7.106
	firstElement := cmplxfp.Multiplycomplex128(cmplxfp.Headcomplex128(es), w.leftBoundary(z))

	// okamoto 7.105
	lastElement := cmplxfp.Multiplycomplex128(cmplxfp.Lastcomplex128(es), w.rightBoundary(z))

	return append(firstElement, append(es, lastElement...)...)
}

func (w SlabWaveguide) rightBoundary(z int) complex128 {
	kdeltaX := complex(0, -1) * w.kRight * complex(w.xDelta, 0)
	return cmplx.Exp(kdeltaX)
}

func (w SlabWaveguide) leftBoundary(z int) complex128 {
	kdeltaX := complex(0, -1) * w.kLeft * complex(w.xDelta, 0)
	return cmplx.Exp(kdeltaX)
}

// Getabcs retorna vazio para todas as geometrias com menos de 5 steps
func (w SlabWaveguide) Getabcs(zIndex int) []ABC {
	result := make([]ABC, 0)

	if w.XSteps >= MINIMALSTEP {
		// okamoto 7.108a
		result = append(result, ABC{complex(0, 0), w.S[zIndex][1] - w.leftBoundary(zIndex), complex(1, 0)})

		for i := 2; i < w.XSteps-2; i++ {
			// okamoto 7.108b
			result = append(result, ABC{complex(1, 0), w.S[zIndex][i], complex(1, 0)})
		}

		// okamoto 7.108c
		result = append(result, ABC{complex(1, 0), w.S[zIndex][w.XSteps-2] - w.rightBoundary(zIndex), complex(0, 0)})
	}

	return result
}
