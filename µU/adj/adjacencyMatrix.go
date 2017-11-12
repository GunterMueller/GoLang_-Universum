package adj

// (c) Christian Maurer   v. 171112 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/col"
  "µU/scr"
)
type (
  pair struct {
       vertex,
         edge Any
              }
  adjacencyMatrix struct {
                         uint "number of rows/columns"
                     vvv,    // pattern vertex
                     eee Any // pattern edge
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
  x.vvv, x.eee = Clone(v), Clone(e)
  x.entry = make ([][]pair, n)
  for i := uint(0); i < n; i++ {
    x.entry[i] = make ([]pair, n)
    for k := uint(0); k < n; k++ {
      x.entry[i][k] = pair { x.vvv, x.eee }
    }
  }
  x.cF, x.cB = scr.StartCols()
  return x
}

func (x *adjacencyMatrix) imp (Y Any) *adjacencyMatrix {
  y, ok := Y.(*adjacencyMatrix)
  if ! ok { TypeNotEqPanic (x, Y) }
  if x.uint != y.uint { ker.Panic ("adj.imp: different size") }
  CheckTypeEq (x.eee, y.eee)
  CheckTypeEq (x.vvv, y.vvv)
  return y
}

func (x *adjacencyMatrix) Empty() bool {
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
      if ! Eq (x.entry[i][k].edge, x.eee) {
        return false
      }
    }
  }
  return true
}

func (x *adjacencyMatrix) Clr() {
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
      x.entry[i][k] = pair { x.vvv, x.eee }
    }
  }
}

func (x *adjacencyMatrix) Eq (Y Any) bool {
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

func (x *adjacencyMatrix) Less (Y Any) bool {
  return false
}

func (x *adjacencyMatrix) Copy (Y Any) {
  y := x.imp (Y)
  x.uint = y.uint
  x.eee, x.vvv = Clone(y.eee), Clone(y.vvv)
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
      x.entry[i][k].vertex = Clone (y.entry[i][k].vertex)
      x.entry[i][k].edge = Clone (y.entry[i][k].edge)
    }
  }
}

func (x *adjacencyMatrix) Clone() Any {
  y := new_(x.uint, x.vvv, x.eee)
  y.Copy (x)
  return y
}

func (x *adjacencyMatrix) Codelen() uint {
  v, e := Codelen(x.vvv), Codelen(x.eee)
  return 4 + (1 + x.uint * x.uint) * (v + e)
}

func (x *adjacencyMatrix) Encode() []byte {
  bs := make ([]byte, x.Codelen())
  v, e := Codelen(x.vvv), Codelen(x.eee)
  copy (bs[:4], Encode (uint32(x.uint)))
  i := uint(4)
  copy (bs[i:i+v], Encode (x.vvv))
  i += v
  copy (bs[i:i+e], Encode (x.eee))
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

func (x *adjacencyMatrix) Decode (bs []byte) {
  v, e := Codelen(x.vvv), Codelen(x.eee)
  x.uint = uint(Decode (uint32(0), bs[:4]).(uint32))
  i := uint(4)
  x.vvv = Decode (x.vvv, bs[i:i+v])
  i += v
  x.eee = Decode (x.eee, bs[i:i+e])
  i += e
  for j := uint(0); j < x.uint; j++ {
    for k := uint(0); k < x.uint; k++ {
      x.entry[j][k].vertex = Decode (x.vvv, bs[i:i+v])
      i += v
      x.entry[j][k].edge = Decode (x.eee, bs[i:i+e])
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
  if ! Eq(x.vvv, y.vvv) || ! Eq(x.eee, y.eee) { return false }
  return true
}

func (x *adjacencyMatrix) Edge (i, k uint, e Any) {
  if i >= x.uint || k >= x.uint { return }
  CheckTypeEq (e, x.eee)
  x.entry[i][k].edge = Clone(e)
}

func (x *adjacencyMatrix) Vertex (i uint) Any {
  return Clone(x.entry[i][i].vertex)
}

func (x *adjacencyMatrix) Val (i, k uint) uint {
  if i >= x.uint || k >= x.uint {
    return 0
  }
  if Eq (x.entry[i][k].edge, x.eee) {
    return 0
  }
  return Val(x.entry[i][k].edge)
}

func (x *adjacencyMatrix) Set (i, k uint, v, e Any) {
  if i >= x.uint || k >= x.uint { return }
  CheckTypeEq (v, x.vvv)
  CheckTypeEq (e, x.eee)
  if i == k {
    x.entry[i][k].edge = Clone(x.eee) // no loops
    x.entry[i][k].vertex = Clone(v)
  } else {
    x.entry[i][k].edge = Clone(e)
    x.entry[i][k].vertex = Clone(x.vvv) // no vertex
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
//          x.entry[i][i].vertex = Clone (y.entry[i][i].vertex)
          x.entry[i][k].edge = Clone (y.entry[i][k].edge)
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
