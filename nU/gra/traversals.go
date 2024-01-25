package gra

// (c) Christian Maurer   v. 231226 - license see nU.go

import
  . "nU/obj"

func (x *graph) Trav (o Op) {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    o (v.any)
  }
}

func (x *graph) TravCond (c CondOp) {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    c (v.any, v.marked)
  }
}

func (x *graph) Trav1Cond (c CondOp) {
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    c (e.any, e.marked)
  }
}

func (x *graph) Trav2Cond (c CondOp2) {
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    c (e.nbPtr0.from.any, e.nbPtr1.from.any, e.marked)
  }
}
