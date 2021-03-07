package ppm

// (c) Christian Maurer   v. 210228 - license see µU.go

import (
  "os/exec"
  "µU/ker"
  . "µU/obj"
  "µU/str"
  "µU/scr"
  "µU/prt"
  "µU/pseq"
)
const
  suffix = ".ppm"

func put (n string) {
  if str.Empty (n) { return }
  str.OffSpc (&n)
  filename := n + suffix
  s := scr.PPMEncode (0, 0, scr.Wd(), scr.Ht())
  file := pseq.New (s)
  file.Name (filename)
  file.Clr()
  file.Put (s)
  file.Fin()
}

func size_(n string) (uint, uint) {
  if str.Empty (n) { return 0, 0 }
  str.OffSpc (&n)
  filename := n + suffix
  l := pseq.Length (filename)
  if l == 0 { return 0, 0 }
  s := make(Stream, l)
  file := pseq.New (s)
  file.Name (filename)
  s = file.Get().(Stream)
  file.Fin()
  return scr.PPMSize (s)
}

func get (n string, x, y uint) {
  if str.Empty (n) { return }
  str.OffSpc (&n)
  filename := n + suffix
  l := pseq.Length (filename)
  if l == 0 { return }
  s := make(Stream, l)
  file := pseq.New (s)
  file.Name (filename)
  s = file.Get().(Stream)
  scr.PPMDecode (s, x, y)
  file.Fin()
}

func print_(n string) {
  put (n)
  e := exec.Command (prt.PrintCommand, "-o", "fit-to-page", n + suffix).Run()
  if e != nil { ker.Panic (e.Error()) }
}
