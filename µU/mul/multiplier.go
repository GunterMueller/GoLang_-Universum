package mul

// (c) Christian Maurer   v. 180902 - license see µU.go

import
  . "µU/obj"

func one (a Any) bool {
  if a == nil { return false }
  switch a.(type) {
  case Multiplier:
    return a.(Multiplier).One()
  case int8:
    return a.(int8) == 1
  case int16:
    return a.(int16) == 1
  case int32:
    return a.(int32) == 1
  case int:
    return a.(int) == 1
  case int64:
    return a.(int64) == 1
  case uint8:
    return a.(uint8) == 1
  case uint16:
    return a.(uint16) == 1
  case uint32:
    return a.(uint32) == 1
  case uint:
    return a.(uint) == 1
  case uint64:
    return a.(uint64) == 1
  case float32:
    return a.(float32) == 1.
  case float64:
    return a.(float64) == 1.
  case complex64:
    a0 := a.(complex64)
    return real(a0) == 1. && imag(a0) == 0.
  case complex128:
    a0 := a.(complex128)
    return real(a0) == 1. && imag(a0) == 0.
  }
  return false
}

func mul (a Any, bs ...Any) Any {
  if a == nil { return nil }
  for _, b := range bs {
    if ! TypeEq(a, b) { return nil }
  }
  switch a.(type) {
  case Multiplier:
    for _, b := range bs {
      a.(Multiplier).Mul (b.(Multiplier))
    }
    return a
  case int8:
    a0 := a.(int8)
    for _, b := range bs {
      a0 *= b.(int8)
    }
    return a0
  case int16:
    a0 := a.(int16)
    for _, b := range bs {
      a0 *= b.(int16)
    }
    return a0
  case int32:
    a0 := a.(int32)
    for _, b := range bs {
      a0 *= b.(int32)
    }
    return a0
  case int:
    a0 := a.(int)
    for _, b := range bs {
      a0 *= b.(int)
    }
    return a0
  case int64:
    a0 := a.(int64)
    for _, b := range bs {
      a0 *= b.(int64)
    }
    return a0
  case uint8:
    a0 := a.(uint8)
    for _, b := range bs {
      a0 *= b.(uint8)
    }
    return a0
  case uint16:
    a0 := a.(uint16)
    for _, b := range bs {
      a0 *= b.(uint16)
    }
    return a0
  case uint32:
    a0 := a.(uint32)
    for _, b := range bs {
      a0 *= b.(uint32)
    }
    return a0
  case uint:
    a0 := a.(uint)
    for _, b := range bs {
      a0 *= b.(uint)
    }
    return a0
  case uint64:
    a0 := a.(uint64)
    for _, b := range bs {
      a0 *= b.(uint64)
    }
    return a0
  case complex64:
/*
    r, i := real(a.(complex64)), imag(a.(complex64))
    for _, b := range bs {
      // TODO
    }
*/
    return nil // TODO
  case complex128:
    return nil // TODO
  }
  return nil // TODO
}

func product (as []Multiplier) Multiplier {
  a0 := as[0]
  for i, a := range as {
    if i > 0 {
      a0.Mul (a)
    }
  }
  return a0
}

func sqr (a Any) Any {
  a.(Multiplier).Mul (a.(Multiplier))
  return a
}

//  Div (y, z Multiplier)

//  DivBy (y Multiplier)
