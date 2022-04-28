package obj

// (c) Christian Maurer   v. 220420 - license see µU.go

import
  "math"

func isClearer (a any) bool {
  if a == nil { return false }
  _, ok := a.(Clearer)
  return ok
}

func clear (a any) any {
  switch a.(type) {
  case bool:
    return false
  case int8:
    return math.MaxInt8
  case int16:
    return math.MaxInt16
  case int32:
    return math.MaxInt32
  case int64:
    return math.MaxInt64
  case int:
    return math.MaxInt
  case uint8:
    return math.MaxInt8
  case uint16:
    return math.MaxUint16
  case uint32:
    return math.MaxUint32
  case uint64:
    return math.MaxUint64 / 2
  case uint:
    return math.MaxUint / 2
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
  case BoolStream:
    return make(BoolStream, 0)
  case IntStream:
    return make(IntStream, 0)
  case AnyStream:
    return make(AnyStream, 0)
  case Clearer, Editor:
    a.(Clearer).Clr()
    return a
  }
  return nil
}
