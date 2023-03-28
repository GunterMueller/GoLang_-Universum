package main

// (c) Christian Maurer   v. 230311 - license see µU.go

import (
  "µU/str"
  "µU/col"
  "µU/scr"
  "µU/box"
  "µU/errh"
  "µU/files"
  . "µU/menue"
  "life/species"
  "life/world"
)
var
  m = world.Mode()

func defined() (string, bool) {
  bx := box.New()
  w := scr.NColumns()
  bx.Wd (w)
  bx.Colours (col.Black(), col.LightWhite())
  bx.Write ("Welt:" + str.New(w - 5), scr.NLines() - 1, 0)
  const n = world.Len
  bx.Wd (n)
  name := str.New (n)
  bx.Colours (col.LightWhite(), col.Black())
  for {
    bx.Edit (&name, scr.NLines() - 1, 6)
    if str.Alphanumeric (name) {
      break
    } else {
      errh.Error0 ("Es dürfen nur Buchstaben und Ziffern im Namen vorkommen")
    }
  }
  str.OffSpc (&name)
  errh.DelHint()
  return name, ! str.Empty (name)
}

func sim() {
  w := world.New()
  for {
    if name, ok := defined(); ok {
      w.Name (name)
      w.Write()
      w.Edit()
    } else {
      break
    }
  }
}

func main() {
  scr.New (0, 0, m)
  scr.ScrColourB (col.LightWhite())
  scr.Cls()
  files.Cds()
  x := New ("Spiel des Lebens")
  game := New ("Game of Life (John Conway)")
  game.Leaf (func() { world.Sys (species.Life); sim() }, true)
  x.Ins (game)
  ecosys := New ("Ökosystem aus Füchsen, Hasen und Pflanzen")
  ecosys.Leaf (func() { world.Sys (species.Eco); sim() }, true)
  x.Ins (ecosys)
  x.Exec()
  scr.Fin()
}
