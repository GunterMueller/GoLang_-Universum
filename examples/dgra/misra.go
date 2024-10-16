package main

// (c) Christian Maurer   v. 241013 - license see µU.go

import (
  "µU/env"
  "µU/scr"
  "µU/ego"
  "µU/dgra"
)

func main() {
  if env.NArgs() != 1 { println ("argument missing"); return }
  scr.NewWH (0, 300, 1, 1); defer scr.Fin()
  g := dgra.G8ringdirord (ego.Me())
  g.SetRoot (0)
  g.Misra()
  done := make (chan int, 1); <-done
}
