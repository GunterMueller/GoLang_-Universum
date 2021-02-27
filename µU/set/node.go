package set

// (c) Christian Maurer  v. 210225 - license see µU.go

import
  . "µU/obj"
type
  balance byte; const (
  leftweighty = balance(iota)
  balanced
  rightweighty
)
type (
  node struct {
              Any
         left,
        right *node
              balance
              }
)

func newNode (a Any) *node {
  n := new (node)
  n.Any = Clone (a)
  n.left, n.right = nil, nil
  n.balance = balanced
  return n
}

func (n *node) num() uint {
  if n == nil {
    return uint(0)
  }
  return n.left.num() + 1 + n.right.num()
}

func (n *node) ex (a Any) (*node, bool) {
  if n == nil {
    return nil, false
  }
  if Eq (a, n.Any) {
    return n, true
  }
  if Less (a, n.Any) {
    return n.left.ex (a)
  }
  return n.right.ex (a) // Less (n.Any, a)
}

func (n *node) minGeq (a Any) *node {
  if n == nil {
    return nil
  }
  if Eq (a, n.Any) {
    return n
  }
  if Less (a, n.Any) {
    if n.left == nil {
      return n
    }
    tmp = n
    return n.left.minGeq (a)
  }
// Less (n.Any, a):
  if n.right == nil {
    return tmp
  }
  return n.right.minGeq (a)
}

func (n *node) next (a Any) *node {
  if n == nil {
    return nil
  }
  if Eq (a, n.Any) {
    if n.right == nil {
      if Less (a, tmp.Any) {
        return tmp
      }
      return n
    }
    return n.right.next (a)
  }
  if Less (a, n.Any) {
    if n.left == nil {
      return n
    }
    tmp = n
    return n.left.next (a)
  }
  // Less (n.Any, a):
  if n.right == nil {
    return tmp
  }
  return n.right.next (a)
}

func (n *node) prev (a Any) *node {
  if n == nil {
    return nil
  }
  if Eq (a, n.Any) {
    if n.left == nil {
      if Less (tmp.Any, a) {
        return tmp
      }
      return n
    }
    return n.left.prev (a)
  }
  if Less (n.Any, a) {
    if n.right == nil {
      return n
    }
    tmp = n
    return n.right.prev (a)
  }
  // Less (a, n.Any):
  if n.left == nil {
    return tmp
  }
  return n.left.prev (a)
}

// Pre: x and x.right are not empty, x is rightweighty,
//      x.right is i) rightweighty or ii) balanced.
// i)  x and x.left are balanced,
// ii) x is leftweighty, x.left is rightweighty.
func (x *node) rotL() *node {
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
func (x *node) rotR() *node {
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
func (x *node) rotLR() *node {
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
func (x *node) rotRL() *node {
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

func (n *node) in (a Any, increased *bool) (*node, *node) {
  if n == nil {
    n = newNode (a)
    *increased = true
    return n, n // second result: the inserted leaf
  }
  var inserted *node
  if Less (a, n.Any) {
    n.left, inserted = n.left.in (a, increased)
    if *increased {
      switch n.balance {
      case leftweighty:
        switch n.left.balance {
        case leftweighty:
          n = n.rotR() // case i)
        case balanced:
          ; // impossible
        case rightweighty:
          n = n.rotLR()
        }
        *increased = false
      case balanced:
        n.balance = leftweighty
      case rightweighty:
        n.balance = balanced
        *increased = false
      }
    }
  } else if Less (n.Any, a) {
    n.right, inserted = n.right.in (a, increased)
    if *increased {
      switch n.balance {
        case rightweighty:
        switch n.right.balance {
        case rightweighty:
          n = n.rotL() // case i)
        case balanced:
          ; // impossible
        case leftweighty:
          n = n.rotRL()
        }
        *increased = false
      case balanced:
        n.balance = rightweighty
      case leftweighty:
        n.balance = balanced
        *increased = false
      }
    }
  } else { // a is already there
    *increased = false
  }
  return n, inserted
}

func (n *node) ins (a Any) (*node, *node) {
  increased:= false
  return n.in (a, &increased)
}

func (n *node) rebalL (decreased *bool) *node {
  if *decreased {
    switch n.balance {
    case leftweighty:
      n.balance = balanced
    case balanced:
      n.balance = rightweighty
      *decreased = false
    case rightweighty:
      if n.right.balance == leftweighty {
        n = n.rotRL()
      } else {
        n = n.rotL()
        if n.balance == leftweighty {
          *decreased = false
        }
      }
    }
  }
  return n
}

func (n *node) rebalR (decreased *bool) *node {
  if *decreased {
    switch n.balance {
    case rightweighty:
      n.balance = balanced
    case balanced:
      n.balance = leftweighty
      *decreased = false
    case leftweighty:
      if n.left.balance == rightweighty {
        n = n.rotLR()
      } else {
        n = n.rotR()
        if n.balance == rightweighty {
          *decreased = false
        }
      }
    }
  }
  return n
}

func (n *node) liftL (y *node, decreased, oneLess *bool) *node {
  if n.right == nil {
    y.Any = Clone (n.Any)
    *decreased, *oneLess = true, true
    n = n.left
  } else {
    n.right = n.right.liftL (y, decreased, oneLess)
    n = n.rebalR (decreased)
  }
  return n
}

func (n *node) liftR (y *node, decreased, oneLess *bool) *node {
  if n.left == nil {
    y.Any = Clone (n.Any)
    *decreased, *oneLess = true, true
    n = n.right
  } else {
    n.left = n.left.liftR (y, decreased, oneLess)
    n = n.rebalL (decreased)
  }
  return n
}

func (n *node) d (a Any, decreased *bool) (*node, bool) {
  oneLess:= false
  if n == nil {
    return n, oneLess
  }
  if Less (a, n.Any) {
    n.left, oneLess = n.left.d (a, decreased)
    n = n.rebalL (decreased)
  } else if Less (n.Any, a) {
    n.right, oneLess = n.right.d (a, decreased)
    n = n.rebalR (decreased)
  } else { // found node to remove
    if n.right == nil {
      *decreased, oneLess = true, true
      n = n.left
    } else if n.left == nil {
      *decreased, oneLess = true, true
      n = n.right
    } else if n.balance == leftweighty {
      n.left = n.left.liftL (n, decreased, &oneLess)
      n = n.rebalL (decreased)
    } else {
      n.right = n.right.liftR (n, decreased, &oneLess)
      n = n.rebalR (decreased)
    }
  }
  return n, oneLess
}

func (n *node) del (a Any) (*node, bool) {
  decreased:= false
  return n.d (a, &decreased)
}

func (n *node) trav (op Op) {
  if n != nil {
    n.left.trav (op)
    op (n.Any)
    n.right.trav (op)
  }
}
