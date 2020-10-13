package btree

// (c) Christian Maurer   v. 201003 - license see µU.go

import (
  . "µU/obj"
  "µU/scr"
  "µU/str"
  "µU/errh"
)
type
  balance byte; const (
  leftweighty = balance(iota)
  balanced
  rightweighty
)
type
  btree struct {
         root Any
         left,
        right *btree
              balance
              }

func new_(a Any) BTree {
  x := new (btree)
  x.root = Clone(a)
  x.left = nil
  x.right = nil
  return x
}

func (x *btree) Root() Any {
  if x == nil { return nil }
  return x.root
}

func (x *btree) Left() BTree {
  if x == nil { return nil }
  return x.left
}

func (x *btree) Right() BTree {
  if x == nil { return nil }
  return x.right
}

func (x *btree) Num() uint {
  if x == nil {
    return uint(0)
  }
  return x.left.Num() + 1 + x.right.Num()
}

func (x *btree) NumPred (p Pred) uint {
  n := uint(0)
  if x == nil { return n }
  if p (x.root) { n ++ }
  return n + x.left.NumPred (p) + x.right.NumPred (p)
}

func (x *btree) Contained (a Any) (BTree, bool) {
  if x == nil { return nil, false }
  if Less (a, x.root) {
    return x.left.Contained (a)
  }
  if Less (x.root, a) {
    return x.right.Contained (a)
  }
  // a and x.Any cannot be distinguished by Less, hence are considered to be equal:
  return x, true
}

func (x *btree) All (p Pred) bool {
  if x == nil { return true }
  if p (x.root) {
    return x.left.All (p) &&
           x.right.All (p)
  }
  return false
}

func (x *btree) First (a Any) BTree {
  if x == nil {
    return nil
  }
  if Less (a, x.root) {
    if x.left == nil {
      return nil
    }
    y := x.left.First (a).(*btree)
    if y == nil {
      return x
    }
    if Less (y.root, x.root) {
      return y
    } else {
      return x
    }
  } else if Less (x.root, a) {
    if x.right == nil {
      return nil
    }
    return x.right.First (a)
  } // see above remark, i.e. Eq (x.root, a)
  return x
}

// Pre: x and x.right are not empty, x is rightweighty,
//      x.right is i) rightweighty or ii) balanced.
// i)  x and x.left are balanced,
// ii) x is leftweighty, x.left is rightweighty.
func (x *btree) rotL() *btree {
  y := x.right
  x.right = y.left
  y.left = x
  x = y
  if x.balance == rightweighty { // case i)
    x.balance = balanced
    x.left.balance = balanced
  } else { // case ii)
    x.balance = leftweighty
    x.left.balance = rightweighty
  }
  return x
}

// dually to rotL
func (x *btree) rotR() *btree {
  y := x.left
  x.left = y.right
  y.right = x
  x = y
  if x.balance == leftweighty {
    x.balance = balanced
    x.right.balance = balanced
  } else {
    x.balance = rightweighty
    x.right.balance = leftweighty
  }
  return x
}

// Pre: x, x.left and x.left.right are not empty, 
//      x is not balanced, 
//      x is leftweighty, x.left is rightweighty.
// x is balanced.
func (x *btree) rotLR() *btree {
  y := x.left
  z := y.right
  y.right = z.left
  z.left = y
  x.left = z.right
  z.right = x
  x = z
  switch x.balance {
  case leftweighty:
    x.left.balance = balanced
    x.right.balance = rightweighty
  case balanced: // exactly the minimal case
    x.left.balance = balanced
    x.right.balance = balanced
  case rightweighty:
    x.left.balance = leftweighty
    x.right.balance = balanced
  }
  x.balance = balanced
  return x
}

// dually to rotLR
func (x *btree) rotRL() *btree {
  y := x.right
  z := y.left
  y.left = z.right
  z.right = y
  x.right = z.left
  z.left = x
  x = z
  switch x.balance {
  case leftweighty: // t was t.right.left before
    x.left.balance = balanced
    x.right.balance = rightweighty
  case balanced: // exactly the minimal case
    x.left.balance = balanced
    x.right.balance = balanced
  case rightweighty:
    x.left.balance = leftweighty
    x.right.balance = balanced
  }
  x.balance = balanced
  return x
}

func (x *btree) in (a Any, increased *bool) (*btree, *btree) {
  if x == nil {
    x = new_(a).(*btree)
    *increased = true
    return x, x // second result: the inserted leaf
  }
  var inserted *btree
  if Less (a, x.root) {
/*/
    if x.left == nil {
      *increased = true
      return new_(a).(*btree), new_(a).(*btree)
    }
/*/
    x.left, inserted = x.left.in (a, increased)
    if *increased {
      switch x.balance {
      case leftweighty:
        switch x.left.balance {
        case leftweighty:
          x = x.rotR() // case i)
        case balanced:
          ; // impossible
        case rightweighty:
          x = x.rotLR()
        }
        *increased = false
      case balanced:
        x.balance = leftweighty
      case rightweighty:
        x.balance = balanced
        *increased = false
      }
    }
  } else if Less (x.root, a) {
    if x.right == nil {
//  errh.Error0 ("Kackhaufen")
      *increased = true
      return new_(a).(*btree), new_(a).(*btree)
    }
    x.right, inserted = x.right.in (a, increased)
    if *increased {
      switch x.balance {
        case rightweighty:
        switch x.right.balance {
        case rightweighty:
          x = x.rotL() // case i)
        case balanced:
          ; // impossible
        case leftweighty:
          x = x.rotRL()
        }
        *increased = false
      case balanced:
        x.balance = rightweighty
      case leftweighty:
        x.balance = balanced
        *increased = false
      }
    }
  } else { // a is already there
    *increased = false
  }
  return x, inserted
}

