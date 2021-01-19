package go_fdmbpm

import (
	"testing"

	"github.com/rennelou/go_fdmbpm/types"
)

const _NX = 1024

func TestHello(t *testing.T) {
	f := types.NewRecurrenceForm(260, _NX, 2048, 1024)
	_ = f.GetAlphasBetas(make([]complex128, _NX-1), 0)
}
