package go_fdmbpm

import "github.com/rennelou/go_fdmbpm/types"

// RecurrenceForm ...
type RecurrenceForm struct {
	dx int
	nx int
	dz int
	nz int

	s [][]complex128
	q [][]complex128
}

type recurrenceFormParams struct {
	lastValue types.ComplexTupla
	values    []types.ComplexTupla
}

// NewRecurrenceForm ...
func NewRecurrenceForm(_dx int, _nx int, _dz int, _nz int) RecurrenceForm {
	_s := make([][]complex128, _nx-1)
	_q := make([][]complex128, _nx-1)

	for i := 0; i < _nz-1; i++ {
		_s[i] = make([]complex128, _nz-1)
		_q[i] = make([]complex128, _nz-1)
	}

	return RecurrenceForm{
		dx: _dx,
		nx: _nx,
		dz: _dz,
		nz: _nz,
		s:  _s,
		q:  _q,
	}
}

// GetAlphasBetas ...
func (f RecurrenceForm) GetAlphasBetas(d []complex128, indexZ int) []types.ComplexTupla {

	zeroTupla := types.ComplexTupla{Alpha: complex(0, 0), Beta: complex(0, 0)}
	initialValue := recurrenceFormParams{zeroTupla, []types.ComplexTupla{zeroTupla}}

	recFunc := types.NewRecFunc(f.dx, initialValue, func(index int, v types.IValue) types.IValue {
		r := v.(recurrenceFormParams)

		a, b, c := f.getabc(index, indexZ)
		alpha := c / (b - a*r.lastValue.Alpha)
		beta := (d[index] + a*r.lastValue.Beta) / (b - r.lastValue.Alpha)

		newTupla := types.ComplexTupla{Alpha: alpha, Beta: beta}
		return recurrenceFormParams{lastValue: newTupla, values: append(r.values, newTupla)}
	})

	return recFunc.ExecRecFunc().(recurrenceFormParams).values
}

func (f RecurrenceForm) getabc(xIndex int, zIndex int) (complex128, complex128, complex128) {
	a, c := complex(1, 0), complex(1, 0)

	boundaryCondition := complex(0, 0)
	b := f.s[xIndex][zIndex] - boundaryCondition

	if xIndex == 0 {
		a = 0
	}
	if xIndex == f.nx-1 {
		c = 0
	}

	return a, b, c
}
