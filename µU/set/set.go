package set

// (c) Christian Maurer  v. 210225 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
)
type
  set struct {
             Any
      anchor,
      actual *node
         num uint
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

func (x *set) check (a Any) {
  CheckTypeEq (x.Any, a)
}

func (x *set) imp (Y Any) *set {
  y, ok := Y.(*set)
  if ! ok { TypeNotEqPanic (x, Y) }
  x.check (y.Any)
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
  x.num = 0
}

func (x *set) Num() uint {
  return x.num
}

func (x *set) Ex (a Any) bool {
  x.check (a)
  if n, c := x.anchor.ex (a); c {
    x.actual = n
    return true
  }
  return false
}

func (x *set) Step (forward bool) {
  if x.Empty() { return }
  if x.actual == nil {
    return
  }
  if forward {
    x.actual = x.anchor.next (x.actual.Any)
  } else {
    x.actual = x.anchor.prev (x.actual.Any)
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
  if x.anchor == nil { return nil }
  if x.actual == nil { ker.Panic ("set.Get error: x.actual == nil") }
  return Clone (x.actual.Any)
}

func (x *set) Put (a Any) {
  x.check (a)
  x.Del()
  x.Ins (a)
}

func (x *set) Ins (a Any) {
  x.check (a)
  if x.anchor == nil {
    x.anchor = newNode (a)
    x.actual = x.anchor
    x.num = 1
  } else {
    var n *node
    x.anchor, n = x.anchor.ins (a)
    if n != nil {
      x.actual = n
      x.num ++
    }
  }
}

func (x *set) Del() Any {
  if x.anchor == nil {
    return nil
  }
  act := x.actual
  toDel := x.actual.Any
  x.Step (true) // to set "actual" one step ahead
  var a Any = nil
  if act == x.actual { // the root to remove is the last right node in x,
                       // "actual" must be reset one position or set to nil, see below
  } else {
    a = Clone (toDel)
  }
  oneLess := false
  x.anchor, oneLess = x.anchor.del (toDel)
  if oneLess {
    if act == x.actual { // the root to remove was the last right node of x
      if x.num == 1 {    // see above
        x.actual = nil   // x is now empty
      } else {
        x.Jump (true)
      }
    } else {
      if x.Ex (a) { // thus the above copy-action to "a": "actual" might have been
                    // rotated off while deleting, with this trick is it restored !
      }
    }
    x.num--
  }
  return Clone (act.Any)
}

func (x *set) ExGeq (a Any) bool {
  tmp = nil
  n := x.anchor.minGeq (a)
  if n == nil {
    return false
  }
  x.actual = n
  return true
}

func (x *set) Trav (op Op) {
  x.anchor.trav (op)
}

func (x *set) Join (Y Collector) {
  y := x.imp (Y)
  y.Trav (func (a Any) { x.Ins (a) })
  y.Clr()
  x.Jump (false)
}

func (x *set) Ordered() bool {
  if x.num <= 1 { return true }
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
  x.anchor, x.num = y.anchor, y.num
  x.Jump (false)
}
