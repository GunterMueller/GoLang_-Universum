package obj

// (c) Christian Maurer   v. 171104 - license see µU.go

import (
  "reflect"
//  "runtime"
//  "math"
  "strconv"
  "µU/ker"
)

func TypeEq (a, b Any) bool {
  x, y := reflect.TypeOf(a), reflect.TypeOf(b)
  if x != y { TypeNotEqPanic (a, b) }
  return x == y
}

func CheckTypeEq (a, b Any) {
  if a == nil && b == nil { return }
  if a == nil && b != nil || b == nil && a != nil {
    TypeNotEqPanic (a, b)
  }
  x, y := reflect.TypeOf(a), reflect.TypeOf(b)
  if x != y {
    TypeNotEqPanic (a, b)
  }
}

func CheckAtomicOrEqualer (a Any) {
  if ! AtomicOrEqualer(a) {
    PanicNotAtomicOrEqualer(a)
  }
}

func CheckAtomicOrObject (a Any) {
  if ! AtomicOrObject(a) {
    PanicNotAtomicOrObject(a)
  }
}

func UintOrValuator (a Any) bool {
  switch a.(type) {
  case byte, uint16, uint32, uint, uint64, Valuator:
    return true
  }
  return false
}

func CheckUintOrValuator (a Any) {
  if ! UintOrValuator(a) {
    PanicNotUintOrValuator(a)
  }
}

func DivBy0Panic() {
  ker.Panic ("division by 0")
}

func bluse (a Any) string {
  t := "nil"
  if a != nil {
    t = reflect.TypeOf(a).String()
  }
  return t
}

func TypeNotEqPanic (a, b Any) {
  ker.Panic ("the types " + bluse(a) + " and " + bluse(b) + " are not equal")
}

func WrongUintParameterPanic (s string, a Any, n uint) {
  ker.Panic ("method " + s + " for object of type " + bluse(a) +
             " got wrong value for " + strconv.FormatUint(uint64(n), 10))
}

func PanicNotAtomicOrEqualer (a Any) {
  ker.Panic ("the type " + bluse(a) + " is neither Atomic nor implements Equaler")
}

func PanicNotAtomicOrObject (a Any) {
  ker.Panic ("the type " + bluse(a) + " is neither Atomic nor implements Object")
}

func PanicNotUintOrValuator (a Any) {
  ker.Panic ("the type " + bluse(a) + " is neither uint nor implements Valuator")
}
