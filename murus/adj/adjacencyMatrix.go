package adj

// (c) murus.org  v. 170810 - license see murus.go

import (
  "murus/ker"
  . "murus/obj"
  "murus/col"
  "murus/scr"
)
type
  entry struct {
               bool
               Any // QuasiValuator
               }
type
  adjacencyMatrix struct {
                         uint "number of rows/columns"
                         Any "empty edge"
                    edge [][]entry
                  cF, cB col.Colour
                         }
var
  char [2]byte = [2]byte { '.', '*' }

func new_(n uint, e Any) AdjacencyMatrix {
  if n == 0 || e == nil { ker.Oops() }
  switch e.(type) {
  case Valuator, uint8, uint16, uint, uint32:
    // ok
  default:
    ker.Oops()
  }
  x := new (adjacencyMatrix)
  x.uint, x.Any = n, Clone(e)
  x.edge = make ([][]entry, n)
  for i := uint(0); i < n; i++ {
    x.edge[i] = make ([]entry, n)
    for k := uint(0); k < n; k++ {
      x.edge[i][k] = entry { false, Clone(x.Any) }
    }
  }
  x.cF, x.cB = scr.StartCols()
  return x
}

func (x *adjacencyMatrix) imp (Y Any) *adjacencyMatrix {
  y, ok := Y.(*adjacencyMatrix)
  if ! ok { TypeNotEqPanic (x, Y) }
  if x.uint != y.uint { ker.Panic ("adj.imp: different size") }
  CheckTypeEq (x.Any, y.Any)
  return y
}

func (x *adjacencyMatrix) Empty() bool {
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
      if x.edge[i][k].bool {
        return false
      }
    }
  }
  return x.Any == nil
}

func (x *adjacencyMatrix) Clr() {
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
      x.edge[i][k] = entry { false, Clone(x.Any) }
    }
  }
}

func (x *adjacencyMatrix) Eq (Y Any) bool {
  y := x.imp (Y)
  if x.Empty() { return y.Empty() }
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
      if x.edge[i][k].bool != y.edge[i][k].bool ||
        ! Eq (x.edge[i][k].Any, y.edge[i][k].Any) {
          return false
      }
    }
  }
  return true
}

func (x *adjacencyMatrix) Less (Y Any) bool {
  y := x.imp (Y)
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
      if x.edge[i][k].bool {
        if y.edge[i][k].bool {
          if ! Eq (x.edge[i][k].Any, y.edge[i][k].Any) {
            return false
          }
        } else {
          return false
        }
      }
    }
  }
  return true
}

func (x *adjacencyMatrix) Copy (Y Any) {
  y := x.imp (Y)
  x.uint, x.Any = y.uint, Clone(y.Any)
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
      x.edge[i][k].bool = y.edge[i][k].bool
      x.edge[i][k].Any = Clone (y.edge[i][k].Any)
    }
  }
}

func (x *adjacencyMatrix) Clone() Any {
  y := new_(x.uint, x.Any)
  y.Copy (x)
  return y
}

func (x *adjacencyMatrix) Codelen() uint {
  q, c := x.uint * x.uint, Codelen(x.Any)
  return 4 + q * (1 + c)
}

func (x *adjacencyMatrix) Encode() []byte {
  bs := make ([]byte, x.Codelen())
  copy (bs[:4], Encode (uint32(x.uint)))
  i, c := uint(4), Codelen(x.Any)
  for j := uint(0); j < x.uint; j++ {
    for k := uint(0); k < x.uint; k++ {
      bs[i] = 0; if x.edge[j][k].bool { bs[i] = 1 }
      i++
      copy (bs[i:i+c], Encode (x.edge[j][k].Any))
      i += c
    }
  }
  return bs
}

func (x *adjacencyMatrix) Decode (bs []byte) {
  x.uint = uint(Decode (uint32(0), bs[:4]).(uint32))
  i, c := uint(4), Codelen(x.Any)
  for j := uint(0); j < x.uint; j++ {
    for k := uint(0); k < x.uint; k++ {
      x.edge[j][k].bool = bs[i] == 1
      i++
      x.edge[j][k].Any = Decode (x.Any, bs[i:i+c])
      i += c
    }
  }
}

func (x *adjacencyMatrix) Colours (f, g col.Colour) {
  x.cF, x.cB = f, g
}

