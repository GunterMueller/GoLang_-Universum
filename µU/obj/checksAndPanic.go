package obj

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  "reflect"
  "strconv"
  "µU/ker"
)

func TypeEq (a, b any) bool {
  x, y := reflect.TypeOf(a), reflect.TypeOf(b)
  if x != y {
    TypeNotEqPanic (a, b)
  }
  return x == y
}

func CheckTypeEq (a, b any) {
  if a == nil && b == nil { return }
  if a == nil && b != nil || b == nil && a != nil {
    TypeNotEqPanic (a, b)
  }
/*/
  TypeEq (a, b)
/*/
  x, y := reflect.TypeOf(a), reflect.TypeOf(b)
  if x != y {
    TypeNotEqPanic (a, b)
  }
}

func CheckEqualerAndComparer (a any) {
  if ! isEqualer(a) && isComparer(a) {
    PanicNotEqualerAndNotComparer(a)
  }
}

func CheckAtomicOrEqualer (a any) {
  if ! AtomicOrEqualer(a) {
    PanicNotAtomicOrEqualer(a)
  }
}

func CheckAtomicOrObject (a any) {
  if ! AtomicOrObject(a) {
    PanicNotAtomicOrObject(a)
  }
}

func UintOrValuator (a any) bool {
  switch a.(type) {
  case byte, uint16, uint32, uint, uint64, Valuator:
    return true
  }
  return false
}

func CheckUintOrValuator (a any) {
  if ! UintOrValuator(a) {
    PanicNotUintOrValuator(a)
  }
}

func DivBy0Panic() {
  ker.Panic ("division by 0")
}

func text (a any) string {
  t := "nil"
  if a != nil {
    t = reflect.TypeOf(a).String()
  }
  return t
}

func TypeNotEqPanic (a, b any) {
  ker.Panic ("the types " + text(a) + " and " + text(b) + " are not equal")
}

func WrongUintParameterPanic (s string, a any, n uint) {
  ker.Panic ("method " + s + " for object of type " + text(a) +
             " got wrong value for " + strconv.FormatUint(uint64(n), 10))
}

func PanicNotAtomicOrEqualer (a any) {
  ker.Panic ("the type " + text(a) + " is neither Atomic nor implements Equaler")
}

func PanicNotEqualerAndNotComparer (a any) {
  ker.Panic ("the type " + text(a) + " does not implement Equaler and Comparer")
}

func PanicNotAtomicOrObject (a any) {
  ker.Panic ("the type " + text(a) + " is neither Atomic nor implements Object")
}

func PanicNotUintOrValuator (a any) {
  ker.Panic ("the type " + text(a) + " is neither uint nor implements Valuator")
}
