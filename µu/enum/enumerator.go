package enum

// (c) Christian Maurer   v. 170918 - license see µu.go

import (
  . "µu/obj"
  "µu/ker"
  "µu/col"
  "µu/enum/internal"
)
type
  enumerator struct {
                    internal.Base
                    }
var
  l, s [NEnums][]string

func new_(e uint8) Enumerator {
  if e >= NEnums { ker.Panic ("enum.New: Parameter >= NEnums") }
  return &enumerator { internal.New (e, [internal.NFormats][]string { s[e], l[e] }) }
}

func (x *enumerator) imp(Y Any) *enumerator {
  y, ok := Y.(*enumerator)
  if ! ok { TypeNotEqPanic (x, Y) }
  CheckTypeEq (x.Typ(), y.Typ())
  return y
}

func (x *enumerator) Colours (f, b col.Colour) {
  x.Base.Colours (f, b)
}

func (x *enumerator) Eq (Y Any) bool {
  return x.Base.Eq (x.imp(Y).Base)
}

func (x *enumerator) Copy (Y Any) {
  x.Base.Copy (x.imp(Y).Base)
}

func (x *enumerator) Clone() Any {
  y := new_(x.Base.Typ()).(*enumerator)
  x.Base.Copy (y.Base)
  return y
}

func (x *enumerator) Less (Y Any) bool {
  return x.Base.Less (x.imp(Y).Base)
}

func (x *enumerator) Set (e uint8) bool {
  return x.Base.Set (uint8(e))
}
