package adj

// (c) Christian Maurer   v. 231110 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/col"
  "µU/vtx"
  "µU/edg"
)
type (
  pair struct {
       vertex vtx.Vertex
         edge edg.Edge
              }
  adjacencyMatrix struct {
                         uint "number of rows/columns"
                       v vtx.Vertex // pattern vertex
                       e edg.Edge // pattern edge
                   entry [][]pair
                  cF, cB col.Colour
                         }
)

func new_(n uint, v vtx.Vertex, e edg.Edge) AdjacencyMatrix {
  if n == 0 || e == nil { ker.Oops() }
  CheckAtomicOrObject (v)
  CheckUintOrValuator (e)
  x := new(adjacencyMatrix)
  x.uint = n
  x.v, x.e = v.Clone().(vtx.Vertex), e.Clone().(edg.Edge)
  x.entry = make ([][]pair, n)
  for i := uint(0); i < n; i++ {
    x.entry[i] = make ([]pair, n)
    for k := uint(0); k < n; k++ {
      x.entry[i][k] = pair {x.v, x.e}
    }
  }
  x.cF, x.cB = col.StartCols()
  return x
}
