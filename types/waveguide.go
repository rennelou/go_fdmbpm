package types

const MINIMALSTEP = 5

// Waveguide ...
type Waveguide struct {
	DX int
	NX int
	DZ int
	NZ int

	S [][]complex128
	Q [][]complex128

	QBoundary []complex128
}

// ABC ...
type ABC struct {
	A complex128
	B complex128
	C complex128
}

// NewWaveguide ...
func NewWaveguide(_dx int, _nx int, _dz int, _nz int) Waveguide {
	_s := make([][]complex128, _nz)
	_q := make([][]complex128, _nz)

	for i := 0; i < _nz; i++ {
		_s[i] = make([]complex128, _nx)
		_q[i] = make([]complex128, _nx)
	}

	return Waveguide{
		DX: _dx,
		NX: _nx,
		DZ: _dz,
		NZ: _nz,
		S:  _s,
		Q:  _q,
	}
}

// Getabcs retorna vazio para todas as geometrias com menos de 5 steps
func (w Waveguide) Getabcs(zIndex int) []ABC {
	boundaryCondition := complex(0, 0)
	result := make([]ABC, 0)

	if w.NX >= MINIMALSTEP {
		// okamoto 7.108a
		result = append(result, ABC{complex(0, 0), w.S[zIndex][0] - boundaryCondition, complex(1, 0)})

		for i := 2; i < w.NX-2; i++ {
			// okamoto 7.108b
			result = append(result, ABC{complex(1, 0), w.S[zIndex][i] - boundaryCondition, complex(1, 0)})
		}

		// okamoto 7.108c
		result = append(result, ABC{complex(1, 0), w.S[zIndex][w.NX-2] - boundaryCondition, complex(0, 0)})
	}

	return result
}
