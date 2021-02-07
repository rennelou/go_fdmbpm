package types

import (
	"github.com/rennelou/go_fdmbpm/types/cmplxfp"
)

// FDMBPM ...
func FDMBPM(w Waveguide, eBoundary []complex128) [][]complex128 {

	result := make([][]complex128, w.ZSteps)
	result[0] = eBoundary

	ds := GetD(eBoundary, w.Q[0])

	for i := 1; i < w.ZSteps; i++ {
		abcs := w.Getabcs(i)
		es := GetRecurrenceForm(GetAlphasBetas(abcs, ds))
		ds = GetD(es, w.Q[i])

		result[i] = es
	}

	return result
}

// GetRecurrenceForm ...
func GetRecurrenceForm(alphasBetas []cmplxfp.AlphaBeta) []complex128 {
	rev := cmplxfp.ReverseAlphaBetas(alphasBetas)

	es := make([]complex128, 0)
	e := complex(0, 0)
	for _, value := range rev {
		e, es = func(lastE complex128, list []complex128) (complex128, []complex128) {

			_e := value.Alpha*lastE + value.Beta // okamoto 7.110

			return _e, append(list, _e)
		}(e, es)
	}

	es = cmplxfp.Reversecomplex128s(es)

	boundaryCondition1, boundaryCondition2 := complex(0, 0), complex(0, 0)

	firstElement := cmplxfp.Multiplycomplex128(cmplxfp.Headcomplex128(es), boundaryCondition1)
	es = append(firstElement, es...) // okamoto 7.106

	lastElement := cmplxfp.Multiplycomplex128(cmplxfp.Lastcomplex128(es), boundaryCondition2)
	es = append(es, lastElement...) // okamoto 7.105

	return es
}

// GetAlphasBetas ...
func GetAlphasBetas(abcs []ABC, ds []complex128) []cmplxfp.AlphaBeta {

	result := make([]cmplxfp.AlphaBeta, 0)
	alpha := complex(0, 0)
	beta := complex(0, 0)
	for i, abc := range abcs {
		alpha, beta, result = func(lastAlpha complex128, lastBeta complex128, list []cmplxfp.AlphaBeta) (complex128, complex128, []cmplxfp.AlphaBeta) {
			a := abc.A
			b := abc.B
			c := abc.C

			_alpha := c / (b - a*lastAlpha)                 // okamoto 7.112a
			_beta := (ds[i] + a*lastBeta) / (b - lastAlpha) // okamoto 7.112b

			return _alpha, _beta, append(list, cmplxfp.AlphaBeta{Alpha: _alpha, Beta: _beta})
		}(alpha, beta, result)
	}

	return result
}

// GetD es e qs precisam ter mesma dimensão
func GetD(es []complex128, qs []complex128) []complex128 {
	result := make([]complex128, 0)

	if len(es) == len(qs) && len(es) >= MINIMALSTEP {
		for i, q := range cmplxfp.DropLastcomplex128(cmplxfp.Restcomplex128(qs)) {

			result = append(result, es[i]+q*es[i+1]+es[i+2]) // okamoto 7.97
		}
	}
	return result
}
