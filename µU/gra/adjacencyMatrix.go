package gra

// (c) Christian Maurer   v. 231225 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/adj"
)

func (x *graph) Matrix() adj.AdjacencyMatrix {
  n := x.Num()
  matrix := adj.New (n, x.vAnchor.any, x.eAnchor.any)
  for v, i := x.vAnchor.nextV, uint(0); v != x.vAnchor; v, i = v.nextV, i + 1 {
    v.time0 = uint32(i) // i.e. v is the i-th vertex of x
  }
  for v, i := x.vAnchor.nextV, uint(0); v != x.vAnchor; v, i = v.nextV, i + 1 {
    matrix.Set (i, i, Clone(v.any), x.eAnchor.any)
    for n := v.nbPtr.nextNb; n != v.nbPtr; n = n.nextNb {
      if n.outgoing {
        matrix.Set (i, uint(n.to.time0), x.vAnchor.any, n.edgePtr.any)
      }
    }
  }
  return matrix
}

func (x *graph) SetMatrix (matrix adj.AdjacencyMatrix) {
  if x.directed == matrix.Symmetric() { ker.Panic ("gra.SetMatrix: x directed, matrix not") }
  m := matrix.Num()
  for i := uint(0); i < m; i++ {
    x.Ins (Clone(matrix.Vertex (i)))
  }
  for i := uint(0); i < m; i++ {
    for k := uint(0); k < m; k++ {
      if k != i {
        val := matrix.Val (i, k)
        if val > 0 && x.ExPred2 (func (a any) bool { return Eq (a, matrix.Vertex (i)) },
                                 func (a any) bool { return Eq (a, matrix.Vertex (k)) }) {
          x.Edge (val)
        }
      }
    }
  }
}
