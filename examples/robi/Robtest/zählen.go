package main

// (c) Christian Maurer   v. 240801
//
// >>> Tiefensuche

import (
  "µU/env"
  . "Robi"
)
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
    if env.E() {
      HinweisAusgeben ("number of blocks=", anzahlKlötze)
    } else {
      HinweisAusgeben ("Anzahl der Klötze =", anzahlKlötze)
    }
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
  if env.E() {
    FehlerMelden ("number of blocks=", anzahlKlötze)
  } else {
    FehlerMelden ("Anzahl der Klötze =", anzahlKlötze)
  }
  Fertig()
}
