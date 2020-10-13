package set

// (c) Christian Maurer   v. 201011 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
)
type
  set struct {
             Any
      anchor,
      actual *tree
             uint "number of objects in set"
        path *node
             }

func new_(a Any) Set {
  CheckAtomicOrObject(a)
  x := new (set)
  x.Any = Clone(a)
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
  x.uint = 0
  x.path = nil
}

func (x *set) Copy (Y Any) {
  y := x.imp (Y)
  x.Clr()
  y.Trav (func (a Any) { x.Ins (a) })
  x.uint = y.uint
}

func (x *set) Clone() Any {
  y := new_(x.Any).(*set)
  x.Trav (func (a Any) { y.Ins (a) })
  y.uint = x.uint
  return y
}

func (x *set) e (y *set, r Rel) bool {
  if x.uint != y.uint { return false }
  if x.anchor == nil { return true }
  xact, yact := x.actual, y.actual
  x.Jump (false)
  y.Jump (false)
  for {
    if r (x.actual.Any, y.actual.Any) {
      if x.Eoc (true) {
        x.actual, y.actual = xact, yact
        return true
      } else {
        x.Step (true)
        y.Step (true)
      }
    } else {
      break
    }
  }
  x.actual, y.actual = xact, yact
  return false
}

func (x *set) Eq (Y Any) bool {
  y := x.imp (Y)
  return x.e (y, Eq)
}

func (x *set) Less (Y Any) bool {
  y := x.imp (Y)
  if ! Less (x.uint, y.uint) { return false }
  if x.anchor == nil { return true }
  return x.All (func (a Any) bool { return y.Ex (a) } )
}

func (x *set) Num() uint {
  return x.uint
}

func (x *set) NumPred (p Pred) uint {
  return x.anchor.numPred (p)
}

func (x *set) Ex (a Any) bool {
  x.check (a)
  if t, c := x.anchor.contained (a); c {
    x.actual = t
    return true
  }
  return false
}

func (x *set) All (p Pred) bool {
  return x.anchor.all (p)
}

func (x *set) Sort() {
  y := new_(x.Any).(*set)
  x.Trav (func (a Any) { y.Ins (a) } )
  x.anchor, x.uint = y.anchor, y.uint
  y.anchor = nil
//  x.actual = y.actual
//  x.Jump (false)
  x.actual = x.anchor
  for x.actual.left != nil {
    x.actual = x.actual.left
  }
}

var old *tree = nil
var neu *tree

func (x *set) Step (forward bool) {
  if x == nil { return }
  if x.anchor == nil { return }
  if x.uint < 2 { return }
  min, max := x.defPath()
  neu = x.actual
  if forward {
    if max { return }
  } else {
    if min { return }
  }
  if x.somethingBelow (forward) {
    x.actual = x.below (forward)
    for {
      if forward {
        if x.actual.left == nil { break }
        x.actual = x.actual.left
      } else {
        if x.actual.right == nil { break }
        x.actual = x.actual.right
      }
    }
  } else {
    for {
      if ! x.abovePointsToCurrent (forward) {
        x.up()
        x.actual = x.pointer()
        return
      }
      x.up()
    }
  }
}

