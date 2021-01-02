package book

// (c) Christian Maurer   v. 201128 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/text"
  "µU/enum"
  "µU/bn"
  "µU/masks"
  "µU/atom"
  "µU/mol"
)
const (
  lenAuthor = 30
  lenTitle = 63
)
type
  order byte; const (
  subject = order(iota)
  author
  nOrders
)
var
  actOrd order
type
  book struct {
              mol.Molecule
              }

func (x *book) imp (Y Any) mol.Molecule {
  y, ok:= Y.(*book)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y.Molecule
}

func new_() Book {
  x := new (book)
  x.Molecule = mol.New (5)

  a := atom.New (enum.New (enum.BookC)) // Gebiet
  a.Colours (col.LightWhite(), col.Blue())
  a.SetFormat (enum.Long)
  x.Ins (a, 1, 1)

  a = atom.New (text.New (lenAuthor)) // Autor
  a.Colours (col.Yellow(), col.Red())
  x.Ins (a, 1, 16)

  a = atom.New (text.New (lenAuthor)) // Koautor
  a.Colours (col.LightWhite(), col.DarkRed())
  x.Ins (a, 1, 49)

  a = atom.New (bn.New (3)) // Nr
  a.Colours (col.LightWhite(), col.DarkCyan())
  x.Ins (a, 4, 10)

  a = atom.New (text.New (lenTitle)) // Titel
  a.Colours (col.LightWhite(), col.DarkGreen())
  x.Ins (a, 4, 16)
/*/
          1         2         3         4         5         6         7
01234567890123456789012345678901234567890123456789012345678901234567890123456789
 Gebiet         Autor/in                         Koautor/in
 ____________   ______________________________   ______________________________

          Nr.   Titel
          ___   _______________________________________________________________
/*/
  m := masks.New()
  m.Ins ("Gebiet",      0,  1)
  m.Ins ("Autor/in",    0, 16)
  m.Ins ("Koautor/in",  0, 49)
  m.Ins ("Nr.",         3, 10)
  m.Ins ("Titel",       3, 16)
  x.SetMask (m)
  return x
}

func (x *book) Eq (Y Any) bool {
  return x.Molecule.Eq (x.imp (Y))
}

func (x *book) Copy (Y Any) {
  x.Molecule.Copy (x.imp(Y))
}

func (x *book) Clone() Any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *book) Less (Y Any) bool {
  y := x.imp(Y)
  xs := x.Component(0).(atom.Atom)
  xa := x.Component(1).(atom.Atom)
  ys := y.Component(0).(atom.Atom)
  ya := y.Component(1).(atom.Atom)
  switch actOrd {
  case subject:
    if xs.Eq (ys) {
      return xa.Less (ya)
    }
    return xs.Less (ys)
  case author:
    if xa.Eq (ya) {
      return xs.Less (ys)
    }
    return xa.Less (ya)
  }
  return false
}

func (x *book) Index() Func {
  return func (a Any) Any {
    return a
  }
}

func (x *book) RotOrder() {
  actOrd = (actOrd + 1) % nOrders
}
