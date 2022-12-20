package R

// (c) Christian Maurer   v. 221213 - license see µU.go

import (
  "math"
  "sort"
  "µU/obj"
  "µU/char"
  "µU/str"
)
const
  FN = 6 // maximal length of function names
var (
  name = [NFunctions]string {"",
                             "sqr", "sqrt",
                             "exp", "exp10", "exp2", "log", "ld", "lg",
                             "sin", "cos", "tan", "cot",
                             "arcsin", "arccos", "arctan", "arccot",
                             "sinh", "cosh", "tanh", "coth",
                             "arsinh", "arcosh", "artanh", "arcoth",
                             "f", "f'", "f\"", "g", "g'", "g\"", "h", "h'", "h\""}
  inv [NFunctions]Function
  deriv [NFunctions]string
)

func init() {
  for f := Function(0); f < NFunctions; f++ {
    inv[f] = Undef
  }
  inv[Sqr] =    Sqrt
  inv[Sqrt] =   Sqr
  inv[Exp] =    Log
  inv[Exp10] =  Lg
  inv[Exp2] =   Ld
  inv[Log] =    Exp
  inv[Ld] =     Exp2
  inv[Lg] =     Exp10
  inv[Sin] =    Arcsin
  inv[Cos] =    Arccos
  inv[Tan] =    Arctan
  inv[Cot] =    Arccot
  inv[Arcsin] = Sin
  inv[Arccos] = Cos
  inv[Arctan] = Tan
  inv[Arccot] = Cot
  inv[Sinh] =   Arsinh
  inv[Cosh] =   Arcosh
  inv[Tanh] =   Artanh
  inv[Coth] =   Arcoth
  inv[Arsinh] = Sinh
  inv[Arcosh] = Cosh
  inv[Artanh] = Tanh
  inv[Arcoth] = Coth
  inv[Arcoth] = Coth

  deriv[Sqrt] =   "0.5*sqrt(x)"
  deriv[Sqr] =    "2*x"
  deriv[Exp] =    "exp(x)"
  deriv[Exp10] =  "log(10)*10*x"
  deriv[Exp2] =   "log(2)*2*x"
  deriv[Log] =    "1/x"
  deriv[Ld] =     "1/log(2)/x"
  deriv[Lg] =     "1/log(10)/x"
  deriv[Sin] =    "cos(x)"
  deriv[Cos] =    "-sin(x)"
  deriv[Tan] =    "1/cos(x)^2"
  deriv[Cot] =    "-1/sin(x)^2"
  deriv[Arcsin] = "1/sqrt(1-x^2)"
  deriv[Arccos] = "-1/sqrt(1-x^2)"
  deriv[Arctan] = "1/(1+x^2)"
  deriv[Arccot] = "-1/(1+x^2)"
  deriv[Sinh] =   "cosh(x)"
  deriv[Cosh] =   "sinh(x)"
  deriv[Tanh] =   "1/cosh(x)^2"
  deriv[Coth] =   "-1/sinh(x)^2"
  deriv[Arsinh] = "1/sqrt(x^2+1)"
  deriv[Arcosh] = "1/sqrt(x^2-1)"
  deriv[Artanh] = "1/(1-x^2)"
  deriv[Arcoth] = "1/(1-x^2)"
  deriv[F] =      "f'(x)"
  deriv[F1] =     "f\"(x)"
  deriv[F2] =     "0" // XXX
  deriv[G] =      "g'(x)"
  deriv[G1] =     "g\"(x)"
  deriv[G2] =     "0" // XXX
  deriv[H] =      "h'(x)"
  deriv[H1] =     "h\"(x)"
  deriv[H2] =     "0" // XXX
}

func funcDefined (s string) (Function, bool) {
  for f := 0; f < NFunctions; f++ {
    if s == name[f] {
      return f, true
    }
  }
  return Undef, false
}

func funcCodelen() uint {
  return 1
}

func funcEncode (f Function) obj.Stream {
  s := make (obj.Stream, 1)
  s[0] = byte(f)
  return s
}

func funcDecode (s obj.Stream) Function {
  if s[0] < NFunctions {
    return Function(s[0])
  }
  return Undef
}

func funcString (f Function) string {
  return name[f]
}

func funcVal (f Function, x float64) float64 {
  switch f {
  case Sqr:
    return x * x
  case Sqrt:
    return math.Sqrt (x)
  case Exp:
    return math.Exp (x)
  case Exp10:
    return math.Exp (x * math.Ln10)
  case Exp2:
    return math.Exp2 (x)
  case Log:
    return math.Log (x)
  case Lg:
    return math.Log10 (x)
  case Ld:
    return math.Log2 (x)
  case Sin:
    return math.Sin (x)
  case Cos:
    return math.Cos (x)
  case Tan:
    return math.Tan (x)
  case Cot:
    return 1 / math.Tan (x)
  case Arcsin:
    return math.Asin (x)
  case Arccos:
    return math.Acos (x)
  case Arctan:
    return math.Atan (x)
  case Arccot:
    return math.Atan (x)
  case Sinh:
    return math.Sinh (x)
  case Cosh:
    return math.Cosh (x) // (math.Exp(x) + math.Exp(-x)) / 2
  case Tanh:
    return math.Tanh (x)
  case Coth:
    return math.Sinh(x) / math.Cosh(x)
    return (math.Exp(x) + math.Exp(-x)) / (math.Exp(x) - math.Exp(-x))
  case Arsinh:
    return math.Asinh (x)
  case Arcosh:
    return math.Acos (x) // math.Log (x + math.Sqrt (x * x - 1))
  case Artanh:
    return math.Atanh (x)
  case Arcoth:
    return math.Log ((x + 1) / (x - 1)) / 2
  }
  return NaN()
}

func inverse (f Function) Function {
  return inv[f]
}

func funcStarted (s string) (Function, uint, bool) {
  a := len(s)
  for f := NFunctions - 1; f > 1; f-- {
    n := name[f] // "gamma"
    l := len(n)
    if l <= a {
      if s[:l] == n {
        s = s[:l]
        return f, uint(l), true
      }
    }
  }
  return Undef, 0, false
}

func isOpSymb (b byte) bool {
  switch b {
  case '+', '-', '*', '/', '^':
    return true
  }
  return false
}

func in (s string, ss []string) bool {
  for i := 0; i < len(ss); i++ {
    if s == ss[i] {
      return true
    }
  }
  return false
}

func strings (s string) []string {
  b := []byte (s)
  for i := 0; i < len(b); i++ {
    if b[i] == '(' || b[i] == ')' || char.IsDigit (b[i]) || isOpSymb (b[i]) {
      b[i] = ' '
    }
  }
  s = string(b)
  n, ss, _ := str.Split (s)
  ts := make([]string, 0)
  for i := uint(0); i < n; i++ {
    if _, ok := funcDefined (ss[i]); ! ok {
      ts = append (ts, ss[i])
    }
  }
  n = uint(len(ts))
  if n == 0 {
    return ts
  }
  sort.Slice (ts, func (i, j int) bool { return str.Less (ts[i], ts[j]) })
  us := make([]string, 1)
  us[0] = ts[0]
  for i := uint(1); i < n; i++ {
    if ! in (ts[i], us) {
      us = append (us, ts[i])
    }
  }
  return us
}

func power (x float64, k uint) float64 {
  if k == 0 {
    return 1
  }
  return x * power (x, k - 1)
}

func derivation (f Function) string {
  return deriv[f]
}
