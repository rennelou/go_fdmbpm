package types

import (
	"math"
)

const MINIMALSTEP = 5
const N0 = complex(1, 0)

// SlabWaveguide ...
type SlabWaveguide struct {
	XSteps int
	ZSteps int

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
func NewSlabWaveguide(_DX float64, _XDelta float64, _DZ float64, _ZDelta float64, k float64, n float64, alpha float64) SlabWaveguide {
	xSteps := int(math.Round(_DX / _XDelta))
	zSteps := int(math.Round((_DZ / _ZDelta)))

	_s := make([][]complex128, zSteps)
	_q := make([][]complex128, zSteps)

	for i := 0; i < zSteps; i++ {
		_s[i] = make([]complex128, xSteps)
		_q[i] = make([]complex128, xSteps)

		guidingSpace := complex(math.Sqrt(k)*math.Sqrt(_XDelta)*(math.Sqrt(n)-math.Sqrt(N0)), 0)
		freeSpace := complex(0, 4*k*N0*math.Sqrt(_XDelta)/_ZDelta)
		loss := complex(0, 2*k*N0*math.Sqrt((_XDelta)*math.Sqrt(alpha)))

		for j := 0; j < xSteps; j++ {
			_s[i][j] = complex(2, 0) - guidingSpace + freeSpace + loss  // okamoto 7.98
			_q[i][j] = complex(-2, 0) + guidingSpace + freeSpace - loss // okamoto 7.99
		}
	}

	return SlabWaveguide{
		XSteps: xSteps,
		ZSteps: zSteps,
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
		es := GetRecurrenceForm(GetAlphasBetas(abcs, ds))
		ds = GetD(es, w.Q[i])

		result[i] = es
	}

	return result
}

// Getabcs retorna vazio para todas as geometrias com menos de 5 steps
func (w SlabWaveguide) Getabcs(zIndex int) []ABC {
	boundaryCondition := complex(0, 0)
	result := make([]ABC, 0)

	if w.XSteps >= MINIMALSTEP {
		// okamoto 7.108a
		result = append(result, ABC{complex(0, 0), w.S[zIndex][0] - boundaryCondition, complex(1, 0)})

		for i := 2; i < w.XSteps-2; i++ {
			// okamoto 7.108b
			result = append(result, ABC{complex(1, 0), w.S[zIndex][i] - boundaryCondition, complex(1, 0)})
		}

		// okamoto 7.108c
		result = append(result, ABC{complex(1, 0), w.S[zIndex][w.XSteps-2] - boundaryCondition, complex(0, 0)})
	}

	return result
}