func (x *adjacencyMatrix) Write (l, c uint) {
  scr.Colours (x.cF, x.cB)
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
//      b := 1; if Eq (x.edge[i][k].Any, x.Any) { b = 0 }
      b := 0; if x.edge[i][k].bool { b = 1 }
      scr.Write1 (char[b], l + i, c + 2 * k)
    }
  }
}

func (x *adjacencyMatrix) Ok() bool {
  for i := uint(0); i < x.uint; i++ {
    if ! Eq (x.edge[i][i], x.Any) {
      return false
    }
  }
  return true
}

func (x *adjacencyMatrix) Loop() uint {
  for i := uint(0); i < x.uint; i++ {
    if Eq (x.edge[i][i], x.Any) {
      return i
    }
  }
  return x.uint
}

func (x *adjacencyMatrix) Num() uint {
  return x.uint
}

func (x *adjacencyMatrix) Val (i, k uint) uint {
  if i < x.uint && k < x.uint {
    return Val(x.edge[i][k].Any)
  }
  return 0
}

func (x *adjacencyMatrix) Edged (i, k uint) bool {
  if i < x.uint && k < x.uint {
    return x.edge[i][k].bool
  }
  return false
}

func (x *adjacencyMatrix) Symmetric() bool {
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
      if x.edge[i][k].bool != x.edge[k][i].bool {
        return false
      }
      if ! Eq (x.edge[i][k].Any, x.edge[k][i].Any) {
        return false
      }
    }
  }
  return true
}

func (x *adjacencyMatrix) Directed() bool {
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
      if i != k || x.edge[i][k].bool == x.edge[k][i].bool {
        return false
      }
    }
  }
  return true
}

func (x *adjacencyMatrix) Equiv (Y AdjacencyMatrix) bool {
  y := x.imp(Y)
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
      if x.edge[i][k].bool && y.edge[i][k].bool {
        if ! Eq (x.edge[i][k].Any, y.edge[i][k].Any) {
          return false
        }
      }
    }
  }
  return true
}

func (x *adjacencyMatrix) Add (Y AdjacencyMatrix) {
//  if ! x.Equiv (Y) { ker.Panic("cannot Add") } // XXX
  y := x.imp(Y)
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
      if ! x.edge[i][k].bool && y.edge[i][k].bool {
        x.edge[i][k].bool = true
        x.edge[i][k].Any = Clone (y.edge[i][k].Any)
      }
    }
  }
}

func (x *adjacencyMatrix) Invert() {
  for i := uint(0); i < x.uint; i++ {
    for k := uint(0); k < x.uint; k++ {
      if i != k {
        x.edge[i][k], x.edge[k][i] = x.edge[k][i], x.edge[i][k]
      }
    }
  }
}

func (x *adjacencyMatrix) Edge (i, k uint) {
  if i >= x.uint || k >= x.uint { return }
  x.edge[i][k].bool = true
  switch x.Any.(type) {
  case Valuator:
    x.edge[i][k].Any.(Valuator).SetVal(1)
  case uint8, uint16, uint32, uint, uint64:
    x.edge[i][k].Any = 1
  }
}

func (x *adjacencyMatrix) Set (i, k uint, v Any) {
  CheckTypeEq (v, x.Any)
  if i >= x.uint || k >= x.uint { return }
  x.edge[i][k].bool = true
  switch x.Any.(type) {
  case Valuator:
    x.edge[i][k].Any = Clone(v)
  case uint8:
    x.edge[i][k].Any = v.(uint8)
  case uint16:
    x.edge[i][k].Any = v.(uint16)
  case uint32:
    x.edge[i][k].Any = v.(uint32)
  case uint:
    x.edge[i][k].Any = v.(uint)
  case uint64:
    x.edge[i][k].Any = v.(uint64)
  }
}

func (x *adjacencyMatrix) Del (i, k uint) {
  if i >= x.uint || k >= x.uint { return }
  x.Set (i, k, 0)
  x.edge[i][k].bool = false
}

func (x *adjacencyMatrix) Full() bool{
  for i := uint(0); i < x.uint; i++ {
    f := false
    for k := uint(0); k < x.uint; k++ {
      f = f || x.edge[i][k].bool
    }
    if ! f { return false }
  }
  return true
}
