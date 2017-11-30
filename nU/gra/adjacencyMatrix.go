package gra

// (c) Christian Maurer   v. 171125 - license see nU.go

import (. "nU/obj"; "nU/adj")

func (x *graph) Matrix() adj.AdjacencyMatrix {
  n := x.Num()
  matrix := adj.New (n, x.vAnchor.Any, x.eAnchor.Any)
  for v, i := x.vAnchor.nextV, uint(0); v != x.vAnchor; v, i = v.nextV, i + 1 {
    v.dist = uint32(i) // i.e. v is the i-th vertex of x
  }
  for v, i := x.vAnchor.nextV, uint(0); v != x.vAnchor; v, i = v.nextV, i + 1 {
    matrix.Set (i, i, Clone(v.Any), x.eAnchor.Any)
    for n := v.nbPtr.nextNb; n != v.nbPtr; n = n.nextNb {
      if n.outgoing {
        matrix.Set (i, uint(n.to.dist), x.vAnchor.Any, n.edgePtr.Any)
      }
    }
  }
  return matrix
}

func (x *graph) SetMatrix (matrix adj.AdjacencyMatrix) {
  if x.bool == matrix.Symmetric() { panic ("gra.SetMatrix: x directed, matrix not") }
  m := matrix.Num()
  for i := uint(0); i < m; i++ {
    x.Ins (Clone(matrix.Vertex (i)))
  }
  for i := uint(0); i < m; i++ {
    for k := uint(0); k < m; k++ {
      if k != i {
        val := matrix.Val (i, k)
        if val > 0 &&
          x.ExPred2 (func (a Any) bool { return Eq (a, matrix.Vertex (i)) },
                     func (a Any) bool { return Eq (a, matrix.Vertex (k)) }) {
          x.Edge (val)
        }
      }
    }
  }
}
