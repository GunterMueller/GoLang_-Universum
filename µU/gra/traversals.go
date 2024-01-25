package gra

// (c) Christian Maurer   v. 231225 - license see µU.go

import
  . "µU/obj"

func (x *graph) True (p Pred) bool {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if ! p (v.any) {
      return false
    }
  }
  return true
}

func (x *graph) TrueMarked (p Pred) bool {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if v.marked {
      if ! p (v.any) {
        return false
      }
    }
  }
  return true
}

func (x *graph) Trav (o Op) {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    o (v.any)
  }
}

func (x *graph) TravCond (o CondOp) {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    o (v.any, v.marked)
  }
}

func (x *graph) Trav1 (o Op) {
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    o (e.any)
  }
}

func (x *graph) Trav1Cond (o CondOp) {
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    o (e.any, e.marked)
  }
}

func (x *graph) Trav1Loc (o Op) {
  for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
    o (n.edgePtr.any)
  }
}

func (x *graph) Trav1Coloc (o Op) {
  for n := x.colocal.nbPtr.nextNb; n != x.colocal.nbPtr; n = n.nextNb {
    o (n.edgePtr.any)
  }
}
