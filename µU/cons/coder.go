package cons

// (c) Christian Maurer   v. 210106 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)

func (X *console) Codelen (w, h uint) uint {
  return 2 * uint(4) + colourdepth * w * h
}

func (X *console) Encode (x, y, w, h uint) []byte {
  s := make (Stream, X.Codelen (w, h))
  i := 2 * uint(4)
  copy (s[:i], Encode4 (uint16(x), uint16(y), uint16(w), uint16(h)))
  for l := X.y; l < X.y + int(h); l++ {
    j := colourdepth * width * uint(l)
    for c := X.x; c < X.x + int(w); c++ {
      copy (s[i:i+3], fbmem[j:j+3])
      i += 3
      j += colourdepth
    }
  }
  return s
}

func (X *console) Decode (s Stream) {
  if s == nil { return }
  if ! visible { return }
  i := 2 * uint(4)
  x0, y0, w, h := Decode4 (s[:i])
  c := col.New()
  for y := int(y0); y < int(y0 + h); y++ {
    for x := int(x0); x < int(x0 + w); x++ {
      c.Set (s[i], s[i+1], s[i+2])
      X.cF, X.codeF = c, c.EncodeInv()
      X.Point (x, y)
      i += 3
    }
  }
}
