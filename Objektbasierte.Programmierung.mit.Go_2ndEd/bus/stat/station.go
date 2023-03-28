package stat

// (c) Christian Maurer   v. 230114 - license see µU.go

import (
  . "µU/obj"
  "µU/time"
  "µU/linewd"
  "µU/str"
  "µU/col"
  "µU/scr"
  "µU/scale"
  "bus/line"
)
const (
  dB =  67.62 // km pro Breitengrad bei 52.5° Breite
  dL = 111.13 // km pro Längengrad
)
type
  station struct {
          bg, lg float64 // Position (Breiten- und Längengrad)
            line line.Line
                 uint // interne Nummer
         umstieg bool
                 string "Name"
    beschriftung byte // Positionierung des Namens auf der Graphik
                 }
const
  max = 32 // maximale Länge der Namen der Bahnhöfe
var
  initialized bool

func init_() {
  if initialized { return }
  initialized = true
  scale.Init()
  scale.SetRange (13.2610 * dB, 52.4550 * dL, 16) // 16)
  scale.Lim (871, 5800, 72, 940, 6)
  scr.Transparence (true)
  scr.MousePointer (true)
}

func (x *station) imp (Y any) *station {
  y, ok := Y.(*station)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func New() Station {
  x := new (station)
  x.beschriftung = '.'
  x.bg, x.lg = 52.5, 13.5
  init_()
  return x
}

func (x *station) Clr() {
  x.line, x.uint = line.Footpath, 0
  x.umstieg = false
  x.string = ""
  x.beschriftung = '.'
  x.bg, x.lg = 52.5, 13.5
}

func (x *station) Empty() bool {
  return x.line == line.Footpath &&
         x.uint == 0
}

func (x *station) Set (l line.Line, n uint, s string, b byte, yy, xx float64) {
  x.line, x.uint, x.string = l, n, s
  switch b {
  case 'o', 'l', 'u', 'r', 'm':
    x.beschriftung = b
  default:
    panic ("affe")
  }
  x.bg, x.lg = xx, yy
}

func (x *station) Pos() (float64, float64) {
  return x.bg, x.lg
}

func (x *station) Line() line.Line {
  return x.line
}

func (x *station) Number() uint {
  return x.uint
}

func (x *station) Umstieg() {
//  x.umstieg = true
}

func (x *station) Equiv (Y any) bool {
  y, ok := Y.(*station)
  if ! ok { return false }
  return x.uint != y.uint &&
         x.bg == y.bg &&
         x.lg == y.lg
}

func (x *station) Renumber (l line.Line, n uint) {
  x.line, x.uint = l, n
  x.umstieg = false
}

func (x *station) EditScale() {
  scale.Edit()
}

func (x *station) UnderMouse() bool {
  xm, ym := scr.MousePosGr()
  xx, yy := scale.Scale (x.bg * dB, x.lg * dL)
  dx, dy := xm - xx, ym - yy
  if dx < 0 { dx = -dx }
  if dy < 0 { dy = -dy }
  const d = 8 // pixel
  return dx <= d && dy <= d
}

func (x *station) Eq (Y any) bool {
  y := x.imp (Y)
  return x.line == y.line &&
         x.uint == y.uint
}

func (x *station) Less (Y any) bool {
  return false
}

func (x *station) Leq (Y any) bool {
  return false
}

func (x *station) Copy (Y any) {
  y := x.imp (Y)
  x.line, x.uint = y.line, y.uint
  x.umstieg = y.umstieg
  x.string = y.string
  x.beschriftung = y.beschriftung
  x.lg, x.bg = y.lg, y.bg
}

func (x *station) Clone() any {
  y := New()
  y.Copy (x)
  return y
}

func (x *station) Write (aktuell bool) {
  xx, yy := scale.Scale (x.bg * dB, x.lg * dL)
  if aktuell {
    scr.Colours (line.Colour[x.line], scr.ScrColB())
  } else { // umstieg
    scr.Colours (col.Black(), scr.ScrColB())
  }
  const r = 1
  lw := scr.ActLinewidth()
  scr.SetLinewidth (linewd.Thin)
  if xx >= r && yy >= r {
    if x.line == line.Zoo || x.line == line.BG {
      scr.ColourF (line.Colour [x.line])
    }
    scr.Circle (xx, yy, r)
  }
  scr.SetLinewidth (lw)
  n := int(str.ProperLen (x.string))
  if n <= 2 { return }
  xn, yn := 0, 0
  w, h := int(scr.Wd1()), int(scr.Ht1())
  switch x.beschriftung {
  case L:
    xn, yn = xx - n * w - w + 1, yy - h / 2
  case R:
    xn, yn = xx + w + 1, yy - h / 2
  case O:
    xn, yn = xx - (n * w) / 2 + 1, yy - h
  case U:
    xn, yn = xx - (n * w) / 2 + 1, yy + h / 2
  }
  if aktuell {
    scr.Colours (line.Colour[x.line], scr.ScrColB())
  } else { // umstieg
    scr.Colours (col.Black(), scr.ScrColB())
    scr.Colours (col.Black(), col.Pink())
  }
  scr.Colours (col.Black(), col.LightWhite())
  x.string = str.Lat1 (x.string)
  str.ReplaceAll (&x.string, '_', " ")
  scr.WriteGr (x.string, xn, yn)
}

func (x *station) Blink() {
  const t = 100
  for i := 0; i < 10; i++ {
    x.Write (true)
    time.Msleep (t)
    x.Write (false)
    time.Msleep (t)
  }
}

func (x *station) Codelen() uint {
  return 0
}

func (x *station) Encode() Stream {
  return make (Stream, 0)
}

func (x *station) Decode (b Stream) {
}
