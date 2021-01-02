package obj

// (c) Christian Maurer   v. 201204 - license see µU.go

import (
  "reflect"
  "µU/ker"
)

func isValuator (a Any) bool {
  if a == nil { return false }
  _, ok := a.(Valuator)
  return ok
}

func val (a Any) uint {
  switch a.(type) {
  case Valuator:
    return (a.(Valuator)).Val()
  case byte:
    return uint(a.(byte))
  case uint16:
    return uint(a.(uint16))
  case uint32:
    return uint(a.(uint32))
  case uint:
    return a.(uint)
  case uint64:
    u := a.(uint64)
    if u < 1<<32 {
      return uint(u)
    } else {
      return uint(u % 1<<32)
    }
  }
  return uint(1)
}

// func intVal (a Any) int { // XXX ?

func setVal (x *Any, n uint) {
  switch (*x).(type) {
  case byte:
    if n < 1<<8 {
      *x = uint(n)
    } else {
      *x = uint(n % 1<<8)
    }
  case uint16:
    if n < 1<<16 {
      *x = uint(n)
    } else {
      *x = uint(n % 1<<16)
    }
  case uint32:
    *x = uint32(n)
  case uint:
    *x = n
  case uint64:
    *x = uint(n)
  case Valuator:
    (*x).(Valuator).SetVal(n)
  default:
    ker.Panic("SetVal not possible for type " + reflect.TypeOf(*x).String())
  }
}
