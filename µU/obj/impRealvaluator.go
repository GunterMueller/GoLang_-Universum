package obj

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  "math"
  "reflect"
  "µU/ker"
)

func isRealValuator (a any) bool {
  if a == nil { return false }
  _, ok := a.(RealValuator)
  return ok
}

func realVal (a any) float64 {
  r := 1.
  switch a.(type) {
  case RealValuator:
    r = (a.(RealValuator)).RealVal()
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
  }
  return r
}

func setRealVal (a *any, r float64) {
  switch (*a).(type) {
  case int8:
    *a = int8(r)
  case int16:
    *a = int16(r)
  case int32:
    *a = int32(r)
  case int:
    *a = int(r)
  case byte:
    if r < 256. {
      *a = byte(math.Trunc(r))
    }
  case uint16:
    *a = uint16(r)
  case uint32:
    *a = uint32(r)
  case uint:
    *a = uint(r)
  case float32:
    *a = float32(r)
  case float64:
    *a = r
  case RealValuator:
    (*a).(RealValuator).SetRealVal(r)
  default:
    ker.Panic (reflect.TypeOf(*a).String() + " has no number-type nor implements RealValuator")
  }
}
