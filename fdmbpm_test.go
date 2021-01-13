package go_fdmbpm

import (
	"testing"
)

func TestHello(t *testing.T) {
	expected := "Vamo que vamo"
	if got := FdmBpm(); got != expected {
		t.Errorf("Hello() = %q, Expected %q", got, expected)
	}
}