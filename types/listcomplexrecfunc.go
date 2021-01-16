package types

type listComplexRecFunc struct {
	kernel    func(int, complex128, []complex128) (complex128, []complex128)
	index     int
	size      int
	lastValue complex128
	values    []complex128
}

// NewListComplexRecFunc ...
func NewListComplexRecFunc(size int, initialValue complex128, kernel func(int, complex128, []complex128) (complex128, []complex128)) listComplexRecFunc {
	return listComplexRecFunc{kernel, 0, size, initialValue, []complex128{initialValue}}
}

// ExecRecFunc ...
func (p listComplexRecFunc) ExecRecFunc() []complex128 {
	it := p
	for it.hasNext() {
		it = it.getNext()
	}
	result := it.getValue()

	return result
}

func (p listComplexRecFunc) hasNext() bool {
	return p.index < p.size-1
}

func (p listComplexRecFunc) getValue() []complex128 {
	return p.values
}

func (p listComplexRecFunc) getNext() listComplexRecFunc {
	value, values := p.kernel(p.index, p.lastValue, p.values)
	nextIndex := p.index + 1
	next := listComplexRecFunc{p.kernel, nextIndex, p.size, value, values}

	return next
}
