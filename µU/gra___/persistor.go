package gra

// (c) Christian Maurer   v. 220204 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/str"
  "µU/pseq"
)

func (x *graph) Name (n string) {
  if str.Empty (n) { ker.Panic ("name is empty") }
  str.OffSpc (&n)
  x.name = n
  x.filename = x.name + "." + suffix
}

func (x *graph) Rename (n string) {
  if str.Empty (n) { return }
  str.OffSpc (&n)
  x.Rename (n)
}

func (x *graph) File() pseq.PersistentSequence {
  return x.file
}

func (x *graph) Load() {
  n := pseq.Length (x.filename)
  if n == 0 { return }
  buf := make(Stream, n)
  x.file = pseq.New (buf)
  x.file.Name (x.filename)
  buf = x.file.Get().(Stream)
  x.Decode (buf)
}

func (x *graph) Store() {
  buf := x.Encode()
  x.file = pseq.New (buf)
  x.file.Name (x.filename)
  x.file.Clr()
  x.file.Put (buf)
}

func (x *graph) Fin() {
  x.file.Fin()
}
