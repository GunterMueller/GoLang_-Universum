package piset

// (c) Christian Maurer   v. 210221 - license see µU.go

import (
  . "µU/obj"
  "µU/ker"
  "µU/errh"
  "µU/str"
  "µU/pseq"
  "µU/buf"
  "µU/set"
  "µU/piset/internal"
)
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

func (x *persistentIndexedSet) Fin() {
  x.PersistentSequence.Fin()
}

func (x *persistentIndexedSet) Offc() bool {
  return x.Empty()
}

func (x *persistentIndexedSet) build() {
  x.Set.Clr()
  i := uint(0)
  x.Buffer = buf.New (i)
  if x.PersistentSequence.Empty() { return }
  x.PersistentSequence.Trav (func (a Any) {
    x.Object = a.(Object).Clone().(Object)
    if x.Object.Empty() {
      x.Buffer.Ins (i)
    } else {
      x.Index.Set (x.Func (x.Object), i)
      index := Clone (x.Index).(internal.Index)
      if index.Pos() != i { errh.Error2 ("i ==", i, "   Pos ==", index.Pos()) }
      x.Set.Ins (index)
    }
    i++
  })
  x.Jump (false)
}

const suffix = "seq"

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
  x.Index.Set (x.Func (a), 0)
  ex := x.Set.Ex (x.Index)
  return ex
}

func (x *persistentIndexedSet) Ins (a Any) {
  if ! IsObject (a) { ker.Panic("piset.Ins: geht nich") }
  if x.Ex (a) || a.(Object).Empty() { return }
  var p uint
  if x.Buffer.Empty() {
    p = x.PersistentSequence.Num()
  } else {
    p = x.Buffer.Get().(uint)
  }
  x.Index.Set (x.Func (a), p)
  x.PersistentSequence.Seek (p)
  x.PersistentSequence.Put (a)
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
  if x.Set.Empty() {
    return
  }
  if ! IsObject (a) { ker.Panic("piset.Put: geht nich") }
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
  x.Buffer.Ins (x.Index.Pos())
/*/
  if ! x.Set.Empty() {
    x.Index = x.Set.Get().(internal.Index)
    x.PersistentSequence.Seek (x.Index.Pos())
  }
/*/
  return x.Object.Clone()
}

func (x *persistentIndexedSet) ExGeq (a Any) bool {
  x.Index.Set (x.Func (a), 0)
  return x.Set.ExGeq (x.Index)
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

func (x *persistentIndexedSet) Ordered() bool {
  return x.Set.Ordered()
}

func (x *persistentIndexedSet) Sort() {
  x.Set.Sort()
}

