package obj

// (c) Christian Maurer   v. 170918 - license see µU.go

import (
  "reflect"
  "runtime"
  "math"
  "strconv"
  "µU/ker"
)

// Equaler ///////////////////////////////////////

func eq (a, b Any) bool {
  if a == nil { return b == nil }
  if ! TypeEq (a, b) {
    return false
  }
  if atomic(a) {
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
  if atomic (a) {
    return a
  }
  switch a.(type) {
  case Equaler:
    return a.(Equaler).Clone()
  case Stream:
    b := make (Stream, len (a.(Stream)))
    copy (b, a.(Stream))
    return b
  case BoolStream:
    n := len(a.(BoolStream))
    b := make(BoolStream, n)
    for i := 0; i < n; i++ {
      b[i] = a.(BoolStream)[i]
    }
    return b
  default:
    ker.Panic ("µU only clones atomic types and objects of type BoolStream, Stream or Equaler")
  }
  return nil
}

// TODO DeepClone

// Object ////////////////////////////////////////

func atomic (a Any) bool {
  switch a.(type) {
  case bool,
       int8,  int16,  int32,  int,  int64,
       uint8, uint16, uint32, uint, uint64,
       float32, float64, complex64, complex128,
       string, Stream: // we treat them also as atomic
    return true
  }
  return false
}

func isObject (a Any) bool {
  _, o := a.(Object)
  _, e := a.(Editor) // Editor implements Object
  return o || e
}
/*
func AtomicOrObject (a Any) bool {
  return atomic (a) || isObject (a)
}
*/
// Comparer //////////////////////////////////////

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
  case Comparer:
    switch b.(type) {
      case Comparer:
      return a.(Comparer).Less (b)
    }
  }
  return false
}

func leq (a, b Any) bool {
  switch a.(type) {
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
  case Comparer:
    switch b.(type) {
    case Comparer:
      switch a.(type) {
      case Equaler:
        switch b.(type) {
        case Equaler:
          return a.(Comparer).Less (b) || a.(Equaler).Eq (b)
        }
      }
    }
  }
  return false
}

// Coder /////////////////////////////////////////

func fail (a Any) {
  ker.Panic ("µU only [en|de]codes atomic types " +
             "and objects of type string, Stream, BoolStream or Coder !")
}

func codelen (a Any) uint {
  if a == nil { return 0 }
  switch a.(type) {
  case bool, int8, uint8:
    return 1
  case int16, uint16:
    return 2
  case int32, uint32, float32:
    return 4
  case int, uint:
    switch runtime.GOARCH {
    case "amd64":
      return 8
    case "386":
      return 4
    default:
      ker.Panic ("$GOARCH not treated")
    }
  case int64, uint64, float64, complex64:
    return 8
  case complex128:
    return 16
  case string:
    return uint(len(a.(string)))
  case BoolStream:
    return uint(len(a.(BoolStream)))
  case Stream:
    return uint(len(a.(Stream)))
  case AnyStream:
    y := uint(0)
    for _, b := range a.(AnyStream) {
      y += uint(codelen(b))
    }
    return y
  case Coder:
    return (a.(Coder)).Codelen()
  }
  fail (a)
  panic("shut up, compiler")
}

