package mol

// (c) Christian Maurer   v. 210414 - license see µU.go

import (
  . "µU/obj"
  "µU/kbd"
  "µU/col"
  "µU/font"
  "µU/atom"
  "µU/masks"
)
type
  molecule struct {
                  uint8 "number of atoms"
             comp []atom.Atom
             l, c []uint
                  masks.MaskSet
                  }

func new_() Molecule {
  x := new(molecule)
  x.comp = make([]atom.Atom, 0)
  x.l, x.c = make([]uint, 0), make([]uint, 0)
  return x
}

func (x *molecule) imp (Y Any) *molecule {
  y, ok := Y.(*molecule)
  if ! ok { TypeNotEqPanic (x, Y) }
/*/
  if x.uint8 != y.uint8 {
    println ("mol.imp: x.uint8 ==", x.uint8,"!= y.uint8 ==", y.uint8);
    TypeNotEqPanic (x, Y)
  }
/*/
  for i := uint8(0); i < y.uint8; i++ {
    if ! x.comp[i].Equiv (y.comp[i]) { TypeNotEqPanic (x.comp[i], y.comp[i]) }
  }
  return y
}

func (x *molecule) Component (n uint) Any {
  if n >= uint(x.uint8) { return nil }
  return x.comp[n]
}

func (x *molecule) Ins (a atom.Atom, l, c uint) {
  x.comp = append (x.comp, a.Clone().(atom.Atom))
  x.l, x.c = append (x.l, l), append (x.c, c)
  x.uint8++
}

func (x *molecule) SetMasks (m masks.MaskSet) {
  x.MaskSet = m
}

func (x *molecule) Empty() bool {
  for i := uint8(0); i < x.uint8; i++ {
    if ! x.comp[i].Empty() {
      return false
    }
  }
  return true
}

func (x *molecule) Clr() {
  for i := uint8(0); i < x.uint8; i++ {
    x.comp[i].Clr()
  }
}

func (x *molecule) Eq (Y Any) bool {
  y := x.imp (Y)
  for i := uint8(0); i < x.uint8; i++ {
    if ! x.comp[i].Eq (y.comp[i]) {
      return false
    }
  }
  return true
}

func (x *molecule) Copy (Y Any) {
//  x.comp = make ([]atom.Atom, x.uint)
//  x.l, x.c = make ([]uint, x.uint), make ([]uint, x.uint)
  y := x.imp (Y)
  for i := uint8(0); i < y.uint8; i++ {
    x.comp[i].Copy (y.comp[i])
    x.l[i], x.c[i] = y.l[i], y.c[i]
  }
// x.MaskSet.Copy (y.MaskSet) // because of type loss in piset.New this does not work
  x.MaskSet = y.MaskSet
}

func (x *molecule) Clone() Any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *molecule) less (y *molecule, n uint) bool {
  if n >= uint(x.uint8) { return false }
  if x.comp[n].Less (y.comp[n]) { return true }
  if x.comp[n].Eq (y.comp[n]) { return x.less (y, n + 1) }
  return false
}

func (x *molecule) Less (Y Any) bool {
  n := uint(0) // TODO
  return x.less (x.imp (Y), n)
}

func (x *molecule) Colours (f, b col.Colour) {
  for i := uint8(0); i < x.uint8; i++ {
     x.comp[i].Colours (f, b)
  }
}

func (x *molecule) Write (l, c uint) {
  if x.MaskSet != nil {
    x.MaskSet.Write (l, c)
  }
  for i := uint8(0); i < x.uint8; i++ {
    if x.l[i] < 512 {
      x.comp[i].Write (l + x.l[i], c + x.c[i])
    }
  }
}

func (x *molecule) Edit (l, c uint) {
  x.Write (l, c)
  i := uint8(0)
  loop:
  for {
    x.comp[i].Edit (l + x.l[i], c + x.c[i])
    switch C, d := kbd.LastCommand(); C {
    case kbd.Esc:
      break loop
    case kbd.Enter:
      if d == 0 {
        if i + 1 < x.uint8 {
          i++
        } else {
          break loop
        }
      } else {
        break loop
      }
    case kbd.Down:
      if i + 1 < x.uint8 {
        i++
      } else {
        i = 0
      }
    case kbd.Up:
      if i > 0 {
        i--
      } else {
        i = x.uint8 - 1
      }
    case kbd.Pos1:
      i = 0
    case kbd.End:
      i = x.uint8 - 1
    case kbd.Tab:
      if d == 0 {
        if i + 1 < x.uint8 {
          i++
        }
      } else {
        if i > 0 {
          i--
        }
      }
    }
  }
}

func (x *molecule) SetFont (f font.Font) {
  for i := uint8(0); i < x.uint8; i++ {
    x.SetFont (f)
  }
}

func (x *molecule) Print (l, c uint) {
  x.MaskSet.Print (l, c)
  for i := uint8(0); i < x.uint8; i++ {
    x.comp[i].Print (x.l[i], x.c[i])
  }
}

func (x *molecule) Codelen() uint {
  c := uint(1)
  for k := uint8(0); k < x.uint8; k++ {
    c += x.comp[k].Codelen()
    c += 2
  }
  if x.MaskSet != nil {
    c += x.MaskSet.Codelen()
  }
  return c
}

func (x *molecule) Encode() Stream {
  s := make (Stream, x.Codelen())
  i, a := uint(0), uint(1)
  s[0] = x.uint8
  i += a
  for k := uint8(0); k < x.uint8; k++ {
    a = x.comp[k].Codelen()
    copy (s[i:i+a], x.comp[k].Encode())
    i += a
    copy (s[i:i+1], Encode(uint8(x.l[k])))
    i++
    copy (s[i:i+1], Encode(uint8(x.c[k])))
    i++
  }
  if x.MaskSet != nil {
    a = x.MaskSet.Codelen()
    copy (s[i:i+a], x.MaskSet.Encode())
  }
  return s
}

func (x *molecule) Decode (s Stream) {
  i, a := uint(0), uint(1)
  x.uint8 = s[0]
  i += a
  for k := uint8(0); k < x.uint8; k++ {
    a = x.comp[k].Codelen()
    x.comp[k].Decode (s[i:i+a])
    i += a
    x.l[k] = uint(Decode (uint8(0), s[i:i+1]).(uint8))
    i++
    x.c[k] = uint(Decode (uint8(0), s[i:i+1]).(uint8))
    i++
  }
  if x.MaskSet != nil {
    a = x.MaskSet.Codelen()
    x.MaskSet.Decode (s[i:i+a])
  }
}
