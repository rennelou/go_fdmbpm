package types

import (
	"github.com/rennelou/go_fdmbpm/types/complxtpl"
)

// FDMBPM ...
func FDMBPM(w Waveguide, eBoundary []complex128) [][]complex128 {

	result := make([][]complex128, w.NZ)
	result[0] = eBoundary

	ds := GetD(eBoundary, w.Q[0])

	for i := 1; i < w.NZ; i++ {
		abcs := w.Getabcs(i)
		es := GetRecurrenceForm(GetAlphasBetas(abcs, ds))
		ds = GetD(es, w.Q[i])

		result[i] = es
	}

	return result
}

// GetRecurrenceForm ...
func GetRecurrenceForm(alphasBetas []complxtpl.ComplexTupla) []complex128 {
	rev := complxtpl.ReverseComplexTuplas(alphasBetas)

	es := make([]complex128, 0)
	e := complex(0, 0)
	for _, value := range rev {
		e, es = func(lastE complex128, list []complex128) (complex128, []complex128) {

			_e := value.Alpha*lastE + value.Beta // okamoto 7.110

			return _e, append(list, _e)
		}(e, es)
	}

	es = complxtpl.Reversecomplex128s(es)

	boundaryCondition1, boundaryCondition2 := complex(0, 0), complex(0, 0)
	es = append(complxtpl.Mapcomplex128(func(c complex128) complex128 {
		return c * boundaryCondition1
	}, head(es)), es...) // okamoto 7.106

	es = append(es, complxtpl.Mapcomplex128(func(c complex128) complex128 {
		return c * boundaryCondition2
	}, last(es))...) // okamoto 7.105

	return es
}

func head(l []complex128) []complex128 {
	if len(l) < 1 {
		return make([]complex128, 0)
	}

	return []complex128{l[0]}
}

func last(l []complex128) []complex128 {
	if len(l) < 1 {
		return make([]complex128, 0)
	}

	return []complex128{l[len(l)-1]}
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
	result := make([]complex128, 0)

	if len(es) == len(qs) && len(es) >= MINIMALSTEP {
		for i, q := range complxtpl.DropLastcomplex128(complxtpl.Restcomplex128(qs)) {

			result = append(result, es[i]+q*es[i+1]+es[i+2]) // okamoto 7.97
		}
	}
	return result
}
