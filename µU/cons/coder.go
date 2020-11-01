package cons

// (c) Christian Maurer   v. 201014 - license see µU.go

import
  . "µU/obj"

func (X *console) Codelen (w, h uint) uint {
  return 4 * uint(4) + w * h * colourdepth
}

func (X *console) Encode (x, y, w, h uint) Stream {
  if w == 0 || h == 0 || w > X.wd || h > X.ht {
    return Stream(nil)
  }
  s := make (Stream, X.Codelen (w, h))
  j := 4 * uint(4)
  copy (s[:j], Encode4 (uint32(x), uint32(y), uint32(w), uint32(h)))
  i := (width * y + x) * colourdepth
  di := width * colourdepth
  dj := w * colourdepth
  for k := y; k < y + h; k++ {
    copy (s[j:j+dj], fbcop[i:i+dj])
    i += di
    j += dj
  }
  return s
}

func (X *console) Decode (s Stream) {
  if s == nil { return }
  if ! visible { return }
  j := 4 * uint(4)
  x, y, w, h := Decode4 (s[:j])
  di := width * colourdepth
  i := uint(y) * di
  i += uint(x) * colourdepth // TODO WHY correction factor: += x * colourdepth * colourdepth / 3   ?
  dj := uint(w) * colourdepth
  for k := uint(0); k < uint(h); k++ {
    copy (fbmem[i:i+dj], s[j:j+dj])
    copy (fbcop[i:i+dj], s[j:j+dj])
    i += di
    j += dj
  }
}
