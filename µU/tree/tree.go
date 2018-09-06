package tree

// (c) Christian Maurer   v. 180817 - license see µU.go

import
  . "µU/obj"
type
  tree struct {
              Any "pattern object"
         root,
       father Any
          son []Tree
}

func new_(a Any) Tree {
  x := new(tree)
  x.Any = Clone(a)
  x.son = make([]Tree, 0)
  return x
}

func (x *tree) imp (Y Any) tree {
  y, ok := Y.(tree)
  if ! ok { panic("XXX") }
  if ! Eq (x.Any, y.Any) { panic("YYY") }
  return y
}

func (x *tree) Eq (Y Any) bool {
  y := x.imp(Y)
  if x.Empty() && y.Empty() { return true }
  if ! Eq (x.root, y.root) { return false }
  for i, s := range x.son {
    if ! s.Eq (y.son[i]) {
      return false
    }
  }
  return true
}

func (x *tree) Copy (Y Any) {
  y := x.imp(Y)
  if y.Empty() {
    x.Clr()
    return
  }
  x.root = y.root
  n := len(y.son)
  x.son = make([]Tree, n)
  for i, _ := range y.son {
    x.son[i].Copy (y.son[i])
  }
}

func (x *tree) Clone() Any {
  y := new_(x.Any)
  y.Copy (x)
  return y
}

func (x *tree) Less (Y Any) bool {
  return false
}

func (x *tree) Empty() bool {
  return x.root == nil
}

func (x *tree) Clr() {
  x.root = nil
  x.son = make([]Tree, 0)
}

func (x *tree) Codelen() uint {
// TODO
  return 0
}

func (x *tree) Encode() Stream {
// TODO
  s := make(Stream, x.Codelen())
  if x.Empty() {

    return s
  }

  return s
}

func (x *tree) Decode (b Stream) {
// TODO

}

func (x *tree) Root() Any {
  return x.root
}

func (x *tree) Father() Any {
  return x.father
}

func (x *tree) Son (i uint) Tree {
  return x.son[i]
}

func (x *tree) Depth() uint {
  if x.Empty() {
    return 0
  }
  n := len(x.son)
  if n == 0 {
    return 1
  }
  d := make([]uint, n)
  for i := 0; i < n; i++ {
    d[i] = x.son[i].(*tree).Depth()
  }
  dp := uint(0)
  for i := 0; i < n; i++ {
    if d[i] > dp {
      dp = d[i]
    }
  }
  return dp
}
