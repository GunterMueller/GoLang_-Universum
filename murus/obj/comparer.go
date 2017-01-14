package obj

// (c) murus.org  v. 150514 - license see murus.go

type
  Comparer interface {

// Pre: x is of the same type as the calling object.
// Returns true, iff the calling object is smaller than x.
  Less (x Any) bool
}

func Less (a, b Any) bool {
  switch a.(type) { case Object:
    switch b.(type) { case Object:
      return a.(Comparer).Less (b)
    }
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
  }
  return false
}

/*
func Leq (a, b Any) bool {
  switch a.(type) {
  case Object:
    switch b.(type) { case Object:
      return a.(Comparer).Less (b) ||
             a.(Equaler).Eq (b)
    }
  case bool:
    return ! a.(bool) && b.(bool) ||
             a.(bool) == b.(bool)
  case byte:
    return a.(byte) <= b.(byte)
  case uint16:
    return a.(uint16) <= b.(uint16)
  case uint32:
    return a.(uint32) <= b.(uint32)
  case uint:
    return a.(uint) <= b.(uint)
  case uint64:
    return a.(uint64) <= b.(uint64)
  case int8:
    return a.(int8) <= b.(int8)
  case int16:
    return a.(int16) <= b.(int16)
  case int32:
    return a.(int32) <= b.(int32)
  case int:
    return a.(int) <= b.(int)
  case int64:
    return a.(int64) <= b.(int64)
  case float32:
    return a.(float32) <= b.(float32)
  case float64:
    return a.(float64) <= b.(float64)
  case string:
    return a.(string) <= b.(string)
  }
  return false
}
*/
