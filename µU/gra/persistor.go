package gra

// (c) Christian Maurer   v. 201011 - license see µU.go

import (
  . "µU/obj"
  "µU/str"
  "µU/pseq"
)

func (x *graph) Name (s string) {
  x.name = s
  str.OffSpc (&x.name)
  if str.Empty (x.name) { x.name = "tmp" } // TODO + pid
  x.filename = x.name + "." + suffix
  n := pseq.Length (x.filename)
  if n > 0 {
    buf := make(Stream, n)
    f := pseq.New (buf, false)
    f.Name (x.filename)
    buf = f.Get().(Stream)
    f.Fin()
    x.Decode (buf)
  }
}

func (x *graph) Rename (s string) {
  x.name = s
  str.OffSpc (&x.name)
  x.filename = x.name + "." + suffix
// rest of implementation TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO
  n := pseq.Length (x.filename)
  if n > 0 {
    buf := make(Stream, n)
    f := pseq.New (buf, false)
    f.Rename (x.name)
//    buf = f.Get().(Stream)
    f.Fin()
//    x.Decode (buf)
  }
}

func (x *graph) Fin() {
  if ! str.Empty (x.name) {
    buf := x.Encode()
    f := pseq.New (buf, false)
    f.Name (x.filename)
    f.Clr()
    f.Put (buf)
    f.Fin()
  }
//  x.Clr()
}
