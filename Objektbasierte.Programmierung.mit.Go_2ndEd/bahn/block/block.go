package block

// (c) Christian Maurer   v. 230306 - license see µU.go

import (
  . "µU/obj"
  "µU/ker"
  "µU/time"
  "µU/col"
  "µU/scr"
  "µU/str"
  "µU/N"
  "µU/seq"
  . "bahn/farbe"
  . "bahn/kilo"
  . "bahn/richtung"
  . "bahn/konstanten"
  s "bahn/signal"
  "bahn/zelle"
)
type
  zustand byte; const (
  frei = zustand(iota)
  besetzt
  befahren
)
type
  block struct {
               uint32 "Nummer"
               Art
               Kilometrierung // Verzweigungsrichtung, falls Weiche
          lage,
      richtung,               // Abzweigungsrichtung, falls Weiche
      stellung Richtung       // Stellung, falls Weiche
               uint           // Länge = Anzahl der Zellen
          l, c uint           // Position am linken Rand
               seq.Sequence   // Folge der Zellen
               zustand
           sig [NK]s.Signal
               }
const (
  m0 = uint32(H)
  dg = uint32(0)
  dk = uint32(1)
  dw = uint32(2)
)
var (
  text = [NArten+1]string {"G",
                           "G", "G",
                           "G", "G",
                           "G", "G",
                           "G", "G",
                           "K",
                           "W",
                           "D",
                           "N"}
  maxWeichennummer = uint(0)
)

func init() {
  N.Colours (col.Yellow(), col.Black())
  for i := 0; i < M; i++ {
    B[i] = new_()
    W[i] = new_()
    D[i] = new_()
  }
}

func new_() Block {
  x := new (block)
  x.uint32 = 0
  x.Art = NArten
  x.lage = Gerade
  x.Sequence = seq.New (zelle.New())
  x.Sequence.Sort()
  x.sig[Mit], x.sig[Gegen] = s.New(), s.New()
  x.zustand = frei
  return x
}

func (x *block) imp (Y any) *block {
  y, ok := Y.(*block)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *block) Empty() bool {
  return x.uint32 == 0
}

func (x *block) Clr() {
  x.uint32 = 0
  x.Art = NArten
  x.lage = ND
  x.richtung = ND
  x.stellung = ND
  x.Kilometrierung = NK
  x.uint = 0
  x.l, x.c = 0, 0
  x.Sequence.Clr()
  x.zustand = frei
  x.sig[Mit].Clr()
  x.sig[Gegen].Clr()
}

func (x *block) Eq (Y any) bool {
  y := x.imp (Y)
  if x.Empty() { return y.Empty() }
  return x.uint32 == y.uint32 &&
            x.Art == y.Art &&
 x.Kilometrierung == y.Kilometrierung &&
           x.lage == y.lage &&
       x.richtung == y.richtung &&
       x.stellung == y.stellung &&
           x.uint == y.uint &&
              x.l == y.l &&
              x.c == y.c &&
       x.Sequence.Eq (y.Sequence) &&
        x.zustand == y.zustand &&
       x.sig[Mit].Eq (y.sig[Mit]) &&
     x.sig[Gegen].Eq (y.sig[Gegen])
}

func (x *block) Less (Y any) bool {
  y := x.imp (Y)
  if x.Empty() || y.Empty () { ker.Oops() }
  return x.c + x.uint - 1 < y.c
}

func (x *block) Leq (Y any) bool {
  y := x.imp (Y)
  if x.Empty() || y.Empty () { ker.Oops() }
  return x.c + x.uint - 1 <= y.c
}

func (x *block) Copy (Y any) {
  y := x.imp (Y)
  x.uint32 = y.uint32
  x.Art = y.Art
  x.lage = y.lage
  x.richtung = y.richtung
  x.stellung = y.stellung
  x.Kilometrierung = y.Kilometrierung
  x.uint = y.uint
  x.l, x.c = y.l, y.c
  x.Sequence.Copy (y.Sequence)
  x.zustand = y.zustand
  x.sig[Mit].Copy (y.sig[Mit])
  x.sig[Gegen].Copy (y.sig[Gegen])
}

