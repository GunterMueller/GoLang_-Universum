package scr

// (c) murus.org  v. 140529 - license see murus.go

import (
  "murus/ker"
  "murus/obj"
  "murus/col"
)

func (X *screen) P6Codelen (w, h uint) uint {
  return ker.T6 + w * h * col.P6
}

func (X *screen) P6Size (bs []byte) (uint, uint) {
  return ker.P6Size (bs)
}

func (X *screen) P6Encode (x, y, w, h uint) []byte {
  bs:= X.Encode (x, y, w, h)
  i:= 4 * 4
  _, _, w1, h1:= obj.Decode4 (bs[:i]); if uint(w1) != w || uint(h1) != h { ker.Panic ("cons.P6Encode bug") }
  cl:= X.P6Codelen (w, h)
  ps:= make ([]byte, cl)
  j:= ker.P6Txt (w, h, ps)
  if w == 0 || h == 0 { return ps[:j] }
  di:= int(col.Depth) // Console
  if underX { di = 4 }
  dj:= col.P6
  for y:= uint(0); y < h; y++ {
    for x:= uint(0); x < w; x++ {
      col.P6Encode (bs[i:i+di], ps[j:j+dj])
      i += di
      j += dj
    }
  }
  return ps
}

func (X *screen) P6Decode (x, y uint, ps []byte) {
  w, h, fix, j:= ker.P6dec (ps)
  if w == 0 || h == 0 || fix != 255 || w > X.Wd() || h > X.Ht() { return }
  i:= 4 * 4
  bs:= make ([]byte, X.Codelen (w, h))
  copy (bs[:i], obj.Encode4 (uint32(x), uint32(y), uint32(w), uint32(h)))
  di:= int(col.Depth) // Console
  if underX { di = 4 }
  dj:= col.P6
  var c col.Colour
  for y:= uint(0); y < h; y++ {
    for x:= uint(0); x < w; x++ {
      col.Decode (&c, ps[j:j+dj])
      cc:= col.Cc(c) // Console
      if underX { cc = obj.Encode (col.Code(c)) }
      copy (bs[i:i+di], cc)
      i += di
      j += dj
    }
  }
  X.Decode (bs)
}
