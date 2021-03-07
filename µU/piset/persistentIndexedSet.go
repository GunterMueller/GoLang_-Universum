package piset

// (c) Christian Maurer   v. 210228 - license see µU.go

import (
  . "µU/obj"
  "µU/str"
  "µU/pseq"
  "µU/buf"
  "µU/set"
  "µU/piset/internal"
)
type
  persistentIndexedSet struct {
                              Object
                              Any "index of Object"
                              pseq.PersistentSequence "file"
                              Func "index function"
                              set.Set "tree"
                              buf.Buffer "position pool"
                              }
const
  suffix = "seq"

func pair (a Any, n uint) internal.Pair {
  return internal.New (a, n)
}

func new_(o Object, f Func) PersistentIndexedSet {
  x := new (persistentIndexedSet)
  x.Object = o.Clone().(Object)
  x.Any = f (o)
  x.PersistentSequence = pseq.New (x.Object)
  x.Func = f
  x.Set = set.New (pair (x.Any, 0))
  x.Buffer = buf.New (uint(0))
  return x
}

func (x *persistentIndexedSet) imp (Y Any) *persistentIndexedSet {
  y, ok := Y.(*persistentIndexedSet)
  if ! ok { TypeNotEqPanic (x, Y) }
  CheckTypeEq (x.Object, y.Object)
  return y
}

func (x *persistentIndexedSet) check (a Any) {
  CheckTypeEq (a, x.Object)
  CheckTypeEq (x.Any, x.Func (x.Object))
}

func (x *persistentIndexedSet) Fin() {
  x.PersistentSequence.Fin()
}

func (x *persistentIndexedSet) Offc() bool {
  return x.Empty()
}

func (x *persistentIndexedSet) build() {
  x.Set.Clr()
  x.Buffer = buf.New (uint(0))
  if x.PersistentSequence.Empty() { return }
  for i := uint(0); i < x.PersistentSequence.Num(); i++ {
    x.PersistentSequence.Seek (i)
    x.Object = x.PersistentSequence.Get().(Object)
    if x.Object.Empty() {
      x.Buffer.Ins (i)
    } else {
      any := x.Func (x.Object)
      x.Set.Ins (pair (any, i))
    }
  }
  x.Jump (false)
}

func (x *persistentIndexedSet) Name (s string) {
  if str.Empty (s) { return }
  x.PersistentSequence.Name (s + "." + suffix)
  x.build()
}

func (x *persistentIndexedSet) Rename (s string) {
  if str.Empty (s) { return }
  x.PersistentSequence.Rename (s + "." + suffix)
}

func (x *persistentIndexedSet) Empty() bool {
  return x.Set.Empty()
}

func (x *persistentIndexedSet) Clr() {
  x.PersistentSequence.Clr()
  x.Set.Clr()
  x.Buffer = buf.New (uint(0))
  x.Object.Clr()
}

func (x *persistentIndexedSet) Num() uint {
  return x.Set.Num()
}

func (x *persistentIndexedSet) Ex (a Any) bool {
  return x.Set.Ex (pair (x.Func (a), 0))
}

func (x *persistentIndexedSet) Ins (a Any) {
  x.check (a)
  if x.Ex (a) || a.(Object).Empty() { return }
  var n uint
  if x.Buffer.Empty() {
    n = x.PersistentSequence.Num()
  } else {
    n = x.Buffer.Get().(uint)
  }
  x.PersistentSequence.Seek (n)
  x.PersistentSequence.Put (a)
  x.Set.Ins (pair (x.Func (a), n))
  x.PersistentSequence.Seek (n)
}

func (x *persistentIndexedSet) Step (forward bool) {
  x.Set.Step (forward)
}

func (x *persistentIndexedSet) Jump (toEnd bool) {
  x.Set.Jump (toEnd)
}

func (x *persistentIndexedSet) Eoc (atEnd bool) bool {
  return x.Set.Eoc (atEnd)
}

func (x *persistentIndexedSet) Get() Any {
  if x.Set.Empty() {
    x.Object.Clr()
    return x.Object
  }
  p := x.Set.Get().(internal.Pair)
  n := p.Pos()
  x.PersistentSequence.Seek (n)
  return x.PersistentSequence.Get().(Object)
}

func (x *persistentIndexedSet) Put (a Any) {
  if x.Set.Empty() {
    return
  }
  x.check (a)
  n := x.Set.Get().(internal.Pair).Pos()
  x.Set.Put (pair (x.Func (a), n))
  x.PersistentSequence.Put (a)
}

func (x *persistentIndexedSet) Del() Any {
  if x.Set.Empty() {
    x.Object.Clr()
    return x.Object
  }
  n := x.Set.Get().(internal.Pair).Pos()
  x.PersistentSequence.Seek (n)
  x.Object = x.PersistentSequence.Get().(Object)
  object := x.Object.Clone().(Object)
  object.Clr()
  x.PersistentSequence.Put (object)
  x.Buffer.Ins (n)
  if ! x.Set.Empty() {
    n := x.Set.Get().(internal.Pair).Pos()
    x.PersistentSequence.Seek (n)
  }
  return x.Object.Clone()
}

func (x *persistentIndexedSet) ExGeq (a Any) bool {
  return x.Set.ExGeq (pair (x.Func (a), 0))
}

func (x *persistentIndexedSet) Trav (op Op) {
  if x.Set.Empty() { return }
  x.PersistentSequence.Jump (false)
  x.Set.Trav (func (a Any) {
    op (a)
    x.PersistentSequence.Put (a)
    x.PersistentSequence.Step (true)
  })
}

func (x *persistentIndexedSet) Join (Y Collector) {
  y := x.imp (Y)
  if y.Set.Empty() { return }
  for i := uint(0); i < y.PersistentSequence.Num(); i++ {
    y.PersistentSequence.Seek (i)
    y.Object = y.PersistentSequence.Get().(Object)
    if ! y.Object.Empty() {
      x.Ins (y.Object)
    }
  }
  x.Jump (false)
  y.Clr()
}

func (x *persistentIndexedSet) Ordered() bool {
  return x.Set.Ordered()
}

func (x *persistentIndexedSet) Sort() {
  x.Set.Sort()
}

