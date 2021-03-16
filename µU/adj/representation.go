package adj

// (c) Christian Maurer   v. 210311 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/col"
)
type (
  pair struct {
       vertex,
         edge Any
              }
  adjacencyMatrix struct {
                         uint "number of rows/columns"
                       v,    // pattern vertex
                       e Any // pattern edge
                   entry [][]pair
                  cF, cB col.Colour
                         }
)

func new_(n uint, v, e Any) AdjacencyMatrix {
  if n == 0 || e == nil { ker.Oops() }
  CheckAtomicOrObject (v)
  CheckUintOrValuator (e)
  x := new(adjacencyMatrix)
  x.uint = n
  x.v, x.e = Clone(v), Clone(e)
  x.entry = make ([][]pair, n)
  for i := uint(0); i < n; i++ {
    x.entry[i] = make ([]pair, n)
    for k := uint(0); k < n; k++ {
      x.entry[i][k] = pair { x.v, x.e }
    }
  }
  x.cF, x.cB = col.StartCols()
  return x
}
