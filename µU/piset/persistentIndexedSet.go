package piset

// (c) Christian Maurer   v. 210509 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/str"
  "µU/kbd"
  "µU/errh"
  "µU/pseq"
  "µU/buf"
  "µU/set"
  "µU/piset/pair"
)
type
  persistentIndexedSet struct {
                              Object "pattern object"
                              Any "index of Object"
                              pseq.PersistentSequence "file"
                              Func "index function"
                              set.Set "pairs of index and position in the file"
                              buf.Buffer "free positions in the file"
                              }
var
  help = []string {" vor-/rückwärts: Pfeiltaste auf-/abwärts",
                   "zum Anfang/Ende: Pos1/End               ",
                   " Eintrag ändern: Enter                  ",
                   "       einfügen: Einfg                  ",
                   "      entfernen: Entf                   ",
                   "       umordnen: F3                     ",
                   "   Programmende: Esc                    " }

func init() {
  for i, h := range (help) { help[i] = str.Lat1 (h) }
}

func new_(o Indexer) PersistentIndexedSet {
  if ! IsIndexer (o) { ker.Oops() }
  x := new (persistentIndexedSet)
  x.Object = o.Clone().(Object)
  x.Func = o.Index()
  x.Any = x.Func (o)
  x.PersistentSequence = pseq.New (x.Object)
  x.Set = set.New (pair.New (x.Any, 0))
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

func (x *persistentIndexedSet) Ex (a Any) bool {
  return x.Set.Ex (pair.New (x.Func (a), 0))
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

func (x *persistentIndexedSet) Get() Any {
  if x.Set.Empty() {
    x.Object.Clr()
    return x.Object
  }
  p := x.Set.Get().(pair.Pair)
  n := p.Pos()
  x.PersistentSequence.Seek (n)
  return x.PersistentSequence.Get().(Object)
}

func (x *persistentIndexedSet) Put (a Any) {
  if x.Set.Empty() {
    return
  }
  x.check (a)
  n := x.Set.Get().(pair.Pair).Pos()
  x.Set.Put (pair.New (x.Func (a), n))
  x.PersistentSequence.Put (a)
}

func (x *persistentIndexedSet) Del() Any {
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

func (x *persistentIndexedSet) ExGeq (a Any) bool {
  return x.Set.ExGeq (pair.New (x.Func (a), 0))
}

func (x *persistentIndexedSet) Trav (op Op) {
  if x.Set.Empty() { return }
  x.PersistentSequence.Jump (false)
  x.Set.Trav (func (a Any) {
    op (a)
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

func TeX (a Any) {
  o := a.(TeXer)
  tex += o.TeX()
}

func (x *persistentIndexedSet) Operate() {
  hint := "Hilfe: F1                  Ende: Esc"
  x.Jump (false)
  if x.Empty() {
    o := x.Object.(Indexer)
    o.Clr()
    for {
      o.(Editor).Edit (0, 0)
      if ! o.Empty() {
        x.Ins (o)
        break
      }
    }
  }
  loop:
  for {
    o := x.Get().(Indexer)
    o0 := Clone(o).(Indexer)
    o.(Editor).Write (0, 0)
    errh.Hint (hint)
    switch k, _ := kbd.Command(); k {
    case kbd.Enter:
      o.(Editor).Edit (0, 0)
      if o.Empty() {
        x.Del()
      } else {
        if ! Eq (o, o0) {
          x.Put (o)
        }
      }
    case kbd.Esc:
      break loop
    case kbd.Up, kbd.Left:
      x.Step (false)
    case kbd.Down, kbd.Right:
      x.Step (true)
    case kbd.Pos1:
      x.Jump (false)
    case kbd.End:
      x.Jump (true)
    case kbd.Ins:
      errh.DelHint()
      o.Clr()
      o.(Editor).Edit (0, 0)
      if ! o.Empty() {
        x.Ins (o)
      }
      errh.Hint (hint)
    case kbd.Del:
      if errh.Confirmed() {
        x.Del()
      }
    case kbd.Act:
      x.Object.(Rotator).Rotate()
      x.Sort()
    case kbd.Help:
      errh.Help (help)
      errh.Hint (hint)
    case kbd.Search:
      o.(Editor).Edit (0, 0)
    case kbd.Print:
      tex = ""
      x.Trav (TeX)
    }
    errh.DelHint()
  }
}
