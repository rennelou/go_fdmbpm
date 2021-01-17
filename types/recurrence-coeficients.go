package types

// ComplexTupla ...
type ComplexTupla struct {
	Alpha complex128
	Beta  complex128
}

// RecurrenceFormParams ...
type RecurrenceFormParams struct {
	LastValue ComplexTupla
	Values    []ComplexTupla
}

// RecurrenceFormFuncWrapper ...
type RecurrenceFormFuncWrapper struct {
	recFunc RecFunc
}

// NewRecurrenceFormFuncWrapper ...
func NewRecurrenceFormFuncWrapper(size int, initialAlpha complex128, initialBeta complex128, kernel func(int, IValue) IValue) RecurrenceFormFuncWrapper {
	c := ComplexTupla{initialAlpha, initialBeta}
	return RecurrenceFormFuncWrapper{NewRecFunc(size, RecurrenceFormParams{c, []ComplexTupla{c}}, kernel)}
}

// ExecRecFunc ...
func (r RecurrenceFormFuncWrapper) ExecRecFunc() []ComplexTupla {
	return r.recFunc.ExecRecFunc().(RecurrenceFormParams).Values
}
