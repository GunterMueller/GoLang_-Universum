package mol

// (c) murus.org  v. 140803 - license see murus.go

import (
  . "murus/obj"
  "murus/kbd"; "murus/col"; "murus/font"
  "murus/atom"; "murus/masks"
)
type
  molecule struct {
                  uint "number of atoms"
             comp []atom.Atom
             l, c []uint
                  masks.MaskSequence
                  }

func newMol() Molecule {
  return new(molecule)
}

func (x *molecule) imp (Y Any) *molecule {
  y, ok := Y.(*molecule)
  if ! ok { TypeNotEqPanic (x, Y) }
  if x.uint != y.uint {
    println ("mol.New: x.uint ==", x.uint,"!= y.uint ==", y.uint); TypeNotEqPanic (x, Y) }
  for i := uint(0); i < y.uint; i++ {
    if ! x.comp[i].Equiv (y.comp[i]) { TypeNotEqPanic (x.comp[i], y.comp[i]) }
  }
  return y
}

func (x *molecule) Num() uint {
  return x.uint
}

func (x *molecule) Component (n uint) Any {
  if n >= x.uint { WrongUintParameterPanic ("Component", x, n) }
  return x.comp[n]
}

func (x *molecule) Ins (a atom.Atom, l, c uint) {
  x.comp = append (x.comp, a.Clone().(atom.Atom))
  x.l, x.c = append (x.l, l), append (x.c, c)
  x.uint ++
}

func (x *molecule) Del (n uint) {
  if n >= x.uint { return }
  for i := uint(n); i + 1 < x.uint; i++ {
    x.comp[i] = x.comp[i + 1]
    x.l[i], x.c[i] = x.l[i + 1], x.c[i + 1]
  }
  x.uint --
  x.comp[x.uint] = nil
}

func (x *molecule) SetMask (m masks.MaskSequence) {
  x.MaskSequence = m
}

func (x *molecule) Empty() bool {
  for i := uint(0); i < x.uint; i++ {
    if ! x.comp[i].Empty() {
      return false
    }
  }
  return true
}

func (x *molecule) Clr() {
  for i := uint(0); i < x.uint; i++ {
    x.comp[i].Clr()
  }
}

func (x *molecule) Eq (Y Any) bool {
  y := x.imp (Y)
  for i := uint(0); i < x.uint; i++ {
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
  for i := uint(0); i < y.uint; i++ {
    x.comp[i].Copy (y.comp[i])
    x.l[i], x.c[i] = y.l[i], y.c[i]
  }
// x.MaskSequence.Copy (y.MaskSequence) // wegen Typverlust in piset.New geht das nicht TODO why not ?
  x.MaskSequence = y.MaskSequence
}

func (x *molecule) Clone() Any {
  y := newMol()
  y.Copy (x)
  return y
}

func (x *molecule) less (y *molecule, n uint) bool {
  if n > x.uint { return false }
  if x.comp[n].Less (y.comp[n]) { return true }
  if x.comp[n].Eq (y.comp[n]) { return x.less (y, n + 1) }
  return false
}

func (x *molecule) Less (Y Any) bool {
  return x.less (x.imp (Y), 0)
}

func (x *molecule) Colours (f, b col.Colour) {
  for i := uint(0); i + 1 < x.uint; i++ {
     x.comp[i].Colours (f, b)
  }
}

func (x *molecule) Write (l, c uint) {
  if x.MaskSequence != nil {
    x.MaskSequence.Write (l, c)
  }
  for i := uint(0); i < x.uint; i++ {
    if x.l[i] < 512 {
      x.comp[i].Write (l + x.l[i], c + x.c[i])
    }
  }
}

func (x *molecule) Edit (l, c uint) {
  x.Write (l, c)
  i := uint(0)
  loop: for {
    x.comp[i].Edit (l + x.l[i], c + x.c[i])
    switch C, d := kbd.LastCommand(); C {
    case kbd.Esc:
      break loop
    case kbd.Enter:
      if d == 0 {
        if i + 1 < x.uint {
          i ++
        } else {
          break loop
        }
      } else {
        break loop
      }
    case kbd.Down:
      if i + 1 < x.uint {
        i ++
      } else {
        i = 0
      }
    case kbd.Up:
      if i > 0 {
        i --
      } else {
        i = x.uint - 1
      }
    case kbd.Pos1:
      i = 0
    case kbd.End:
      i = x.uint - 1
    }
  }
}

func (x *molecule) SetFont (f font.Font) {
  for i := uint(0); i < x.uint; i++ {
    x.SetFont (f)
  }
}

func (x *molecule) Print (l, c uint) {
//  x.MaskSequences.Print (l, c)
  for i := uint(0); i < x.uint; i++ {
    x.comp[i].Print (x.l[i], x.c[i])
  }
}

func (x *molecule) Codelen() uint {
  c := uint(4)
  for k := uint(0); k < x.uint; k++ {
    c += x.comp[k].Codelen()
  }
  return c
}

func (x *molecule) Encode() []byte {
  b := make ([]byte, x.Codelen())
  i, a := uint(0), uint(4)
  copy (b[i:i+a], Encode (uint32(x.uint)))
  i += a
  for k := uint(0); k < x.uint; k++ {
    a = x.comp[k].Codelen()
    copy (b[i:i+a], x.comp[k].Encode())
    i += a
  }
  return b
}

func (x *molecule) Decode (b []byte) {
  i, a := uint(0), uint(4)
  x.uint = uint(Decode (uint32(0), b[i:i+a]).(uint32))
  i += a
  for k := uint(0); k < x.uint; k++ {
    a = x.comp[k].Codelen()
    x.comp[k].Decode (b[i:i+a])
    i += a
  }
}
