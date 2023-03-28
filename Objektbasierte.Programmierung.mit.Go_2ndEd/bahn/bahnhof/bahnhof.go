package bahnhof

// (c) Christian Maurer   v. 230305 - license see µU.go

// TODO Signalschaltungen überprüfen; Hp2 auch, wenn Fahrstraße ablenkend ist

import (
  "os"
  "sync"
  "µU/ker"
  "µU/time"
  "µU/env"
  . "µU/obj"
  "µU/kbd"
  "µU/scr"
  "µU/errh"
//  "µU/pseq"
  "µU/gra"
  "bahn/signal"
  b "bahn/block"
  "bahn/fahrweg"
  "bahn/netz"
  . "bahn/farbe"
  . "bahn/kilo"
  . "bahn/richtung"
)
type
  bahnhof struct {
                 gra.Graph
                 Kilometrierung
                 }
const (
  A = netz.A
  suffix = ".bs"
)
var (
  nachbar [b.H]uint
  mutex sync.Mutex
  fahrende int
)

func init() {
  netz.MeinBahnhof = netz.Eisenhausen
  if env.NArgs() > 0 {
    netz.MeinBahnhof = env.N (1)
  }
  for i := uint(0); i < b.H; i++ {
    nachbar[i] = 0
  }
}

func new_() Bahnhof {
  x := new (bahnhof)
  x.Graph = gra.New (true, uint(0), nil)
  x.Kilometrierung = Mit
  return x
}

func (x *bahnhof) Empty() bool {
  return x.Graph.Empty()
}

func (x *bahnhof) Clr() {
  x.Graph.Clr()
}

func (x *bahnhof) Codelen() uint {
  return 4 +
         4 +
         x.Graph.Codelen()
}

func (x *bahnhof) Encode() Stream {
  s := make (Stream, x.Codelen())
  i, a := uint32(0), uint32(4)
  s[i] = byte(x.Kilometrierung)
  i += a
  k := uint32(x.Graph.Codelen())
  copy (s[i:i+a], Encode (k))
  i += a
  copy (s[i:i+k], x.Graph.Encode())
  return s
}

func (x *bahnhof) Decode (s Stream) {
  i, a := uint32(0), uint32(4)
  x.Kilometrierung = Kilometrierung(s[i])
  i += a
  k := Decode (uint32(0), s[i:i+a]).(uint32)
  i += a
  x.Graph.Decode (s[i:i+k])
}

func (x *bahnhof) codelenBlocks() uint {
  return b.M * b.New().Codelen()
}

func (x *bahnhof) encodeBlocks() Stream {
  s := make(Stream, x.codelenBlocks())
  i, a := uint(0), b.New().Codelen()
  for j := 0; j < b.M; j++ {
    copy (s[i:i+a], Encode (b.B[j]))
    i += a
  }
  return s
}

func (x *bahnhof) decodeBlocks (s Stream) {
  i, a := uint(0), b.New().Codelen()
  for j := 0; j < b.M; j++ {
    b.B[j].Decode (s[i:i+a])
    b.B[j].Ausgeben (b.B[j].Farbe())
    i += a
  }
}

func check (t string, bk b.Block, n uint) {
  s := bk.Encode()
  bk1 := b.New()
  bk1.Decode (s)
  if ! bk.Eq (bk1) { ker.Panic1 ("Codierfehler bei " + t, n) }
}

func (x *bahnhof) gleis (n uint, a b.Art, lage Richtung, l, z, s, sn uint,
                         gt signal.Typ, g Kilometrierung, gst signal.Stellung, gsn uint,
                         mt signal.Typ, m Kilometrierung, mst signal.Stellung, msn uint){
  if l == 0 { ker.Oops() }
  if sn >= s + l { println (n, sn, s, l); ker.Oops() }
  bn := b.New()
  bn.GleisDefinieren (n, a, lage, l, z, s, sn, gt, g, gst, z, gsn, mt, m, mst, z, msn)
  x.Graph.Ins (bn.Nummer())
  check ("Gleis", bn, n)
}

func (x *bahnhof) gleisBesetzen (n uint) {
  b.B[n].Besetzen()
}

func (x *bahnhof) knick (n uint, k Kilometrierung, r Richtung, z, s uint) {
  bn := b.New()
  bn.KnickDefinieren (n, k, r, z, s)
  x.Graph.Ins (bn.Nummer())
  check ("Knick", bn, n)
}

