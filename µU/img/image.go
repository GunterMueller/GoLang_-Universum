package img

// (c) Christian Maurer   v. 201228 - license see µU.go

import (
  "os/exec"
  . "µU/obj"
  "µU/str"
  "µU/scr"
  "µU/prt"
  "µU/errh"
  "µU/pseq"
)
const
  suffix = ".ppm"

func put (n string, x, y, w, h uint) {
  if str.Empty (n) { return }
  str.OffSpc (&n)
  filename := n + suffix
  if scr.UnderX() { errh.Hint (errh.ToWait) }
  s := scr.P6Encode (x, y, w, h)
  if scr.UnderX() { errh.DelHint() }
  file := pseq.New (s)
  file.Name (filename)
  file.Clr()
  file.Put (s)
  file.Fin()
}

func put1 (n string) {
  put (n, scr.X(), scr.Y(), scr.Wd(), scr.Ht())
}

func size_(n string) (uint, uint) {
  if str.Empty (n) { return 0, 0 }
  str.OffSpc (&n)
  filename := n + suffix
  l := pseq.Length (filename)
  if l == 0 { return 0, 0 }
  s := make (Stream, l)
  file := pseq.New (s)
  file.Name (filename)
  s = file.Get().(Stream)
  file.Fin()
  return scr.P6Size (s)
}

func get (n string, x, y uint) {
  if str.Empty (n) { return }
  if ! scr.UnderX() { x += scr.X(); y += scr.Y() }
  str.OffSpc (&n)
  filename := n + suffix
  l := pseq.Length (filename)
  if l == 0 { return }
  s := make (Stream, l)
  file := pseq.New (s)
  file.Name (filename)
  s = file.Get().(Stream)
  file.Fin()
  scr.P6Decode (s, x, y)
}

func print_(n string, x, y, w, h uint) {
  put (n, x, y, w, h)
  exec.Command (prt.PrintCommand, "-o", "fit-to-page", n + suffix).Run()
}

func print1 (n string) {
  print_(n, 0, 0, scr.Wd(), scr.Ht())
}
