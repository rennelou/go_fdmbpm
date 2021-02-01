package types

import (
	"math"
)

const MINIMALSTEP = 5
const N0 = complex(1, 0)

// Waveguide ...
type Waveguide struct {
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

// NewWaveguide _uDX, _uDZ, _uXDelta, _uZDelta em micrometro
func NewWaveguide(_uDX float64, _uXDelta float64, _uDZ float64, _uZDelta float64, k float64, n float64, alpha float64) Waveguide {
	xSteps := int(math.Round(_uDX / _uXDelta))
	zSteps := int(math.Round((_uDZ / _uZDelta)))

	nXDelta := _uXDelta * 1000
	nZDelta := _uZDelta * 1000

	guidingSpace := complex(math.Sqrt(k)*math.Sqrt(nXDelta)*(math.Sqrt(n)-math.Sqrt(N0)), 0)
	freeSpace := complex(0, 4*k*N0*math.Sqrt(nXDelta)/nZDelta)
	loss := complex(0, 2*k*N0*math.Sqrt((nXDelta)*math.Sqrt(alpha)))

	_s := make([][]complex128, zSteps)
	_q := make([][]complex128, zSteps)

	for i := 0; i < zSteps; i++ {
		_s[i] = make([]complex128, xSteps)
		_q[i] = make([]complex128, xSteps)

		for j := 0; j < xSteps; j++ {
			_s[i][j] = complex(2, 0) - guidingSpace + freeSpace + loss  // okamoto 7.98
			_q[i][j] = complex(-2, 0) + guidingSpace + freeSpace - loss // okamoto 7.99
		}
	}

	return Waveguide{
		XSteps: xSteps,
		ZSteps: zSteps,
		S:      _s,
		Q:      _q,
	}
}

// Getabcs retorna vazio para todas as geometrias com menos de 5 steps
func (w Waveguide) Getabcs(zIndex int) []ABC {
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
