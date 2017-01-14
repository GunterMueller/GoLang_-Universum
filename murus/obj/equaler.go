package obj

// (c) murus.org  v. 170101 - license see murus.go

// TODO DeepClone

import (
  "reflect"
)
type
  Equaler interface { // x denotes the calling object.

// Returns true, iff the x has the same type as y
// and coincides with it in all its value[s].
  Eq (y Any) bool

// If y has the same type as x, then x.Eq(y) (y is unchanged).
  Copy (y Any)

// Returns a clone of x, i.e. x.Eq(x.Clone()).
  Clone() Any
}

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

func Eq (a, b Any) bool {
//  println ("obj.Eq")
  if a == nil { return b == nil }
  if ! TypeEq (a, b) {
    return false
  }
  if X, ok:= a.(Object); ok {
    return X.Eq (b.(Object))
  }
  if x, ok:= a.([]byte); ok {
    for i:= 0; i < len (x); i++ {
      if a.([]byte)[i] != b.([]byte)[i] {
        return false
      }
      return true
    }
  }
  if Atomic (a) {
    return a == b
  }
  return reflect.DeepEqual (a, b)
}

func Clone (a Any) Any {
  if a == nil {
    return nil
  }
  if Atomic (a) {
    return a
  }
  switch a.(type) {
  case Object:
    return a.(Object).Clone()
/*
  case bool:
    return a.(bool)
  case string:
    return a.(string)
  case uint8:
    return a.(uint8)
  case uint16:
    return a.(uint16)
  case uint32:
    return a.(uint32)
  case uint:
    return = a.(uint)
  case uint64:
    return a.(uint64)
  case int8:
    return a.(int8)
  case int16:
    return a.(int16)
  case int32:
    return a.(int32)
  case int:
    return a.(int)
  case int64:
    return a.(int64)
  case float32:
    return a.(float32)
  case float64:
    return a.(float64)
  case complex64:
    return a.(complex64)
  case complex128:
    return a.(complex128)
*/
  case []byte:
    b:= make ([]byte, len (a.([]byte)))
    copy (b, a.([]byte))
    return b
  default:
    println ("obj/equaler: Clone.default: TODO")
  }
  return nil
}