func (x *set) Jump (toEnd bool) {
  if x.anchor == nil { return }
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
    x.anchor = leaf (a)
    x.actual = x.anchor
    x.uint = 1
  } else {
    var t *tree
    x.anchor, t = x.anchor.ins (a)
    if t != nil {
      x.actual = t
      x.uint++
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
  var tmp Any = nil
  if act == x.actual { // the root to remove is the last right node in x,
                       // "actual" must be reset one position or set to nil, see below
  } else {
    tmp = Clone (toDel)
  }
  oneLess := false
  x.anchor, oneLess = x.anchor.del (toDel)
  if oneLess {
    if act == x.actual { // the root to remove was the last right node of x
      if x.uint == 1 {    // see above
        x.actual = nil   // x is now empty
      } else {
        x.Jump (true)
      }
    } else {
      if x.Ex (tmp) { // thus the above copy-action to "tmp": "actual" might have been
                      // rotated off while deleting, with this trick is it restored !
      }
    }
    x.uint--
  }
  return Clone (act.Any)
}

func (x *set) ExPred (p Pred, f bool) bool {
  t := x.anchor.exPred (p)
  if t == nil {
    return false
  }
  x.actual = t
  return true
}

func (x *set) StepPred (p Pred, f bool) bool {
  if x == nil { return false }
  xact := x.actual
  for ! x.Eoc (f) {
    x.Step (f)
    if x.Eoc (f) { break }
    if p (x.actual.Any) {
      return true
    }
  }
  x.actual = xact
  return false
}

func (x *set) ExGeq (a Any) bool {
  t := x.anchor.first (a)
  if t == nil {
    return false
  }
  x.actual = t
  return true // XXX ? XXX ? XXX ? XXX ? XXX ? XXX ? XXX ? XXX ? XXX ? XXX ? XXX ? XXX ? XXX
}

func (x *set) Trav (op Op) {
  x.anchor.trav (op)
}

func (x *set) Filter (Y Iterator, p Pred) {
  y := x.imp (Y)
  if x.anchor == nil { return }
  y.Clr()
  x.Trav (func (a Any) { if p (a) { y.Ins (a) } })
  y.Jump (false)
}

func (x *set) Split (Y Iterator) {
  y := x.imp (Y)
  y.Clr()
  if x.anchor == nil { return }
  x1 := new_(x.Any).(*set)
  b := x.actual.Any
  x.Trav (func (a Any) { if Less (a, b) { x1.Ins (a) } else { y.Ins (a) } })
  x.anchor, x.uint = x1.anchor, x1.uint
  x.Jump (false)
  y.Jump (false)
}

func (x *set) Cut (Y Iterator, p Pred) {
  y := x.imp (Y)
  y.Clr()
  if x.anchor == nil { return }
  x1 := new_(x.Any).(*set)
  x.Trav (func (a Any) { if p (a) { y.Ins (a) } else { x1.Ins (a) } })
  x.anchor, x.uint = x1.anchor, x1.uint
  x.Jump (false)
  y.Jump (false)
}

func (x *set) ClrPred (p Pred) {
  if x.anchor == nil { return }
  y := new_(x.Any).(*set)
  x.Trav (func (a Any) { if ! p (a) { y.Ins (a) } })
  x.anchor, x.uint = y.anchor, y.uint
  x.Jump (false)
}

func (x *set) Join (Y Iterator) {
  y := x.imp (Y)
  y.Trav (func (a Any) { x.Ins (a) })
  y.Clr()
  x.Jump (false)
}

func (x *set) Codelen() uint {
  n := uint(4)
  x.Trav (func (a Any) { n += 4 + Codelen (a) })
  return n
}

func (x *set) Encode() []byte {
  b := make ([]byte, x.Codelen())
  copy (b[:4], Encode (x.uint))
  i := uint(4)
  x.Trav (func (a Any) {
            k := Codelen (a)
            copy (b[i:i+4], Encode (k))
            i += 4
            copy (b[i:i+k], Encode (a))
            i += k
          })
  return b
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

func (x *set) Decode (b []byte) {
  x.Clr()
  n := Decode (uint(0), b[:4]).(uint)
  i := uint(4)
  for j := uint(0); j < n; j++ {
    k := Decode (uint(0), b[i:i+4]).(uint)
    i += 4
    a := Decode (Clone (x.Any), b[i:i+k])
    i += k
    x.Ins (a)
  }
}

func (x *set) Write (x0, x1, y, dy uint, f func (Any) string) {
  x.anchor.write (x0, x1, y, dy, f)
}

func (x *set) Write1 (f func (Any) string) {
  x.anchor.write1 (0, f)
}
