package obj

// (c) Christian Maurer   v. 171112 - license see µU.go

type
  Clearer interface {

// Returns true, iff the calling object is empty. What "empty" actually means,
// depends on the very semantics of the type of the objects considered.
// If that type is e.g. a collector, empty means "containing no objects"; otherwise it is
// normally an object with undefined value, represented by strings consisting only of spaces. 
  Empty() bool

// The calling object is empty.
  Clr()
}

func isClearer (a Any) bool {
  _, c := a.(Clearer)
  return c
}

/*
func clear (a Any) Any {
  switch a.(type) {
  case bool:
    return false
  case int8:
    return math.MaxInt8
  case int16:
    return math.MaxInt16
  case int32:
    return math.MaxInt32
  case int:
    return 0 // XXX
  case int64:
    return math.MaxInt64
  case uint8:
    return math.MaxInt8
  case uint16:
    return math.MaxUint16
  case uint32:
    return math.MaxUint32
  case uint:
    return 0 // XXX
  case uint64:
    return math.MaxUint64
  case float32:
    return math.MaxFloat32
  case float64:
    return math.MaxFloat64
  case complex64:
    return 0 // TODO
  case complex128:
    return 0 // TODO
  case string:
    return ""
  case Stream:
    return make(Stream, 0)
  case Clearer, Editor:
    a.(Clearer).Clr()
    return a
  }
  return nil
}
*/