func (x *block) Clone() any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *block) Codelen() uint {
  k := x.Sequence.Codelen() // maximal NSpalten * zelle.New().Codelen()
  return 4 + // x.uint32
         1 + // x.art
         1 + // x.Kilometrierung
         1 + // x.lage
         1 + // x.richtung
         1 + // x.stellung
         1 + // x.uint
         1 + // x.l
         1 + // x.c
         4 + // x.Sequence.Codelen()
         k + // Sequence.Codelen()
         1 + // x.zustand
         2 * x.sig[Mit].Codelen()
}

func (x *block) Encode() Stream {
  s := make (Stream, x.Codelen())
  i, a := uint(0), uint(4)
  copy (s[i:i+a], Encode(x.uint32))
  i += a
  s[i] = byte(x.Art)
  i++
  s[i] = byte(x.Kilometrierung)
  i++
  s[i] = byte(x.lage)
  i++
  s[i] = byte(x.richtung)
  i++
  s[i] = byte(x.stellung)
  i++
  s[i] = byte(x.uint)
  i++
  s[i] = byte(x.l)
  i++
  s[i] = byte(x.c)
  i++
  k := x.Sequence.Codelen()
  copy (s[i:i+a], Encode(uint32(k)))
  i += a
  a = k
  copy (s[i:i+a], x.Sequence.Encode())
  i += a
  s[i] = byte(x.zustand)
  i++
  a = x.sig[Mit].Codelen()
  copy (s[i:i+a], x.sig[Mit].Encode())
  i += a
  copy (s[i:i+a], x.sig[Gegen].Encode())
  return s
}

func (x *block) Decode (s Stream) {
  i, a := uint(0), uint(4)
  x.uint32 = Decode(uint32(0), s[i:i+a]).(uint32)
  i += a
  x.Art = Art(s[i])
  i++
  x.Kilometrierung = Kilometrierung(s[i])
  i++
  x.lage = Richtung(s[i])
  i++
  x.richtung = Richtung(s[i])
  i++
  x.stellung = Richtung(s[i])
  i++
  x.uint = uint(s[i])
  i++
  x.l = uint(s[i])
  i++
  x.c = uint(s[i])
  i++
  k := uint(Decode(uint32(0), s[i:i+a]).(uint32))
  i += a
  a = k
  x.Sequence.Decode (s[i:i+a])
  i += a
  x.zustand = zustand(s[i])
  i++
  a = x.sig[Mit].Codelen()
  x.sig[Mit].Decode (s[i:i+a])
  i += a
  x.sig[Gegen].Decode (s[i:i+a])
}

func (x *block) String() string {
  if x.Art == NArten {
    return ""
  }
  return text[x.Art] + N.String (uint(x.uint32 % m0))
}

func (x *block) Defined (s string) bool { // wird nicht benutzt
  str.OffSpc (&s)
  if len(s) > 5 { return false }
  var d uint32
  switch s[0] {
  case 'G':
    x.Art = Dfg
//          AsM
//          AsG
//          EfM
//          EfG
//          AfM
//          AfG
//          EAM
//          EAG
    d = dg
  case 'K':
    x.Art = Knick 
    d = dk
  case 'W':
    x.Art = Weiche
    d = dw
  case 'D':
    x.Art = DKW
    d = dw
  default:
    return false
  }
  if k, ok := N.Natural (s[1:]); ok {
    x.uint32 = m0 * d + uint32(k)
    return true
  } 
  return false
}

func (x *block) Schräglage() Richtung {
  return x.lage
}

func (x *block) Numerieren (n uint) {
  if x.uint32 > 0 {
    x.uint32 = uint32(n)
  }
}

func (x *block) Art_() Art {
  return x.Art
}

func (x *block) Nummer() uint {
  return uint(x.uint32)
}

func (x *block) Nummerkurz() uint {
  return uint(x.uint32 % m0)
}

func (X *block) check (t string, n uint) {
  s := X.Encode()
  Y := new_()
  Y.Decode (s)
  if ! Y.Eq (X) { ker.Panic1 ("Code-Fehler bei " + t, n) }
}

