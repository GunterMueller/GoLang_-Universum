package obj

// (c) Christian Maurer   v. 190805 - license see µU.go

func isComparer (a Any) bool {
  if a == nil { return false }
  _, ok := a.(Comparer)
  return ok
}

/*/
func sless (a, b Stream) {
  for i := 0; i < len(b); i++ {
    if a[i]
  }
}
/*/

func less (a, b Any) bool {
  switch a.(type) {
  case bool:
    return ! a.(bool) && b.(bool)
  case byte:
    return a.(byte) < b.(byte)
  case uint16:
    return a.(uint16) < b.(uint16)
  case uint32:
    return a.(uint32) < b.(uint32)
  case uint:
    return a.(uint) < b.(uint)
  case uint64:
    return a.(uint64) < b.(uint64)
  case int8:
    return a.(int8) < b.(int8)
  case int16:
    return a.(int16) < b.(int16)
  case int32:
    return a.(int32) < b.(int32)
  case int:
    return a.(int) < b.(int)
  case int64:
    return a.(int64) < b.(int64)
  case float32:
    return a.(float32) < b.(float32)
  case float64:
    return a.(float64) < b.(float64)
  case string:
    return a.(string) < b.(string)
/*
  case BoolStream:
    a.(BoolStream)  b.(BoolStream)
  case Stream:
    a.(Stream)  b.(Stream)
  case IntStream:
    a.(IntStream)  b.(IntStream)
  case UintStream:
    a.(UintStream)  b.(UintStream)
  case AnyStream:
    a.(AnyStream)  b.(AnyStream)
*/
  case Comparer:
    switch b.(type) {
    case Comparer:
      return a.(Comparer).Less (b)
    }
  }
  return false
}
