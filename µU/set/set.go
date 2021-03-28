package set

// (c) Christian Maurer  v. 210321 - license see µU.go

import
  . "µU/obj"
type
  balance byte; const (
  leftweighty = balance(iota)
  balanced
  rightweighty
)
type
  node struct {
          Any "content of the node"
         left,
        right *node
              balance
              }
type
  set struct {
             Any "pattern object"
      anchor,
      actual *node
             uint "number of objects in the set"
             }
var
  tmp *node

func new_(a Any) Set {
  CheckEqualerAndComparer (a)
  x := new (set)
  x.Any = Clone(a)
  x.anchor, x.actual = nil, nil
  return x
}

func (x *set) imp (Y Any) *set {
  y, ok := Y.(*set)
  if ! ok { TypeNotEqPanic (x, Y) }
  CheckTypeEq (x.Any, y.Any)
  return y
}

func (x *set) Offc() bool {
  return x.actual == nil
}

func (x *set) Empty() bool {
  return x.anchor == nil
}

func (x *set) Clr() {
  x.anchor, x.actual = nil, nil
  x.uint = 0
}

func (x *set) Num() uint {
  n := x.uint
//  if num (x.anchor) != n { ker.Oops() }
  return n
}

func (x *set) Ex (a Any) bool {
  CheckTypeEq (x.Any, a)
  if n, ok := ex (&(x.anchor), a); ok {
    x.actual = n
    return true
  }
  return false
}

func (x *set) Step (forward bool) {
  if x.Empty() || x.actual == nil {
    return
  }
  if forward {
    x.actual = next (x.anchor, x.actual.Any)
  } else {
    x.actual = prev (x.anchor, x.actual.Any)
  }
}

func (x *set) Jump (toEnd bool) {
  if x.Empty() { return }
  x.actual = x.anchor
  for {
    if toEnd {
      if x.actual.right == nil {
        break
      }
      x.actual = x.actual.right
    } else {
      if x.actual.left == nil {
        break
      }
      x.actual = x.actual.left
    }
  }
}

func (x *set) Eoc (forward bool) bool {
  t := x.anchor
  for t != nil {
    if t == x.actual {
      if forward {
        return t.right == nil
      } else {
        return t.left == nil
      }
    }
    if forward {
      t = t.right
    } else {
      t = t.left
    }
  }
  return false
}

func (x *set) Get() Any {
  if x.anchor == nil || x.actual == nil {
    return nil
  }
  return Clone (x.actual.Any)
}

func (x *set) Put (a Any) {
  CheckTypeEq (x.Any, a)
  x.Del()
  x.Ins (a)
}

func (x *set) Ins (a Any) {
  CheckTypeEq (x.Any, a)
  if x.anchor == nil {
    x.anchor = newNode (a)
    x.actual = x.anchor
    x.uint = 1
  } else {
    increased := false
    n := ins (&(x.anchor), a, &increased)
    if n != nil {
      x.actual = n
      x.uint++
    }
  }
}

func (x *set) Del() Any {
  if x.anchor == nil {
    return nil
  }
  act := x.actual
  toDelete := x.actual.Any
  x.Step (true) // to set "actual" to the node containing
                // the next largest object, iff such exists
  var a Any
  if act == x.actual { // the node to delete is the node with
    a = nil            // the largest object in x, so "actual"
                       // must be set to the node containing
                       // the next smallest object, see below
  } else {
    a = Clone (toDelete)
  }
  decreased := false
  if del (&(x.anchor), toDelete, &decreased) { // the object
                         // to be deleted was found and deleted
    if act == x.actual { // the node to delete was the last
                         // right node of x
      if x.uint == 1 {   // see above
        x.actual = nil   // x is now empty
      } else {
        x.Jump (true)    // "actual" is the last right node
      }
    } else { // the node with the next largest object exists
      if x.Ex (a) { // thus the above copy-action to "a":
                    // "actual" might have been rotated off
                    // while deleting, with this trick
                    // it is found again.
      }
    }
    x.uint--
  }
  return Clone (act.Any)
}

func (x *set) ExGeq (a Any) bool {
  tmp = nil
  n := minGeq (&(x.anchor), a)
  if n == nil {
    return false
  }
  x.actual = n
  return true
}

func (x *set) Trav (op Op) {
  trav (x.anchor, op)
}

func (x *set) Join (Y Collector) {
  y := x.imp (Y)
  y.Trav (func (a Any) { x.Ins (a) })
  y.Clr()
  x.Jump (false)
}

func (x *set) Ordered() bool {
  if x.uint <= 1 { return true }
  x.Jump (false)
  result, first, o := true, true, x.actual.Any
  x.Trav (func (a Any) {
            if first {
              first = false
            } else {
              if ! Less (o, a) {
                result = false
              }
            }
            o = a
          })
  return result
}

func (x *set) Sort() {
  if x.Num() <= 1 { return }
  y := new_(x.Any).(*set)
  x.Trav (func (a Any) { y.Ins (a) } )
  x.anchor, x.uint = y.anchor, y.uint
  x.Jump (false)
}
