package obj

// (c) Christian Maurer   v. 220126 - license see µU.go

import
  "math"

func isZValuator (a Any) bool {
  if a == nil { return false }
  _, ok := a.(ZValuator)
  return ok
}

func zVal (a Any) int {
  z := 1
  switch a.(type) {
  case ZValuator:
    z = (a.(ZValuator)).Val()
  case bool:
    if ! a.(bool) {
      z = 0
    }
  case int8:
    z = int(a.(int8))
  case int16:
    z = int(a.(int16))
  case int32:
    z = int(a.(int32))
  case int:
    z = int(a.(int))
  case byte:
    z = int(a.(byte))
  case uint16:
    z = a.(int)
  case uint32:
    z = int(a.(uint32))
  case uint:
    z = int(a.(uint))
  case float32:
    z = int(math.Trunc (float64(a.(float32) + 0.5)))
  case float64:
    z = int(math.Trunc (a.(float64) + 0.5))
  case complex64:
//    c := a.(complex64)
//    z = math.Trunc (math.Sqrt(float64(real(c) * real(c) + imag(c) * imag(c))) + 0.5)
  case complex128:
//    c := a.(complex128)
//    z = math.Trunc (math.Sqrt(real(c) * real(c) + imag(c) * imag(c)) + 0.5)
  case string:
    // TODO sum of bytes of the string ? Hash-Code ?
  }
  return z
}
