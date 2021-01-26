package complxtpl

import "math"

//go:generate gofp -destination complexfp.go -pkg complxtpl -type "complex128"

//go:generate gofp -destination complxtplfp.go -pkg complxtpl -type "ComplexTupla"
type ComplexTupla struct {
	Alpha complex128
	Beta  complex128
}

// IsEquals ...
func IsEquals(c1 []complex128, c2 []complex128) bool {
	if len(c1) != len(c2) {
		return false
	}
	for i := 0; i < len(c1); i++ {
		if complexAbs(c1[i])-complexAbs(c2[i]) > 1e-6 {
			return false
		}
	}

	return true
}

func complexAbs(x complex128) float64 { return math.Hypot(real(x), imag(x)) }
