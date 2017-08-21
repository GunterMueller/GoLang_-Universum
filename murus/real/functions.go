package real

// (c) murus.org  v. 140803 - license see murus.go

// >>> lots of things TODO, particularly format

import
  "math"
const
  FN = 6 // maximal length of function names
var (
  name = [NFunctions]string { "",
           "int", "fract", "-", "rec", "sqr", "sqrt",
           "exp", "exp10", "exp2", "log", "ld", "lg",
           "sin", "cos", "tan", "cot", "arcsin", "arccos", "arctan", "arccot",
           "sinh", "cosh", "tanh", "coth", "arsinh", "Arcosh", "artanh", "Arcoth",
           "gamma", "f", "f'", "f\"", "g", "g'", "g\"", "h", "h'", "h\"" }
  inv [NFunctions]Function
//  deriv [NFunctions]string
)

func init() {
  for f := Function(0); f < NFunctions; f++ {
    inv[f] = Undef
  }
  inv[Chs] =    Chs
  inv[Rec] =    Rec
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
  inv[Gamma]  = Undef
/*
  deriv[Floor] =  "0"
  deriv[Fract] =  "?"
  deriv[Chs] =    "-1"
  deriv[Rec] =    "-1/x2"
  deriv[Sqrt] =   "1/2/sqrt(x)"
  deriv[Sqr] =    "2*x"
  deriv[Exp] =    "ex"
  deriv[Exp10] =  "Log10*10x"
  deriv[Exp2] =   "Log2*2x"
  deriv[Log] =    "1/x"
  deriv[Ld] =     "1/Log2/x"
  deriv[Lg] =     "1/Log10/x"
  deriv[Sin] =    "cos(x)"
  deriv[Cos] =    "-(sin(x))"
  deriv[Tan] =    "1/(cos(x))2"
  deriv[Cot] =    "-1/(sin(x))2"
  deriv[Arcsin] = "1/sqrt(1-x2)"
  deriv[Arccos] = "-1/sqrt(1-x2)"
  deriv[Arctan] = "1/(1+x2)"
  deriv[Arccot] = "-1/(1+x2)"
  deriv[Sinh] =   "cosh(x)"
  deriv[Cosh] =   "sinh(x)"
  deriv[Tanh] =   "1/(cosh(x))2"
  deriv[Coth] =   "-1/(sinh(x))2"
  deriv[Arsinh] = "1/sqrt(x2+1)"
  deriv[Arcosh] = "1/sqrt(x2-1)"
  deriv[Artanh] = "1/(1-x2)"
  deriv[Arcoth] = "1/(1-x2)"
//  deriv[Gamma] =  "undef"
  deriv[F] =      "f'(x)"
  deriv[F1] =     "f\"(x)"
  deriv[F2] =     "f3(x)"
  deriv[G] =      "g'(x)"
  deriv[G1] =     "g\"(x)"
  deriv[G2] =     "g3(x)"
  deriv[H] =      "h'(x)"
  deriv[H1] =     "h\"(x)"
  deriv[H2] =     "h3(x)"
*/
}

/*
func FuncDefine (s string) Function {
  p := uint(0)
  return FunctionContained (s, f, &p) // Blödsinn
}

func FunctionContained (s string, f *Function, p *uint) bool { // Blödsinn
  a := *p
  if ! letter (s[p]) {
    *p = a
    return false
  }
  for i := 0; i < FN; i++ { Name[i] = "" }
  i := 0
  Name := ""
  for {
    if i > FN { break }
    c := s[*p]
    if letterOrChar (c) {
      Name[i] = c // Blödsinn
      i++
      *p++
    } else {
      break
    }
  }
  for fu := Function(0); fu < NFunctions; fu++ {
    if Name == name[fu]) {
      *f = fu
      return true
    }
  }
  *p = a
  return false
}
*/

func FunctionString (f Function) string {
  return name [f]
}

func FunctionVal (f Function, x float64) float64 {
  switch f {
  case Floor:
    i, _ := math.Modf (x)
    return i
  case Fract:
    _, f := math.Modf (x)
    return f
  case Chs:
    return -x
  case Rec:
    if x == 0 { // TODO
      return math.NaN()
    }
    return 1 / x
  case Sqr:
    return x * x
  case Sqrt:
    return math.Sqrt (x)
  case Exp:
    return math.Exp (x)
  case Exp10:
    return math.Exp (x * math.Ln10)
  case Exp2:
    return math.Exp (x * math.Ln2)
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
    return (math.Exp (x) + math.Exp (-x)) / 2
  case Tanh:
    return math.Tanh (x)
  case Coth:
    return (math.Exp (x) + math.Exp (-x)) / (math.Exp (x) - math.Exp (-x))
  case Arsinh:
    return math.Asinh (x)
  case Arcosh:
    return math.Log (x + math.Sqrt (x * x - 1))
  case Artanh:
    return math.Atanh (x)
  case Arcoth:
    return math.Log ((x + 1) / (x - 1)) / 2
  case Gamma:
    return math.Gamma (x)
  }
  return math.NaN()
}

func CodelenFunc() uint {
  return 1
}

func EncodeFunc (f Function) []byte {
  b := make ([]byte, 1)
  b[0] = byte(f)
  return b
}

func DecodeFunc (b []byte) Function {
  if b[0] < NFunctions {
    return Function(b[0])
  }
  return Undef
}

func Inverse (f Function) Function {
  return inv[f]
}

/*
func Derivation (f Function) string {
  return deriv [f]
}
*/