func (X *block) GleisDefinieren (n uint, a Art, lage Richtung, Länge, l, c, sn uint,
                    gt s.Typ, g Kilometrierung, gst s.Stellung, gz, gsn uint,
                    mt s.Typ, m Kilometrierung, mst s.Stellung, mz, msn uint) {
  if n == 0 { ker.Panic ("Gleis: n == 0") }
  X.uint32 = uint32(n)
  Nr = append (Nr, uint(X.uint32))
  X.Art = a
  X.Kilometrierung = NK // spielt keine Rolle
  X.lage = lage
  X.richtung, X.stellung = ND, ND // spielen keine Rolle
  X.uint = Länge
  X.l, X.c = l, c
  l0, c0 := l, c
  for u := X.uint; u > 0; u-- {
    z := zelle.New()
    z.Gleis (n, X.lage, l, c)
    if c == sn {
      z.Numerieren (uint(X.uint32))
    } else {
      z.Numerieren (0)
//     z.Numerieren (uint(X.uint32)) // zum Testen
    }
    X.Sequence.Ins (z)
    c++
    switch lage {
    case Links:
      l--
    case Rechts:
      l++
    } 
  }
  if X.Art == AsM {
    l, c = X.l, X.c + X.uint - 1
    z := zelle.New()
    z.Prellbock (Mit, l, c)
    X.Sequence.Ins (z)
  }
  if X.Art == AsG {
    l, c = X.l, X.c
    z := zelle.New()
    z.Prellbock (Gegen, l, c)
    X.Sequence.Ins (z)
  }
  if X.Sequence.Num() >= NSpalten { ker.Panic ("das Gleis ist zu lang") }
  B[n] = X
  X.sig[Gegen].Definieren (n, gt, g, gst, gz, gsn)
  X.sig[Mit].Definieren (n, mt, m, mst, mz, msn)
  l, c = l0, c0
  x, y := int(c) * W1, Y0 + int(l) * H1 - H2
  x0, ln := x, int(X.Sequence.Num())
  do := true
  for k := Mit; k < NK; k++ {
    if k == Mit {
      switch X.Art {
      case Dfg:
        x = x0 + W1 * ln
        do = true
      case AsM:
        do = false
      case AsG:
        x = x0 + W1 * (ln - 1)
        do = true
      case EfM:
        x = x0 + W1 * ln
        do = true
      case EfG:
        do = false
      case AfM:
        do = false
      case AfG:
        x = x0 + W1 * ln
      case EAM:
        do = false
      case EAG:
        x = x0 + W1 * ln
        do = true
      }
    } else { // k == Gegen
      switch X.Art {
      case Dfg:
        x = x0
      case AsM:
        x = x0
        do = true
      case AsG:
        do = false
      case EfM:
        do = false
      case EfG:
        x = x0
        do = true
      case AfM:
        x = x0
        do = true
      case AfG:
        do = false
      case EAM:
        x = x0
        do = true
      case EAG:
        do = false
      }
    }
    switch X.lage {
    case Links:
      if k == Mit {
        y -= H2
      } else {
        y += H1
      }
    case Rechts:
      if k == Mit {
        y += H2
      } else {
        y -= H1
      }
    }
    if do {
      X.ins (n, k, x, y)
    }
  }
  X.check ("Gleis", n)
}

func (x *block) IstGleis() bool {
  switch x.Art {
  case Knick, Weiche, DKW:
    return false
  }
  return true
}

func (x *block) IstDurchfahrgleis() bool {
  return x.Art == Dfg
}

func (x *block) IstEinfahrgleis() bool {
  switch x.Art {
  case EfM, EfG:
    return true
  }
  return false
}

func (x *block) IstAusfahrgleis() bool {
  switch x.Art {
  case AfM, AfG:
    return true
  }
  return false
}

func (x *block) IstEinAusfahrgleis() bool {
  switch x.Art {
  case EAM, EAG:
    return true
  }
  return false
}

func (x *block) IstAbstellgleis (k Kilometrierung) bool {
  if x.Art == Weiche || x.Art == DKW { ker.Oops() }
  if k == Mit {
    return x.Art == AsM
  }
  return x.Art == AsG
}

