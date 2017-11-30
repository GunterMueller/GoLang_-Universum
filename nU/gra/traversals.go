package gra

// (c) Christian Maurer   v. 171122 - license see nU.go

import . "nU/obj"

func (x *graph) Trav (o Op) {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    o (v.Any)
  }
}
