package piset

// (c) Christian Maurer   v. 190805 - license see µU.go

import (
  . "µU/obj"
  "µU/ker"
  "µU/str"
  "µU/pseq"
  "µU/buf"
  "µU/set"
  "µU/piset/internal"
)
const
  suffix = "seq"
type
  persistentIndexedSet struct {
                              Object
                              pseq.PersistentSequence "file"
                              internal.Index
                              Func "index func"
                              set.Set "tree"
                              buf.Buffer "position pool"
                              }

func new_(o Object, f Func) PersistentIndexedSet {
  x := new (persistentIndexedSet)
  x.Object = o.Clone().(Object)
  x.PersistentSequence = pseq.New (x.Object)
  x.Func, x.Index = f, internal.New (f (o))
  x.Set = set.New (x.Index)
  x.Buffer = buf.New (uint(0))
  return x
}

func (x *persistentIndexedSet) imp (Y Any) *persistentIndexedSet {
  y, ok := Y.(*persistentIndexedSet)
  if ! ok { TypeNotEqPanic (x, Y) }
  CheckTypeEq (x.Object, y.Object)
  return y
}

func (x *persistentIndexedSet) object (Y Any) Object {
  y, ok := Y.(Object)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *persistentIndexedSet) Fin() {
  x.PersistentSequence.Fin()
}

func (x *persistentIndexedSet) Offc() bool {
  return x.Empty()
}