func encode (a Any) Stream {
  if a == nil { return nil }
  var bs Stream
  switch a.(type) {
  case bool:
    bs = make (Stream, 1)
    if a.(bool) { bs[0]++ }
  case int8:
    bs = make (Stream, 1)
    bs[0] = a.(byte)
  case int16:
    n, x := 2, a.(int16)
    bs = make (Stream, n)
    for i := 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case int32:
    n, x := 4, a.(int32)
    bs = make (Stream, n)
    for i := 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case int:
    n, x := codelen(int(0)), a.(int)
    bs = make (Stream, n)
    for i := uint(0); i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case int64:
    n, x := 8, a.(int64)
    bs = make (Stream, n)
    for i := 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case uint8:
    bs = make (Stream, 1)
    bs[0] = a.(uint8)
  case uint16:
    n, x := 2, a.(uint16)
    bs = make (Stream, n)
    for i := 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case uint32:
    n, x := 4, a.(uint32)
    bs = make (Stream, n)
    for i := 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case uint:
    x := a.(uint)
    bs = make (Stream, C0)
    for i := uint(0); i < C0; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case uint64:
    n, x := 8, a.(uint64)
    bs = make (Stream, n)
    for i := 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case float32:
    n, x := 4, math.Float32bits (a.(float32))
    bs = make (Stream, n)
    for i := 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case float64:
    n, x := 8, math.Float64bits (a.(float64))
    bs = make (Stream, n)
    for i := 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case complex64:
    n, c := 8, a.(complex64)
    x, y := math.Float32bits (real(c)), math.Float32bits (imag(c))
    bs = make (Stream, n)
    for i := 0; i < n / 2; i++ {
      bs[i] = byte(x)
      bs[i + n/2] = byte(y)
      x >>= 8; y >>= 8
    }
  case complex128:
    n, c := 16, a.(complex128)
    x, y := math.Float64bits (real(c)), math.Float64bits (imag(c))
    bs = make (Stream, n)
    for i := 0; i < n / 2; i++ {
      bs[i] = byte(x)
      bs[i + n/2] = byte(y)
      x >>= 8; y >>= 8
    }
  case string:
    return (Stream)(a.(string))
  case BoolStream:
    n := len (a.(BoolStream))
    ys := make (Stream, n)
    for i := 0; i < n; i++ {
      ys[i] = 0; if a.(BoolStream)[i] { ys[i] = 1 }
    }
    copy (bs, ys)
  case Stream:
    return a.(Stream)
  case AnyStream:
    as := a.(AnyStream)
    n := uint(len(as))
    c := C0
    for j := uint(0); j < n; j++ {
      c += C0 + 1 + codelen(as[j])
    }
    bs = make (Stream, c)
    copy (bs[:C0], encode(n))
    i := C0
    for j := uint(0); j < n; j++ {
      g := gödel(as[j])
      copy(bs[i:i+1], encode(g))
      i++
      k := codelen(as[j])
      copy(bs[i:i+C0], encode(k))
      i += C0
      copy(bs[i:i+k], encode(as[j]))
      i += k
    }
  case Coder:
    return a.(Coder).Encode()
  default:
    fail (a)
  }
  return bs
}

func gödel (a Any) byte {
  if a == nil { return 0 }
  switch a.(type) {
  case bool:
    return 1
  case int8:
    return 2
  case int16:
    return 3
  case int32:
    return 4
  case int:
    return 5
  case int64:
    return 6
  case uint8:
    return 7
  case uint16:
    return 8
  case uint32:
    return 9
  case uint:
    return 10
  case uint64:
    return 11
  case float32:
    return 12
  case float64:
    return 13
  case complex64:
    return 14
  case complex128:
    return 15
  case string:
    return 16
  case Object:
    return 254
  }
  return 255
}

func degödel (b byte) Any {
  switch b {
  case 0:
    return nil
  case 1:
    return false
  case 2:
    return int8(0)
  case 3:
    return int16(0)
  case 4:
    return int32(0)
  case 5:
    return 0
  case 6:
    return int64(0)
  case 7:
    return uint8(0)
  case 8:
    return uint16(0)
  case 9:
    return uint32(0)
  case 10:
    return uint(0)
  case 11:
    return uint64(0)
  case 12:
    return float32(0)
  case 13:
    return float64(0)
/*
  case 14:
    return complex64(0, 0)
  case 15:
    return complex128(0, 0)
*/
  case 16:
    return ""
  }
  return nil
}

func chk (b Stream, n int) {
  if len(b) < n { // != n {
    ker.Panic ("µU/obj/coder.go chk: len(b) == " + strconv.Itoa(len(b)) + " < n == " + strconv.Itoa(n))
  }
}

