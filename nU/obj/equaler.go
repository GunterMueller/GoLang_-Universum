package obj

// (c) Christian Maurer   v. 180813 - license see nU.go

import "reflect"

type Equaler interface {

// Returns true, iff the x has the same type as y
// and coincides with it in all its value[s].
  Eq (y Any) bool

// If y has the same type as x, then x.Eq(y) (y is unchanged).
  Copy (y Any)

// Returns a clone of x, i.e. x.Eq(x.Clone()).
  Clone() Any
}

// Pre: a and b are atomic or implement Equaler.
// Returns true, if a and b are equal.
func Eq (a, b Any) bool { return eq(a,b) }

// Pre: a is atomic or implements Equaler.
// Returns a clone of a.
func Clone (a Any) Any { return clone(a) }

// Returns true, iff a is atomic or implements Equaler.
func IsEqualer (a Any) bool { return isEqualer(a) }

// Returns true, iff a is atomic or implements Equaler.
func AtomicOrEqualer (a Any) bool { return Atomic(a) || isEqualer(a) }

func isEqualer (a Any) bool {
  _, e := a.(Equaler)
  return e
}

func eq (a, b Any) bool {
  if a == nil { return b == nil }
  if b == nil { return a == nil }
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
      if ! eq (b.(IntStream)[i], y) {
        return false
      }
      return true
    }
  case UintStream:
    n := len(a.(UintStream))
    if n != len(b.(UintStream)) { return false }
    for i, y := range a.(UintStream) {
      if ! eq (b.(UintStream)[i], y) {
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
  case Stream:
    b := make (Stream, len (a.(Stream)))
    copy (b, a.(Stream))
    return b
  case IntStream:
    n := len(a.(IntStream))
    b := make(IntStream, n)
    for i := 0; i < n; i++ {
      b[i] = a.(IntStream)[i]
    }
    return b
  case UintStream:
    n := len(a.(UintStream))
    b := make(UintStream, n)
    for i := 0; i < n; i++ {
      b[i] = a.(UintStream)[i]
    }
    return b
  case AnyStream:
    n := len(a.(AnyStream))
    b := make(AnyStream, n)
    for i := 0; i < n; i++ {
      b[i] = a.(AnyStream)[i]
    }
    return b
  default:
    panic ("nU only clones atomic types and objects of type string, __Stream or Equaler")
  }
  return nil
}
