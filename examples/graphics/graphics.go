package main

// (c) Christian Maurer   v. 241005 - license see µU.go

import (
  "µU/scr"
  . "examples/internal"
)

func main() {
  s := scr.NewWH (0, 0, 1594, 1150); defer s.Fin()
  s.Cls()
  s.Go (Draw, 3, -6, 2, 0, 0, 0, 0, 0, 1)
}
