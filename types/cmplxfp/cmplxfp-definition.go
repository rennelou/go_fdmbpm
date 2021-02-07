package cmplxfp

import "math"

//go:generate gofp -destination complexfp.go -pkg cmplxfp -type "complex128"

//go:generate gofp -destination complxtplfp.go -pkg cmplxfp -type "AlphaBeta"
type AlphaBeta struct {
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

func Multiplycomplex128(l []complex128, c complex128) []complex128 {
	return Mapcomplex128(func(a complex128) complex128 {
		return a * c
	}, l)
}

func Headcomplex128(l []complex128) []complex128 {
	if len(l) < 1 {
		return make([]complex128, 0)
	}

	return []complex128{l[0]}
}

func Lastcomplex128(l []complex128) []complex128 {
	if len(l) < 1 {
		return make([]complex128, 0)
	}

	return []complex128{l[len(l)-1]}
}