func decode (a Any, bs Stream) Any {
  if a == nil { return nil }
  switch a.(type) {
  case bool:
    chk (bs, 1)
    return bs[0] > 0
  case int8:
    chk (bs, 1)
    return int8(bs[0])
  case int16:
    n, x := 2, int16(0)
    chk (bs, n)
    for i := n - 1; i >= 0; i-- {
      x <<= 8
      x += int16(bs[i])
    }
    return x
  case int32:
    n, x := 4, int32(0)
    chk (bs, n)
    for i := n - 1; i >= 0; i-- {
      x <<= 8
      x += int32(bs[i])
    }
    return x
  case int:
    n, x := int(codelen(0)), 0
    chk (bs, int(n))
    for i := n - 1; i >= 0; i-- {
      x <<= 8
      x += int(bs[i])
    }
    return x
  case int64:
    n, x := 8, int64(0)
    chk (bs, n)
    for i := n - 1; i >= 0; i-- {
      x <<= 8
      x += int64(bs[i])
    }
    return x
  case uint8:
    chk (bs, 1)
    return bs[0]
  case uint16:
    n, x := 2, uint16(0)
    chk (bs, n)
    for i := n - 1; i >= 0; i-- {
      x <<= 8
      x += uint16(bs[i])
    }
    return x
  case uint32:
    n, x := 4, uint32(0)
    chk (bs, n)
    for i := n - 1; i >= 0; i-- {
      x <<= 8
      x += uint32(bs[i])
    }
    return x
  case uint:
    n, x := int(codelen(0)), uint(0)
    chk (bs, int(n))
    for i := n - 1; i >= 0; i-- {
      x <<= 8
      x += uint(bs[i])
    }
    return x
  case uint64:
    n, x := 8, uint64(0)
    chk (bs, n)
    for i := n - 1; i >= 0; i-- {
      x <<= 8
      x += uint64(bs[i])
    }
    return x
  case float32:
    n, x := 4, uint32(0)
    chk (bs, n)
    for i := n - 1; i >= 0; i-- {
      x <<= 8
      x += uint32(bs[i])
    }
    return math.Float32frombits (x)
  case float64:
    n, x := 8, uint64(0)
    chk (bs, n)
    for i := n - 1; i >= 0; i-- {
      x <<= 8
      x += uint64(bs[i])
    }
    return math.Float64frombits (x)
  case complex64:
    n, x, y := 8, uint32(0), uint32(0)
    chk (bs, n)
    for i := n / 2 - 1; i >= 0; i-- {
      x <<= 8; y <<= 8
      x += uint32(bs[i])
      y += uint32(bs[i + n/2])
    }
    return complex (math.Float32frombits (x), math.Float32frombits (y))
  case complex128:
    n, x, y := 16, uint64(0), uint64(0)
    chk (bs, n)
    for i := n / 2 - 1; i >= 0; i-- {
      x <<= 8; y <<= 8
      x += uint64(bs[i])
      y += uint64(bs[i + n/2])
    }
    return complex (math.Float64frombits (x), math.Float64frombits (y))
  case string:
    return string(bs)
  case BoolStream:
    n := len(bs)
    y := make(BoolStream, n)
    for i := 0; i < n; i++ {
      if bs[i] == 1 { y[i] = true }
    }
    return y
  case Stream:
    return bs
    copy (a.(Stream), bs)
  case AnyStream:
    n := decode(uint(0), bs[:C0]).(uint)
    as := make(AnyStream, n)
    i := C0
    for j := uint(0); j < n; j++ {
      g := degödel(bs[i])
      i++
      k := decode(uint(0), bs[i:i+C0]).(uint)
      i += C0
      as[j] = decode(g, bs[i:i+k])
      i += k
    }
    return as
  case Coder:
    chk (bs, int(a.(Coder).Codelen()))
    a.(Coder).Decode (bs)
  default:
    fail (a)
  }
  return a
}

