package zelle

// (c) Christian Maurer   v. 230109 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/col"
  "µU/scr"
  "µU/N"
  . "bahn/kilo"
  . "bahn/richtung"
  . "bahn/farbe"
  . "bahn/konstanten"
)
const
  max = 15
type
  art byte; const (
  gleis = art(iota)
  knick
  weiche
  dkw
  prellbock
  nk
)
type
  zelle struct {
               art
               uint32 "Nummer"
               Kilometrierung
          lage,
      richtung,
      stellung Richtung
          z, s uint
   letzteFarbe col.Colour
               }
var (
  alle [NZeilen+1][NSpalten+1]*zelle
  ZJustierung = H2 / 4 - 1 // zur vertikalen Zentrierung der Nummern innerhalb einer Zelle
)

func init() {
  for z := 0; z < NZeilen; z++ {
    for s := 0; s < NSpalten; s++ {
      alle[z][s] = new_().(*zelle)
    }
  }
}

func (x *zelle) imp (Y any) *zelle {
  y, ok := Y.(*zelle)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func new_() Zelle {
  x := new (zelle)
  x.art = nk
  x.Kilometrierung = Mit
  x.lage = Gerade
  x.richtung = Gerade
  x.stellung = Gerade // Grundstellung
  x.letzteFarbe = Hintergrundfarbe
  return x
}

func (x *zelle) Empty() bool {
  return x.art == nk
}

func (x *zelle) Clr() {
  x.art = nk
}

func (x *zelle) Eq (Y any) bool {
  y := x.imp (Y)
  return x.z == y.z && x.s == y.s &&
         x.art == y.art &&
         x.Kilometrierung == y.Kilometrierung &&
         x.lage == y.lage && x.richtung == y.richtung &&
         x.stellung == y.stellung &&
         x.uint32 == y.uint32
}

func (x *zelle) Copy (Y any) {
  y := x.imp(Y)
  x.z, x.s = y.z, y.s
  x.art, x.Kilometrierung = y.art, y.Kilometrierung
  x.lage, x.richtung, x.stellung = y.lage, y.richtung, y.stellung
  x.uint32 = y.uint32
  x.letzteFarbe = y.letzteFarbe
  alle[x.z][x.s] = x
}

func (x *zelle) Clone() any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *zelle) Less (Y any) bool {
  return x.s < x.imp (Y).s
}

func (x *zelle) Leq (Y any) bool {
  return x.s <= x.imp (Y).s
}

func (x *zelle) definieren (a art, n uint, k Kilometrierung, l, r, st Richtung, z, s uint) {
  if ! alle[z][s].Empty() { ker.Panic2 ("schon alle:", z, ",", s) }
  x.art = a
  x.uint32 = uint32(n)
  x.Kilometrierung = k
  x.lage, x.richtung = l, r
  x.stellung = st
  x.letzteFarbe = Hintergrundfarbe
  x.z, x.s = z, s
  alle[z][s] = x
}

func (x *zelle) Gleis (n uint, l Richtung, z, s uint) {
  x.definieren (gleis, n, NK, l, ND, ND, z, s)
}

func (x *zelle) Knick (n uint, k Kilometrierung, a Richtung, z, s uint) {
  x.definieren (knick, n, k, ND, a, ND, z, s)
}

func (x *zelle) Prellbock (k Kilometrierung, z, s uint) {
  x.art = prellbock
  x.Kilometrierung = k
  x.z, x.s = z, s
}

func (x *zelle) Weiche (n uint, k Kilometrierung, l, r, st Richtung, z, s uint) {
  x.definieren (weiche, n, k, l, r, st, z, s)
}

func (x *zelle) DKW (n uint, l Richtung, r Richtung, z, s uint) {
  if l == Gerade { ker.Oops() }
  x.art = dkw
  x.uint32 = uint32(n)
  x.z, x.s = z, s
  x.lage = l
  x.stellung = Gerade
  x.definieren (dkw, n, NK, l, Gerade, r, z, s)
  alle[z][s] = x
}

func (x *zelle) Numerieren (n uint) {
  x.uint32 = uint32(n)
}

func (x *zelle) HatPosition (z, s uint) bool {
  return alle[z][s] == x
}

func (x *zelle) Pos() (uint, uint) {
  return x.z, x.s
}