func (x *bahnhof) weiche (n uint, k Kilometrierung, l, r, st Richtung, z, s uint) {
  if l == r { ker.Panic ("l == r") }
  if r == Gerade { ker.Panic ("r == Gerade") }
  if st == Entgegen (r) { ker.Panic1 ("st = Gegenrichtung (r)", n) }
  bn := b.New()
  bn.WeicheDefinieren (n, k, l, r, st, z, s)
  x.Graph.Ins (bn.Nummer())
  check ("Weiche", bn, n)
}

func (x *bahnhof) dkw (n uint, l, st Richtung, z, s uint) {
  bn := b.New()
  bn.DKWDefinieren (n, l, st, z, s)
  x.Graph.Ins (bn.Nummer())
  check ("DKW", bn, n)
}

func (x *bahnhof) verbinden() {
  for i := uint(0); i < b.NPaare(); i++ {
    n, n1 := b.Paar (i)
    k, k1 := b.B[n].Nummer(), b.B[n1].Nummer()
// println (k, k1) // zum Testen der Verbindungen bei der Konstruktion von Bahnhöfen
    if x.Graph.Ex2 (k, k1) {
      x.Graph.Edge (nil)
    }
  }
}

func (x *bahnhof) Fahrweglänge() uint {
  return uint(len(x.Graph.ShortestPath()))
}

func (x *bahnhof) Fahrwegblock (i uint) b.Block {
  p := x.Graph.ShortestPath()
  if i >= uint(len(p)) { ker.Oops() }
  return b.B[p[i].(uint)]
}

func (x *bahnhof) erzeugt (start, ziel b.Block) (Kilometrierung, fahrweg.Weg, bool) {
  w := fahrweg.New()
  s, z := start.Nummer(), ziel.Nummer()
  if s == 0 || z == 0 { ker.Shit() }
  if ! x.Graph.Ex2 (s, z) { ker.Oops() } // start colocal, ziel local
  x.Graph.FindShortestPathPred (func (a any) bool {
                     ba := b.B[a.(uint)]
                     if ! ba.Frei() {
                       return false
                     }
                     return true
                   })
  if x.Graph.NumMarked() <= 1 {
    x.Graph.Inv()
    if ! x.Graph.Ex2 (s, z) { ker.Oops() } // start colocal, ziel local
    x.Graph.FindShortestPathPred (func (a any) bool {
                       ba := b.B[a.(uint)]
                       if ! ba.Frei() {
                         return false
                       }
                       return true
                     })
    if x.Graph.NumMarked() <= 1 {
      errh.Error2 ("Es gibt keinen Fahrweg von Gleis", start.Nummerkurz(),
                                         "nach Gleis", ziel.Nummerkurz())
      start.Besetzen()
      return NK, fahrweg.New(), false
    }
  }
//  f := "" // zum Testen des Fahrwegs
  for i := uint(0); i < x.Fahrweglänge(); i++ {
    bf := x.Fahrwegblock (i)
    w.Insert (bf.Nummer())
    bf.Ausgeben (Zugfarbe)
//    f += " " + bf.String()
  }
//  errh.Error0 (f)
  if start.Less (ziel) {
    x.Kilometrierung = Mit
  } else {
    x.Kilometrierung = Gegen
  }
  start.Besetzen()
  h := signal.Hp1
  if w.Ablenkend() { h = signal.Hp2 }
  start.SignalStellen (x.Kilometrierung, h)
  if ! ziel.IstAusfahrgleis() && ! ziel.IstEinAusfahrgleis() {
    ziel.SignalStellen (x.Kilometrierung, signal.Hp0)
  }
  return x.Kilometrierung, w, true
}

func (x *bahnhof) fahren (w fahrweg.Weg, k Kilometrierung) {
  if w.Num() == 0 {
    fahrende--
    return
  }
  for i := uint(0); i < w.Num(); i++ {
    b.B[w.Nr(i)].Befahren()
  }
  n1 := w.Num() - 1
  for i := uint(0); i < w.Num(); i++ {
    time.Msleep (500 * b.B[w.Nr(i)].Länge())
    b.B[w.Nr(i)].SignalStellen (k, signal.Hp0)
    if i < n1 {
      if i == 0 {
        if b.B[w.Nr(0)].IstEinfahrgleis() || b.B[w.Nr(0)].IstEinAusfahrgleis() {
          von := nachbar[b.B[w.Nr(0)].Nummer()]
          nach := netz.MeinBahnhof
          n := A * von + nach
          if netz.EinfahrtBesetzt (n) {
            netz.EinfahrtFreigeben (n)
          }
        }
      }
      b.B[w.Nr(i)].Freigeben()
    } else { // b ist letzter Block von w
      von := netz.MeinBahnhof
      nach := nachbar[b.B[w.Nr(n1)].Nummer()]
      if b.B[w.Nr(n1)].IstAusfahrgleis() || b.B[w.Nr(n1)].IstEinAusfahrgleis() {
        n := A * von + nach
        for {
          if netz.EinfahrtBesetzt (n) {
            time.Sleep (1)
          } else {
            break
          }
        }
        b.B[w.Nr(n1)].Freigeben()
        netz.EinfahrtBesetzen (n)
      } else {
        time.Sleep (1)
        b.B[w.Nr(n1)].Besetzen()
      }
    }
  }
  fahrende--
}

