package obj

// (c) murus.org  v. 161229 - license see murus.go

import (
  "runtime"
  "math"
  "strconv"
  "murus/ker"
)
const
  z32 = uint32(0)
type
  Coder interface {

// Returns the number of bytes, that are needed
// to serialize x uniquely revertibly.
  Codelen() uint

// x.Eq (x.Decode (x.Encode())
  Encode() []byte

// Pre: b is result of y.Encode() for some y of the type of x.
// x.Eq(y); x.Encode() == b, i.e. those slices coincide.
  Decode (b []byte)
}

func fail (a Any) {
  ker.Panic ("murus only [en|de]codes atomic types and objects of string, []byte, []bool or Object !")
}

func Codelen (a Any) uint {
  if a == nil { return 0 }
  switch a.(type) { case Object:
    return (a.(Object)).Codelen()
  case string:
    return uint(len(a.(string)))
  case []bool:
    return uint(len(a.([]bool)))
  case []byte:
    return uint(len(a.([]byte)))
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
  }
  fail (a)
  panic("shut up, compiler")
}

func chk (b []byte, n int) {
  if len(b) < n { // != n {
    ker.Panic ("murus/obj/coder.go chk: len(b) == " + strconv.Itoa(len(b)) + " < n == " + strconv.Itoa(n))
  }
}

func Encode (a Any) []byte {
  if a == nil { return nil }
  var bs []byte
  switch a.(type) {
  case bool:
    bs = make ([]byte, 1)
    if a.(bool) { bs[0]++ }
  case int8:
    bs = make ([]byte, 1)
    bs[0] = a.(byte)
  case int16:
    n, x:= 2, a.(int16)
    bs = make ([]byte, n)
    for i:= 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case int32:
    n, x:= 4, a.(int32)
    bs = make ([]byte, n)
    for i:= 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case int:
    n, x:= Codelen(int(0)), a.(int)
    bs = make ([]byte, n)
    for i:= uint(0); i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case int64:
    n, x:= 8, a.(int64)
    bs = make ([]byte, n)
    for i:= 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case uint8:
    bs = make ([]byte, 1)
    bs[0] = a.(uint8)
  case uint16:
    n, x:= 2, a.(uint16)
    bs = make ([]byte, n)
    for i:= 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case uint32:
    n, x:= 4, a.(uint32)
    bs = make ([]byte, n)
    for i:= 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case uint:
    n, x:= Codelen(uint(0)), a.(uint)
    bs = make ([]byte, n)
    for i:= uint(0); i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case uint64:
    n, x:= 8, a.(uint64)
    bs = make ([]byte, n)
    for i:= 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case float32:
    n, x:= 4, math.Float32bits (a.(float32))
    bs = make ([]byte, n)
    for i:= 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case float64:
    n, x:= 8, math.Float64bits (a.(float64))
    bs = make ([]byte, n)
    for i:= 0; i < n; i++ {
      bs[i] = byte(x)
      x >>= 8
    }
  case complex64:
    n, c:= 8, a.(complex64)
    x, y:= math.Float32bits (real(c)), math.Float32bits (imag(c))
    bs = make ([]byte, n)
    for i:= 0; i < n / 2; i++ {
      bs[i] = byte(x)
      bs[i + n/2] = byte(y)
      x >>= 8; y >>= 8
    }
  case complex128:
    n, c:= 16, a.(complex128)
    x, y:= math.Float64bits (real(c)), math.Float64bits (imag(c))
    bs = make ([]byte, n)
    for i:= 0; i < n / 2; i++ {
      bs[i] = byte(x)
      bs[i + n/2] = byte(y)
      x >>= 8; y >>= 8
    }
  case string:
    return ([]byte)(a.(string))
  case []byte:
    bs = make ([]byte, len (a.([]byte)))
    copy (bs, a.([]byte))
  case []bool:
    n := len (a.([]bool))
    ys := make ([]byte, n)
    for i := 0; i < n; i++ {
      ys[i] = 0; if a.([]bool)[i] { ys[i] = 1 }
    }
    copy (bs, ys)
  case Object:
    return a.(Object).Encode()
  default:
    fail (a)
  }
  return bs
}

