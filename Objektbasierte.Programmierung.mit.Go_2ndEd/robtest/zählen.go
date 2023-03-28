package main

// (c) Christian Maurer   v. 230309 - license see µU.go
//
// >>> Tiefensuche

import
  . "robi"
var
  anzahlKlötze uint

func weiter() bool {
  return ! VorRand() && ! VorMauer()
}

func suchen() {
  var linksWeiter, geradeWeiter, rechtsWeiter bool
  if Markiert() {
    return
  }
  Markieren()
  if ! Leer() {
    anzahlKlötze++
    HinweisAusgeben ("Anzahl der Klötze =", anzahlKlötze)
  }
  LinksDrehen(); linksWeiter = weiter(); RechtsDrehen()
  geradeWeiter = weiter()
  RechtsDrehen(); rechtsWeiter = weiter(); LinksDrehen()
  x, y := Pos()
  if linksWeiter {
    LinksDrehen()
    Laufen()
    suchen()
    RechtsDrehen()
  }
  Set (x, y)
  if geradeWeiter {
    Laufen()
    suchen()
  }
  Set (x, y)
  if rechtsWeiter {
    RechtsDrehen()
    Laufen()
    suchen()
    LinksDrehen()
  }
}

func main() {
  Laden ("Labyrinth")
  suchen()
  FehlerMelden ("Anzahl der Klötze =", anzahlKlötze)
  Fertig()
}
