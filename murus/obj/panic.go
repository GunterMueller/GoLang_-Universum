package obj

// (c) murus.org  v. 160310 - license see murus.go

import (
  "reflect"
  "strconv"
  "murus/ker"
)

func DivBy0Panic() {
  ker.Panic ("division by 0")
}

func TypeNotEqPanic (a, b Any) {
  ker.Panic ("the types " + reflect.TypeOf(a).String() +
                  " and " + reflect.TypeOf(b).String() + " are not equal")
//  ker.Panic ("the types " + x.String() + " and " + y.String() + " are not equal")
}

func WrongUintParameterPanic (s string, a Any, n uint) {
  ker.Panic ("method " + s +
             " for object of type " + reflect.TypeOf(a).String() +
             " got wrong value for " + strconv.FormatUint(uint64(n), 10))
}

func PanicNotAtomicOrObject (a Any) {
  ker.Panic ("the type " + reflect.TypeOf(a).String() + " is neither Atomic nor implements Object")
}

func PanicNotUintOrValuator (a Any) {
  ker.Panic ("the type " + reflect.TypeOf(a).String() + " is neither uint nor implements Valuator")
}
