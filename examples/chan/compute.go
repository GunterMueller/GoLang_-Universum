package main

// (c) Christian Maurer   v. 241005 - license see µU.go

// >>> Einführendes Beispiel zum Botschaftenaustausch:
//     Auftraggeber wartet auf Addition, Addierer führt sie durch

import (
  "µU/scr"
  "µU/N"
)

func order (K chan uint) {
  scr.Write ("1. Zahl:", 0, 0)
  scr.Write ("2. Zahl:", 1, 0)
  scr.Write ("  Summe:", 2, 0)
  x := uint(0)
  y := uint(0)
  var s uint
  N.SetWd (4)
  for {
    N.Edit (&x, 0, 12)
    K <- x
    if x == 0 { break }
    N.Edit (&y, 1, 12)
    K <- y
    s = <-K
    N.Write (s, 2, 12)
  }
}

func add (K chan uint, f chan bool) {
  var x, y uint
  for {
    x = <-K
    if x == 0 { break }
    y = <-K
    K <- x + y
  }
  f <- true
}

func main() {
  scr.NewWH (0, 0, 800, 600); defer scr.Fin()
  c := make (chan uint)
  done := make (chan bool)
  go order (c) // Auftraggeber
  go add (c, done) // Addierer
  <-done
}