func (X *block) KnickDefinieren (n uint, k Kilometrierung, r Richtung, z, s uint) {
  if n == 0 { ker.Panic ("Knick: n == 0") }
  X.uint32 = m0 * dk + uint32(n)
  Nr = append (Nr, uint(X.uint32))
  X.Art = Knick
  X.Kilometrierung = k
  X.lage = ND
  X.richtung, X.stellung = r, ND
  X.uint = 1
  X.l, X.c = z, s
  c := zelle.New()
  c.Knick (n, k, r, z, s)
  X.Sequence.Ins (c)
  B[X.uint32] = X
  x, y := int(s) * W1, Y0 + int(z) * H1 - H2
  x0, y0 := x, y
  for k := Mit; k < NK; k++ {
    if k == Mit {
      if X.Kilometrierung == Mit {
        x = x0 + W1
        if X.richtung == Links {
          y = y0 - H2
        } else {
          y = y0 + H2
        }
      } else {
        x, y = x0 + W1, y0
      }
    } else {
      if X.Kilometrierung == Mit {
        if X.richtung == Links {
          x, y = x0, y0
        } else {
          x, y = x0, y0
        }
      } else {
        x = x0
        if X.richtung == Links {
          y = y0 + H2
        } else {
          y = y0 - H2
        }
      }
    }
    X.ins (uint(X.uint32), k, x, y)
  }
  X.check ("Knick", n)
}

func (X *block) IstKnick() bool {
  return X.Art == Knick
}

func (X *block) WeicheDefinieren (n uint, k Kilometrierung, l, r, st Richtung, z, s uint) {
  if n == 0 { ker.Panic ("Weiche: n == 0") }
  X.uint32 = m0 * dw + uint32(n)
  if maxWeichennummer < uint(X.uint32) {
    maxWeichennummer = uint(X.uint32)
  }
  Nr = append (Nr, uint(X.uint32))
  X.Art = Weiche
  X.Kilometrierung = k
  X.lage = l
  X.richtung, X.stellung = r, st
  X.uint = 1
  X.l, X.c = z, s
  W[n] = X
  c := zelle.New()
  c.Weiche (n, k, l, r, st, z, s)
  X.Sequence.Ins (c)
  B[X.uint32] = X
  x, y := int(s) * W1, Y0 + int(z) * H1 - H2
  x0, y0 := x, y
  for k := Mit; k < NK; k++ {
    if k == Mit {
      x = x0 + W1
      switch X.lage {
      case Links:
        y = y0 - H2
      case Rechts:
        y = y0 + H2
      }
    } else {
      x = x0
      switch X.lage {
      case Links:
        y = y0 + H2
      case Rechts:
        y = y0 - H2
      }
    }
    X.ins (uint(X.uint32), k, x, y)
  }
  if X.Kilometrierung == Mit {
    x = x0 + W1
    switch X.lage {
    case Links, Rechts:
      y = y0
    case Gerade:
      switch X.richtung {
      case Links:
        y = y0 - H2
      case Rechts:
        y = y0 + H2
      }
    }
  } else {
    x = x0
    switch X.lage {
    case Links, Rechts:
      y = y0
    case Gerade:
      switch X.richtung {
      case Links:
        y = y0 + H2
      case Rechts:
        y = y0 - H2
      }
    }
  }
  X.ins (uint(X.uint32), k, x, y)
  X.check ("Weiche", n)
}

func (x *block) IstWeiche() bool {
  return x.Art == Weiche
}

func (X *block) DKWDefinieren (n uint, l, r Richtung, z, s uint) {
  if l == Gerade { ker.Oops() }
  X.Art = DKW
  X.uint32 = m0 * dw + uint32(n)
  Nr = append (Nr, uint(X.uint32))
  X.Kilometrierung = NK
  X.lage = l
  X.richtung, X.stellung = ND, r
  X.uint = 1
  X.l, X.c = z, s
  D[n] = X
  c := zelle.New()
  c.DKW (n, l, r, z, s)
  X.Sequence.Ins (c)
  B[X.uint32] = X
  x, y := int(s) * W1, Y0 + int(z) * H1 - H2
  X.ins (uint(X.uint32), Gegen, x, y)
  X.ins (uint(X.uint32), Mit, x + W1, y)
  if l == Rechts {
    X.ins (uint(X.uint32), Gegen, x, y - H2)
    X.ins (uint(X.uint32), Mit, x + W1, y + H2)
  } else {
    X.ins (uint(X.uint32), Gegen, x, y + H2)
    X.ins (uint(X.uint32), Mit, x + W1, y - H2)
  }
  X.check ("DKW", n)
}

func (x *block) IstDKW() bool {
  return x.Art == DKW
}

func (x *block) Weichenrichtung() Richtung {
  if ! x.IstWeiche() { ker.Oops() }
  return x.richtung
}

