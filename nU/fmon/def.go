package fmon

// (c) Christian Maurer   v. 171125 - license see nU.go

import . "nU/obj"

type FarMonitor interface {

// Vor.: i < Anzahl der Monitorfunktionen von x.
// Der Wert von a ist mit dem Aufruf zum Anbieter gesendet.
// Die Monitorfunltion fs(_, i) ist auf dem Anbieter
// ausgeführt; ggf. war der aufrufende Prozess solange
// blockiert, bis das i-te Prädikat ps(-, i) true war.
// Liefert den Wert, den der Anbieter zurückgeschickt hat.
  F (a Any, i uint) Any

// Alle von x belegten Netzkanäle sind geschlossen.
  Fin()
}

// Vor.: fs und ps sind in ihrem 2. Argument für alle i < n
//       definiert. h ist der Name eines Rechners, der in
//       /etc/hosts enthalten oder per DNS erreichbar ist.
//       Die Ports p..p+n-1 sind nicht von einem
//       Netzwerkdienst benutzt. Im Fall s == true ist
//       New von einem Prozess auf Rechner h aufgerufen.
// Liefert einen fernen Monitor mit n Monitorfunktionen.
// Für alle i < n sind fs(_,i) die i-te Monitorfunktion
// und ps(_, i) das zugehörige Prädikat, mit dem
// entschieden wird, ob die Monitorfunktion ausgeführt
// werden kann.
// h ist der Name des Anbieters, der den Monitor ausführt
// und p..p+n-1 sind die Ports, die von der TCP-IP-
// Verbindung zwischen Anbieter und Kunden benutzt wird.
// Die gebrauchten Netzkanäle sind geöffnet.
// Der ferne Monitor läuft genau dann als Anbieter,
// wenn s == true gilt; andernfalls als Kunde.
// Für a == nil können unterschiedlich große Objekte
// mit einer Codelänge <= 65536 gesendet und empfangen
// werden. In diesem Fall liefert der Anbieter als Wert
// eines Monitoraufrufs eine Bytefolge vom Typ Stream
// und der Aufrufer muss diese Folge selber decodieren;
// andernfalls wird das von den Empfangsoperationen
// der verwendeten Netzwerkkanäle erledigt.
func New (a Any, n uint, fs FuncSpectrum, ps PredSpectrum,
          h string, p uint16, s bool) FarMonitor {
  return new_(a,n,fs,ps,h,p,s)
}
