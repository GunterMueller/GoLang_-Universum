package pseq

// (c) Christian Maurer   v. 220715 - license see µU.go
//
//  >  The following problem is not yet solved: Access to psequences
//     is only possible, if the rights are set correspondingly.
//     At the moment clients are not protected from trying to
//     access persistent Sequences without having the rights to.

import (
  "reflect"
  "sort"
  . "µU/ker"
  . "µU/obj"
  "µU/str"
  "µU/errh"
  "µU/pseq/internal"
)

type
  persistentSequence struct {
       name, tmpName string
             ordered bool
                     any "pattern object"
              object any
                file internal.File
        owner, group,
      size, pos, num uint
           buf, buf1 Stream
                     }
const (
  null = uint(0)
  one  = uint(1)
)
var
  wasRead, wasWritten int

func (x *persistentSequence) read (bs Stream) {
  r, _ := x.file.Read (bs[:x.size])
  wasRead = r
}

func (x *persistentSequence) write (bs Stream) {
  w, _ := x.file.Write (bs[:x.size])
  if uint(w) < x.size {
    wasWritten = w
  }
}

func new_(a any) PersistentSequence {
  switch a.(type) {
  case Equaler, Coder:
    // ok
  default:
    if Atomic (a) || Streamic (a) {
      // ok
    } else {
      Panic ("neither Atomic nor Streamic, but " + reflect.TypeOf(a).String())
    }
  }
  x := new(persistentSequence)
  x.any = Clone(a)
  x.object = Clone(a)
  x.num = null
  x.size = Codelen(a)
  x.file = internal.New()
  x.buf = make (Stream, x.size)
  x.buf1 = make (Stream, x.size)
  return x
}

func (x *persistentSequence) imp (a any) *persistentSequence {
  y, ok := a.(*persistentSequence)
  CheckTypeEq (x.any, y.any)
  if ! ok || x.size != y.size {
    TypeNotEqPanic (x, a)
  }
  if x.file == nil || y.file == nil { Panic ("pseq-error: file = nil") }
  if x == y { Panic ("pseq error: x == y") }
  return y
}

func (x *persistentSequence) Fin() {
  x.file.Fin()
}

func length (n string) uint {
  return uint(internal.DirectLength (n))
}

func Erase (n string) {
  internal.Erase (n)
}

func (x *persistentSequence) Name (n string) {
  if str.Empty (n) { Panic ("name is empty") }
  str.OffSpc (&n)
  x.name = n
  x.file.Name (x.name)
//    $USER 
  x.pos = 0
  if x.file.Empty() {
    x.num = 0
  } else {
    x.num = x.file.Length() / x.size
  }
  x.tmpName = x.name + "-tmp"
//  tmpName.Temporieren()
}

func (x *persistentSequence) Rename (n string) {
  if str.Empty (n) || n == x.name {
    return
  }
  f := new_(byte(0))
  f.Name (n)
  if ! f.Empty() {
    Panic ("a file with the name " + n + " already exists")
    f.Fin()
  }
  x.name = n
  x.file.Rename (x.name)
}

func (x *persistentSequence) Empty() bool {
  if str.Empty (x.name) { return true }
  return x.file.Empty()
}

func (x *persistentSequence) Clr() {
  if x.file == nil { Panic ("pseq.Clr: file = nil") }
  x.file.Clr()
  x.pos = 0
  x.num = 0
}

func equal (as, bs Stream) bool {
  if len (as) != len (bs) { return false }
  for i, a := range (as) {
    if a != bs[i] { return false }
  }
  return true
}

func (x *persistentSequence) Num() uint {
  if x.num != x.file.Length() / x.size { Panic ("num != Num") }
  return uint(x.num)
}

/*/
func (x *persistentSequence) NumPred (p Pred) uint {
  n := uint(0)
  if x.num == 0 { return 0 }
  x.file.Seek (0)
  for i := null; i < x.num; i++ {
    x.read (x.buf)
    if p (x.buf) {
      n++
    }
  }
  return n
}
/*/

func (x *persistentSequence) Ex (a any) bool {
  CheckTypeEq (x.any, a)
  if x.num == 0 { return false }
  x.file.Seek (0)
  for i := null; i < x.num; i++ {
    x.read (x.buf)
    if equal (x.buf, Encode (a)) {
      x.pos = i
      return true
    }
  }
  return false
}

func (x *persistentSequence) Step (forward bool) {
  if forward {
    if x.pos * x.size < x.file.Length() {
      x.pos++
    }
  } else if x.pos > 0 {
    x.pos--
  }
}

func (x *persistentSequence) Seek (n uint) {
  x.pos = n
}