func (x *zelle) Schräglage (k Kilometrierung) (Richtung, uint, uint) {
  var a Richtung
  if x.art == gleis && k == x.Kilometrierung {
    a = x.richtung
  } else {
    a = x.lage
  }
  return a, x.z, x.s
}

func (x *zelle) IstGleis() bool {
  return x.art == gleis
}

func (x *zelle) IstWeiche() (Kilometrierung, bool) {
  if x.art != weiche { return NK, false }
  if x.richtung == x.lage { ker.Oops() }
  return x.Kilometrierung, true
}

func (x *zelle) IstDKW() (Kilometrierung, bool) {
  if x.art != dkw { return NK, false }
  if x.richtung == x.lage { ker.Oops() }
  return x.Kilometrierung, true
}

func (x *zelle) String() string {
  switch x.art {
  case gleis:
    return "Gleis"
  case knick:
    return "Knick"
  case weiche:
    return "Weiche"
  case dkw:
    return "DKW"
  case prellbock:
    return "Prellbock"
  }
  return "Nix"
}

func (x *zelle) Kilo() Kilometrierung {
  return x.Kilometrierung
}

func (x *zelle) change (zelle, z, s uint) {
  x.z, x.s = z, s
  alle[z][s] = x
}

func (x *zelle) relativeLageDesAbzweigs() Richtung {
  if x.art != weiche { ker.Oops() }
  if x.lage == Gerade {
    return x.richtung
  }
  return Entgegen (x.lage)
}

func line (x, y, x1, y1 int) {
  scr.Line (x, y, x1, y1)
}

func (X *zelle) ganzAusgeben (l Richtung, c col.Colour) {
  scr.ColourF (c)
  x, y := int(X.s) * W1, Y0 + int(X.z) * H1
  x1 := x + W1
  switch l {
  case Links:
    line (x, y,      x1, y - H1)
  case Gerade:
    line (x, y - H2, x1, y - H2)
  case Rechts:
    line (x, y - H1, x1, y)
  }
}

func (X *zelle) halbAusgeben (k Kilometrierung, r Richtung, c col.Colour) {
  scr.ColourF (c)
  x, y := int(X.s) * W1, Y0 + int(X.z) * H1
  if k == Mit {
    switch r {
    case Links:
      line (x + W2, y - H2, x + W1, y - H1)
    case Gerade:
      line (x + W2, y - H2, x + W1, y - H2)
    case Rechts:
      line (x + W2, y - H2, x + W1, y)
    }
  } else {
    switch r {
    case Links:
      line (x, y,      x + W2, y - H2)
    case Gerade:
      line (x, y - H2, x + W2, y - H2)
    case Rechts:
      line (x, y - H1, x + W2, y - H2)
    }
  }
}

func (x *zelle) knickAusgeben (c col.Colour) {
  scr.ColourF (c)
  x.halbAusgeben (Gegenrichtung (x.Kilometrierung), Gerade, c)
  x.halbAusgeben (x.Kilometrierung, x.richtung, c)
}

func (X *zelle) prellbockAusgeben (c col.Colour) {
  x, y := int(X.s) * W1 + W2, Y0 + int(X.z) * H1 - H2
  dx := W2; if X.Kilometrierung == Gegen { dx = -W2 }
  scr.ColourF (c)
  scr.Line (x + dx, y - 5, x + dx, y + 5)
}

func (x *zelle) weicheAusgeben (c col.Colour) {
  k := x.Kilometrierung
  if x.lage == Gerade {
    if x.stellung == Gerade {
      x.halbAusgeben (k, x.richtung, Nichtfarbe)
      x.halbAusgeben (k, x.lage, c)
    } else {
      x.halbAusgeben (k, x.lage, Nichtfarbe)
      x.halbAusgeben (k, x.stellung, c)
    }
  } else {
    if x.stellung == Gerade {
      x.halbAusgeben (k, x.stellung, Nichtfarbe)
      x.halbAusgeben (k, x.lage, c)
    } else {
      x.halbAusgeben (k, x.lage, Nichtfarbe)
      x.halbAusgeben (k, Gerade, c)
    }
  }
  x.halbAusgeben (Gegenrichtung (k), x.lage, c)
}

