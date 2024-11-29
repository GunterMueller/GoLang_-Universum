package adj

// (c) Christian Maurer   v. 241014 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/col"
  "µU/scr"
)
type (
  pair struct {
       vertex,
         edge any
              }
  adjacencyMatrix struct {
                         uint "number of rows/columns"
                       v,    // pattern vertex
                       e any // pattern edge
                   entry [][]pair
                  cF, cB col.Colour
                         }
)

func new_(n uint, v, e any) AdjacencyMatrix {
  if n == 0 || e == nil { ker.PrePanic() }
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

func (x *adjacencyMatrix) imp (Y any) *adjacencyMatrix {
  y, ok := Y.(*adjacencyMatrix)
  if ! ok { TypeNotEqPanic (x, Y) }
  if x.uint != y.uint { ker.Panic ("adj.imp: different size") }
  CheckTypeEq (x.e, y.e)
  CheckTypeEq (x.v, y.v)
  return y
}

func (x *adjacencyMatrix) Empty() bool {
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
      if ! Eq (x.entry[i][k].edge, x.e) {
        return false
      }
    }
  }
  return true
}

func (x *adjacencyMatrix) Clr() {
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
      x.entry[i][k] = pair { x.v, x.e }
    }
  }
}

func (x *adjacencyMatrix) Eq (Y any) bool {
  y := x.imp (Y)
  if x.Empty() { return y.Empty() }
  for i := uint(0); i < x.uint; i++ {
    if ! Eq (x.Vertex(i), y.Vertex(i)) { return false }
    for k := uint(0); k < x.uint; k++ {
      if ! Eq (x.entry[i][k].edge, y.entry[i][k].edge) ||
         ! Eq (x.entry[i][k].vertex, y.entry[i][k].vertex) {
//      if i != k && x.Val(i,k) != y.Val(i,k) {
        return false
      }
    }
  }
  return true
}

func (x *adjacencyMatrix) Less (Y any) bool {
  return false
}

func (x *adjacencyMatrix) Leq (Y any) bool {
  return false
}

func (x *adjacencyMatrix) Copy (Y any) {
  y := x.imp (Y)
  x.uint = y.uint
  x.e, x.v = Clone(y.e), Clone(y.v)
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
      x.entry[i][k].vertex = Clone (y.entry[i][k].vertex)
      x.entry[i][k].edge = Clone (y.entry[i][k].edge)
    }
  }
}

func (x *adjacencyMatrix) Clone() any {
  y := new_(x.uint, x.v, x.e)
  y.Copy (x)
  return y
}

func (x *adjacencyMatrix) Codelen() uint {
  v, e := Codelen(x.v), Codelen(x.e)
  return 4 + (1 + x.uint * x.uint) * (v + e)
}

func (x *adjacencyMatrix) Encode() Stream {
  bs := make (Stream, x.Codelen())
  v, e := Codelen(x.v), Codelen(x.e)
  copy (bs[:4], Encode (uint32(x.uint)))
  i := uint(4)
  copy (bs[i:i+v], Encode (x.v))
  i += v
  copy (bs[i:i+e], Encode (x.e))
  i += e
  for j := uint(0); j < x.uint; j++ {
    for k := uint(0); k < x.uint; k++ {
      copy (bs[i:i+v], Encode (x.entry[j][k].vertex))
      i += v
      copy (bs[i:i+e], Encode (x.entry[j][k].edge))
      i += e
    }
  }
  return bs
}

func (x *adjacencyMatrix) Decode (bs Stream) {
  v, e := Codelen(x.v), Codelen(x.e)
  x.uint = uint(Decode (uint32(0), bs[:4]).(uint32))
  i := uint(4)
  x.v = Decode (x.v, bs[i:i+v])
  i += v
  x.e = Decode (x.e, bs[i:i+e])
  i += e
  for j := uint(0); j < x.uint; j++ {
    for k := uint(0); k < x.uint; k++ {
      x.entry[j][k].vertex = Decode (x.v, bs[i:i+v])
      i += v
      x.entry[j][k].edge = Decode (x.e, bs[i:i+e])
      i += e
    }
  }
}

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
  if ! x.Equiv (Y) { ker.Panic("cannot Add") }
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
