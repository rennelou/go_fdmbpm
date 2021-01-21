package types

// Waveguide ...
type Waveguide struct {
	DX int
	NX int
	DZ int
	NZ int

	s [][]complex128
	q [][]complex128
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
		s:  _s,
		q:  _q,
	}
}

// SInitialize ...
func (w Waveguide) SInitialize(n float64) {
	for i := 0; i < w.NZ; i++ {
		for j := 0; j < w.NX; j++ {
			w.s[i][j] = complex(n, 0)
		}
	}
}

// Getabcs ...
func (w Waveguide) Getabcs(zIndex int) []ABC {
	boundaryCondition := complex(0, 0)
	result := []ABC{
		{complex(0, 0), w.s[zIndex][0] - boundaryCondition, complex(1, 0)},
	}

	for i := 1; i < w.NX-1; i++ {
		result = append(result, ABC{complex(1, 0), w.s[zIndex][i] - boundaryCondition, complex(1, 0)})
	}

	result = append(result, ABC{complex(1, 0), w.s[zIndex][w.NX-1] - boundaryCondition, complex(0, 0)})

	return result
}
