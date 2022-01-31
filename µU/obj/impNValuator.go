package obj

// (c) Christian Maurer   v. 220126 - license see µU.go

import
  "math"

func isNValuator (a Any) bool {
  if a == nil { return false }
  _, ok := a.(NValuator)
  return ok
}

func nVal (a Any) uint {
  n := uint(1)
  switch a.(type) {
  case NValuator:
    n = (a.(NValuator)).Val()
  case bool:
    if ! a.(bool) {
      n = uint(0)
    }
  case int8:
    n = uint(a.(int8))
  case int16:
    n = uint(a.(int16))
  case int32:
    n = uint(a.(int32))
  case int:
    n = uint(a.(int))
  case byte:
    n = uint(a.(byte))
  case uint16:
    n = a.(uint)
  case uint32:
    n = a.(uint)
  case uint:
    n = a.(uint)
  case float32:
    n = uint(math.Trunc (float64(a.(float32) + 0.5)))
  case float64:
    n = uint(math.Trunc (a.(float64) + 0.5))
  case complex64:
//    c := a.(complex64)
//    n = math.Trunc (math.Sqrt(float64(real(c) * real(c) + imag(c) * imag(c))) + 0.5)
  case complex128:
//    c := a.(complex128)
//    n = math.Trunc (math.Sqrt(real(c) * real(c) + imag(c) * imag(c)) + 0.5)
  case string:
    // TODO sum of bytes of the string ? Hash-Code ?
  }
  return n
}
