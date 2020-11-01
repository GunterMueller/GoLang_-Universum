package img

// (c) Christian Maurer   v. 201014 - license see µU.go

import (
  "os/exec"
  . "µU/obj"
  "µU/time"
  "µU/str"
  "µU/scr"
  "µU/prt"
  "µU/errh"
  "µU/pseq"
)
const
  suffix = ".ppm"
var
  toPrint bool

func put (n string, x, y, w, h uint) {
  if str.Empty (n) { return }
  str.OffSpc (&n)
  filename := n + suffix
//  if scr.UnderX() { errh.Hint (errh.ToWait) }
  buf := scr.P6Encode (x, y, w, h)
  if scr.UnderX() { errh.DelHint() }
  file := pseq.New (buf)
  file.Name (filename)
  file.Clr()
  file.Put (buf)
  file.Fin()
  if ! toPrint {
    exec.Command ("pnmtopng", filename + suffix, ">", n, ".png").Run()
    time.Msleep (100)
    exec.Command ("rm", filename)
  }
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
  buf := make (Stream, l)
  file := pseq.New (buf)
  file.Name (filename)
  buf = file.Get().(Stream)
  file.Fin()
  return scr.P6Size (buf)
}

func get (n string, x, y uint) {
  if str.Empty (n) { return }
  if ! scr.UnderX() { x += scr.X(); y += scr.Y() }
  str.OffSpc (&n)
  filename := n + suffix
  l := pseq.Length (filename)
  if l == 0 { return }
  buf := make (Stream, l)
  file := pseq.New (buf)
  file.Name (filename)
  buf = file.Get().(Stream)
  file.Fin()
  scr.P6Decode (x, y, buf)
}

func print_(x, y, w, h uint) {
  toPrint = true
  filename := "tmp"
  put (filename, x, y, w, h)
  toPrint = false
//  exec.Command (prt.PrintCommand, "-o", "landscape", "-o", "fit-to-page", filename + suffix).Run()
  exec.Command (prt.PrintCommand, "-o", "fit-to-page", filename + suffix).Run()
}

func print1() {
  print_(0, 0, scr.Wd(), scr.Ht())
}
