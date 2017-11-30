package nchan

// (c) Christian Maurer   v. 171125 - license see nU.go

import . "nU/obj"

const Port0 = uint16(49152) // erster "privater" Port

type NetChannel interface { // Kanäle zum netzweiten Austausch
                            // von Objekten zwischen Prozessen

  Chan() (chan Any, chan Any)

// Vor.: a ist vom Typ von x.
// Das Objekt a ist auf x zum Kommunikationspartner
// des aufrufenden Prozesses gesendet.
  Send (a Any)

// Liefert einen Slice von Bytes, wenn x nicht an einen Typ
// gebunden ist. In diesem Fall muss der Empfänger diesen
// Slice in ein Objekt des passenden Typs decodieren.
// Liefert andernfalls das Objekt vom Typ von x,
// das vom Kommunikationspartner auf x gesendet wurde, falls
// es vollständig übertragen wurde, andernfalls nil.
// Der aufrufende Prozess war ggf. solange
// blockiert, bis ein Objekt empfangen wurde.
  Recv() Any

// Alle von x benutzten Netzkanäle sind geschlossen.
  Fin()
}

// h0 always denotes the name of the host running the calling process.

// For all constructors for the first parameter a the following holds:
//      For a == nil, arbitrary objects of Codelen <= 1<<32 can be passed.
//      In this case, calls of Recv() return only slices of bytes, so
//      it is up to the receiver to Decode the object wanted to receive.
//
// Pre: h is contained in /etc/hosts or denotes a DNS-resolvable host.
//      h is different from h0.
//      p > 0; p is not used on h0 or h by a network service.
//      The communication partner calls New with
//      - an object of the same type as the type of a as 1st argument,
//      - with h0 as 2nd argument and
//      - an identical value of the 3rd argument.
// Returns a asynchronous 1:1 channel for messages of the type of a
// between h0 and h over port p.
// p is now used on h0 and h by a network service.
func NewD (a Any, h string, p uint16) NetChannel { return newd(a,h,p) }


// Vor.: me != i; me ist die Identität des aufrufenden
//       Prozesses und i ist die Identität des Prozesses,
//       der auf dem Rechner mit dem Namen h läuft (h ist
//       in /etc/hosts eingetragen oder per DNS auflösbar).
//       Port0 + p < 65536; Port0 + p ist nicht von einem
//       Netzwerkdienst auf den beteiligten Rechnern belegt.
//       Der Kommunikationspartner ruft New mit einem Objekt
//       des gleichen Typs wie dem von a und dem gleichen Port
//       auf, aber mit vertauschten Werten von me und i und
//       dem Namen des Rechners, auf dem der aufrufende Prozess
//       läuft.
// Liefert einen asynchronen Kanal für Botschaften zwischen
// dem Rechner, auf dem der aufrufende Prozess läuft,
// und einem Prozess, der auf h läuft, als
// Kommunikationspartner des aufrufenden Prozesses.
// Port0 + p ist jetzt von einem Netzwerkdienst
// auf diesen beiden Rechnern belegt.
// Für a == nil können unterschiedlich große Objekte
// einer Codelänge <= 65536 ausgetauscht werden.
// In diesem Fall liefern Aufrufe von Recv() Slices von Bytes,
// die der Empfänger selber decodieren muss, was allerdings
// voraussetzt, dass beide Kommunikationspartner den Kanal für
// den Austausch von Botschaften des gleichen Typs benutzen.
  func New (a Any, me, i uint, h string, p uint16) NetChannel { return new_(a,me,i,h,p) }

// Siehe obige Funktion.
// Dieser Konstruktor wird für die Erzeugung eines fernen Monitors gebraucht:
// h ist genau dann der anbietende Monitor, wenn s den Wert true hat. 
  func NewN (a Any, h string, p uint16, s bool) NetChannel { return newn(a,h,p,s) }

// Wichtig: Bei aufeinanderfolgenden Aufrufen von New
//          muss zur Vermeidung von Verklemmungen
//          die korrekte Paarung eingehalten werden!

// Vor.: i, j < n; a < 2.
// Liefert verschiedene Zahlen < NPorts(n,a).
func Port (n, i, j, a uint) uint16 { return port(n,i,j,a) }
func NPorts (n, a uint) uint { return nPorts(n,a) }
