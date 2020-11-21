package real

// (c) Christian Maurer   v. 201106 - license see µU.go

import (
  "math"
  "µU/obj"
)
var
  opSymbol = [NOperations]byte { 'n', '+', '-', '*', '/', '^' } // ' ' statt 'n' ?

func OpDefined (op *Operation, b byte) bool {
  for o:= Operation(0); o < NOperations; o++ {
    if b == opSymbol[o] {
      *op = o
      return true
    }
  }
  return false
}

func addSubOp (op *Operation, b byte) bool {
  for o:= Operation(Add); o <= Sub; o ++ {
    if b == opSymbol[o] {
      *op = o
      return true
    }
  }
  return false
}

func mulDivOp (op *Operation, b byte) bool {
  for o:= Operation(Mul); o <= Div; o ++ {
    if b == opSymbol[o] {
      *op = o
      return true
    }
  }
  return false
}

func powOp (op *Operation, b byte) bool {
  if b == opSymbol[Pow] {
    *op = Pow
    return true
  }
  return false
}

func opString (op Operation) string {
  return string(opSymbol[op])
}

func opCodelen() uint {
  return 1
}

func opEncode (op Operation) obj.Stream {
  s := make (obj.Stream, 1)
  s[0] = opSymbol[op]
  return s
}

func opDecode (s obj.Stream) Operation {
  if s[0] < NOperations {
    return Operation(s[0])
  }
  return NoOp
}

func opVal (op Operation, x, y float64) float64 {
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

func isOp (b byte) (Operation, bool) {
  for op := 1; op < NOperations; op++ {
    if b == opSymbol[op] {
      return op, true
    }
  }
  return NoOp, false
}

func opStarted (s string) (Operation, bool) {
  for op := 1; op < NOperations; op++ {
    if s[0] == opSymbol[op] {
      return op, true
    }
  }
  return NoOp, false
}
