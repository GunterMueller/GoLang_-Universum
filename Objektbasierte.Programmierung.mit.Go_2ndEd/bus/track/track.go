package track

// (c) Christian Maurer   v. 230114 - license see µU.go

import (
  . "µU/obj"
  "µU/linewd"
  "µU/col"
  "µU/scr"
  "µU/scale"
  "bus/line"
)
const
  dB, dL = 67.62, 111.13 // km pro Breiten- bzw. Längengrad bei 52.5° Breite
type
  track struct {
               line.Line
  x, y, x1, y1 float64 // Positionen der Bahnhöfe
               uint // mittlere Fahrzeit in Minuten
               }
const (
  lgMax = 13.9084 // Strausberg Nord
  lgMin = 13.0671 // Potsdam
  lgM   = 13.4877
  bgMax = 52.7536 // Oranienburg
  bgMin = 52.2964 // Königs-Wusterhausen
  bgM   = 52.4938
)

func init() {
  scale.SetRange (lgM * dB, bgM * dL, 16)
  scale.Lim (lgMin * dB, bgMin * dL, lgMax * dB, bgMax * dL, 18)
}

func New() Track {
  return new (track)
}

func (x *track) imp (Y any) *track {
  y, ok := Y.(*track)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *track) Def (l line.Line, n uint) {
  x.Line, x.uint = l, n
}

func (x *track) Empty() bool {
  return x.uint > 0
}

func (x *track) Clr() {
  x.uint = 0
}

func (x *track) Eq (Y any) bool {
  y := x.imp (Y)
  return x.Line == y.Line &&
         x.x == y.x && x.y == y.y &&
         x.x1 == y.x1 && x.y1 == y.y1 &&
         x.uint == y.uint
}

func (x *track) Less (Y any) bool {
  return false
}

func (x *track) Leq (Y any) bool {
  return false
}

func (x *track) Copy (Y any) {
  y := x.imp (Y)
  x.Line = y.Line
  x.x, x.y = y.x, y.y
  x.x1, x.y1 = y.x1, y.y1
  x.uint = y.uint
}

func (x *track) Clone() any {
  y := New()
  y.Copy (x)
  return y
}

func (x *track) Val() uint {
  return x.uint
}

func (x *track) SetVal (n uint) {
  if n < 10 {
    x.uint = n
  }
}

func (x *track) SetPos (x0, y0, x1, y1 float64) {
  x.x, x.y, x.x1, x.y1 = x0, y0, x1, y1
}

func (x *track) Write (aktuell bool) {
  if aktuell || x.Line == line.Footpath {
    scr.ColourF (line.Colour [x.Line])
  } else {
    scr.ColourF (col.Black())
  }
  if aktuell {
    scr.SetLinewidth (linewd.Thicker)
  } else {
    scr.SetLinewidth (linewd.Thin)
  }
  x0, y0 := scale.Scale (x.x * dB, x.y * dL)
  x1, y1 := scale.Scale (x.x1 * dB, x.y1 * dL)
  if aktuell {
    scr.SetLinewidth (linewd.Thicker) 
  } else {
    scr.SetLinewidth (linewd.Thin) 
  }
// if x.Line == line.BG { return }
  if x.Line == line.Zoo || x.Line == line.BG {
    f := line.Colour [x.Line]
    scr.ColourF (f)
    scr.SetLinewidth (linewd.Fat) 
  }
  scr.Line (x0, y0, x1, y1)
  scr.SetLinewidth (linewd.Thin) 
}

func (x *track) Codelen() uint {
  return 0
}

func (x *track) Encode() Stream {
  return make (Stream, x.Codelen())
}

func (x *track) Decode (s Stream) {
}
