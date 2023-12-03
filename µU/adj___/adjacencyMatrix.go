package adj

// (c) Christian Maurer   v. 231110 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/col"
  "µU/scr"
  "µU/vtx"
  "µU/edg"
)

func (x *adjacencyMatrix) Num() uint {
  return x.uint
}

func (x *adjacencyMatrix) Equiv (Y AdjacencyMatrix) bool {
  y := x.imp(Y)
  if x.uint != y.uint { return false }
  if ! Eq (x.v, y.v) || ! Eq (x.e, y.e) { return false }
  return true
}

func (x *adjacencyMatrix) Edge (i, k uint, e edg.Edge) {
  if i >= x.uint || k >= x.uint { return }
  CheckTypeEq (e, x.e)
  x.entry[i][k].edge = e.Clone().(edg.Edge)
}

func (x *adjacencyMatrix) Vertex (i uint) vtx.Vertex {
  return x.entry[i][i].vertex.Clone().(vtx.Vertex)
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

func (x *adjacencyMatrix) Set (i, k uint, v vtx.Vertex, e edg.Edge) {
  if i >= x.uint || k >= x.uint { return }
  if i == k {
    x.entry[i][i].vertex = v.Clone().(vtx.Vertex)
    x.entry[i][i].edge = x.e.Clone().(edg.Edge) // no loops
  } else {
    x.entry[i][k].vertex = x.v.Clone().(vtx.Vertex) // no vertex
    x.entry[i][k].edge = e.Clone().(edg.Edge)
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
  if ! x.Equiv (Y) { ker.Panic("cannot Add") }
  y := x.imp(Y)
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
      if i != k {
        if x.Val (i, k) == 0 && y.Val (i, k ) > 0 {
          x.entry[i][k].edge = y.entry[i][k].edge.Clone().(edg.Edge)
        }
        if Eq (x.entry[i][i].vertex, x.v) {
          x.entry[i][i].vertex = y.entry[i][i].vertex.Clone().(vtx.Vertex)
          x.entry[i][i].vertex = y.entry[i][i].vertex.Clone().(vtx.Vertex)
        }
        if Eq (x.entry[k][k].vertex, x.v) {
          x.entry[k][k].vertex = y.entry[k][k].vertex.Clone().(vtx.Vertex)
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

func (x *adjacencyMatrix) Colours (f, g col.Colour) {
  x.cF, x.cB = f, g
}

func (x *adjacencyMatrix) Cols() (col.Colour, col.Colour) {
  return x.cF, x.cB
}

func (x *adjacencyMatrix) Write (l, c uint) {
  scr.Colours (x.cF, x.cB)
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
      val := Val (x.entry[i][k].edge)
      if val > 0 {
        scr.WriteNat (val, l + i, c + 2 * k)
      } else if i == k {
        scr.Write ("*", l + i, c + 2 * k)
      } else {
        scr.Write (".", l + i, c + 2 * k)
      }
    }
  }
}
