package main

// (c) Christian Maurer   v. 221110 - license see µU.go
//
// >>> Ein Beispiel, wie fest definierte Figuren in eine eTafel eingebunden werden können.
// >>> Diese eTafel kann natürlich mit dem Programm epen bearbeitet werden.

import (
  . "µU/obj"
  "µU/env"
  "µU/str"
  "µU/col"
  "µU/scr"
  "µU/box"
  "µU/seq"
  "µU/pseq"
  "µU/files"
  "µU/fig2"
)

func store (s seq.Sequence, name string) {
  const suffix = ".epn"
  pseq.Erase (name)
  n := s.Codelen()
  buf := make (Stream, n)
  buf = s.Encode()
  file := pseq.New (buf)
  file.Name (name + suffix)
  file.Put (buf)
  file.Fin()
}

func main() {
  const lenName = 10
  scr.NewWH (80, 0, 1600, 1000); defer scr.Fin()
  files.Cd (env.Val ("GOSRC") + "/.epen")
  figures := seq.New (fig2.New())

  figures.Ins (fig2.NewRectangle (100, 100, 300, 200, true, col.Red()))
  figures.Ins (fig2.NewCircle (400, 200, 60, true, col.Green()))
  figures.Ins (fig2.NewEllipse(600, 300, 100, 50, true, col.LightBlue()))
  figures.Ins (fig2.NewPolygon([]int{100, 200, 200, 400, 300},
                               []int{400, 300, 400, 400, 500}, true, col.Yellow()))

  figures.Trav (func (a any) { a.(fig2.Figure2).Write() })
  bx := box.New()
  bx.Wd (lenName)
  bx.Colours (col.HeadF(), col.HeadB())
  name := env.Arg(1)
  if str.Empty (name) { name = "temp" }
  bx.Edit (&name, 0, 0)
  str.OffSpc (&name)
  store (figures, name)
}
