package enum

// (c) Christian Maurer   v. 201010 - license see µU.go

import (
  . "µU/obj"
  "µU/ker"
  "µU/col"
  "µU/enum/internal"
)
type
  enumerator struct {
                    internal.Base
                    }
var
  l, s [NTypes][]string

func init() {
  N = make([]uint8, NTypes)
  N[Title] = uint8(len(l[Title]))
  N[AudioC] = uint8(len(l[AudioC]))
  N[BookC] = uint8(len(l[BookC]))
  N[AudioMedium] = uint8(len(l[AudioMedium]))
  N[Religion] = uint8(len(l[Religion]))
  N[Subject] = uint8(len(l[Subject]))
  N[Wortart] = NWortarten
  N[Casus] = NCasus
  N[Genus] = NGenera
  N[Persona] = NPersonae
  N[Numerus] = NNumeri
  N[Tempus] = NTempora
  N[Modus] = NModi
  N[GenusVerbi] = NGeneraVerbi
  N[Comparatio] = NComparationes
}

func new_(t uint8) Enumerator {
  if t >= NTypes { ker.Panic ("enum.New: Parameter >= NEnums") }
  return &enumerator { internal.New (t, [internal.NFormats][]string { s[t], l[t] }) }
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

func (x *enumerator) Found (ss string, t Type, f Format) uint8 {
  var s0 []string
  if f == Short {
    s0 = s[t]
  } else {
    s0 = l[t]
  }
  for i := uint8(0); i < N[t]; i++ {
    if ss == s0[i] {
      return i
    }
  }
  return uint8(0)
}
