package scr

// (c) Christian Maurer   v. 220119 - license see µU.go

import (
  "µU/ker"
  "µU/obj"
  "µU/char"
)
var (
  ppmheader string
  ppmHeaderLength uint
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

func ppmSize (s obj.Stream) (uint, uint) {
  w, h, _, _ := ppmHeaderData (s)
  return w, h
}

func ppmHeader (w, h uint) string {
  s := "P6 " + string_(w) + " " + string_(h) + " 255" + string(byte(10))
  ppmheader = s
  ppmHeaderLength = uint(len(s))
  return s
}

func ppmHeaderData (s obj.Stream) (uint, uint, uint, int) {
  p := string(s[:2]); if p != "P6" { ker.Panic ("wrong ppm-header: " + p + "<") }
  i := 3
  if s[i] == '#' {
    for {
      i++
      if s[i] == byte(10) {
        i++
        break
      }
    }
  }
  w, dw := number (s[i:])
  i += dw + 1
  h, dh := number (s[i:])
  i += dh + 1
  m, dm := number (s[i:])
  i += dm
  return w, h, m, i + 1
}

func ppmCodelen (w, h uint) uint {
  ppmHeader (w, h)
  return ppmHeaderLength + 3 * w * h
}
