package real

// (c) Christian Maurer  v. 140803 - license see murus.go

import
  "math"
var
  opSymbol = [NOperations]byte { 'n', '+', '-', '*', '/', '^' } // 'n' <- ' ' ?

func OpDefined (op *Operation, c byte) bool {
  for o:= Operation(0); o < NOperations; o++ {
    if c == opSymbol[o] {
      *op = o
      return true
    }
  }
  return false
}

func AddSubOpDefined (op *Operation, c byte) bool {
  for o:= Operation(Add); o <= Sub; o ++ {
    if c == opSymbol[o] {
      *op = o
      return true
    }
  }
  return false
}

func MulDivOpDefined (op *Operation, c byte) bool {
  for o:= Operation(Mul); o <= Div; o ++ {
    if c == opSymbol[o] {
      *op = o
      return true
    }
  }
  return false
}

func PowOpDefined (op *Operation, c byte) bool {
  if c == opSymbol[Pow] {
    *op = Pow
    return true
  }
  return false
}

func AddSubOp (op Operation) bool {
  return op == Add || op == Sub
}

func MulDivOp (op Operation) bool {
  return op == Mul || op == Div
}

func PowOp (op Operation) bool {
  return op == Pow
}

func OpString (op Operation) string {
  return string(opSymbol[op])
}

func CodelenOp() uint {
  return 1
}

func EncodeOp (op Operation) []byte {
  b:= make ([]byte, 1)
  b[0] = opSymbol[op]
  return b
}

func DecodeOp (b []byte) Operation {
  if b[0] < NOperations {
    return Operation(b[0])
  }
  return NoOp
}

func OpVal (op Operation, x, y float64) float64 {
  switch op {
  case Add:
    return x + y
  case Sub:
    return x - y
  case Mul:
    return x * y
  case Div:
    if y == 0 { // TODO
      return math.NaN()
    }
    return x / y
  case Pow:
    return math.Pow (x, y)
  }
  return math.NaN()
}
