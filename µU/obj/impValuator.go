package obj

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  "reflect"
  "µU/ker"
)

func isValuator (a any) bool {
  if a == nil { return false }
  switch a.(type) {
  case Valuator:
    return true
  case byte, uint16, uint32, uint, uint64:
    return true
  }
  return false
}

func val (a any) uint {
  n := uint(1)
  switch a.(type) {
  case Valuator:
    n = (a.(Valuator)).Val()
/*/
  case int8:
    n = uint(a.(int8))
    if a.(int8) >= 0 {
      n = uint(a.(int8))
    } else {
      n = uint(-a.(int8))
    }
  case int16:
    if a.(int16) >= 0 {
      n = uint(a.(int16))
    } else {
      n = uint(-a.(int16))
    }
  case int32:
    if a.(int32) >= 0 {
      n = uint(a.(int32))
    } else {
      n = uint(-a.(int32))
    }
  case int64:
    if a.(int64) >= 0 {
      n = uint(a.(int64))
    } else {
      n = uint(-a.(int64))
    }
  case int:
    if a.(int) >= 0 {
      n = uint(a.(int))
    } else {
      n = uint(-a.(int))
    }
/*/
  case byte:
    n = uint(a.(byte))
  case uint16:
    n = uint(a.(uint16))
  case uint32:
    n = uint(a.(uint32))
  case uint64:
    n = uint(a.(uint64))
  case uint:
    n = a.(uint)
  case float32:
    if a.(float32) >= 0 {
      n = uint(a.(float32))
    } else {
      n = uint(-a.(float32))
    }
  case float64:
    if a.(float64) >= 0 {
      n = uint(a.(float64))
    } else {
      n = uint(-a.(float64))
    }
  }
  return n
}

func setVal (a *any, n uint) {
  switch (*a).(type) {
  case int8:
    *a = uint(n % 1<<8)
  case int16:
    *a = uint(n % 1<<16)
  case int32:
    *a = uint(n % 1<<32)
  case int64:
    *a = int64(n)
  case int:
    *a = int(n)
  case byte:
    *a = uint(n % 1<<8)
  case uint16:
    *a = uint(n % 1<<16)
  case uint32:
    *a = uint(n % 1<<32)
  case uint:
    *a = n
  case uint64:
    *a = uint(n)
  case float32:
    *a = uint(float64(n % 1<<32))
  case float64:
    *a = uint(float64(n))
  case Valuator:
    (*a).(Valuator).SetVal(n)
  default:
    ker.Panic (reflect.TypeOf(*a).String() + " has no uint-type nor implements Valuator")
  }
}
