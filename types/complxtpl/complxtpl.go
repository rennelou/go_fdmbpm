package complxtpl

import "math"

//go:generate gofp -destination complexfp.go -pkg complxtpl -type "complex128"

//go:generate gofp -destination complxtplfp.go -pkg complxtpl -type "ComplexTupla"
type ComplexTupla struct {
	Alpha complex128
	Beta  complex128
}

// IsEquals ...
func (c1 ComplexTupla) IsEquals(c2 ComplexTupla) bool {
	return (complexAbs(c1.Alpha-c2.Alpha) < 1e-6) && complexAbs(c1.Beta-c2.Beta) < 1e-6
}

func complexAbs(x complex128) float64 { return math.Hypot(real(x), imag(x)) }
