package main

// (c) Christian Maurer   v. 210404 - license see µU.go

import (
  "µU/mode"
  "µU/scr"
  "µU/errh"
  "mini/prog"
)

func main () {
  scr.New (0, 0, mode.VGA); defer scr.Fin()
  program := prog.New()
  program.GetLines()
  fail, n := program.Parse()
  if fail == "" {
    program.Write()
    program.Edit()
    program.Run()
  } else {
    errh.Error (fail + " <- fehlerhafte Programmzeile Nr.", n + 1) // TODO call Editor
  }
}
