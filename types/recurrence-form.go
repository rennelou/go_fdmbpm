package types

type recurrenceFormParams struct {
	lastValue ComplexTupla
	values    []ComplexTupla
}

// GetAlphasBetas ...
func GetAlphasBetas(w Waveguide, d []complex128, indexZ int) []ComplexTupla {

	zeroTupla := ComplexTupla{Alpha: complex(0, 0), Beta: complex(0, 0)}
	initialValue := recurrenceFormParams{zeroTupla, []ComplexTupla{}}

	recFunc := NewRecFunc(w.NX, initialValue, func(index int, v IValue) IValue {
		r := v.(recurrenceFormParams)

		a, b, c := w.Getabc(index, indexZ)
		alpha := c / (b - a*r.lastValue.Alpha)
		beta := (d[index] + a*r.lastValue.Beta) / (b - r.lastValue.Alpha)

		newTupla := ComplexTupla{Alpha: alpha, Beta: beta}
		return recurrenceFormParams{lastValue: newTupla, values: append(r.values, newTupla)}
	})

	return recFunc.ExecRecFunc().(recurrenceFormParams).values
}
