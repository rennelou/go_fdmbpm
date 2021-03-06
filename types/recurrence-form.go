package types

import (
	"github.com/rennelou/go_fdmbpm/types/cmplxfp"
)

// GetRecurrenceForm ...
func GetRecurrenceForm(alphasBetas []cmplxfp.AlphaBeta) []complex128 {
	es := make([]complex128, 0)
	e := complex(0, 0)
	for _, value := range cmplxfp.ReverseAlphaBetas(alphasBetas) {
		e, es = func(lastE complex128, list []complex128) (complex128, []complex128) {

			_e := value.Alpha*lastE + value.Beta // okamoto 7.110

			return _e, append(list, _e)
		}(e, es)
	}

	return cmplxfp.Reversecomplex128s(es)
}

// GetAlphasBetas ...
func GetAlphasBetas(abcs []ABC, ds []complex128) []cmplxfp.AlphaBeta {

	result := make([]cmplxfp.AlphaBeta, 0)
	value := cmplxfp.AlphaBeta{Alpha: complex(0, 0), Beta: complex(0, 0)}
	for i, abc := range abcs {
		value, result = func(lastValue cmplxfp.AlphaBeta, list []cmplxfp.AlphaBeta) (cmplxfp.AlphaBeta, []cmplxfp.AlphaBeta) {
			a := abc.A
			b := abc.B
			c := abc.C

			_alpha := c / (b - a*lastValue.Alpha)                       // okamoto 7.112a
			_beta := (ds[i] + a*lastValue.Beta) / (b - lastValue.Alpha) // okamoto 7.112b

			newValue := cmplxfp.AlphaBeta{Alpha: _alpha, Beta: _beta}
			return newValue, append(list, newValue)
		}(value, result)
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
