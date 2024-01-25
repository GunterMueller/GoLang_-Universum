package obj

// (c) Christian Maurer   v. 220702 - license see nU.go

import (
  "reflect"
  "strconv"
)

func TypeEq (a, b any) bool {
  x, y := reflect.TypeOf(a), reflect.TypeOf(b)
  if x != y { TypeNotEqPanic (a, b) }
  return x == y
}

func CheckTypeEq (a, b any) {
  if a == nil && b == nil { return }
  if a == nil && b != nil || b == nil && a != nil {
    TypeNotEqPanic (a, b)
  }
  x, y := reflect.TypeOf(a), reflect.TypeOf(b)
  if x != y {
    TypeNotEqPanic (a, b)
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
  panic ("division by 0")
}

func TypeNotEqPanic (a, b any) {
  panic ("the types " + reflect.TypeOf(a).String() +
              " and " + reflect.TypeOf(b).String() + " are not equal")
}

func WrongUintParameterPanic (s string, a any, n uint) {
  panic ("method " + s +
             " for object of type " + reflect.TypeOf(a).String() +
             " got wrong value for " + strconv.FormatUint(uint64(n), 10))
}

func PanicNotAtomicOrObject (a any) {
  panic ("the type " + reflect.TypeOf(a).String() + " is neither Atomic nor implements Object")
}

func PanicNotUintOrValuator (a any) {
  panic ("the type " + reflect.TypeOf(a).String() + " is neither uint nor implements Valuator")
}
