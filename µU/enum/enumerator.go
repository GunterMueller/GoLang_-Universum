package enum

// (c) Christian Maurer   v. 221021 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/col"
  "µU/kbd"
  "µU/str"
  "µU/scr"
  "µU/box"
  "µU/errh"
  "µU/pbox"
  "µU/sel"
  "µU/pseq"
)
const
  suffix = ".dat"
type
  enumerator struct {
        name, namek string // names of files
               s, k []string // strings and their shortcuts
                    uint // length of s
              w, wk uint // length of the s[i] and k[i]
                    string // selected s[i]
               f, b col.Colour
              hintk string
                    }
var (
  bx = box.New()
  pbx = pbox.New()
)

func new_(w uint) Enum {
  if w == 0 { ker.PrePanic() }
  x := new (enumerator)
  x.name = str.New (x.w)
  x.namek = ""
  x.s = make([]string, 0)
  x.w = w
  x.wk = 0
  x.string = str.New (x.w)
  x.f, x.b = col.LightWhite(), col.Black()
  return x
}

func newk (w, k uint) Enum {
  x := new_(w).(*enumerator)
  x.namek = x.name + "k"
  x.k = make([]string, 0)
  x.wk = k
  return x
}

func (x *enumerator) imp (Y any) *enumerator {
  y, ok := Y.(*enumerator)
  if ! ok { TypeNotEqPanic(x, Y) }
  return y
}

func (x *enumerator) Set (ss ...string) {
  x.uint = uint(len(ss))
  if x.uint == 0 { ker.PrePanic() }
  x.s = make([]string, x.uint)
  for i := uint(0); i < x.uint; i++ {
    if str.Empty(ss[i]) { ker.PrePanic() }
    x.s[i] = str.Lat1(ss[i])
    str.Norm (&x.s[i], x.w)
  }
}

func (x *enumerator) Setk (sk ...string) {
  if uint(len(sk)) != x.uint { ker.PrePanic() }
  x.k = make([]string, x.uint)
  x.hintk = ""
  for i := uint(0); i < x.uint; i++ {
    str.OffSpc(&sk[i])
    x.k[i] = str.Lat1(sk[i])
    x.hintk += x.k[i]
    x.hintk += " "
  }
  x.hintk += " (Auswahl: F1)"
}

func (x *enumerator) SetEdit (name string, l, c uint) {
  x.s = make([]string, 0)
  w := M
  if c + w > scr.NColumns() - 1 {
    w = scr.NColumns() - c - 1
  }
  x.w = w
  bx.Wd (w)
  bx.Colours (col.LightWhite(), col.Black())
  errh.Hint ("Zeichenketten eingeben")
  for {
    s := str.New (x.w)
    bx.Wd (x.w)
    bx.Edit (&s, l, c)
    if str.Empty (s) {
      errh.Error0 ("Zeichenkette leer")
    } else {
      if k, _ := kbd.LastCommand(); k == kbd.Esc {
        break
      }
      x.s = append (x.s, s)
    }
  }
  x.uint = uint(len(x.s))
  errh.DelHint()
  file := pseq.New (x.string)
  file.Name (x.name + suffix)
  file.Clr()
  s := string(byte(x.w))
  str.Norm (&s, x.w)
  file.Seek (0)
  file.Put (s)
  for i := uint(0); i < x.uint; i++ {
    file.Seek (i + 1)
    file.Put (x.s[i])
  }
}

func (x *enumerator) SetEditk (name string, l, c uint) {
  if x.wk == 0 { ker.PrePanic() }
  x.SetEdit (name, l, c)
  bx.Wd (x.wk)
  for {
    for {
      s := str.New (x.wk)
      bx.Edit (&s, l, c)
      str.OffSpc (&s)
      if uint(len(s)) == x.wk {
        x.k = append (x.k, s)
        break
      } else {
        errh.Error0 ("darf keine Leerzeichen enthalten")
      }
    }
  }
  filek := pseq.New (str.New (x.wk))
  filek.Name (x.namek + suffix)
  filek.Clr()
  s2 := string(byte(x.wk))
  str.Norm (&s2, x.wk)
  filek.Seek (0)
  filek.Put (s2)
  for i := uint(0); i < x.uint; i++ {
    filek.Seek (i + 1)
    filek.Put (x.k[i])
  }
}

