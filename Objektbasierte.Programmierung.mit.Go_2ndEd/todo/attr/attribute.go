package attr

// (c) Christian Maurer   v. 221021 - license see µU.go

import (
  . "µU/obj"
  "µU/env"
  "µU/str"
  "µU/font"
  "µU/col"
  "µU/box"
  "µU/pbox"
  "µU/errh"
  "µU/pseq"
  "µU/files"
)
type
  attribute struct {
                   Attr
                   bool "marked"
                   }
var (
  attrs = pseq.New (str.New(Wd+1))
  nAttrs uint
  txt = make([]string, 1)
  Hilfetext string
  bx, setbx = box.New(), box.New()
  pbx = pbox.New()
  cF, cB, cFm, cBm = col.LightWhite(), col.Green(), col.LightWhite(), col.Green()
)

func init() {
  files.Cd (env.Gosrc() + "/todo/")
  attrs.Name ("Terminattribute.kfg")
  nAttrs = attrs.Num() + 1
  txt[0] = str.New (Wd)
  for i := uint(1); i < nAttrs; i++ {
    attrs.Seek (i-1)
    s := attrs.Get().(string)
    str.Norm (&s, Wd)
    txt = append (txt, s)
  }
  for a := Attr(1); a < nAttrs; a++ { Hilfetext += txt[a] + " " }
  bx.Wd (Wd)
  setbx.Wd (uint(nAttrs))
  setbx.Colours (cF, cB)
}

func (x *attribute) imp (Y any) *attribute {
  y, ok := Y.(*attribute)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func New() Attribute {
  x := new (attribute)
  x.Attr = 0
  return x
}

func (x *attribute) Empty() bool {
  return x.Attr == 0
}

func (x *attribute) Clr() {
  x.Attr = 0
}

func (x *attribute) Copy (Y any) {
  y := x.imp (Y)
  x.Attr = y.Attr
  x.bool = y.bool
}

func (x *attribute) Clone() any {
  y := New().(*attribute)
  y.Attr = x.Attr
  return y
}

func (x *attribute) Eq (Y any) bool {
  return x.Attr == x.imp (Y).Attr
}

func (x *attribute) Less (Y any) bool {
  return x.Attr < x.imp (Y).Attr
}

func (x *attribute) Leq (Y any) bool {
  return x.Attr <= x.imp (Y).Attr
}

func (x *attribute) Mark (m bool) {
  x.bool = m
}

func (x *attribute) Marked() bool {
  return x.bool
}

func (x *attribute) Colours (f, b col.Colour) {
  bx.Colours (f, b)
}

func (x *attribute) Write (l, c uint) {
  if x.bool {
    bx.Colours (cFm, cBm)
  } else {
    bx.Colours (cF, cB)
  }
  s := txt[x.Attr]
  bx.Write (s, l, c)
}

func (x *attribute) SetFont (f font.Font) {
// dummy
}

func (x *attribute) Print (l, c uint) {
  pbx.Print (txt[x.Attr], l, c)
}

func (x *attribute) Edit (l, c uint) {
  T := txt[x.Attr]
  if T == "" {
    errh.Error0 ("aha")
    return
  }
  for {
    bx.Edit (&T, l, c)
    for a := uint(0); a < nAttrs; a++ {
      if T[0] == txt[a][0] {
        x.Attr = a
        x.Write (l, c)
        return
      }
    }
    errh.Error0 (Hilfetext)
  }
}

func (x *attribute) Codelen() uint {
  return 1
}

func (x *attribute) Encode() Stream {
  s := make (Stream, 1)
  s[0] = byte(x.Attr)
  return s
}

func (x *attribute) Decode (s Stream) {
  if s[0] < byte(nAttrs) {
    x.Attr = Attr(s[0])
  } else {
    x.Attr = 0
  }
}

func (x *attribute) String() string {
  return txt[x.Attr]
}

func (x *attribute) Defined (s string) bool {
  for a := uint(0); a < nAttrs; a++ {
    if s == txt[a] {
      return true
    }
  }
  return false
}
