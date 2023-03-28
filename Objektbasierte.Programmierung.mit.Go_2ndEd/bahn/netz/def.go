package netz

// (c) Christian Maurer   v. 230305 - license see µU.go

const
  M = 11 // maximale Länge der Bahnhofsnamen
const (
  Bahnheim = uint(iota)
  Bahnhausen
  Bahnstadt
  Eisenheim
  Eisenstadt
  Eisenhausen
  Server
  N
  A = N - 1
)
var (
  MeinBahnhof uint
  MeinName string
)

func Name (n uint) string { return name(n) }

func AnzahlNachbarn (n uint) uint { return anzahlNachbarn(n) }

func Nachbar (n, i uint) uint { return nachbar(n,i) }

func Aktivieren() { aktivieren() }

func EinfahrtFreigeben (n uint) { einfahrtFreigeben(n) }

func EinfahrtBesetzen (n uint) { einfahrtBesetzen(n) }

func EinfahrtBesetzt (n uint) bool { return einfahrtBesetzt(n) }
