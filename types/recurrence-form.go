package types

// GetAlphasBetas ...
func GetAlphasBetas(w Waveguide, d []complex128, indexZ int) []ComplexTupla {

	abcs := w.Getabcs(indexZ)

	result := make([]ComplexTupla, 0)
	alpha := complex(0, 0)
	beta := complex(0, 0)
	for i, abc := range abcs {
		alpha, beta, result = func(lastAlpha complex128, lastBeta complex128, list []ComplexTupla) (complex128, complex128, []ComplexTupla) {
			a := abc.A
			b := abc.B
			c := abc.C

			_alpha := c / (b - a*lastAlpha)
			_beta := (d[i] + a*lastBeta) / (b - lastAlpha)

			return _alpha, _beta, append(list, ComplexTupla{Alpha: _alpha, Beta: _beta})
		}(alpha, beta, result)
	}

	return result
}
