package go_fdmbpm

import (
	"testing"
)

func TestHello(t *testing.T) {
	f := NewRecurrenceForm(260, 1024, 2048, 1024)
	_ = f.GetAlphasBetas(make([]complex128, f.nx-1), 0)
}
