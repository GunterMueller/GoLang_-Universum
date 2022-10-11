package adj

// (c) Christian Maurer   v. 220702 - license see nU.go

import (. "nU/obj"; "nU/col"; "nU/scr")

func (x *adjacencyMatrix) Num() uint {
  return x.uint
}

func (x *adjacencyMatrix) Equiv (Y AdjacencyMatrix) bool {
  y := x.imp(Y)
  if x.uint != y.uint { return false }
  if ! Eq(x.v, y.v) || ! Eq(x.e, y.e) { return false }
  return true
}

func (x *adjacencyMatrix) Edge (i, k uint, e any) {
  if i >= x.uint || k >= x.uint { return }
  CheckTypeEq (e, x.e)
  x.entry[i][k].edge = Clone(e)
}

func (x *adjacencyMatrix) Vertex (i uint) any {
  return Clone(x.entry[i][i].vertex)
}

func (x *adjacencyMatrix) Val (i, k uint) uint {
  if i >= x.uint || k >= x.uint {
    return 0
  }
  if Eq (x.entry[i][k].edge, x.e) {
    return 0
  }
  return Val(x.entry[i][k].edge)
}

func (x *adjacencyMatrix) Set (i, k uint, v, e any) {
  if i >= x.uint || k >= x.uint { return }
  CheckTypeEq (v, x.v)
  CheckTypeEq (e, x.e)
  if i == k {
    x.entry[i][i].vertex = Clone(v)
    x.entry[i][i].edge = Clone(x.e) // no loops
  } else {
    x.entry[i][k].vertex = Clone(x.v) // no vertex
    x.entry[i][k].edge = Clone(e)
  }
}

func (x *adjacencyMatrix) Symmetric() bool {
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
      if Val(x.entry[i][k].edge) != Val(x.entry[k][i].edge) {
        return false
      }
    }
  }
  return true
}

func (x *adjacencyMatrix) Add (Y AdjacencyMatrix) {
  y := x.imp(Y)
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
      if i != k {
        if x.Val (i, k) == 0 && y.Val (i, k ) > 0 {
          x.entry[i][k].edge = Clone (y.entry[i][k].edge)
        }
        if Eq (x.entry[i][i].vertex, x.v) {
          x.entry[i][i].vertex = Clone (y.entry[i][i].vertex)
        }
        if Eq (x.entry[k][k].vertex, x.v) {
          x.entry[k][k].vertex = Clone (y.entry[k][k].vertex)
        }
      }
    }
  }
}

func (x *adjacencyMatrix) Full() bool {
  for i := uint(0); i < x.uint; i++ {
    full := false
    for k := uint(0); k < x.uint; k++ {
      full = full || x.Val (i, k) > 0
    }
    if ! full {
      return false
    }
  }
  return true
}

func (x *adjacencyMatrix) Write() {
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
      if i == k {
        scr.ColourF (col.Magenta())
      } else {
        scr.ColourF (col.Yellow())
      }
      scr.WriteNat (Val(x.entry[i][k].edge), i, 2 * k)
    }
  }
}
