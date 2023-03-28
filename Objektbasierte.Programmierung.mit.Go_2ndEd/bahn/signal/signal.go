package signal

// (c) Christian Maurer   v. 230107 - license see µU.go

import (
  . "µU/obj"
//  "µU/col"
  "µU/scr"
  . "bahn/kilo"
  . "bahn/konstanten"
  . "bahn/farbe"
)
type
  signal struct {
                uint // Nummer des Gleises
                Typ
                Kilometrierung
                Stellung
  zeile, spalte uint
                }
var (
  halbeHoehe  = H1 / 2
  halbeBreite = W1 / 2
)

func new_() Signal {
  x := new (signal)
  x.Typ = NT
  x.Kilometrierung = NK
  x.Stellung = NS
  x.zeile, x.spalte = NZeilen, NSpalten
  x.zeile, x.spalte = 9, 3
  return x
}

func (x *signal) imp (Y any) *signal {
  y, ok := Y.(*signal)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *signal) Empty() bool {
  return x.uint == 0 ||
         x.Typ == NT ||
         x.Kilometrierung == NK ||
         x.Stellung == NS
}

func (x *signal) Clr() {
  x.Typ = NT
}

func (x *signal) Eq (Y any) bool {
  y := x.imp (Y)
  return x.Typ == y.Typ &&
         x.Kilometrierung == y.Kilometrierung &&
//         x.Stellung == y.Stellung &&
         x.zeile == y.zeile && x.spalte == y.spalte
}

func (x *signal) Less (Y any) bool {
  return false
}

func (x *signal) Leq (Y any) bool {
  return false
}

func (x *signal) Copy (Y any) {
  y := x.imp (Y)
  x.uint = y.uint
  x.Typ = y.Typ
  x.Kilometrierung = y.Kilometrierung
  x.Stellung = y.Stellung
  x.zeile, x.spalte = y.zeile, y.spalte
}

func (x *signal) Clone() any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *signal) Codelen() uint {
  return 4 +
         3 * 1
}

func (x *signal) Encode() Stream {
  b := make (Stream, x.Codelen())
  i, a := uint32(0), uint32(4)
  n := uint32(x.zeile << 16 + x.spalte)
  copy (b[i:i+a], Encode (n))
  i += a
  b[i] = byte(x.Typ)
  i++
  b[i] = byte(x.Kilometrierung)
  i++
  b[i] = byte(x.Stellung)
  return b
}

func (x *signal) Decode (b Stream) {
  i, a := uint32(0), uint32(4)
  n := uint(Decode (uint32(0), b[i:i+a]).(uint32))
  x.zeile, x.spalte = n >> 16, n % (1 << 16)
  i += a
  x.Typ = Typ(b[i])
  i++
  x.Kilometrierung = Kilometrierung(b[i])
  i++
  x.Stellung = Stellung(b[i])
}

func (x *signal) Definieren (n uint, t Typ, k Kilometrierung, st Stellung, z, s uint) {
  x.uint = n
  x.Typ = t
  x.Kilometrierung = k
  x.Stellung = st
  x.zeile, x.spalte = z, s
}

func (x *signal) Signaltyp() Typ {
  return x.Typ
}

func (x *signal) Stellen (s Stellung) {
  x.Stellung = s
  switch x.Typ {
  case T0:
    return
  case T1:
    if s == Hp2 { s = Hp1 }
    x.Stellung = s
  case T2:
    x.Stellung = s
  case NT:
    x.Stellung = NS
  }
  x.Ausgeben()
}

func (X *signal) Ausgeben() {
  if X.Empty() { return }
  x, y := int(X.spalte) * W1, Y0 + int(X.zeile) * H1 - H2 / 2
  f, f1 := Haltfarbe, Hintergrundfarbe
  switch X.Stellung {
  case Hp0:
    ;
  case Hp1:
    f = Fahrtfarbe
  case Hp2:
    f, f1 = Fahrtfarbe, Langsamfahrtfarbe
  case NS:
    f = Hintergrundfarbe
  }
  const r = 3
  x, y  = int(X.spalte) * W1, Y0 + int(X.zeile) * H1 - H2 / 2
  if X.Kilometrierung == Gegen {
    x += W1 / 3
    y -= H1 / 2
  }
  scr.ColourF (f)
  scr.CircleFull (x, y, r)
  if X.Typ == T2 {
    scr.ColourF (f1)
    if X.Kilometrierung == Mit {
      x -= 2 * r
    } else {
      x += 2 * r
    }
    scr.CircleFull (x, y, r)
  }
}
