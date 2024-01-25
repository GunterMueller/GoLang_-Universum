package nchan

// (c) Christian Maurer   v. 220702 - license see nU.go

const
  Port0 = uint16(49152) // erster "privater" Port
type
  NetChannel interface { // Kanäle zum netzweiten Austausch
                         // von Objekten zwischen Prozessen

  Chan() (chan any, chan any)

// Vor.: a ist vom Typ von x.
// Liefert genau dann nil, wenn das Objekt a ist auf x zum
// Kommunikationspartner des aufrufenden Prozesses gesendet ist.
  Send (a any)

// Liefert einen Slice von Bytes, wenn x nicht an einen Typ
// gebunden ist. In diesem Fall muss der Empfänger diesen
// Slice in ein Objekt des passenden Typs decodieren.
// Liefert andernfalls das Objekt vom Typ von x, das vom
// Kommunikationspartner auf x gesendet wurde, falls es
// es vollständig übertragen wurde, andernfalls das
// Musterobjekt von x. Der aufrufende Prozess war ggf.
// solange blockiert, bis ein Objekt empfangen wurde.
  Recv() any

// Alle von x benutzten Netzkanäle sind geschlossen.
  Fin()
}

// h0 ist der Rechner, auf dem der aufrufende Prozess
// läuft und h ist der Rechner mit dem Namen n.
// Vor.: me != i; me ist die Identität des aufrufenden
//       Prozesses und i ist die Identität des Prozesses,
//       der auf dem Rechner mit dem Namen n läuft (n ist
//       in /etc/hosts eingetragen oder per DNS auflösbar).
//       Port0 + p < 65536; Port0 + p ist nicht von
//       einem Netzwerkdienst auf h0 oder h belegt.
//       Der Kommunikationspartner ruft New mit einem
//       Objekt des gleichen Typs wie dem von a und dem
//       gleichen Port auf, aber mit vertauschten Werten
//       von me und i und dem Namen von h0 anstelle von n.
// Liefert einen asynchronen Kanal für Botschaften zwischen
// den beiden beteiligten Rechnern. Port0 + p ist jetzt
// von einem Netzwerkdienst auf beiden Rechnern belegt.
// a ist das Musterobjekt dieses Kanals.
// Für a == nil können unterschiedlich große Objekte
// gesendet werden. In diesem Fall liefern Aufrufe von Recv()
// Slices von Bytes, die der Empfänger selber decodieren muss.
  func New (a any, me, i uint, n string, p uint16) NetChannel { return new_(a,me,i,n,p) }

// Siehe obige Funktion.
// Dieser Konstruktor wird für die Erzeugung eines fernen Monitors gebraucht:
// h ist genau dann der anbietende Monitor, wenn s den Wert true hat. 
  func NewN (a any, n string, p uint16, s bool) NetChannel { return newn(a,n,p,s) }

// Wichtig: Bei aufeinanderfolgenden Aufrufen von New
//          muss zur Vermeidung von Verklemmungen
//          die korrekte Paarung eingehalten werden!

// Vor.: i, j < n; a < 2.
// Liefert verschiedene Zahlen < NPorts(n,a).
func Port (n, i, j, a uint) uint16 { return port(n,i,j,a) }
func NPorts (n, a uint) uint { return nPorts(n,a) }
