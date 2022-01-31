package obj

// (c) Christian Maurer   v. 220128 - license see µU.go

import (
  "math"
  "strconv"
  "µU/ker"
)

func isCoder (a Any) bool {
  if a == nil { return false }
  _, ok := a.(Coder)
  return ok
}

func fail (a Any) {
  ker.Panic ("µU only [en|de]codes atomic types and objects of type string, ..Stream or Coder !")
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
    return C0
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
  case IntStream:
    return C0 + C0 * uint(len(a.(IntStream)))
  case UintStream:
    return C0 + C0 * uint(len(a.(UintStream)))
  case AnyStream:
    y := C0
    for _, b := range a.(AnyStream) {
      y += uint(codelen(b))
    }
    return y
  case Object:
    return (a.(Object)).Codelen()
  }
  fail (a)
  panic("shut up, compiler")
}

func encode (a Any) Stream {
  if a == nil {
    return nil
  }
  var s Stream
  switch a.(type) {
  case bool:
    s = make (Stream, 1)
    if a.(bool) { s[0]++ }
  case int8:
    s = make (Stream, 1)
    s[0] = a.(byte)
  case int16:
    n, x := 2, a.(int16)
    s = make (Stream, n)
    for i := 0; i < n; i++ {
      s[i] = byte(x)
      x >>= 8
    }
  case int32:
    n, x := 4, a.(int32)
    s = make (Stream, n)
    for i := 0; i < n; i++ {
      s[i] = byte(x)
      x >>= 8
    }
  case int:
    n, x := codelen(int(0)), a.(int)
    s = make (Stream, n)
    for i := uint(0); i < n; i++ {
      s[i] = byte(x)
      x >>= 8
    }
  case int64:
    n, x := 8, a.(int64)
    s = make (Stream, n)
    for i := 0; i < n; i++ {
      s[i] = byte(x)
      x >>= 8
    }
  case uint8:
    s = make (Stream, 1)
    s[0] = a.(uint8)
  case uint16:
    n, x := 2, a.(uint16)
    s = make (Stream, n)
    for i := 0; i < n; i++ {
      s[i] = byte(x)
      x >>= 8
    }
  case uint32:
    n, x := 4, a.(uint32)
    s = make (Stream, n)
    for i := 0; i < n; i++ {
      s[i] = byte(x)
      x >>= 8
    }
  case uint:
    x := a.(uint)
    s = make (Stream, C0)
    for i := uint(0); i < C0; i++ {
      s[i] = byte(x)
      x >>= 8
    }
  case uint64:
    n, x := 8, a.(uint64)
    s = make (Stream, n)
    for i := 0; i < n; i++ {
      s[i] = byte(x)
      x >>= 8
    }
  case float32:
    n, x := 4, math.Float32bits (a.(float32))
    s = make (Stream, n)
    for i := 0; i < n; i++ {
      s[i] = byte(x)
      x >>= 8
    }
  case float64:
    n, x := 8, math.Float64bits (a.(float64))
    s = make (Stream, n)
    for i := 0; i < n; i++ {
      s[i] = byte(x)
      x >>= 8
    }
  case complex64:
    n, c := 8, a.(complex64)
    x, y := math.Float32bits (real(c)), math.Float32bits (imag(c))
    s = make (Stream, n)
    for i := 0; i < n / 2; i++ {
      s[i] = byte(x)
      s[i + n/2] = byte(y)
      x >>= 8; y >>= 8
    }
  case complex128:
    n, c := 16, a.(complex128)
    x, y := math.Float64bits (real(c)), math.Float64bits (imag(c))
    s = make (Stream, n)
    for i := 0; i < n / 2; i++ {
      s[i] = byte(x)
      s[i + n/2] = byte(y)
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
    copy (s, ys)
  case Stream:
    return a.(Stream)
  case IntStream:
    is := a.(IntStream)
    n := len(is)
    c := int(C0)
    s = make(Stream, c * (n + 1))
    copy (s[:c], encode(n))
    i := c
    for j := 0; j < n; j++ {
      copy (s[i:i+c], encode(is[j]))
      i += c
    }
  case UintStream:
    u := a.(UintStream)
    n := uint(len(u))
    c := C0
    s = make(Stream, c * (n + 1))
    copy (s[:c], encode(n))
    i := c
    for j := uint(0); j < n; j++ {
      copy (s[i:i+c], encode(u[j]))
      i += c
    }
  case AnyStream:
    as := a.(AnyStream)
    n := uint(len(as))
    c := C0
    for j := uint(0); j < n; j++ {
      c += C0 + codelen(as[j])
    }
    s = make (Stream, c)
    copy (s[:C0], encode(n))
    i := C0
    for j := uint(0); j < n; j++ {
//      bs[i] = gödel(as[j])
//      i++
      k := codelen(as[j])
      copy(s[i:i+C0], encode(k))
      i += C0
      copy(s[i:i+k], encode(as[j]))
      i += k
    }
  case Object:
    return a.(Object).Encode()
  default:
    fail (a)
  }
  return s
}

/*/
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
  case BoolStream:
    return 17
  case Stream:
    return 18
  case IntStream:
    return 19
  case UintStream:
    return 20
  case AnyStream:
    return 21
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
  case 14:
    return complex64(0)
  case 15:
    return complex128(0)
  case 16:
    return ""
  case 17:
    return new(BoolStream)
  case 18:
    return new(Stream)
  case 19:
    return new(IntStream)
  case 21:
    return new(UintStream)
  case 22:
    return new(AnyStream)
  case 255:
    return Object.Clr()
  }
  return nil
}
/*/

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
    n, x := int(C0), 0
    chk (bs, n)
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
    n, x := int(C0), uint(0)
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
  case IntStream:
    n := decode (0, bs[:C0]).(int)
    us := make(IntStream, n)
    i := C0
    for j := 0; j < n; j++ {
      us[j] = decode (0, bs[i:i+C0]).(int)
      i += C0
    }
    return us
  case UintStream:
    n := decode (uint(0), bs[:C0]).(uint)
    us := make(UintStream, n)
    i := C0
    for j := uint(0); j < n; j++ {
      us[j] = decode (uint(0), bs[i:i+C0]).(uint)
      i += C0
    }
    return us
  case AnyStream:
    n := decode (uint(0), bs[:C0]).(uint)
    as := make(AnyStream, n)
    i := C0
    for j := uint(0); j < n; j++ {
      k := decode (uint(0), bs[i:i+C0]).(uint)
      i += C0
      as[j] = bs[i:i+k] // any client has to decode this Stream himself
      i += k
    }
    return as
  case Coder:
    chk (bs, int(a.(Coder).Codelen()))
    a.(Coder).Decode (bs)
  case Object:
    chk (bs, int(a.(Object).Codelen()))
    a.(Object).Decode (bs)
  default:
    fail (a)
  }
  return a
}

func encode2 (a, b int) Stream {
  s := make (Stream, 16)
  copy (s[0:8], encode (a))
  copy (s[8:16], encode (b))
  return s
}

func decode2 (s Stream) (int, int) {
  a := decode (0, s[0:8]).(int)
  b := decode (0, s[8:16]).(int)
  return a, b
}

func encodes (as AnyStream, c []uint) Stream {
  n := uint(0)
  for _, b:= range c {
    n += b
  }
  s := make (Stream, n)
  a := uint(0)
  for i, x := range as {
    copy (s[a:a+c[i]], encode(x))
    a += c[i]
  }
  return s
}

func decodes (b Stream, as AnyStream, c []uint) {
  a := uint(0)
  for i, x := range as {
    as[i] = decode (x, b[a:a+c[i]])
    a += c[i]
  }
}
