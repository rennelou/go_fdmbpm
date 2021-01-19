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

// NewWaveguide ...
func NewWaveguide(_dx int, _nx int, _dz int, _nz int) Waveguide {
	_s := make([][]complex128, _nx)
	_q := make([][]complex128, _nx)

	for i := 0; i < _nx; i++ {
		_s[i] = make([]complex128, _nz)
		_q[i] = make([]complex128, _nz)
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
	for i := 0; i < w.NX; i++ {
		for j := 0; j < w.NZ; j++ {
			w.s[i][j] = complex(n, 0)
		}
	}
}

// Getabc ...
func (w Waveguide) Getabc(xIndex int, zIndex int) (complex128, complex128, complex128) {
	a, c := complex(1, 0), complex(1, 0)

	boundaryCondition := complex(0, 0)
	b := w.s[xIndex][zIndex] - boundaryCondition

	if xIndex == 0 {
		a = 0
	}
	if xIndex == w.NX-1 {
		c = 0
	}

	return a, b, c
}