func (x *persistentIndexedSet) build() {
  x.Set.Clr()
  i := uint(0)
  x.Buffer = buf.New(i)
  if x.PersistentSequence.Empty() { return }
  x.PersistentSequence.Trav (func (a Any) {
    x.Object = a.(Object).Clone().(Object)
//    e, ok := x.Object.(Editor); if ok { e.Write (0, 0) }
    if x.Object.Empty() {
      x.Buffer.Ins (i)
    } else {
      x.Index.Set (x.Func (x.Object), i)
      x.Set.Ins (x.Index)
    }
    i ++
  })
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

func (x *persistentIndexedSet) Less (Y Object) bool {
  return x.Set.Less (x.imp (Y).Set)
}

func (x *persistentIndexedSet) Num() uint {
  return x.Set.Num()
}

func (x *persistentIndexedSet) NumPred (p Pred) uint {
  n := uint(0)
  x.Trav (PredOp2Op(p, func (a Any) { n++ }))
  return n
}

func (x *persistentIndexedSet) Ex (a Any) bool {
  x.Index.Set (x.Func (a), 0)
  return x.Set.Ex (x.Index)
}

func (x *persistentIndexedSet) Ins (a Any) {
  object := x.object (a)
  if object.Empty() { return }
  if x.Ex (object) { return }
  var p uint
  if x.Buffer.Num() == 0 {
    p = x.PersistentSequence.Num()
  } else {
    p = x.Buffer.Get().(uint)
  }
  x.Index.Set (x.Func (object), p)
  x.PersistentSequence.Seek (p)
  x.PersistentSequence.Put (object)
  x.Set.Ins (x.Index)
  x.PersistentSequence.Seek (p)
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
  x.Index = x.Set.Get().(internal.Index)
  x.PersistentSequence.Seek (x.Index.Pos())
  return x.PersistentSequence.Get().(Object)
}

func (x *persistentIndexedSet) Put (a Any) {
  if x.Set.Empty() { return }
  x.Index.Set (x.Func (a), x.Index.Pos())
  x.Set.Put (x.Index)
  x.PersistentSequence.Put (a)
}

func (x *persistentIndexedSet) Del() Any {
  if x.Set.Empty() {
    x.Object.Clr()
    return x.Object
  }
  x.Index = x.Set.Get().(internal.Index)
  x.PersistentSequence.Seek (x.Index.Pos())
  x.Object = x.PersistentSequence.Get().(Object)
  object := x.Object.Clone().(Object)
  object.Clr()
  x.PersistentSequence.Put (object)
  x.Set.Del()
  x.Buffer.Ins (x.Index.Pos())
/*
  if x.Set.Empty() {
  } else {
    x.Index = x.Set.Get().(internal.Index)
    x.PersistentSequence.Seek (x.Index.Pos())
  }
*/
  return x.Object.Clone()
}

func (x *persistentIndexedSet) Sort() {
  if x == nil { ker.Panic ("piset.Sort: x == nil") }
  if x.Set == nil { ker.Panic ("piset.Sort: x.Set == nil") }
  x.Set.Sort()
  x.Jump (false)
}

func (x *persistentIndexedSet) Ordered() bool {
  ordered := true
  if x.Set.Empty() { return ordered }
  first := true
  var former Object
  x.Set.Trav (func (a Any) {
    x.PersistentSequence.Seek (a.(internal.Index).Pos())
    x.Object = x.PersistentSequence.Get().(Object)
    if first {
      first = false
    } else {
      if ! Leq (former, x.Object) {
        ordered = false
      }
    }
    former = x.Object
  })
  return ordered
}

func (x *persistentIndexedSet) ExGeq (a Any) bool {
  x.Index.Set (x.Func (a), 0)
  return x.Set.ExGeq (x.Index)
}

func (x *persistentIndexedSet) All (p Pred) bool {
  if x.Set.Empty() { return true }
  defer x.Jump (false)
  x.Set.Jump (false)
  for i := uint(0); i < x.Set.Num(); i++ {
    x.Index = x.Set.Get().(internal.Index)
    x.PersistentSequence.Seek (x.Index.Pos())
    if ! p (x.PersistentSequence.Get().(Object)) {
      return false
    }
    x.Set.Step (true)
  }
  return true
}

func (x *persistentIndexedSet) ExPred (p Pred, f bool) bool {
  x.Set.Jump (! f)
  for i := uint(0); i < x.Set.Num(); i++ {
    x.Index = x.Set.Get().(internal.Index)
    x.PersistentSequence.Seek (x.Index.Pos())
    if p (x.PersistentSequence.Get().(Object)) {
      return true
    }
    x.Set.Step (f)
  }
  return false
}

func (x *persistentIndexedSet) StepPred (p Pred, f bool) bool {
  if x.Set.Eoc (f) { return false }
  x.Set.Step (f)
  for {
    x.Index = x.Set.Get().(internal.Index)
    x.PersistentSequence.Seek (x.Index.Pos())
    if p (x.PersistentSequence.Get().(Object)) {
      return true
    }
    if x.Set.Eoc (f) { break }
    x.Set.Step (f)
  }
  return false
}

func (x *persistentIndexedSet) Trav (op Op) {
  if x.Set.Empty() { return }
  first := true
  var former Object
  x.Set.Trav (func (a Any) {
    x.PersistentSequence.Seek (a.(internal.Index).Pos())
    x.Object = x.PersistentSequence.Get().(Object)
    old := x.Object.Clone().(Object)
    op (x.Object)
    if ! x.Object.Eq (old) { x.PersistentSequence.Put (x.Object) }
    if first {
      first = false
    } else {
      if ! Leq (former, x.Object) { ker.Panic ("Pre of Trav not met: op is not monotone") }
    }
    former = x.Object.Clone().(Object)
  })
  x.Jump (false)
}

func (x *persistentIndexedSet) Filter (Y Iterator, p Pred) {
  if x.Set.Empty() { return }
  y := x.imp (Y)
  y.Clr()
  x.Set.Trav (func (a Any) {
    x.PersistentSequence.Seek (a.(internal.Index).Pos())
    x.Object = x.PersistentSequence.Get().(Object)
    if ! x.Object.Empty() && p (x.Object) {
      y.Ins (x.Object)
    }
  })
  y.Jump (false)
}

func (x *persistentIndexedSet) Cut (Y Iterator, p Pred) {
  if x.Set.Empty() { return }
  y := x.imp (Y)
  y.Clr()
  x.Set.Trav (func (a Any) {
    x.PersistentSequence.Seek (a.(internal.Index).Pos())
    x.Object = x.PersistentSequence.Get().(Object)
    if p (x.Object) {
      y.Ins (x.Object)
      x.Object.Clr()
      x.PersistentSequence.Put (x.Object)
    }
  })
  x.build()
  y.Jump (false)
}

func (x *persistentIndexedSet) ClrPred (p Pred) {
  if x.Set.Empty() { return }
  x.Set.Trav (func (a Any) {
    x.PersistentSequence.Seek (a.(internal.Index).Pos())
    x.Object = x.PersistentSequence.Get().(Object)
    if p (x.Object) {
      object := x.Object.Clone().(Object)
      object.Clr()
      x.PersistentSequence.Put (object)
    }
  })
  x.build()
}

func (x *persistentIndexedSet) Split (Y Iterator) {
  y := x.imp(Y)
  y.Clr()
  for {
    i := x.Set.Get().(internal.Index).Pos()
    x.PersistentSequence.Seek(i)
    x.Object = x.PersistentSequence.Get().(Object)
    y.Ins (x.Object)
    x.Object.Clr()
    x.PersistentSequence.Put (x.Object)
    if x.Set.Eoc(true) { break }
    x.Set.Step(true)
  }
  x.build()
  y.Jump(false)
}

func (x *persistentIndexedSet) Join (Y Iterator) {
  y := x.imp (Y)
  if y.Set.Empty() { return }
  y.Set.Trav (func (a Any) {
    y.PersistentSequence.Seek (a.(internal.Index).Pos())
    y.Object = y.PersistentSequence.Get().(Object)
    if ! y.Object.Empty() {
      x.Ins (y.Object)
    }
  })
  x.Jump (false)
  y.Clr()
}
