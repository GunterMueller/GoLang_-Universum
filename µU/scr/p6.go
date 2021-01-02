package scr

// (c) Christian Maurer   v. 201231 - license see µU.go

import (
  "µU/obj"
  "µU/char"
  "µU/col"
)
const
  lf = string(byte(10))

func string_(n uint) string {
  if n == 0 { return "0" }
  var s string
  for s = ""; n > 0; n /= 10 {
    s = string(n % 10 + '0') + s
  }
  return s
}

func number (s obj.Stream) (uint, int) {
  n := uint(0)
  i := 0
  for char.IsDigit (s[i]) { i++ }
  for j := 0; j < i; j++ {
    n = 10 * n + uint(s[j] - '0')
  }
  return n, i
}

func (X *screen) P6Header (w, h uint) string {
  return "P6 " + string_(w) + " " + string_(h) + " 255" + lf
}

func (X *screen) P6Codelen (w, h uint) uint {
  return uint(len(X.P6Header (w, h))) + h * 3 * 4 * (w + 1)
}

func (X *screen) P6Size (s obj.Stream) (uint, uint) {
  return P6Size (s)
}

func (X *screen) P6Encode (x, y, w, h uint) obj.Stream {
  bs := X.Encode (x, y, w, h)
  i := 4 * 4
  cl := X.P6Codelen (w, h)
  s := make (obj.Stream, cl)
  head := X.P6Header (w, h)
  j := len(head)
  copy (s[:j], obj.Stream(head))
  if w == 0 || h == 0 { return s[:j] }
  di := int(col.Depth()) // console
  if underX {
    di = 4
  }
  for y := uint(0); y < h; y++ {
    for x := uint(0); x < w; x++ {
      for k := 0; k < 3; k++ {
        s[j+k] = bs[i+k]
      }
      i += di
      j += 3
    }
  }
  return s
}

func headerData (s obj.Stream) (uint, uint, uint, int) {
  if string(s[:2]) != "P6" { panic ("wrong ppm-header: " + string(s[:2])) }
  i := 3
  w, dw := number (s[i:])
  i += dw + 1
  h, dh := number (s[i:])
  i += dh + 1
  m, dm := number (s[i:])
  i += dm
  return w, h, m, i
}

func (X *screen) P6Decode (s obj.Stream, x, y uint) {
  w, h, _, j := headerData (s)
  if w == 0 || h == 0 || w > X.Wd() || h > X.Ht() { return }
  i := 4 * 4
  bs := make (obj.Stream, X.Codelen (w, h))
  copy (bs[:i], obj.Encode4 (uint32(x), uint32(y), uint32(w), uint32(h)))
  di := int(col.Depth()) // console
  if underX {
    di = 4
  }
  c := col.New()
  for y := uint(0); y < h; y++ {
    for x := uint(0); x < w; x++ {
      c.Decode (s[j:j+3])
      cs := c.Cstream() // Console
      if underX {
        cs = obj.Encode (c.Code())
      }
      copy (bs[i:i+di], cs)
      i += di
      j += 3
    }
  }
  X.Decode (bs)
}
