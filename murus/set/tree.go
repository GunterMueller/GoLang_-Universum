package set

// (c) murus.org  v. 140101 - license see murus.go

import (
  . "murus/obj"
)
type
  balance byte; const (
  leftweighty = balance(iota)
  balanced
  rightweighty
)
type (
  tree struct {
              Any
         left,
        right *tree
              balance
              }
)

func leaf (a Any) *tree {
  t:= new (tree)
  t.Any = Clone (a)
  t.left, t.right = nil, nil
  t.balance = balanced
  return t
}

func (x *tree) num() uint {
  if x == nil {
    return uint(0)
  }
  return x.left.num() + 1 + x.right.num()
}

func (x *tree) numPred (p Pred) uint {
  n:= uint(0)
  if x == nil { return n }
  if p (x.Any) { n ++ }
  return n + x.left.numPred (p) + x.right.numPred (p)
}

func (x *tree) contained (a Any) (*tree, bool) {
  if x == nil { return nil, false }
  if Less (a, x.Any) {
    return x.left.contained (a)
  }
  if Less (x.Any, a) {
    return x.right.contained (a)
  }
  // a and x.Any cannot be distinguished by Less, hence are considered to be equal:
  return x, true
}

func (x *tree) all (p Pred) bool {
  if x == nil { return true }
  if p (x.Any) {
    return x.left.all (p) &&
           x.right.all (p)
  }
  return false
}

func (x *tree) first (a Any) *tree {
  if x == nil { return nil }
  if Less (a, x.Any) {
    y:= x.left.first (a)
    if y == nil {
      return x
    }
    if Less (y.Any, x.Any) {
      return y
    } else {
      return x
    }
  } else if Less (x.Any, a) {
    return x.right.first (a)
  } // see above remark
  return x
}

// Pre: x and x.right are not empty, x is rightweighty,
//      x.right is i) rightweighty or ii) balanced.
// i)  x and x.left are balanced,
// ii) x is leftweighty, x.left is rightweighty.
func (x *tree) rotL() *tree {
  y:= x.right
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
func (x *tree) rotR() *tree {
  y:= x.left
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

// Pre: t, t.left and t.left.right are not empty, 
//      t is not balanced, 
//      t is leftweighty, t.left is rightweighty.
// t is balanced.
func (x *tree) rotLR() *tree {
  y:= x.left
  z:= y.right
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
func (x *tree) rotRL() *tree {
  y:= x.right
  z:= y.left
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

func (x *tree) in (a Any, increased *bool) (*tree, *tree) {
  if x == nil {
    x = leaf (a)
    *increased = true
    return x, x // second result: the inserted leaf
  }
  var inserted *tree
  if Less (a, x.Any) {
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
  } else if Less (x.Any, a) {
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

func (x *tree) ins (a Any) (*tree, *tree) {
  increased:= false
  return x.in (a, &increased)
}

func (x *tree) rebalL (decreased *bool) *tree {
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

func (x *tree) rebalR (decreased *bool) *tree {
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

func (x *tree) liftL (y *tree, decreased, oneLess *bool) *tree {
  if x.right == nil {
    y.Any = Clone (x.Any)
    *decreased, *oneLess = true, true
    x = x.left
  } else {
    x.right = x.right.liftL (y, decreased, oneLess)
    x = x.rebalR (decreased)
  }
  return x
}

func (x *tree) liftR (y *tree, decreased, oneLess *bool) *tree {
  if x.left == nil {
    y.Any = Clone (x.Any)
    *decreased, *oneLess = true, true
    x = x.right
  } else {
    x.left = x.left.liftR (y, decreased, oneLess)
    x = x.rebalL (decreased)
  }
  return x
}

func (x *tree) d (a Any, decreased *bool) (*tree, bool) {
  oneLess:= false
  if x == nil {
    return x, oneLess
  }
  if Less (a, x.Any) {
    x.left, oneLess = x.left.d (a, decreased)
    x = x.rebalL (decreased)
  } else if Less (x.Any, a) {
    x.right, oneLess = x.right.d (a, decreased)
    x = x.rebalR (decreased)
  } else { // found tree to remove
    if x.right == nil {
      *decreased, oneLess = true, true
      x = x.left
    } else if x.left == nil {
      *decreased, oneLess = true, true
      x = x.right
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

func (x *tree) del (a Any) (*tree, bool) {
  decreased:= false
  return x.d (a, &decreased)
}

func (x *tree) exPred (p Pred) *tree {
  if x == nil { return nil }
  l:= x.left.exPred (p)
  if l != nil { return l }
  r:= x.right.exPred (p)
  if r != nil { return r }
  if p (x.Any) {
    return x
  }
  return nil
}

func (x *tree) trav (op Op) {
  if x != nil {
    x.left.trav (op)
    op (x.Any)
    x.right.trav (op)
  }
}

func (x *tree) split (p Pred) (*tree, *tree) {
  return nil, nil
}
