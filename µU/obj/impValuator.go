package obj

// (c) Christian Maurer   v. 210309 - license see µU.go

import (
  "reflect"
  "µU/ker"
)

func isValuator (a Any) bool {
  switch a.(type) {
  case Valuator:
    return true
  case byte, uint16, uint32, uint, uint64:
    return true
  }
  return false
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
    return a.(uint)
  }
  return uint(1)
}

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
    if n < 1<<32 {
      *x = uint(n)
    } else {
      *x = uint(n % 1<<32)
    }
  case uint:
    *x = n
  case uint64:
    *x = uint(n)
  case Valuator:
    (*x).(Valuator).SetVal(n)
  default:
    ker.Panic(reflect.TypeOf(*x).String() + " has no uint-type nor implements Valuator")
  }
}
