package main

// (c) Christian Maurer   v. 230924 - license see µU.go

// >>> Serielles Netzwerk als Filter: Sieb des Eratosthenes

import (
  "µU/col"
  "µU/scr"
  "µU/errh"
)
const
  n = 313

func generieren (aus chan uint) {
  a := uint(2)
  aus <- a
  a++
  for {
    aus <- a
    a += 2
  }
}

func filtrieren (ein, aus, weg chan uint) {
  Primzahl := <-ein
  weg <- Primzahl
  var a uint
  for {
    a = <-ein
    if (a % Primzahl) != 0 {
      aus <- a
    }
  }
}

func konsumieren (ein chan uint) {
  for { <-ein }
}

func ausgeben (ein chan uint, fertig chan bool) {
  Z, S := uint(0), uint(0)
  var a uint
  for i := 1; i < n; i++ {
    a = <-ein
    scr.Colours (col.Yellow(), col.DarkGray())
    scr.WriteNat (a, Z, S)
    if S + 9 < scr.NColumns() {
      S += 6
    } else {
      Z++
      S = 0
    }
  }
  fertig <- true
}

func main() {
  scr.NewWH (0, 0, 800, 600); defer scr.Fin()
  var Roehre [n]chan uint
  for i := 0; i < n; i++ {
    Roehre[i] = make (chan uint)
  }
  Hahn := make (chan uint)
  fertig := make (chan bool)
  go generieren (Roehre [0]) // Generator
  for i := 1; i < n; i++ {
    go filtrieren (Roehre[i-1], Roehre[i], Hahn) // Filter [i]
  }
  go konsumieren (Roehre[n - 1]) // Konsument
  go ausgeben (Hahn, fertig) // Ausgabe
  <-fertig
  errh.Error2 ("habe fertig die ersten", n, "Primzahlen", 0)
}
