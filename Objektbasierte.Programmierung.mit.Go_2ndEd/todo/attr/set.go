package attr

// (c) Christian Maurer   v. 221021 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/scr"
)
type
  set struct {
           m []bool
             }

func newSet() AttrSet {
  s := new(set)
  s.m = make([]bool, nAttrs)
  return s
}

func (x *set) imp (Y any) *set {
  y, ok := Y.(*set)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *set) Empty() bool {
  for a := uint(0); a < nAttrs; a++ {
    if x.m[a] { return false }
  }
  return true
}

func (x *set) Clr() {
  for a := uint(1); a < nAttrs; a++ {
    x.m[a] = false
  }
}

func (x *set) Eq (Y any) bool {
  y := x.imp (Y)
  for a := uint(0); a < nAttrs; a++ {
    if x.m[a] != y.m[a] {
      return false
    }
  }
  return true
}

func (x *set) Less (Y any) bool {
  return false
}

func (x *set) Leq (Y any) bool {
  return false
}

func (x *set) Copy (Y any) {
  y := x.imp (Y)
  for a := uint(0); a < nAttrs; a++ {
    x.m[a] = y.m[a]
  }
}

func (x *set) Clone() any {
  y := newSet()
  y.Copy (x)
  return y
}

func (x *set) Ins (A Attribute) {
  a := A.(*attribute)
  x.m[a.Attr] = true
}

func (x *set) Codelen() uint {
  return uint(nAttrs)
}

func (x *set) Encode() Stream {
  s := make (Stream, x.Codelen())
  for a := uint(0); a < nAttrs; a++ {
    s[a] = 1
    if ! x.m[a] { s[a] = 0 }
  }
  return s
}

func (x *set) Decode (s Stream) {
  for a := uint(0); a < nAttrs; a++ {
    x.m[a] = s[a] == 1
  }
}

func (x *set) Write (l, c uint, w bool) {
  if x.Empty() {
    scr.Clr (l, c, uint(nAttrs), 1)
    return
  }
  t := ""
  for a := Attr(1); a < nAttrs; a++ {
    if x.m[a] {
      t += string(txt[a][0])
    }
  }
  setbx.Write (t, l, c)
  if w {
    scr.ColourB (col.Red())
    scr.Write (" ", l, c - 1)
  }
}