/*/
func (x *bahnhof) put (name string) {
  s := x.encodeBlocks()
  n := len(s)
  t := make(Stream, n)
  f := pseq.New (t)
  f.Name (name + suffix)
  f.Put (s)
  f.Fin()
}
/*/

func (x *bahnhof) Start() uint {
  start := uint(0)
  scr.MousePointer (true)
  for {
    hinweis := true
    loop:
    for {
      if hinweis {
        errh.Hint ("Startgleis anklicken        Betrieb einstellen: Esc")
      }   
      switch c, _ := kbd.Command(); c {
      case kbd.Here:
        break loop
      case kbd.Esc:
        errh.Hint ("Das Programm wird beendet.")
        for fahrende > 0 {
          time.Sleep (1)
        }
        errh.DelHint()
/*/
        x.put (netz.MeinName)
/*/
        scr.Fin()
        os.Exit (1)
      default:
        hinweis = false
      }
    }
    start = b.Gefunden()
    bs := b.B[start]
    ok := bs.IstGleis() && bs.Besetzt()
    if ok && ( bs.IstEinfahrgleis() || bs.IstEinAusfahrgleis()) {
      bs.Blinken()
      break
    }
    if ok && bs.Besetzt() {
      bs.Blinken()
      break
    }
  }
  errh.DelHint()
  b.B[start].Besetzen()
  return start
}

func (X *bahnhof) StartZiel() (uint, uint) {
  var ziel uint
  startAnklicken:
  start := X.Start()
  scr.MousePointer (true)
  for {
    hinweis := true
    loop:
    for {
      if hinweis {
        errh.Hint ("Zielgleis anklicken         anderes Startgleis: Esc")
      }
      switch c, _ := kbd.Command(); c {
      case kbd.Esc, kbd.Back:
        goto startAnklicken
      case kbd.Here:
        break loop
      default:
        hinweis = false
      }
    }
    errh.DelHint()
    ziel = b.Gefunden()
    bz := b.B[ziel]
    ok := ziel != start && bz.IstGleis() && ! bz.Besetzt()
    if ok && ! bz.IstEinfahrgleis() {
      break
    }
  }
  return start, ziel
}

func (x *bahnhof) schalten (w fahrweg.Weg, k Kilometrierung) {
  var st Richtung
  for i := uint(0); i + 0 < w.Num(); i++ {
    n := w.Nr (i)
    bn := b.B[n]
    n %= b.H
    if bn.IstWeiche() {
      v, l, r := bn.Verzweigungsrichtung(), bn.Schräglage(), bn.Weichenrichtung()
      z := bn.Zeile()
      if k == Mit {
        z1 := b.B[w.Nr(i+1)].Zeile()
        if v == Mit {
          if l == Gerade {
            if z1 == z { st = Gerade } else { st = r }
          } else {
            if z1 == z { st = r } else { st = Gerade }
          }
        } else { // v == Gegen
          z1 := b.B[w.Nr(i-1)].Zeile()
          if l == Gerade {
            if z1 == z { st = Gerade } else { st = r }
          } else {
            if z1 == z { st = r } else { st = Gerade }
          }
        }
      } else { // k == Gegen
        if v == Mit {
          z1 := b.B[w.Nr(i-1)].Zeile()
          if l == Gerade {
            if z1 == z { st = Gerade } else { st = r }
          } else {
            if z1 == z { st = r } else { st = Gerade }
          }
        } else { // v == Gegen 
          z1 := b.B[w.Nr(i+1)].Zeile()
          if l == Gerade {
            if z1 == z { st = Gerade } else { st = r }
          } else {
            if z1 == z { st = r } else { st = Gerade }
          }
        }
      }
      b.W[n].Stellen (st)
    }
    if bn.IstDKW() {
      l := bn.Schräglage()
      z := bn.Zeile()
      if k == Mit {
        z1, z2 := b.B[w.Nr(i+1)].Zeile(), b.B[w.Nr(i-1)].Zeile()
        if l == Links {
          if z2 == z {
            if z1 == z { st = Gerade } else { st = Links }
          } else {
            if z1 == z { st = Links } else { st = Gerade }
          }
        } else {
          if z2 == z {
            if z1 == z { st = Gerade } else { st = Links }
          } else {
            if z1 == z { st = Links } else { st = Gerade }
          }
        }
      } else { // k == Gegen
        z1, z2 := b.B[w.Nr(i-1)].Zeile(), b.B[w.Nr(i+1)].Zeile()
        if l == Links {
          if z2 == z {
            if z1 == z { st = Gerade } else { st = Links }
          } else {
            if z1 == z { st = Links } else { st = Gerade }
          }
        } else {
          if z2 == z {
            if z1 == z { st = Gerade } else { st = Links }
          } else {
            if z1 == z { st = Links } else { st = Gerade }
          }
        }
      }
      b.D[n].Stellen (st)
    }
    if bn.IstGleis() {
      bn.SignalStellen (k, signal.Hp1)
    }
  }
  for i := uint(0); i < w.Num(); i++ {
    b.B[w.Nr(i)].Besetzen()
  }
}