func (x *enumerator) Get (name string) {
  filename := name + suffix
  if pseq.Length (filename) == 0 {
    ker.Panic ("die Datei " + filename + " existiert nicht")
  }
  file := pseq.New (x.string)
  file.Name (filename)
  n := file.Num()
  if n < 4 {
    ker.Panic ("die Datei " + filename + " ist zu kurz")
  }
  file.Seek (0)
  s := file.Get().(string)
  if x.w != uint(s[0]) { ker.Oops() }
  x.s = make([]string, x.uint - 1)
  for i := uint(0); i < x.uint; i++ {
    file.Seek (i + 1)
    x.s[i] = file.Get().(string)
  }
  if x.wk > 0 {
    filename = x.namek + suffix
    if pseq.Length (filename) == 0 {
      ker.Panic ("die Datei " + filename + " existiert nicht")
    }
    filek := pseq.New (str.New(x.wk))
    file.Name (filename)
    nk := filek.Num()
    if nk != n { ker.Shit() }
    filek.Seek (0)
    sk := file.Get().(string)
    if x.wk != uint(sk[0]) { ker.Oops() }
    x.k = make([]string, x.uint - 1)
    for i := uint(0); i < x.uint; i++ {
      filek.Seek (i + 1)
      x.k[i] = filek.Get().(string)
    }
  }
}

func (x *enumerator) Empty() bool {
  return str.Empty (x.string)
}

func (x *enumerator) Clr() {
  x.string = str.New (x.w)
}

func (x *enumerator) Copy (Y any) {
  y := x.imp(Y)
  x.name, x.namek = y.name, y.namek
  x.s = make([]string, x.uint)
  if x.wk > 0 {
    x.k = make([]string, x.uint)
  }
  for i := uint(0); i < x.uint; i++ {
    x.s[i] = y.s[i]
    if x.wk > 0 {
      x.k[i] = y.k[i]
    }
  }
  x.uint = y.uint
  x.w, x.wk = y.w, y.wk
  x.string = y.string
  x.f.Copy (y.f)
  x.b.Copy (y.b)
}

func (x *enumerator) Clone() any {
  y := new_(x.w)
  y.Copy (x)
  return y
}

func (x *enumerator) Eq (Y any) bool {
  y := x.imp(Y)
  return x.name == y.name && x.namek == y.namek &&
//       x.s[i]
//       x.k[i]
         x.uint == y.uint &&
         x.w == y.w && x.wk == y.wk &&
         x.string == y.string
}

func (x *enumerator) Less (Y any) bool {
  return str.Less (x.string, x.imp(Y).string)
}

func (x *enumerator) Leq (Y any) bool {
  return x.Less (Y) || x.Eq (Y)
}

func (x *enumerator) Codelen() uint {
  return x.w
}

func (x *enumerator) Encode() Stream {
  return Stream (x.string)
}

func (x *enumerator) Decode (s Stream) {
  x.string = string(s[0:x.w])
}

func (x *enumerator) Colours (f, b col.Colour) {
  x.f, x.b = f, b
}

func (x *enumerator) Cols() (col.Colour, col.Colour) {
  return x.f, x.b
}

func (x *enumerator) Write (l, c uint) {
  bx.Wd (x.w)
  bx.Colours (x.f, x.b)
  bx.Write (x.string, l, c)
}

func (x *enumerator) Select (l, c uint) {
  i := uint(0)
  sel.Select1 (x.s, x.uint, x.w, &i, l, c, x.f, x.b)
  if i == uint(len(x.s)) { // selection cancelled
    x.Clr()
  } else {
    x.string = x.s[i]
    x.Write (l, c)
  }
}

func (x *enumerator) Edit (l, c uint) {
  x.string = str.New (x.w)
  x.Write (l, c)
  bx.Wd (x.w)
  bx.Colours (x.f, x.b)
  if x.wk > 0 {
    errh.Hint (x.hintk)
  }
  for {
    s := str.New (x.w)
    bx.Edit (&s, l, c)
    if str.Empty (s) {
      switch cmd, d := kbd.LastCommand(); cmd {
      case kbd.Esc:
        x.Clr()
        goto X
      case kbd.Enter:
        if d == 0 {
          x.Clr()
          goto X
        } else {
          x.Select (l, c)
          goto X
        }
      case kbd.Help:
        x.Select (l, c)
        goto X
      }
    }
    for i := uint(0); i < x.uint; i++ {
      if s == x.s[i] {
        x.string = s
        goto X
      }
      if x.wk > 0 {
        if str.ProperLen (s) == x.wk && s[:x.wk] == x.s[i] {
          x.string = s
          goto X
        }
      }
    }
    errh.Error0 ("falsche Eingabe")
  }
X:
  x.Write (l, c)
  if x.wk > 0 {
    errh.DelHint()
  }
}

func (x *enumerator) String() string {
  return x.string
}

func (x *enumerator) TeX() string {
  t := x.String()
  t = str.UTF8 (t)
  str.OffSpc1 (&t)
  return t
}

func (x *enumerator) Defined (s string) bool {
  if str.Empty (s) {
    x.Clr()
    return true
  }
  s = str.Lat1 (s)
  str.Norm (&s, x.w)
  for i := uint(0); i < x.Num(); i++ {
    if s == x.s[i] {
      x.string = s
      return true
    }
  }
  return false
}

func (x *enumerator) Print (l, c uint) {
  pbx.Print (x.string, l, c)
}

func (x *enumerator) Num() uint {
  return x.uint
}

func (x *enumerator) Width() uint {
  return x.w
}
