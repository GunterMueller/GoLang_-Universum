package main

// (c) Christian Maurer   v. 171018 - license see µU.go

import (
  "µU/mode"
  "µU/scr"
  "µU/errh"
  "µU/img"
)

func main() {
  scr.New(0, 0, mode.XGA); defer scr.Fin()
  const a = 60
  var p = []string { "Platon", "Sokrates", "Aristoteles", "Cicero", "Pythagoras",
                     "Diogenes", "Thales", "Epikur", "Heraklit", "Anaxagoras",
                     "Protagoras", "Demokrit", "Theophrast" }
  var x, y uint
  for _, s := range p {
    img.Get (s, x, y)
    errh.Error0 (s)
    if x + 2 * a < scr.Wd() {
      x += a
    } else {
      x = 0; y = a
    }
  }
  errh.Error ("So viele Philosophen:", uint(len(p)))
}
