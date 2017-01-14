package obj

// (c) murus.org  v. 161216 - license see murus.go

type
  Valuator interface {

// Returns the value of x.
  Val() uint

// Returns true, iff x can be given a value. In that case x.Val() == n.
// Returns otherwise false.
  SetVal (n uint) bool
}

func Val (a Any) uint {
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

// func IntVal (a Any) int { // XXX to be implemented

func CheckUintOrValuator (a Any) {
  if ! UintOrValuator(a) {
    PanicNotUintOrValuator(a)
  }
}

// Returns true, iff a is of type uint or implements Valuator.
func UintOrValuator (a Any) bool {
  switch a.(type) {
  case uint, Valuator:
    return true
  }
  return false
}

func SetVal (x *Any, n uint) {
// TODO check *x Quasivaluator, if not, ker.Panic ?
  switch (*x).(type) {
  case Valuator:
    (*x).(Valuator).SetVal(n)
  case uint:
    *x = n
  }
}
