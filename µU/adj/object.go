package adj

// (c) Christian Maurer   v. 171112 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
)

func (x *adjacencyMatrix) imp (Y Any) *adjacencyMatrix {
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
  x.e, x.v = Clone(y.e), Clone(y.v)
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
      x.entry[i][k].vertex = Clone (y.entry[i][k].vertex)
      x.entry[i][k].edge = Clone (y.entry[i][k].edge)
    }
  }
}

func (x *adjacencyMatrix) Clone() Any {
  y := new_(x.uint, x.v, x.e)
  y.Copy (x)
  return y
}

func (x *adjacencyMatrix) Codelen() uint {
  v, e := Codelen(x.v), Codelen(x.e)
  return 4 + (1 + x.uint * x.uint) * (v + e)
}

func (x *adjacencyMatrix) Encode() []byte {
  bs := make ([]byte, x.Codelen())
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

func (x *adjacencyMatrix) Decode (bs []byte) {
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
