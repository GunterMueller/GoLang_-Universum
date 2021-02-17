package set

// (c) Christian Maurer   v. 210215 - license see µU.go

// >>> The implementations of Ex, ExGeq have linear complexity, which is  very bad.
// >>> Therefor, this implementation will eventually be replaced by one using B-trees.

import (
  "sort"
  . "µU/obj"
)
type
  set struct {
             Any "pattern object"
         all []Any
             int // ordinal number of the actual object; -1 iff set is empty
             }

func new_(a Any) Set {
  CheckAtomicOrObject (a)
  x := new (set)
  x.Any = Clone (a)
  x.all = make([]Any, 0)
  x.int = -1
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
  return x.int == len (x.all)
}

func (x *set) Empty() bool {
  return len (x.all) == 0
}

func (x *set) Clr() {
  x.all = make([]Any, 0)
  x.int = -1
}

func (x *set) Num() uint {
  return uint(len(x.all))
}

func (x *set) Ex (a Any) bool {
  x.check (a)
  for i := 0; i < len(x.all); i++ {
    if Eq (a, x.all[i]) {
      x.int = i
      return true
    }
  }
  return false
}

func (x *set) Step (forward bool) {
  if x.Empty() { return }
  if forward {
    if x.int + 1 < len(x.all) {
      x.int++
    }
  } else {
    if x.int > 0 {
      x.int--
    }
  }
}

func (x *set) Jump (toEnd bool) {
  if x.Empty() { return }
  if toEnd {
    x.int = len(x.all) - 1
  } else {
    x.int = 0
  }
}

func (x *set) Eoc (forward bool) bool {
  if x.Empty() {
    return false
  }
  if forward {
    return x.int + 1 == len(x.all)
  }
  return x.int == 0
}

func (x *set) Get() Any {
  if x.Empty() {
    return nil
  }
  if x.int == -1 { return nil }
  return Clone (x.all[x.int])
}

func (x *set) Put (a Any) {
  x.check (a)
  if x.Empty() {
    x.all = make([]Any, 1)
    x.all[0] = Clone (a)
    x.int = 0
    return
  }
  x.Del()
  x.Ins (a)
}

func (x *set) Ins (a Any) {
  if x.Empty() {
    x.Put (a)
  }
  x.check (a)
  n := len(x.all)
  if n == 1 {
    if Eq (a, x.all[0]) {
      return
    }
    if Less (a, x.all[0]) {
      x.all = []Any {Clone(a), x.all[0]}
    } else {
      x.all = []Any {x.all[0], Clone(a)}
    }
    return
  }
  if x.ExGeq (a) {
    i := x.int
    if Eq (a, x.all[i]) {
      return
    }
    z := x.all[i]
    if i + 1 == len(x.all) && Less (a, z) {
       x.all = append (append (x.all[:i], a), z)
      return
    }
    if i + 1 < len(x.all) {
      zs := make([]Any, i)
      for j := 0; j < i; j++ {
        zs[j] = x.all[j]
      }
      xs := append (zs, Clone(a))
      ys := x.all[i:]
      x.all = append (xs, ys...)
      return
    }
  }
  if Eq (a, x.all[n-1]) {
    return
  }
  x.all = append (x.all, Clone(a))
}

func (x *set) Del() Any {
  if x.Empty() { return nil }
  a := x.all[x.int]
  if len(x.all) == 1 {
    x.Clr()
    return a
  }
  if x.int + 1 == len(x.all) {
    x.all = x.all[:x.int]
    x.int--
  } else {
    x.all = append (x.all[:x.int], x.all[x.int+1:]...)
  }
  return a
}

func (x *set) ExGeq (a Any) bool {
  for i, b := range x.all {
    if Leq (a, b) {
      x.int = i
      return true
    }
  }
  return false
}

func (x *set) Trav (op Op) {
  for _, a := range x.all {
    op (a)
  }
}

func (x *set) Join (Y Collector) {
  y, ok := Y.(*set)
  if ! ok { TypeNotEqPanic (x, Y) }
  x.check (y.Any)
  for _, a := range y.all {
    x.Ins (a)
  }
}

func (x *set) Ordered() bool {
  n := len(x.all)
  if n < 2 {
    return true
  }
  for i := 2; i < n; i++ {
    if Less (x.all[i], x.all[i-1]) {
      return false
    }
  }
  return true
}

func (x *set) less (i, j int) bool {
  return Less (x.all[i], x.all[j])
}

func (x *set) Sort() {
  sort.Slice (x.all, x.less)
}
