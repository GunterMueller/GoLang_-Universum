package obj

// (c) Christian Maurer   v. 171125 - license see nU.go

import "reflect"

type Valuator interface {

// Returns the value of x.
  Val() uint

// Returns true, iff x can be given a value. In that case x.Val() == n.
// Returns otherwise false.
  SetVal (n uint) bool
}

func Val (a Any) uint { return val(a) }

func SetVal (x *Any, n uint) { setVal(x,n) }

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
