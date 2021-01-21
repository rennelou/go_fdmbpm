package types

import (
	"github.com/rennelou/go_fdmbpm/types/complxtpl"
)

func GetRecurrenceForm(alphasBetas []complxtpl.ComplexTupla) []complex128 {
	rev := complxtpl.ReverseComplexTuplas(alphasBetas)

	eList := make([]complex128, 0)
	e := complex(0, 0)
	for _, value := range rev {
		e, eList = func(lastE complex128, list []complex128) (complex128, []complex128) {
			_e := value.Alpha*lastE + value.Beta

			return _e, append(list, _e)
		}(e, eList)
	}

	return complxtpl.Reversecomplex128s(eList)
}

// GetAlphasBetas ...
func GetAlphasBetas(w Waveguide, d []complex128, indexZ int) []complxtpl.ComplexTupla {

	abcs := w.Getabcs(indexZ)

	result := make([]complxtpl.ComplexTupla, 0)
	alpha := complex(0, 0)
	beta := complex(0, 0)
	for i, abc := range abcs {
		alpha, beta, result = func(lastAlpha complex128, lastBeta complex128, list []complxtpl.ComplexTupla) (complex128, complex128, []complxtpl.ComplexTupla) {
			a := abc.A
			b := abc.B
			c := abc.C

			_alpha := c / (b - a*lastAlpha)
			_beta := (d[i] + a*lastBeta) / (b - lastAlpha)

			return _alpha, _beta, append(list, complxtpl.ComplexTupla{Alpha: _alpha, Beta: _beta})
		}(alpha, beta, result)
	}

	return result
}