func (x *btree) Ins (a Any) (BTree, BTree) {
  increased := false
  return x.in (a, &increased)
}

func (x *btree) rebalL (decreased *bool) *btree {
  if *decreased {
    switch x.balance {
    case leftweighty:
      x.balance = balanced
    case balanced:
      x.balance = rightweighty
      *decreased = false
    case rightweighty:
      if x.right.balance == leftweighty {
        x = x.rotRL()
      } else {
        x = x.rotL()
        if x.balance == leftweighty {
          *decreased = false
        }
      }
    }
  }
  return x
}

func (x *btree) rebalR (decreased *bool) *btree {
  if *decreased {
    switch x.balance {
    case rightweighty:
      x.balance = balanced
    case balanced:
      x.balance = leftweighty
      *decreased = false
    case leftweighty:
      if x.left.balance == rightweighty {
        x = x.rotLR()
      } else {
        x = x.rotR()
        if x.balance == rightweighty {
          *decreased = false
        }
      }
    }
  }
  return x
}

func (x *btree) liftL (y BTree, decreased, oneLess *bool) *btree {
  if x.right == nil {
    y.(*btree).root = Clone (x.Root())
    *decreased, *oneLess = true, true
    x = x.left
  } else {
    x.right = x.right.liftL (y, decreased, oneLess)
    x = x.rebalR (decreased)
  }
  return x
}

func (x *btree) liftR (y BTree, decreased, oneLess *bool) *btree {
  if x.left == nil {
    y.(*btree).root = Clone (x.root)
    *decreased, *oneLess = true, true
    x = x.right
  } else {
    x.left = x.left.liftR (y, decreased, oneLess)
    x = x.rebalL (decreased)
  }
  return x
}

func (x *btree) del (a Any, decreased *bool) (*btree, bool) {
  oneLess := false
  if x == nil {
    return x, oneLess
  }
  if Less (a, x.root) {
    x.left, oneLess = x.left.del (a, decreased)
    x = x.rebalL (decreased)
  } else if Less (x.root, a) {
    x.right, oneLess = x.right.del (a, decreased)
    x = x.rebalR (decreased)
  } else { // found btree to remove
    if x.right == nil {
      *decreased, oneLess = true, true
      x = x.left // .(*btree)
    } else if x.left == nil {
      *decreased, oneLess = true, true
      x = x.right // .(*btree)
    } else if x.balance == leftweighty {
      x.left = x.left.liftL (x, decreased, &oneLess)
      x = x.rebalL (decreased)
    } else {
      x.right = x.right.liftR (x, decreased, &oneLess)
      x = x.rebalR (decreased)
    }
  }
  return x, oneLess
}

func (x *btree) Del (a Any) (BTree, bool) {
  decreased := false
  return x.del (a, &decreased)
}

func (x *btree) ExPred (p Pred) BTree {
  if x == nil { return nil }
  l := x.left.ExPred (p)
  if l != nil { return l }
  r := x.right.ExPred (p)
  if r != nil { return r }
  if p (x.root) {
    return x
  }
  return nil
}

func (x *btree) Trav (op Op) {
  if x != nil {
    x.left.Trav (op)
    op (x.root)
    x.right.Trav (op)
  }
}

func (x *btree) split (p Pred) (*btree, *btree) {
  errh.Error0 ("Split is not yet implemented")
  return nil, nil
}

func (x *btree) Write (x0, x1, y, dy uint, f func (Any) string) {
  if x == nil { return }
  xm := (x0 + x1) / 2
  y1 := int(y + scr.Ht1() / 2) - 1
  if x.left != nil {
    scr.Line (int(xm), y1, int(x0 + xm) / 2, y1 + int (dy))
  }
  if x.right != nil {
    scr.Line (int(xm), y1, int(xm + x1) / 2, y1 + int (dy))
  }
  scr.WriteGr (f (x.Root()), int(xm - scr.Wd1()), int(y))
  x.left.Write (x0, xm, y + dy, dy, f)
  x.right.Write (xm, x1, y + dy, dy, f)
}

func (x *btree) Write1 (d uint, f func (Any) string) {
  if x == nil { return }
  x.right.Write1 (d + 1, f)
  println (str.New (6 * d) + f (x.Root()))
  x.left.Write1 (d + 1, f)
}
