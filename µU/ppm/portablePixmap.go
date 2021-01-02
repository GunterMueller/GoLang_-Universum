package ppm

// (c) Christian Maurer   v. 201231 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/scr"
  "µU/errh"
  "µU/pseq"
)
const (
  lf = byte(10)
  suffix = ".ppm"
)
type
  ppm struct {
     header string
  w, h, max,
     lh, ls int
            Stream
     fg, bg col.Colour
            pseq.PersistentSequence
       name string
            }

func new_() PPM {
  p := new(ppm)
  p.Stream = make(Stream, 0)
  return p
}

func (p *ppm) Set (w, h uint) {
  p.w, p.h, p.max = 3 * int(w), int(h), 255
  p.header = scr.P6Header (w, h)
  p.lh = len(p.header)
  p.ls = p.lh + p.w * p.h + 1
  p.Stream = make(Stream, p.ls)
  copy (p.Stream[:p.lh], Stream(p.header))
  p.Stream[p.ls-1] = lf
  p.PersistentSequence = pseq.New (p.Stream)
}

func (p *ppm) ColourF (c col.Colour) {
  p.fg = c
}

func (p *ppm) Point (x, y int) {
  if x < 0 || 3 * x >= p.w || y < 0 || y >= p.h { return }
  i := p.lh + y * p.w + 3 * x
  copy (p.Stream[i:i+3], p.fg.Encode())
}

func (p *ppm) Name (n string) {
  p.name = n
  p.PersistentSequence.Name (p.name + suffix)
}

func (p *ppm) Rename (n string) {
  if n == p.name { return }
  p.name = n
  p.PersistentSequence.Rename (p.name + suffix)
}

func (p *ppm) Write() {
  c, c0 := col.New(), col.New()
  i := p.lh
  for y := 0; y < p.h; y++ {
    for x := 0; x < p.w / 3; x++ {
      c.Decode (p.Stream[i:i+3])
      if ! c.Eq (c0) {
        scr.ColourF (c)
        c0.Copy (c)
      }
      scr.Point (x, y)
      i += 3
    }
  }
}

func (p *ppm) Get (n string) {
  errh.Hint (errh.ToWait)
  p.ls = int(pseq.Length (n + suffix))
  s := make(Stream, p.ls)
  p.PersistentSequence = pseq.New (s)
  p.Name (n)
  p.Stream = p.PersistentSequence.Get().(Stream)
  w, h, m, i := scr.HeaderData (p.Stream)
  p.w, p.h, p.max, p.lh = 3 * int(w), int(h), int(m), i + 1
  errh.DelHint()
}

func (p *ppm) Put() {
  p.PersistentSequence.Put (p.Stream)
}

func (p *ppm) Fin() {
  p.PersistentSequence.Fin()
}