func (x *block) Verzweigungsrichtung() Kilometrierung {
  if ! x.IstWeiche() { ker.Oops() }
  return x.Kilometrierung
}

func (x *block) Stellen (s Richtung) {
  switch x.Art {
  case Weiche, DKW:
    if x.uint != 1 { ker.Shit() }
    x.stellung = s
    x.Sequence.Seek (0)
    z := x.Sequence.Get().(zelle.Zelle)
    z.Stellen (s)
    x.Sequence.Put (z)
  default:
    ker.Oops()
  }
  f := Besetztfarbe
  if  x.Frei() {
    f = Freifarbe
  }
  x.Ausgeben (f)
}

/*/
func (x *block) Umstellen() {
  switch x.Art {
  case Weiche:
    if x.stellung == Gerade {
      x.stellung = x.richtung
    } else {
      x.stellung = Gerade
    }
  case DKW:
    if x.stellung == Gerade {
      x.stellung = Links // oder Rechts, ist egal
    } else {
      x.stellung = Gerade
    }
  default:
    ker.Oops()
  }
  x.Sequence.Seek (0)
  z := x.Sequence.Get().(zelle.Zelle)
  z.Stellen (x.stellung)
  x.Sequence.Put (z)
}
/*/

func (x *block) Pos() (uint, uint) {
  return x.l, x.c
}

func (x *block) Zeile() uint {
  return x.l
}

func (x *block) Stellung() Richtung {
  if ! x.IstWeiche() && ! x.IstDKW() { ker.Oops() }
  return x.stellung
}

func (x *block) Signaltyp (k Kilometrierung) s.Typ {
  return x.sig[k].Signaltyp()
}

func (x *block) SignalStellen (k Kilometrierung, st s.Stellung) {
  if x.Art == Knick || x.Art == Weiche || x.Art == DKW {
    return // ker.Panic (text[x.Art])
  }
  x.sig[k].Stellen (st)
}

func (x *block) Belegt (l, c uint) bool {
  return x.Sequence.ExPred (func (a any) bool {
                              return a.(zelle.Zelle).HatPosition (l, c)
                            }, true)
}

func (x *block) Freigeben() {
  x.zustand = frei
  x.Ausgeben (Freifarbe)
}

func (x *block) Frei() bool {
  return x.zustand == frei
}

func (x *block) Besetzen() {
  x.zustand = besetzt
  x.Ausgeben (Besetztfarbe)
}

func (x *block) AnkunftBesetzen() {
  x.Besetzen()
  x.Blinken()
}

func (x *block) Befahren() {
  x.zustand = befahren
  x.Ausgeben (Zugfarbe)
}

func (x *block) Besetzt() bool {
  return x.zustand == besetzt || x.zustand == befahren
}

func (x *block) Farbe() col.Colour {
  switch x.zustand {
  case frei:
    return Freifarbe
  case besetzt:
    return Besetztfarbe
  case befahren:
    return Zugfarbe
  }
  return Nichtfarbe
}

func (x *block) Blinken() {
  return
  f := x.Farbe()
  const t = 100
  for i := 0; i < 10; i++ {
    x.Ausgeben (Nichtfarbe)
    time.Msleep (t)
    x.Ausgeben (f)
    time.Msleep (t)
  }
}

func (x *block) Ausgeben (f col.Colour) {
  if x.Empty() { return }
  scr.Lock1()
  x.Sequence.Trav (func (a any) { a.(zelle.Zelle).Ausgeben (f) })
  for k := Mit; k < NK; k++ {
    x.sig[k].Ausgeben()
  }
  scr.Unlock1()
}

func (x *block) Länge() uint {
  return x.uint
}

func (X *block) UnterMaus() bool {
  return X.Sequence.ExPred (func (a any) bool { return a.(zelle.Zelle).UnterMaus() }, true)
}

func (x *block) Verzweigt (k Kilometrierung) bool {
  switch x.Art {
  case Weiche:
    return x.Kilometrierung == k
  case DKW:
    return true
  }
  return false
}

func gefunden() uint {
  for i := uint(0); i < uint(len(Nr)); i++ {
    n := Nr[i]
    b := B[n].(*block)
    if b.UnterMaus() {
      return n
    }
  }
  return 0
}
