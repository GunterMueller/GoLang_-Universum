package real

// (c) Christian Maurer   v. 201014 - license see µU.go

import (
  "math"
  "µU/obj"
)
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

func EncodeOp (op Operation) obj.Stream {
  s := make (obj.Stream, 1)
  s[0] = opSymbol[op]
  return s
}

func DecodeOp (s obj.Stream) Operation {
  if s[0] < NOperations {
    return Operation(s[0])
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
