package gra

// (c) Christian Maurer   v. 171227 - license see nU.go

import . "nU/obj"

func (x *graph) Trav (o Op) {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    o (v.Any)
  }
}

func (x *graph) travCond (c CondOp) {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    c (v.Any, v.bool)
  }
}

func (x *graph) trav2Cond (c CondOp2) {
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    c (e.nbPtr0.from.Any, e.nbPtr1.from.Any, e.bool)
  }
}
