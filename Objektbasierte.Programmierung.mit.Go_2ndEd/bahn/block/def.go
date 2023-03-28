package block

// (c) Christian Maurer   v. 230305 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  . "bahn/kilo"
  . "bahn/richtung"
  "bahn/signal"
)
const
  H = 100
type
  Art byte; const (
  Dfg = Art(iota) // Durchfahrgleis
  AsM // AbstellgleisMit
  AsG // AbstellgleisGegen
  EfM // EinfahrgleisMit
  EfG // EinfahrgleisGegen
  AfM // AusfahrgleisMit
  AfG // AusfahrgleisGegen
  EAM // EinAusfahrgleisMit
  EAG // EinAusfahrgleisGegen
  Knick
  Weiche
  DKW
  NArten
)
type
  Block interface {

  Object

  Stringer

// Vor.: x ist nicht leer.
// Liefert die Nummer von x modulo M.
  Nummerkurz() uint

// Liefert die Nummer von x.
  Nummer() uint

// Liefert die Schräglage von x.
  Schräglage() Richtung

// Vor.: z < AnzahlZeilen, s + l <= AnzahlSpalten, s < sn < s + l.
//       Die Position (z, s) ist noch nicht durch einen Block belegt.
// x beginnt bei (z, s) und verläuft gerade nach rechts (a = gerade: horizontal,
// a = links: diagonal steigend, a = rechts: diagonal fallend) mit der Spaltenlänge l. 
// x hat die Nummer n, sie erscheint n in der Spalte sn.
// x hat die Signale ...
//  GleisDefinieren (n uint, a Art, lage Richtung, l, z, s, sn uint,
//            gn uint, gt signal.Typ, g Kilometrierung, gst signal.Stellung, gz, gsn uint,
//            mn uint, mt signal.Typ, m Kilometrierung, mst signal.Stellung, mz, msn uint)
  GleisDefinieren (n uint, a Art, lage Richtung, l, z, s, sn uint,
                   gt signal.Typ, g Kilometrierung, gst signal.Stellung, gz, gsn uint,
                   mt signal.Typ, m Kilometrierung, mst signal.Stellung, mz, msn uint)

// Liefert genau dann true, wenn x ein Gleis ist.
  IstGleis() bool

// Liefert genau dann true, wenn x ein Durchfahrgleis ist.
  IstDurchfahrgleis() bool

// Liefert genau dann true, wenn x ein Einfahrgleis ist.
  IstEinfahrgleis() bool

// Liefert genau dann true, wenn x ein Ausfahrgleis ist.
  IstAusfahrgleis() bool

// Liefert genau dann true, wenn x ein EinAusfahrgleis ist.
  IstEinAusfahrgleis() bool

  KnickDefinieren (n uint, k Kilometrierung, r Richtung, z, s uint)

  IstKnick() bool

// Vor.: z < AnzahlZeilen, s < AnzahlSpalten, l ungleich r, r = Links oder Rechts.
//       Die Position (z, s) ist noch nicht durch einen Block belegt.
// x ist nicht leer. x hat die Nummer n. 
// k ist die Kilometrierung, in der sich die Weiche verzweigt.
// l ist die Lage des durchgehenden Astes der Weiche
// (l = Gerade/Links/Rechts: horizontal/diagonal steigend/diagonal fallend). 
// x ist für r == Links bzw. Rechts eine Links-/ bzw. Rechtsweiche
// mit der Stellung st an der Position (z, s).
  WeicheDefinieren (n uint, k Kilometrierung, l, r, st Richtung, z, s uint)

// Liefert genau dann true, wenn x eine Weiche ist.
  IstWeiche() bool

// Vor.: l != Gerade.
// x ist eine DKW mit der Nummer n, der Schräglage l,  und der Position (z, s).
  DKWDefinieren (n uint, l, r Richtung, z, s uint)

// Liefert genau dann true, wenn x eine DKW ist.
  IstDKW() bool

// Vor.: x ist eine Weiche.
// Liefert die Richtung des abzweigenden Astes von x.
  Weichenrichtung() Richtung

// Vor.: x ist eine Weiche.
// Liefert die Kilometrierungsrichtung, in die x verzweigt ist.
  Verzweigungsrichtung() Kilometrierung

// Liefert die Position am linken Rand von x.
  Pos() (uint, uint)
  Zeile() uint

// Vor.: x ist eine Weiche oder eine DKW.
// x hat die Stellung r.
  Stellen (r Richtung)

// Vor.: x ist eine Weiche oder eine DKW.
// Die Stellung von x hat gewechselt.
/////////////////////////////////////////////////////////////////////////////  Umstellen()

// Vor.: x ist eine Weiche.
// Liefert die Stellung von x.
  Stellung() Richtung 

// Wenn x in Richtung k ein Signal hat, hat es die Stellung s. 
// Das Signal ist ausgegeben.
  SignalStellen (k Kilometrierung, s signal.Stellung)

// x ist ausgegeben.
// Ist x ein Gleis mit einer Nummer > 0, ist diese Nummer mit ausgegeben.
  Ausgeben (f col.Colour)

// Vor.: x ist nicht leer. l < AnzahlZeilen, c < AnzahlSpalten.
// Liefert genau dann true, wenn x die Position (l, c) belegt.
  Belegt (l, c uint) bool

// x ist nicht besetzt.
  Freigeben()

// Liefert genau dann true, wenn x nicht besetzt ist.
  Frei() bool

// Vor.: x ist nicht leer.
// x ist mit einem stehenden Zug besetzt.
  Besetzen()

// Vor.: x ist nicht leer.
// x ist mit einem fahrenden Zug besetzt.
  Befahren()

// Vor.: x ist nicht leer.
// Der Block ist mit einem stehenden Zug besetzt und blinkt.
  AnkunftBesetzen()

// Liefert genau dann true, wenn x mit einem stehenden oder fahrenden Zug besetzt ist.
  Besetzt() bool

// Liefert die Farbe von x je nach Zustand frei, besetzt oder befahren.
  Farbe() col.Colour

// Vor.: x ist nicht leer.
// Liefert den Typ des Signals von x in direction k, falls es eins gibt;
// andernfalls NT.
  Signaltyp (k Kilometrierung) signal.Typ

// x blinkt einen Augenblick lang.
  Blinken()

// Liefert die Spaltenlänge von x.
  Länge() uint

// Liefert genau dann true, wenn der Mauszeiger auf x zeigt.
  UnterMaus() bool

// Liefert genau dann true, wenn x eine Weiche oder DKW mit der Verzweigungsrichtung k ist.
  Verzweigt (k Kilometrierung) bool
}

var
  Nr []uint
const
  M = 300
var
  B, W, D [M]Block

// Liefert einen neuen leeren Block.
func New() Block { return new_() }

// Liefert die Anzahl der Paare.
func NPaare() uint { return nPaare() }

// Vor.: i < NPaare()
// Liefert das i-te Paar.
func Paar (i uint) (uint, uint) { return paar(i) }

// Liefert genau dann die Nummer des Blocks, der unter der Maus liegt;
// in diesem Fall ist sie > 0. liefert andernfalls 0.
func Gefunden() uint { return gefunden() }