func Decode (a Any, bs []byte) Any {
  if a == nil { return nil }
  switch a.(type) {
  case bool:
    chk (bs, 1)
    return bs[0] > 0
  case int8:
    chk (bs, 1)
    return int8(bs[0])
  case int16:
    n, x:= 2, int16(0)
    chk (bs, n)
    for i:= n - 1; i >= 0; i-- {
      x <<= 8
      x += int16(bs[i])
    }
    return x
  case int32:
    n, x:= 4, int32(0)
    chk (bs, n)
    for i:= n - 1; i >= 0; i-- {
      x <<= 8
      x += int32(bs[i])
    }
    return x
  case int:
    n, x:= int(Codelen(0)), 0
    chk (bs, int(n))
    for i:= n - 1; i >= 0; i-- {
      x <<= 8
      x += int(bs[i])
    }
    return x
  case int64:
    n, x:= 8, int64(0)
    chk (bs, n)
    for i:= n - 1; i >= 0; i-- {
      x <<= 8
      x += int64(bs[i])
    }
    return x
  case uint8:
    chk (bs, 1)
    return bs[0]
  case uint16:
    n, x:= 2, uint16(0)
    chk (bs, n)
    for i:= n - 1; i >= 0; i-- {
      x <<= 8
      x += uint16(bs[i])
    }
    return x
  case uint32:
    n, x:= 4, z32
    chk (bs, n)
    for i:= n - 1; i >= 0; i-- {
      x <<= 8
      x += uint32(bs[i])
    }
    return x
  case uint:
    n, x:= int(Codelen(0)), uint(0)
    chk (bs, int(n))
    for i:= n - 1; i >= 0; i-- {
      x <<= 8
      x += uint(bs[i])
    }
    return x
  case uint64:
    n, x:= 8, uint64(0)
    chk (bs, n)
    for i:= n - 1; i >= 0; i-- {
      x <<= 8
      x += uint64(bs[i])
    }
    return x
  case float32:
    n, x:= 4, z32
    chk (bs, n)
    for i:= n - 1; i >= 0; i-- {
      x <<= 8
      x += uint32(bs[i])
    }
    return math.Float32frombits (x)
  case float64:
    n, x:= 8, uint64(0)
    chk (bs, n)
    for i:= n - 1; i >= 0; i-- {
      x <<= 8
      x += uint64(bs[i])
    }
    return math.Float64frombits (x)
  case complex64:
    n, x, y:= 8, z32, z32
    chk (bs, n)
    for i:= n / 2 - 1; i >= 0; i-- {
      x <<= 8; y <<= 8
      x += uint32(bs[i])
      y += uint32(bs[i + n/2])
    }
    return complex (math.Float32frombits (x), math.Float32frombits (y))
  case complex128:
    n, x, y:= 16, uint64(0), uint64(0)
    chk (bs, n)
    for i:= n / 2 - 1; i >= 0; i-- {
      x <<= 8; y <<= 8
      x += uint64(bs[i])
      y += uint64(bs[i + n/2])
    }
    return complex (math.Float64frombits (x), math.Float64frombits (y))
  case string:
    return string(bs)
  case []byte:
    copy (a.([]byte), bs)
  case []bool:
    n := len(bs)
    y := make([]bool, n)
    for i := 0; i < n; i++ {
      if bs[i] == 1 { y[i] = true }
    }
    return y
  case Object:
    chk (bs, int(a.(Object).Codelen()))
    a.(Object).Decode (bs)
  default:
    fail (a)
  }
  return a
}

func Encode4 (a, b, c, d uint32) []byte {
  bs:= make ([]byte, 4 * 4)
  copy (bs[:4], Encode (a))
  copy (bs[4:8], Encode (b))
  copy (bs[8:12], Encode (c))
  copy (bs[12:16], Encode (d))
  return bs
}

func Decode4 (bs []byte) (uint32, uint32, uint32, uint32) {
  a:= Decode (z32, bs[:4]).(uint32)
  b:= Decode (z32, bs[4:8]).(uint32)
  c:= Decode (z32, bs[8:12]).(uint32)
  d:= Decode (z32, bs[12:16]).(uint32)
  return a, b, c, d
}

func Encodes (as []Any, cs []uint) []byte {
  l:= uint(0); for _, b:= range cs { l += b }
  bs:= make ([]byte, l)
  a:= uint(0)
  for i, x:= range as {
    copy (bs[a:a+cs[i]], Encode(x))
    a += cs[i]
  }
  return bs
}

func Decodes (bs []byte, as []Any, cs []uint) {
  l:= uint(0); for _, b:= range cs { l += b }
  a:= uint(0)
  for i, x:= range as {
    as[i] = Decode (x, bs[a:a+cs[i]])
    a += cs[i]
  }
}