func (x *persistentSequence) Jump (forward bool) {
  if forward {
    x.Seek (uint(x.num))
  } else {
    x.Seek (0)
  }
}

func (x *persistentSequence) Offc() bool {
  return x.Pos() == x.Num()
}

func (x *persistentSequence) Eoc (forward bool) bool {
  if forward {
    return (x.pos + 1) * x.size == x.file.Length()
  }
  return x.pos == 0
}

func (x *persistentSequence) Pos() uint {
  return uint(x.pos)
}

func (x *persistentSequence) Get() any {
  x.file.Seek (x.pos * x.size)
  x.read (x.buf)
//  return Clone (Decode (x.object, x.buf))
  return Decode (x.object, x.buf)
}

func (x *persistentSequence) Put (a any) {
  CheckTypeEq (x.any, a)
  x.file.Seek (x.pos * x.size)
  x.write (Encode (a))
  x.num = x.file.Length() / x.size
}

func (x *persistentSequence) ins (a any) {
  n, p := x.Num(), x.Pos()
  if p >= n {
    p = n
  } else {
    for j := n; j > p; j-- {
      x.Seek (j - 1)
      b := x.Get()
      x.Seek (j)
      x.Put (b)
    }
  }
  x.Seek (p)
  x.Put (a)
  x.Seek (p + 1)
}

func (x *persistentSequence) insOrd (a any) {
  inserted := false
  n := x.Num()
  for i := uint(0); i < n; i++ {
    x.Seek (i)
    b := x.Get()
    if Eq (a, b) {
      return
    }
    if Less (a, b) {
      for j := n; j > i; j-- {
        x.Seek (j - 1)
        b = x.Get()
        x.Seek (j)
        x.Put (b)
      }
      x.Seek (i)
      x.Put (a)
      inserted = true
      break
    }
  }
  if ! inserted {
    x.Seek (n)
    x.Put (a)
  }
}

func (x *persistentSequence) Ins (a any) {
  CheckTypeEq (x.any, a)
  if x.ordered {
    x.insOrd (a)
  } else {
    x.ins (a)
  }
}

func (x *persistentSequence) Del() any {
  if x.num == 0 || x.pos >= x.num {
    return nil
  }
  n := x.num
  y := new_(x.any).(*persistentSequence)
  y.Name (x.tmpName)
  y.Clr()
  x.file.Seek (0)
  var a any
  for i := null; i < x.num; i++ {
    x.read (x.buf)
    if i == x.pos {
      a = Decode (Clone (x.object), x.buf)
    } else {
      y.write (x.buf)
    }
  }
  if x.pos == x.num - 1 && x.pos > 0 {
    x.pos--
  }
  p := x.pos
  x.file.Clr()
  y.file.Rename (x.name)
  y.file.Fin()
  y.Fin()
  x.file.Name (x.name)
  x.pos = p
  x.num = x.file.Length() / x.size // x.num--
  if x.num + 1 != n {
    errh.Error2 ("what to devil", uint(x.num + 1), "is here loose", uint(n))
  }
  if x.num != x.Num() { Panic ("pseq.Del: num bug") }
  return a
}

func (x *persistentSequence) Ordered() bool {
  if x.Num() <= 1 { return true }
  x.file.Seek (0)
  x.read (x.buf1)
  for i := one; i < x.num; i++ {
    x.read (x.buf)
    if Less (x.buf, x.buf1) {
      return false
    }
    copy (x.buf1, x.buf)
  }
  return true
}

func (x *persistentSequence) Sort() {
  if x.ordered { return }
  n := x.Num()
  if n <= 1 { return }
  s := make([]any, 0)
  for i := uint(0); i < n; i++ {
    x.Seek(i)
    s = append (s, x.Get())
  }
  sort.Slice (s, func (k, n int) bool { return Less (s[k], s[n]) })
  for i := uint(0); i < n; i++ {
    x.Seek (i)
    x.Put (s[i])
  }
  x.ordered = true
}

func (x *persistentSequence) ExGeq (a any) bool {
  if ! x.ordered { Panic ("x is not ordered") }
  n := x.Num()
  if n == 0 { return false }
// XXX not efficient TODO binary search
/*/
  x.Seek (n/2)
  b := x.Get().(any)
  if a < b {
    search first half
  } else if Eq (a, b) {
    return true
  } else {
   search second half
  }
/*/
  for i := uint(0); i < n; i++ {
    x.Seek(i)
    if ! Less (x.Get(), a) {
      return true
    }
  }
  return false
}

