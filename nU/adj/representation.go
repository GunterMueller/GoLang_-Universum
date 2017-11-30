package adj

// (c) Christian Maurer   v. 171125 - license see nU.go

import . "nU/obj"

type (
  pair struct {
       vertex,
         edge Any
              }
  adjacencyMatrix struct {
                         uint "number of rows/columns"
                    v, e Any // pattern vertex and edge
                   entry [][]pair
                         }
)

func new_(n uint, v, e Any) AdjacencyMatrix {
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
  return x
}
