package go_fdmbpm

import "github.com/rennelou/go_fdmbpm/types"

const _DX = 260 // micrometros
const _NX = 1024
const _DELTAX = _DX / _NX

const _DZ = 2048 // micrometros
const _NZ = 1024
const _DELTAZ = _DZ / _NZ

var s = make([][]complex128, _NX-1)
var q = make([][]complex128, _NX-1)

func init() {
	for i := 0; i < _NX-1; i++ {
		s[i] = make([]complex128, _NZ-1)
		q[i] = make([]complex128, _NZ-1)
	}
}

// FdmBpm ...
func FdmBpm() string {
	_ = getAlphasBetas(make([]complex128, _NX-1), 0, _NX)

	return "Vamo que vamo"
}

func getAlphasBetas(d []complex128, indexZ int, sizeX int) []types.ComplexTupla {

	recFunc := types.NewRecurrenceFormFuncWrapper(_NX, complex(0, 0), complex(0, 0), func(index int, v types.IValue) types.IValue {
		r := v.(types.RecurrenceFormParams)

		a, b, c := getabc(index, indexZ)
		alpha := c / (b - a*r.LastValue.Alpha)
		beta := (d[index] + a*r.LastValue.Beta) / (b - r.LastValue.Alpha)

		newTupla := types.ComplexTupla{alpha, beta}
		return types.RecurrenceFormParams{newTupla, append(r.Values, newTupla)}
	})

	return recFunc.ExecRecFunc()
}

func getabc(xIndex int, zIndex int) (complex128, complex128, complex128) {
	a, c := complex(1, 0), complex(1, 0)

	boundaryCondition := complex(0, 0)
	b := s[xIndex][zIndex] - boundaryCondition

	if xIndex == 0 {
		a = 0
	}
	if xIndex == _NX-1 {
		c = 0
	}

	return a, b, c
}
