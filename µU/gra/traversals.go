package gra

// (c) Christian Maurer   v. 171112 - license see µU.go

import
  . "µU/obj"

func (x *graph) True (p Pred) bool {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if ! p (v.Any) {
      return false
    }
  }
  return true
}

func (x *graph) TrueSub (p Pred) bool {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if v.bool {
      if ! p (v.Any) {
        return false
      }
    }
  }
  return true
}

func (x *graph) Trav (o Op) {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    o (v.Any)
  }
}

func (x *graph) TravCond (o CondOp) {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    o (v.Any, v.bool)
  }
}

func (x *graph) Trav1 (o Op) {
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    o (e.Any)
  }
}

func (x *graph) Trav1Cond (o CondOp) {
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    o (e.Any, e.bool)
  }
}

func (x *graph) Trav1Loc (o Op) {
  for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
    o (n.edgePtr.Any)
  }
}

func (x *graph) Trav1Coloc (o Op) {
  for n := x.colocal.nbPtr.nextNb; n != x.colocal.nbPtr; n = n.nextNb {
    o (n.edgePtr.Any)
  }
}