func encode4 (a, b, c, d uint32) Stream {
  bs := make (Stream, 4 * 4)
  copy (bs[:4], encode (a))
  copy (bs[4:8], encode (b))
  copy (bs[8:12], encode (c))
  copy (bs[12:16], encode (d))
  return bs
}

func decode4 (bs Stream) (uint32, uint32, uint32, uint32) {
  a := decode (uint32(0), bs[:4]).(uint32)
  b := decode (uint32(0), bs[4:8]).(uint32)
  c := decode (uint32(0), bs[8:12]).(uint32)
  d := decode (uint32(0), bs[12:16]).(uint32)
  return a, b, c, d
}

// TODO Spec
func Encodes (as AnyStream, cs []uint) Stream {
  l := uint(0)
  for _, b:= range cs {
    l += b
  }
  bs := make (Stream, l)
  a := uint(0)
  for i, x := range as {
    copy (bs[a:a+cs[i]], encode(x))
    a += cs[i]
  }
  return bs
}

// TODO Spec
func Decodes (bs Stream, as AnyStream, cs []uint) {
  l := uint(0)
  for _, b:= range cs {
    l += b
  }
  a := uint(0)
  for i, x := range as {
    as[i] = decode (x, bs[a:a+cs[i]])
    a += cs[i]
  }
}

// Valuator //////////////////////////////////////

func val (a Any) uint {
  switch a.(type) {
  case Valuator:
    return (a.(Valuator)).Val()
  case byte:
    return uint(a.(byte))
  case uint16:
    return uint(a.(uint16))
  case uint32:
    return uint(a.(uint32))
  case uint:
    return a.(uint)
  case uint64:
    u := a.(uint64)
    if u < 1<<32 {
      return uint(u)
    } else {
      return uint(u % 1<<32)
    }
  }
  return uint(1)
}

// func intVal (a Any) int { // XXX ?

func setVal (x *Any, n uint) {
  switch (*x).(type) {
  case byte:
    if n < 1<<8 {
      *x = uint(n)
    } else {
      *x = uint(n % 1<<8)
    }
  case uint16:
    if n < 1<<16 {
      *x = uint(n)
    } else {
      *x = uint(n % 1<<16)
    }
  case uint32:
    *x = uint32(n)
  case uint:
    *x = n
  case uint64:
    if n < 1<<32 {
      *x = uint(n)
    } else {
      *x = uint(n % 1<<32)
    }
  case Valuator:
    (*x).(Valuator).SetVal(n)
  default:
    ker.Panic("SetVal not possible for type " + reflect.TypeOf(*x).String())
  }
}

// RealValuator //////////////////////////////////

func realVal (a Any) float64 {
  var r float64 = 1.0
  switch a.(type) { case RealValuator:
    r = (a.(RealValuator)).RealVal()
  case bool:
    if ! a.(bool) {
      r = 0.0
    }
  case int8:
    r = float64(a.(int8))
  case int16:
    r = float64(a.(int16))
  case int32:
    r = float64(a.(int32))
  case int:
    r = float64(a.(int))
  case byte:
    r = float64(a.(byte))
  case uint16:
    r = float64(a.(uint16))
  case uint32:
    r = float64(a.(uint32))
  case uint:
    r = float64(a.(uint))
  case float32:
    r = math.Trunc (float64(a.(float32) + 0.5))
  case float64:
    r = math.Trunc (a.(float64) + 0.5)
  case complex64:
    c := a.(complex64)
    r = math.Trunc (math.Sqrt(float64(real(c) * real(c) + imag(c) * imag(c))) + 0.5)
  case complex128:
    c := a.(complex128)
    r = math.Trunc (math.Sqrt(real(c) * real(c) + imag(c) * imag(c)) + 0.5)
  case string:
    // TODO sum of bytes of the string ? Hash-Code ?
  }
  return r
}
