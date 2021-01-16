package go_fdmbpm

import "github.com/rennelou/go_fdmbpm/types"

const _DX = 260 // micrometros
const _NX = 1024
const _DELTA_X int64 = _DX / _NX

const _DZ = 2048 // micrometros
const _NZ = 1024
const _DELTA_Z int64 = _DZ / _NZ

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
	alphas := getAlphas(0, _NX)
	_ = getBetas(make([]complex128, _NX-1), alphas, 0, _NX-1)

	return "Vamo que vamo"
}

func getAlphas(indexZ int, sizeX int) []complex128 {

	recGetAlpha := types.NewListComplexRecFunc(_NX, complex(0, 0), func(index int, lastValue complex128, alphas []complex128) (complex128, []complex128) {
		a, b, c := getCoefficients(indexZ, index)
		alpha := c / (b - a*lastValue)

		return alpha, append(alphas, alpha)
	})

	return recGetAlpha.ExecRecFunc()
}

func getBetas(d []complex128, alphas []complex128, indexZ int, sizeX int) []complex128 {

	recGetBeta := types.NewListComplexRecFunc(_NX, complex(0, 0), func(index int, lastValue complex128, betas []complex128) (complex128, []complex128) {
		a, b, _ := getCoefficients(indexZ, index)
		dI := d[toIndexX(index)]
		alpha := alphas[toIndexX(index-1)]

		beta := (dI + a*lastValue) / (b - alpha)

		return beta, append(betas, beta)
	})

	return recGetBeta.ExecRecFunc()
}

func getCoefficients(indexZ int, indexX int) (complex128, complex128, complex128) {
	a, c := complex(1, 0), complex(1, 0)

	boundaryCondition := complex(0, 0)
	b := s[toIndexX(indexX)][toIndexX(indexZ)] - boundaryCondition

	if indexX == 0 {
		a = 0
	}
	if indexX == _NX-1 {
		c = 0
	}

	return a, b, c
}

func toIndexX(index int) int {
	if index < 0 {
		return 0
	}

	if index > _NX-1 {
		return _NX - 1
	}

	return index
}

func toIndexZ(index int) int {
	if index < 0 {
		return 0
	}

	if index > _NZ-1 {
		return _NZ - 1
	}

	return index
}
