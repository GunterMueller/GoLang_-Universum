package scr

// (c) Christian Maurer   v. 210107 - license see µU.go

import (
  "µU/obj"
  "µU/char"
  "µU/col"
)

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

func (X *screen) PPMHeader (w, h uint) string {
  s := "P6 " + string_(w) + " " + string_(h) + " 255" + string(byte(10))
  X.ppmheader = s
  X.lh = uint(len(s))
  return s
}

func (X *screen) PPMCodelen (w, h uint) uint {
  X.PPMHeader (w, h)
  return X.lh + 3 * w * h
}

func (X *screen) PPMSize (s obj.Stream) (uint, uint) {
  w, h, _, _ := ppmHeaderData (s)
  return w, h
}

func (X *screen) PPMEncode (x0, y0, w, h uint) obj.Stream {
  s := X.Encode (x0, y0, w, h)
  return append (obj.Stream(X.PPMHeader (w, h)), s[2*4:]...)
}

func ppmHeaderData (s obj.Stream) (uint, uint, uint, int) {
  p := string(s[:2]); if p != "P6" { panic ("wrong ppm-header: " + p) }
  i := 3
  w, dw := number (s[i:])
  i += dw + 1
  h, dh := number (s[i:])
  i += dh + 1
  m, dm := number (s[i:])
  i += dm
  return w, h, m, i + 1
}

func (X *screen) PPMDecode (s obj.Stream, x0, y0 uint) {
  w, h, _, j := ppmHeaderData (s)
  if w == 0 || h == 0 || w > X.Wd() || h > X.Ht() { return }
  i := 4 * uint(2)
  l := i + 3 * w * h
  e := make(obj.Stream, l)
  copy (e[:i], obj.Encode4 (uint16(x0), uint16(y0), uint16(w), uint16(h)))
  if underX {
    c := col.New()
    for y := uint(0); y < h; y++ {
      for x := uint(0); x < w; x++ {
        c.Decode (s[j:j+3])
        copy (e[i:i+3], obj.Encode (c.Code()))
        i += 3
        j += 3
      }
    }
  } else { // console
    copy (e[i:], s[j:])
  }
  X.Decode (e)
}
