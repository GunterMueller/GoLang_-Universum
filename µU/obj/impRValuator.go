package obj

// (c) Christian Maurer   v. 220126 - license see µU.go

import
  "math"

func isRValuator (a Any) bool {
  if a == nil { return false }
  _, ok := a.(RValuator)
  return ok
}

func rVal (a Any) float64 {
  var r float64 = 1.0
  switch a.(type) {
  case RValuator:
    r = (a.(RValuator)).Val()
  case bool:
    if ! a.(bool) {
      r = 0.0
    }
  case int8:
    r = float64(a.(int8))
  case int16:
    r = float64(a.(int16))
  case int32:
    r = float64(a.(int32))
  case int:
    r = float64(a.(int))
  case byte:
    r = float64(a.(byte))
  case uint16:
    r = float64(a.(uint16))
  case uint32:
    r = float64(a.(uint32))
  case uint:
    r = float64(a.(uint))
  case float32:
    r = float64(a.(float32))
  case float64:
    r = a.(float64)
  case complex64:
    c := a.(complex64)
    r = math.Sqrt(float64(real(c) * real(c) + imag(c) * imag(c)))
  case complex128:
    c := a.(complex128)
    r = math.Sqrt(real(c) * real(c) + imag(c) * imag(c))
  case string:
    // TODO sum of bytes of the string ? Hash-Code ?
  }
  return r
}
