package types

// IValue ...
type IValue interface{}

// RecFunc ...
type RecFunc struct {
	kernel func(int, IValue) IValue
	index  int
	size   int
	value  IValue
}

// NewRecFunc ...
func NewRecFunc(size int, initialValue IValue, kernel func(int, IValue) IValue) RecFunc {
	return RecFunc{kernel, 0, size, initialValue}
}

// ExecRecFunc ...
func (p RecFunc) ExecRecFunc() interface{} {
	it := p
	for it.hasNext() {
		it = it.getNext()
	}

	return p.value
}

func (p RecFunc) hasNext() bool {
	return p.index < p.size-1
}

func (p RecFunc) getNext() RecFunc {
	value := p.kernel(p.index, p.value)
	nextIndex := p.index + 1
	next := RecFunc{p.kernel, nextIndex, p.size, value}

	return next
}
