package set

// (c) Christian Maurer  v. 210321 - license see µU.go

import
  . "µU/obj"
type
  pointer = *node

func newNode (a Any) pointer {
  x := new (node)
  x.Any = Clone (a)
  x.left, x.right = nil, nil
  x.balance = balanced
  return x
}

func num (x *node) uint {
  if x == nil {
    return uint(0)
  }
  return num (x.left) + 1 + num (x.right)
}

func ex (x *pointer, a Any) (*node, bool) {
  if *x == nil {
    return nil, false
  }
  if Eq (a, (*x).Any) {
    return *x, true
  }
  if Less (a, (*x).Any) {
    return ex (&(*x).left, a)
  }
  return ex (&(*x).right, a) // Less ((*x).Any, a)
}

func minGeq (x *pointer, a Any) *node {
  if *x == nil {
    return nil
  }
  if Eq (a, (*x).Any) {
    return *x
  }
  if Less (a, (*x).Any) {
    if (*x).left == nil {
      return *x
    }
    tmp = *x
    return minGeq (&((*x).left), a)
  }
// Less ((*x).Any, a):
  if (*x).right == nil {
    return tmp
  }
  return minGeq (&((*x).right), a)
}

func next (x pointer, a Any) pointer {
  if x == nil {
    return nil
  }
  if Eq (a, x.Any) {
    if x.right == nil {
      if Less (a, tmp.Any) {
        return tmp
      }
      return x
    }
    return next (x.right, a)
  }
  if Less (a, x.Any) {
    if x.left == nil {
      return x
    }
    tmp = x
    return next (x.left, a)
  }
  // Less (x.Any, a):
  if x.right == nil {
    return tmp
  }
  return next (x.right, a)
}

func prev (x pointer, a Any) pointer {
  if x == nil {
    return nil
  }
  if Eq (a, x.Any) {
    if x.left == nil {
      if Less (tmp.Any, a) {
        return tmp
      }
      return x
    }
    return prev (x.left, a)
  }
  if Less (x.Any, a) {
    if x.right == nil {
      return x
    }
    tmp = x
    return prev (x.right, a)
  }
  // Less (a, x.Any):
  if x.left == nil {
    return tmp
  }
  return prev (x.left, a)
}

// Pre: *x and (*x).right are not empty, *x is rightweighty,
//      (*x).right is i) rightweighty or ii) balanced.
// i)  *x and (*p).left are balanced,
// ii) *x is leftweighty, (*x).left is rightweighty.
func rotL (x *pointer) {
  y := (*x).right
  (*x).right = (*y).left
  (*y).left = *x
  *x = y
  if (*x).balance == rightweighty { // case i)
    (*x).balance = balanced
    (*x).left.balance = balanced
  } else { // case ii)
    (*x).balance = leftweighty
    (*x).left.balance = rightweighty
  }
}

// Pre: *x and (*x).left are not empty, *x is leftweighty,
//      (*x).left is i) leftweighty or ii) balanced.
// i)  *x and (*x).right are balanced,
// ii) *x is rightweighty, (*x).right is leftweighty.
func rotR (x *pointer) {
  y := (*x).left
  (*x).left = (*y).right
  (*y).right = *x
  *x = y
  if (*x).balance == leftweighty { // case i)
    (*x).balance = balanced
    (*x).right.balance = balanced
  } else { // case ii)
    (*x).balance = rightweighty
    (*x).right.balance = leftweighty
  }
}

// Pre: *x, (*x).left and (*x).left.right are not empty, 
//      (*x) is not balanced, 
//      (*x) is leftweighty, (*x).left is rightweighty.
// *x is balanced.
func rotLR (x *pointer) {
  y := (*x).left
  z := y.right
  y.right = z.left
  z.left = y
  (*x).left = z.right
  z.right = *x
  *x = z
  switch (*x).balance {
  case leftweighty:
    (*x).left.balance = balanced
    (*x).right.balance = rightweighty
  case balanced:
    (*x).left.balance = balanced
    (*x).right.balance = balanced
  case rightweighty:
    (*x).left.balance = leftweighty
    (*x).right.balance = balanced
  }
  (*x).balance = balanced
}

