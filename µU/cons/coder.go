package cons

// (c) Christian Maurer   v. 201228 - license see µU.go

import
  . "µU/obj"

func (X *console) Codelen (w, h uint) uint {
  return 4 * uint(4) + w * h * colourdepth
}

func (X *console) Encode (x, y, w, h uint) Stream {
  if w == 0 || h == 0 || w > X.wd || h > X.ht {
    return make(Stream, 0)
  }
  s := make (Stream, X.Codelen (w, h))
  i, k := 4 * uint(4), uint(0)
  copy (s[:i], Encode4 (uint32(x), uint32(y), uint32(w), uint32(h)))
  j :=  colourdepth * (width * y + x)
  dj := colourdepth * width
  dw := colourdepth * w
  for n := y; n < y + h; n++ {
    k = i - 16
    copy (s[i:i+dw], fbcop[k:k+dw])
    i += dw
    j += dj
  }
  return s
}

func (X *console) Decode (s Stream) {
  if s == nil { return }
  if ! visible { return }
  j := 4 * uint(4)
  x4, y4, w4, h4 := Decode4 (s[:j])
  x, y, w, h := uint(x4), uint(y4), uint(w4), uint(h4)
  i := colourdepth * (width * y + x)
  di := colourdepth * width
  dw := colourdepth * w
  for n := uint(0); n < h; n++ {
    copy (fbmem[i:i+dw], s[j:i+dw])
    copy (fbcop[i:i+dw], s[j:i+dw])
    i += di
    j += dw
  }
}