func (X *zelle) writeDKW (c col.Colour) {
  x, y := int(X.s) * W1 + W2, Y0 + int(X.z) * H1
  if X.lage == Links {
    if X.stellung == Gerade {
      scr.ColourF (Nichtfarbe)
      line (x - W2, y, x + W2, y - H2)
      line (x - W2, y - H2,  x + W2, y - H1)
      X.ganzAusgeben (Gerade, c)
      X.ganzAusgeben (Links, c)
    } else {
      X.ganzAusgeben (Gerade, Nichtfarbe)
      X.ganzAusgeben (Links, Nichtfarbe)
      scr.ColourF (c)
      line (x - W2, y, x + W2, y - H2)
      line (x - W2, y - H2, x + W2, y - H1)
    }
  } else {
    if X.stellung == Gerade {
      scr.ColourF (Nichtfarbe)
      line (x - W2, y - H2, x + W2, y)
      line (x - W2, y - H1, x + W2, y - H2)
      X.ganzAusgeben (Gerade, c)
      X.ganzAusgeben (Rechts, c)
    } else {
      X.ganzAusgeben (Gerade, Nichtfarbe)
      X.ganzAusgeben (Rechts, Nichtfarbe)
      scr.ColourF (c)
      line (x - W2, y - H2, x + W2, y)
      line (x - W2, y - H1, x + W2, y - H2)
    }
  }
}

func (X *zelle) Nummer() uint {
  return uint(X.uint32)
}

func (X *zelle) nummerAusgeben (c col.Colour) {
  if X.uint32 == 0 { return }
  x, y := int(X.s) * W1, Y0 + int(X.z) * H1
  scr.Colours (c, Hintergrundfarbe)
  scr.WriteGr (N.String (uint(X.uint32)), x + (W1 - 2 * 8) / 2, y - H1 + H2 / 3)
}

func (x *zelle) Ausgeben (c col.Colour) {
  scr.ColourF (c)
  switch x.art {
  case gleis:
    x.ganzAusgeben (x.lage, c)
    x.nummerAusgeben (c)
  case knick:
    x.knickAusgeben (c) 
  case weiche:
    x.weicheAusgeben (c)
  case dkw:
    x.writeDKW (c)
  case prellbock:
    x.ganzAusgeben (x.lage, c)
    x.prellbockAusgeben (c)
  }
}

func (x *zelle) Stellen (s Richtung) {
  switch x.art {
  case weiche, dkw:
    x.stellung = s
  default:
    ker.Oops()
  }
  x.Ausgeben (Nichtfarbe)
//  x.Ausgeben (x.letzteFarbe); println ("bluse ausgegeben")
}

func (x *zelle) Stellung() Richtung {
  switch x.art {
  case gleis, knick:
    ker.Oops()
  case weiche:
    return x.stellung
  case dkw:
    return x.stellung
  }
  return Gerade
}

func (x *zelle) Codelen() uint {
  return 4 +
         5 +
         4 +
         x.letzteFarbe.Codelen()
}

func (x *zelle) Encode() Stream {
  b := make (Stream, x.Codelen())
  i, a := uint32(0), uint32(4)
  n := x.z << 8 + x.s
  copy (b[i:a], Encode (n))
  i += a
  b[i] = byte(x.art)
  i++
  b[i] = byte(x.Kilometrierung)
  i++
  b[i] = byte(x.lage)
  i++
  b[i] = byte(x.richtung)
  i++
  b[i] = byte(x.stellung)
  i++
  copy (b[i:i+a], Encode (x.uint32))
  i += a
  a = uint32(x.letzteFarbe.Codelen())
  copy (b[i:i+a], x.letzteFarbe.Encode())
  return b
}

func (x *zelle) Decode (b Stream) {
  i, a := uint(0), uint(4)
  n := uint(Decode (uint32(0), b[i:a]).(uint32))
  x.z, x.s = n >> 8, n % 256
  i += a
  x.art = art(b[i])
  i++
  x.Kilometrierung = Kilometrierung(b[i])
  i++
  x.lage = Richtung(b[i])
  i++
  x.richtung = Richtung(b[i])
  i++
  x.stellung = Richtung(b[i])
  i++
  x.uint32 = Decode (x.uint32, b[i:i+a]).(uint32)
  i += a
  a = x.letzteFarbe.Codelen()
  x.letzteFarbe.Decode (b[i:i+a])
}

func (X *zelle) UnterMaus() bool {
  x, y := scr.MousePosGr()
  if int(X.s) * W1 <= x && x <= int(X.s + 1) * W1 {
    if int(X.z + 1) * H1 <= y && y <= int(X.z + 2) * H1 {
      return true
    }
  }
  return false
}