// Pre: (*x), (*x).right and (*x).right.left are not empty, 
//      (*x) is not balanced, 
//      (*x) is rightweighty, (*x).right is leftweighty.
// (*x) is balanced.
func rotRL (x *pointer) {
  y := (*x).right
  z := y.left
  y.left = z.right
  z.right = y
  (*x).right = z.left
  z.left = *x
  *x = z
  switch (*x).balance {
  case leftweighty:
    (*x).left.balance = balanced
    (*x).right.balance = rightweighty
  case balanced:
    (*x).left.balance = balanced
    (*x).right.balance = balanced
  case rightweighty:
    (*x).left.balance = leftweighty
    (*x).right.balance = balanced
  }
  (*x).balance = balanced
}

func ins (x *pointer, a Any, increased *bool) pointer {
  if *x == nil {
    *x = newNode (a)
    *increased = true
    return *x
  }
  var inserted pointer
  if Less (a, (*x).Any) {
    inserted = ins (&((*x).left), a, increased)
    if *increased {
      switch (*x).balance {
      case leftweighty:
        switch (*x).left.balance {
        case leftweighty:
          rotR (x) // case i)
        case balanced:
          // impossible
        case rightweighty:
          rotLR (x)
        }
        *increased = false
      case balanced:
        (*x).balance = leftweighty
      case rightweighty:
        (*x).balance = balanced
        *increased = false
      }
    }
  } else if Less ((*x).Any, a) {
    inserted = ins (&((*x).right), a, increased)
    if *increased {
      switch (*x).balance {
        case rightweighty:
        switch (*x).right.balance {
        case rightweighty:
          rotL (x) // case i)
        case balanced:
          // impossible
        case leftweighty:
          rotRL (x)
        }
        *increased = false
      case balanced:
        (*x).balance = rightweighty
      case leftweighty:
        (*x).balance = balanced
        *increased = false
      }
    }
  } else { // Eq (a, (*x).Any), i.e., a is already there
    *increased = false
  }
  return inserted
}

func rebalL (x *pointer, decreased *bool) {
  if *decreased {
    switch (*x).balance {
    case leftweighty:
      (*x).balance = balanced
    case balanced:
      (*x).balance = rightweighty
      *decreased = false
    case rightweighty:
      if (*x).right.balance == leftweighty {
        rotRL (x)
      } else {
        rotL (x)
        if (*x).balance == leftweighty {
          *decreased = false
        }
      }
    }
  }
}

func rebalR (x *pointer, decreased *bool) {
  if *decreased {
    switch (*x).balance {
    case rightweighty:
      (*x).balance = balanced
    case balanced:
      (*x).balance = leftweighty
      *decreased = false
    case leftweighty:
      if (*x).left.balance == rightweighty {
        rotLR (x)
      } else {
        rotR (x)
        if (*x).balance == rightweighty {
          *decreased = false
        }
      }
    }
  }
}

func liftL (x *pointer, y pointer, decreased, oneLess *bool) {
  if (*x).right == nil {
    y.Any = Clone ((*x).Any)
    *decreased, *oneLess = true, true
    *x = (*x).left
  } else {
    liftL (&((*x).right), y, decreased, oneLess)
    rebalR (x, decreased)
  }
}

func liftR (x *pointer, y pointer, decreased, oneLess *bool) {
  if (*x).left == nil {
    y.Any = Clone ((*x).Any)
    *decreased, *oneLess = true, true
    *x = (*x).right
  } else {
    liftR (&((*x).left), y, decreased, oneLess)
    rebalL (x, decreased)
  }
}

func del (x *pointer, a Any, decreased *bool) bool {
  oneLess := false
  if *x == nil {
    return oneLess
  }
  if Less (a, (*x).Any) {
    oneLess = del (&((*x).left), a, decreased)
    rebalL (x, decreased)
  } else if Less ((*x).Any, a) {
    oneLess = del (&((*x).right), a, decreased)
    rebalR (x, decreased)
  } else { // found node to remove
    if (*x).right == nil {
      *decreased, oneLess = true, true
      *x = (*x).left
    } else if (*x).left == nil {
      *decreased, oneLess = true, true
      *x = (*x).right
    } else if (*x).balance == leftweighty {
      liftL (&((*x).left), *x, decreased, &oneLess)
      rebalL (x, decreased)
    } else {
      liftR (&((*x).right), *x, decreased, &oneLess)
      rebalR (x, decreased)
    }
  }
  return oneLess
}

func trav (x *node, op Op) {
  if x != nil {
    trav (x.left, op)
    op (x.Any)
    trav (x.right, op)
  }
}