func (x *persistentSequence) Trav (op Op) {
  if x.num != x.Num() { Panic ("pseq.Trav: num bug") }
  b := x.file.Length() == 0
  if b { if x.num != 0 || ! x.Empty() { Panic ("pseq.Trav: oops") } }
  x.file.Seek (0)
  for i := null; i < x.num; i++ {
    x.read (x.buf)
    if uint(wasRead) < x.size {
      copy (x.buf, Encode (x.any)) // provisorial
    }
    x.object = Decode (Clone (x.any), x.buf)
    op (x.object)
    if ! equal (x.buf, Encode (x.object)) {
      copy (x.buf, Encode (x.object))
      x.file.Seek (i * x.size)
      x.write (x.buf)
      x.file.Seek (i * x.size)
    }
  }
  x.file.Fin()
}

/*/
func (x *persistentSequence) Filter (Y Collector, p Pred) {
  y := x.imp (Y)
  if y == nil { return }
  if x.num == 0 { return }
  x.file.Seek (0)
  y.Clr()
  y.pos = 0
  for i := null; i < x.num; i++ {
    x.read (x.buf)
    if p (Decode (Clone (x.any), x.buf)) {
      y.write (x.buf)
      y.pos++
    }
  }
  y.file.Fin()
  if x.num != x.Num() { Panic ("pseq.Filter: x.num bug") }
  if y.num != y.Num() { Panic ("pseq.Filter: y.num bug") }
}

func (x *persistentSequence) Cut (Y Collector, p Pred) {
  y := x.imp (Y)
  if y == nil { return }
  y.Clr()
  if x.name == y.name { return }
  x1 := new_(x.any).(*persistentSequence)
  x1.Name (x.tmpName)
  x1.Clr()
  x.file.Seek (0)
  x.pos = 0
  for i := null; i < x.num; i++ {
    x.read (x.buf)
    if p (Decode (Clone (x.any), x.buf)) {
      y.write (x.buf)
      y.pos++
    } else {
      x1.write (x.buf)
      x.pos++
    }
  }
  x.file.Clr()
  x1.file.Rename (x.name)
  x1.file.Fin()
  x1.Fin()
  x.file.Name (x.name)
  y.file.Fin()
  if x.num != x.Num() { Panic ("pseq.Cut: x.num bug") }
  if y.num != y.Num() { Panic ("pseq.Cut: y.num bug") }
}
/*/

func (x *persistentSequence) concatenate (Y PersistentSequence) {
  y := x.imp (Y)
  if y.num == 0 { return }
/*
  if x.num == 0 {
    should be more effective: // TODO
    rename ...
    y.Name -> x.Name
  }
*/
  x.file.Seek (x.num * x.size)
  y.file.Seek (0)
  for i := null; i < y.num; i++ {
    y.read (x.buf)
    x.write (x.buf)
  }
  x.file.Fin()
  x.num = x.file.Length() / x.size
  y.Clr()
}

func (x *persistentSequence) join (Y PersistentSequence) {
  y := x.imp (Y)
  if y.num == 0 { return }
/*
  if x.num == 0 {
    more effective: see concatenate
  }
*/
  ps := new_(x.any).(*persistentSequence)
  ps.Name (x.tmpName)
  ps.Clr()
  x.file.Seek (0)
  y.file.Seek (0)
  y.read (y.buf)
  i, j := null, null
  if x.num > 0 {
    x.read (x.buf)
    for {
      if Less (x.buf, y.buf) {
        ps.write (x.buf)
        i++
        if i < x.num {
          x.read (x.buf)
        } else {
          break
        }
      } else {
       if Less (y.buf, x.buf) {
          ps.write (y.buf)
          j++
          if j < y.num {
            y.read (y.buf)
          } else {
            break
          }
        } else {
          ps.write (y.buf)
          i++
          if i < x.num {
            x.read (x.buf)
          }
          j++
          if j < y.num {
            y.read (y.buf)
          }
          if i == x.num || j == y.num {
            break
          }
        }
      }
    }
  }
  for {
    if i == x.num { break }
    ps.write (x.buf)
    i++
    if i < x.num {
      x.read (x.buf)
    }
  }
  for {
    if j == y.num { break }
    ps.write (y.buf)
    j++
    if j < y.num {
      y.read (y.buf)
    }
  }
  x.file.Clr()
  x.num = x.file.Length() / x.size
  y.Clr()
  ps.file.Rename (x.name)
  ps.Fin()
}

func (x *persistentSequence) Join (Y Collector) {
  y := x.imp (Y)
  if y == nil { return }
  if x.ordered {
    x.join (y)
  } else {
    x.concatenate (y)
  }
}
