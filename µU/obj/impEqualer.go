package obj

// (c) Christian Maurer   v. 180902 - license see µU.go

import(
  "reflect"
  "µU/ker"
)

func isEqualer (a Any) bool {
  if a == nil { return false }
  _, ok := a.(Equaler)
  return ok
}

func eq (a, b Any) bool {
  if a == nil { return b == nil }
  if ! TypeEq (a, b) {
    return false
  }
  if Atomic(a) {
    return a == b
  }
  switch a.(type) {
  case Equaler:
    return a.(Equaler).Eq (b.(Equaler))
  case BoolStream:
    n := len(a.(BoolStream))
    if n != len(b.(BoolStream)) { return false }
    for i, y := range a.(BoolStream) {
      if b.(BoolStream)[i] != y {
        return false
      }
      return true
    }
  case Stream:
    n := len(a.(Stream))
    if n != len(b.(Stream)) { return false }
    for i, y := range a.(Stream) {
      if b.(Stream)[i] != y {
        return false
      }
      return true
    }
  case IntStream:
    n := len(a.(IntStream))
    if n != len(b.(IntStream)) { return false }
    for i, y := range a.(IntStream) {
      if b.(IntStream)[i] != y {
        return false
      }
      return true
    }
  case UintStream:
    n := uint(len(a.(IntStream)))
    if n != uint(len(b.(UintStream))) { return false }
    for i, y := range a.(UintStream) {
      if b.(UintStream)[i] != y {
        return false
      }
      return true
    }
  case AnyStream:
    n := len(a.(AnyStream))
    if n != len(b.(AnyStream)) { return false }
    for i, y := range a.(AnyStream) {
      if ! eq (b.(AnyStream)[i], y) {
        return false
      }
      return true
    }
  case *Any:
    return eq (a, b)
  }
  return reflect.DeepEqual (a, b)
}

func clone (a Any) Any {
  if a == nil {
    return nil
  }
  if Atomic (a) {
    return a
  }
  switch a.(type) {
  case Equaler:
    return a.(Equaler).Clone()
  case BoolStream:
    n := len(a.(BoolStream))
    b := make(BoolStream, n)
    for i := 0; i < n; i++ {
      b[i] = a.(BoolStream)[i]
    }
    return b
  case Stream:
    b := make (Stream, len (a.(Stream)))
    copy (b, a.(Stream))
    return b
  case IntStream:
    b := make (IntStream, len (a.(IntStream)))
    copy (b, a.(IntStream))
    return b
  case UintStream:
    b := make (UintStream, len (a.(UintStream)))
    copy (b, a.(UintStream))
    return b
  case AnyStream:
    b := make (AnyStream, len (a.(AnyStream)))
    copy (b, a.(AnyStream))
    return b
  default:
    ker.Panic ("µU only clones atomic types and objects of type string or _Stream or Equaler")
  }
  return nil
}

// TODO deepClone
