package piset

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/str"
  "µU/pseq"
  "µU/buf"
  "µU/set"
  "µU/piset/pair"
)
type
  persistentIndexedSet struct {
                              Object "pattern object"
                              any "index of Object"
                              pseq.PersistentSequence "file"
                              Func "index function"
                              set.Set "pairs of index and position in the file"
                              buf.Buffer "free positions in the file"
                              }

func new_(o Indexer) PersistentIndexedSet {
  if ! IsIndexer (o) { ker.Oops() }
  x := new (persistentIndexedSet)
  x.Object = o.Clone().(Object)
  x.Func = o.Index()
  x.any = x.Func (o)
  x.PersistentSequence = pseq.New (x.Object)
  x.Set = set.New (pair.New (x.any, 0))
  x.Buffer = buf.New (uint(0))
  return x
}

func (x *persistentIndexedSet) imp (Y any) *persistentIndexedSet {
  y, ok := Y.(*persistentIndexedSet)
  if ! ok { TypeNotEqPanic (x, Y) }
  CheckTypeEq (x.Object, y.Object)
  return y
}

func (x *persistentIndexedSet) check (a any) {
  CheckTypeEq (a, x.Object)
  CheckTypeEq (x.any, x.Func (x.Object))
}

func (x *persistentIndexedSet) Fin() {
  x.PersistentSequence.Fin()
}

func (x *persistentIndexedSet) Offc() bool {
  return x.Empty()
}

func (x *persistentIndexedSet) Name (s string) {
  if str.Empty (s) { return }
  x.PersistentSequence.Name (s + ".seq")
  x.Set.Clr()
  x.Buffer = buf.New (uint(0))
  if x.PersistentSequence.Empty() { return }
  for i := uint(0); i < x.PersistentSequence.Num(); i++ {
    x.PersistentSequence.Seek (i)
    x.Object = x.PersistentSequence.Get().(Object)
    if x.Object.Empty() {
      x.Buffer.Ins (i)
    } else {
      x.Set.Ins (pair.New (x.Func (x.Object), i))
    }
  }
  x.Jump (false)
}

func (x *persistentIndexedSet) Rename (s string) {
  if str.Empty (s) { return }
  x.PersistentSequence.Rename (s + ".seq")
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

func (x *persistentIndexedSet) Ex (a any) bool {
  return x.Set.Ex (pair.New (x.Func (a), 0))
}

func (x *persistentIndexedSet) ExPred (p Pred) bool {
  x.Jump (false)
  for {
    o := x.Get()
    if p (o) {
      return true
    }
    if x.Eoc (true) {
      return false
    }
    x.Step (true)
  }
  return false
}

func (x *persistentIndexedSet) Ins (a any) {
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
  x.Set.Ins (pair.New (x.Func (a), n))
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

func (x *persistentIndexedSet) Get() any {
  if x.Set.Empty() {
    x.Object.Clr()
    return x.Object
  }
  p := x.Set.Get().(pair.Pair)
  n := p.Pos()
  x.PersistentSequence.Seek (n)
  return x.PersistentSequence.Get().(Object)
}

func (x *persistentIndexedSet) Put (a any) {
  if x.Set.Empty() {
    return
  }
  x.check (a)
  n := x.Set.Get().(pair.Pair).Pos()
  x.Set.Put (pair.New (x.Func (a), n))
  x.PersistentSequence.Put (a)
}

func (x *persistentIndexedSet) Del() any {
  if x.Set.Empty() {
    x.Object.Clr()
    return x.Object
  }
  n := x.Set.Get().(pair.Pair).Pos()
  x.PersistentSequence.Seek (n)
  x.Object = x.PersistentSequence.Get().(Object)
  object := x.Object.Clone().(Object)
  object.Clr()
  x.PersistentSequence.Put (object)
  x.Buffer.Ins (n)
  if ! x.Set.Empty() {
    n := x.Set.Get().(pair.Pair).Pos()
    x.PersistentSequence.Seek (n)
  }
  return x.Object.Clone()
}

func (x *persistentIndexedSet) ExGeq (a any) bool {
  return x.Set.ExGeq (pair.New (x.Func (a), 0))
}

func (x *persistentIndexedSet) Trav (op Op) {
  if x.Set.Empty() { return }
  x.PersistentSequence.Jump (false)
  x.Set.Trav (func (a any) {
    p := a.(pair.Pair).Pos()
    x.PersistentSequence.Seek (p)
    x.Object = x.PersistentSequence.Get().(Object)
    object := x.Object.Clone().(Object)
    op (object)
/*/
    x.PersistentSequence.Put (a)
    x.PersistentSequence.Step (true)
/*/
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

var
  tex string

func texClr() {
  tex = ""
}

func teX (a any) {
  tex += a.(TeXer).TeX()
}

func (x *persistentIndexedSet) TeX() string {
  defer texClr()
  return tex
}
