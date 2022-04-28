package internal

// (c) Christian Maurer   v. 220420 - license see µU.go

import
  . "µU/obj"
type
  heap struct {
              any "pattern object"
  left, right *heap
              }

func new_() Heap {
  return nil
}

// Pre: n > 0.
// Returns the greatest power of 2 <= n.
func g (n uint) uint {
  if n == 1 {
    return n
  }
  return 2 * g (n / 2)
}

// Pre: n > 0.
// Returns true, iff the last node of x with n nodes is contained
// in the left subheap of x, and in this case the number of nodes
// in the left, otherwise in the right subheap of x.
func f (n uint) (bool, uint) {
  left := true
  if n == 1 {
    return left, 0
  }
  a := g (n)
  b := n - a
  left = b < a / 2
  if left {
    b += a / 2
  }
  return left, b
}

func (x *heap) Ins (a any, n uint) Heap {
  if n == 1 {
    x = new (heap)
    x.any = Clone (a)
    x.left, x.right = nil, nil
  } else {
    left, k := f (n)
    if left {
      x.left = x.left.Ins (a, k).(*heap)
    } else {
      x.right = x.right.Ins (a, k).(*heap)
    }
  }
  return x
}

func (x *heap) swap (l bool) {
  if l {
    if x.left != nil {
      if Less (x.left.any, x.any) {
        x.any, x.left.any = x.left.any, x.any
      }
    }
  } else if x.right != nil {
    if Less (x.right.any, x.any) {
      x.any, x.right.any = x.right.any, x.any
    }
  }
}

func (x *heap) Lift (n uint) {
  if n > 0 {
    left, k := f (n)
    if left {
      x.left.Lift (k)
    } else {
      x.right.Lift (k)
    }
    x.swap (left)
  }
}

// Pre: n == number of objects in x > 0.
// Returns the former pointer to the n-th node of x,
// and this pointer is now nil.
func (x *heap) last (n uint) *heap {
  switch n {
  case 1:
    return x
  case 2:
    y := x.left
    x.left = nil
    return y
  case 3:
    y := x.right
    x.right = nil
    return y
  }
  left, k := f (n)
  if left {
    return x.left.last (k)
  }
  return x.right.last (k)
}

func (x *heap) Del (n uint) (Heap, any) {
  y := x.last (n)
  switch n {
  case 1:
    y = nil
  case 2:
    // see above
  case 3:
    y.left = x.left
  default:
    y.left = x.left
    y.right = x.right
  }
  return y, x.any
}

func (x *heap) Sift (n uint) {
  if x.left != nil {
    if x.right == nil {
      if Less (x.left.any, x.any) {
        x.swap (true)
      }
    } else { // x.left != nil && x.right != nil
      if Less (x.any, x.left.any) && Less (x.any, x.right.any) {
        return
      }
      if Less (x.left.any, x.right.any) {
        x.swap (true)
        x.left.Sift (n)
      } else {
        x.swap (false)
        x.right.Sift (n)
      }
    }
  }
}

func (x *heap) Get() any {
  if x == nil {
    return nil
  }
  return Clone (x.any)
}
