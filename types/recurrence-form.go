package types

import (
	"github.com/rennelou/go_fdmbpm/types/complxtpl"
)

// FDMBPM ...
func FDMBPM(w Waveguide, dInitials []complex128) [][]complex128 {

	result := make([][]complex128, w.NZ)
	ds := dInitials
	for i := 0; i < w.NZ; i++ {
		abcs := w.Getabcs(i)
		es := GetRecurrenceForm(GetAlphasBetas(abcs, ds))
		ds = GetD(es, w.q[i])

		result[i] = es
	}

	return result
}

// GetRecurrenceForm ...
func GetRecurrenceForm(alphasBetas []complxtpl.ComplexTupla) []complex128 {
	rev := complxtpl.ReverseComplexTuplas(alphasBetas)

	eList := make([]complex128, 0)
	e := complex(0, 0)
	for _, value := range rev {
		e, eList = func(lastE complex128, list []complex128) (complex128, []complex128) {

			_e := value.Alpha*lastE + value.Beta // okamoto 7.110

			return _e, append(list, _e)
		}(e, eList)
	}

	boundaryCondition1, boundaryCondition2 := complex(0, 0), complex(0, 0)

	eList = append(eList, eList[len(eList)-1]*boundaryCondition1) // okamoto 7.106

	eList = complxtpl.Reversecomplex128s(eList)
	eList = append(eList, eList[len(eList)-1]*boundaryCondition2) // okamoto 7.105

	return complxtpl.Reversecomplex128s(eList)
}

// GetAlphasBetas ...
func GetAlphasBetas(abcs []ABC, ds []complex128) []complxtpl.ComplexTupla {

	result := make([]complxtpl.ComplexTupla, 0)
	alpha := complex(0, 0)
	beta := complex(0, 0)
	for i, abc := range abcs {
		alpha, beta, result = func(lastAlpha complex128, lastBeta complex128, list []complxtpl.ComplexTupla) (complex128, complex128, []complxtpl.ComplexTupla) {
			a := abc.A
			b := abc.B
			c := abc.C

			_alpha := c / (b - a*lastAlpha)                 // okamoto 7.112a
			_beta := (ds[i] + a*lastBeta) / (b - lastAlpha) // okamoto 7.112b

			return _alpha, _beta, append(list, complxtpl.ComplexTupla{Alpha: _alpha, Beta: _beta})
		}(alpha, beta, result)
	}

	return result
}

// GetD es e qs precisam ter mesma dimensão
func GetD(es []complex128, qs []complex128) []complex128 {
	var result []complex128
	for i, q := range complxtpl.DropLastcomplex128(complxtpl.Restcomplex128(qs)) {

		result = append(result, es[i]+q*es[i+1]+es[i+2]) // okamoto 7.97
	}
	return result
}
