package go_fdmbpm

const _DX = 260 // micrometros
const _NX = 1024
const _DELTA_X int64 = _DX / _NX

const _DZ = 2048 // micrometros
const _NZ = 1024
const _DELTA_Z int64 = _DZ / _NZ

var s = make([][]complex128, _NX - 1)
var q = make([][]complex128, _NX - 1)

func init() {
    for i := 1; i < _NX; i++ {
        s[i] = make([]complex128, _NZ - 1)
        q[i] = make([]complex128, _NZ - 1)
    }
}

func FdmBpm() string {
    return "Vamo que vamo"
}

//TODO transformar getAlphas e getBetas em funções maps com getAlpha e getBeta
func getAlphas(s_m []complex128, size_x int) []complex128 {
    alphas := make([]complex128, size_x - 1)
    for index := 1; index < size_x; index++ {
        a, b, c := getCoefficients(s_m[index - 1], index)

        if (index == 1) {
            alphas[index - 1] = 1 / b
        } else {
            alphas[index - 1] = c / (b - a * alphas[index - 2])
        }
    }

    return alphas
}

func getBetas(s_m []complex128, d []complex128, alphas []complex128, size_x int) []complex128 {
    betas := make([]complex128, size_x - 1)
    for index := 1; index < size_x; index++ {
        a, b, _ := getCoefficients(s_m[index - 1], index)

        if (index == 1) {
            betas[index - 1] = d[index - 1] / b
        } else {
            betas[index - 1] = (d[index - 1] + a * betas[index - 2]) / (b - alphas[index - 2])
        }
    }

    return betas
}

func getCoefficients(s_m_i complex128, index int) (complex128, complex128, complex128) {
   a, c := complex(1, 0), complex(1, 0)

   boundaty_condition := complex(0, 0)
   b := s_m_i - boundaty_condition

   if (index == 1) {
      a = 0
   }
   if (index == _NX - 1) {
      c = 0
   }

   return a, b, c
}