func (x *bahnhof) einfahrgleis (n uint) uint {
  for i := uint(1); i < 40; i++ {
    if nachbar[i] == n && (b.B[i].IstEinfahrgleis() || b.B[i].IstEinAusfahrgleis()) {
      return i
    }
  }
  return 0
}

func (x *bahnhof) einfahren (i uint) {
  von := netz.Nachbar (netz.MeinBahnhof, i)
  n := A * von + netz.MeinBahnhof
  g := x.einfahrgleis (von)
  for {
    if b.B[g].Frei() {
      if netz.EinfahrtBesetzt(n) {
        b.B[g].Besetzen()
      }
    }
    time.Msleep (100)
  }
}

func (x *bahnhof) betreiben() {
  if netz.MeinBahnhof != netz.Server {
    for i := uint(0); i < netz.AnzahlNachbarn (netz.MeinBahnhof); i++ {
      go x.einfahren (i)
    }
  }
  var start, ziel b.Block
  for {
    s, z := x.StartZiel()
    start, ziel = b.B[s], b.B[z]
    if ! start.Besetzt() { ker.Oops() }
    if k, r, ok := x.erzeugt (start, ziel); ok && r.Num() > 0 {
      mutex.Lock()
      x.schalten (r, k)
      fahrende++
      go x.fahren (r, k) 
      mutex.Unlock()
    }
  }
}

/*/
func (x *bahnhof) get (name string) {
  filename := name + suffix
  f := pseq.New (b.New())
  f.Name (filename)
  n := f.Num()
  for i := uint(0); i < n; i++ {
    f.Seek (i)
    b.B[i] = f.Get().(b.Block)
    b.B[i].Ausgeben (b.B[i].Farbe())
    time.Sleep (2)
  }
  f.Fin()
  x.verbinden() // ?
}
/*/

func (x *bahnhof) ausgeben() {
  for i := uint(0); i < b.M; i++ {
    a := b.B[i]
    if ! a.Empty() {
      f := Freifarbe
      if a.Besetzt() {
        f = Besetztfarbe
      }
      a.Ausgeben (f)
    }
  }
}

var (
  sekunde uint
  warten = true
)

func aufServerWarten() {
  for warten {
    if s, _ := time.SecondsSinceUnix(); s > sekunde + 1 {
      ker.Panic ("Server nicht gestartet")
    }
    time.Msleep (100)
  }
}

func (x *bahnhof) Betreiben() {
  sekunde, _ = time.SecondsSinceUnix()
  go aufServerWarten()
  netz.Aktivieren()
  warten = false
  name := netz.MeinName
  scr.Name (name)
/*/
  if pseq.Length (name + suffix) > 0 {
    x.get (netz.MeinName)
  } else {
/*/
  switch netz.MeinBahnhof {
  case netz.Bahnheim:
    x.bahnheim()
  case netz.Bahnhausen:
    x.bahnhausen()
  case netz.Bahnstadt:
    x.bahnstadt()
  case netz.Eisenheim:
    x.eisenheim()
  case netz.Eisenstadt:
    x.eisenstadt()
  case netz.Eisenhausen:
    x.eisenhausen()
  }
/*/
  }
/*/
  x.ausgeben()
  x.betreiben()
}
