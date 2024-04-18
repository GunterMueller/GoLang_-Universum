package set

// (c) Christian Maurer   v. 240318 - license see µU.go

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
          any "content of the node"
         left,
        right *node
              balance
              }
type
  set struct {
             any "pattern object"
      anchor,
      actual *node
             uint "number of objects in the set"
             }
var
  tmp *node

func new_(a any) Set {
  CheckEqualerAndComparer (a)
  x := new (set)
  x.any = Clone(a)
  x.anchor, x.actual = nil, nil
  return x
}

func (x *set) imp (Y any) *set {
  y, ok := Y.(*set)
  if ! ok { TypeNotEqPanic (x, Y) }
  CheckTypeEq (x.any, y.any)
  return y
}

func (x *set) Eq (Y any) bool {
  xs := make ([]any, 0)
  x.Trav (func (a any) {xs = append (xs, a)})
  y := x.imp (Y)
  ys := make ([]any, 0)
  y.Trav (func (a any) {ys = append (ys, a)})
  n := len(xs)
  if len(ys) != n {
    return false
  }
  for i := 0; i < n; i++ {
    if ! Eq (xs[i], ys[i]) {
      return false
    }
  }
  return true
}

func (x *set) Copy (Y any) {
  x.Clr()
  y := x.imp (Y)
  y.Trav (func (a any) {x.Ins (a)})
}

func (x *set) Clone() any {
  y := new_(x.any).(*set)
  y.Copy (x)
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
  return n
}

func (x *set) Ex (a any) bool {
  CheckTypeEq (x.any, a)
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
    x.actual = next (x.anchor, x.actual.any)
  } else {
    x.actual = prev (x.anchor, x.actual.any)
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

func (x *set) Get() any {
  if x.anchor == nil || x.actual == nil {
    return nil
  }
  return Clone (x.actual.any)
}

func (x *set) Put (a any) {
  CheckTypeEq (x.any, a)
  x.Del()
  x.Ins (a)
}

func (x *set) Ins (a any) {
  CheckTypeEq (x.any, a)
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

func (x *set) Del() any {
  if x.anchor == nil {
    return nil
  }
  act := x.actual
  toDelete := x.actual.any
  x.Step (true) // to set "actual" to the node containing
                // the next largest object, iff such exists
  var a any
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
  return Clone (act.any)
}

func (x *set) ExGeq (a any) bool {
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
  y.Trav (func (a any) { x.Ins (a) })
  y.Clr()
  x.Jump (false)
}

func (x *set) Ordered() bool {
  if x.uint <= 1 { return true }
  x.Jump (false)
  result, first, o := true, true, x.actual.any
  x.Trav (func (a any) {
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
  y := new_(x.any).(*set)
  x.Trav (func (a any) { y.Ins (a) } )
  x.anchor, x.uint = y.anchor, y.uint
  x.Jump (false)
}
