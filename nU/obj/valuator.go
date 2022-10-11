package obj

// (c) Christian Maurer   v. 220702 - license see nU.go

import "reflect"

type Valuator interface {

// Liefert den Wert von x.
  Val() uint

// Liefert genau dann true, wenn x ein Wert gegeben werden kann;
// in diesem Fall ist x.Val() = n.
  SetVal (n uint) bool
}

func Val (a any) uint { return val(a) }

func SetVal (x *any, n uint) { setVal(x,n) }

func val (a any) uint {
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

func setVal (x *any, n uint) {
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
    if n < 1<<32 {
      *x = uint(n)
    } else {
      *x = uint(n % 1<<32)
    }
  case Valuator:
    (*x).(Valuator).SetVal(n)
  default:
    panic("SetVal not possible for type " + reflect.TypeOf(*x).String())
  }
}
