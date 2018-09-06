package pseq

// (c) Christian Maurer   v. 170509 - license see µU.go

// >>> still a lot of things TODO

import (
  "io"
// "reflect"
  . "µU/ker"
  . "µU/obj"
  "µU/str"
//  "µU/errh"
//  "µU/seq"
  "µU/pseq/internal"
)

const (
  null = uint64(0)
  one  = uint64(1)
)
type
  persistentSequence struct {
       name, tmpName string
         emptyObject Any
                     Any
                file internal.File
        owner, group uint
      size, pos, num uint64
           buf, buf1 []byte
             ordered bool
                     }
//var
//  filenames seq.Sequence
/*
  Among others the following problems are not yet solved:
  1. Not more than 1 psequence must be (re-)named with the same name.
     Help: Put names at (re-)defining into the sequence "filenames" and
           remove them at terminating.
  2. Access to psequences is only possible, if the rights are named correspondingly.
     At the moment clients are not protected from trying to access persistent Sequences
     without having the rights to.
  3. The following trivial handling of read/write-errors should be replaced by an error-based concept.
*/

var
  wasRead, wasWritten int

func (x *persistentSequence) read (bs []byte) {
  r, _ := x.file.Read (bs[:x.size])
  wasRead = r
}

func (x *persistentSequence) write (bs []byte) {
  w, _ := x.file.Write (bs[:x.size])
  if uint64(w) < x.size {
    wasWritten = w
  }
}

func (x *persistentSequence) check (a Any) {
  CheckTypeEq (x.emptyObject, a)
}

func new_(a Any) PersistentSequence {
  switch a.(type) {
  case Equaler, Coder:
    ; // ok
  default:
    if Atomic(a) {
      ; // ok
    } else {
      panic("not Atomic or Equaler and Coder")
    }
  }
  x := new(persistentSequence)
  x.emptyObject = Clone(a)
  x.Any = Clone(a)
  x.num = null
  x.size = uint64(Codelen(a))
  x.file = internal.New()
  x.buf = make ([]byte, x.size)
  x.buf1 = make ([]byte, x.size)
  x.ordered = false
  return x
}

func (x *persistentSequence) imp (a Any) *persistentSequence {
  y, ok := a.(*persistentSequence)
  x.check (y.emptyObject)
  if ! ok || x.size != y.size {
    TypeNotEqPanic (x, a)
  }
  if x.file == nil || y.file == nil { Panic ("pseq-error: file = nil") }
  if x == y { Panic ("pseq error: x == y") }
  return y
}

func (x *persistentSequence) Fin() {
///*
//  n := str.Length (x.name)
//  if filenames.Ex (x.name, n) {
//    filenames.Del()
//  } else {
//    Fehler
//  }
///*
  x.file.Fin()
}

func Length (n string) uint { // < -- uint64 !
  return uint(internal.DirectLength (n))
}

func Erase (n string) {
  internal.Erase (n)
}

/*
func accessible (Name string, Zugriff Zugriffe) bool {
  return file.accessible (Name, VAL (Zugriffe, ORD (Zugriff)))
}
*/

func (x *persistentSequence) Name (N string) {
//  if ! files.Defined (N) { Fehler }
  x.name = N
//  str.DelSpaces (&x.name)
//  n := str.Length (x.name)
//  if filenames.Ex (x.name, n) {
//    // Fehlersituation, siehe oben Bemerkung 1.
//    Fehler
//  }
  x.file.Name (x.name)
//    $USER 
  x.pos = 0
  x.num = x.file.Length() / x.size
  x.tmpName = x.name + "-tmp"
//  tmpName.Temporieren()
}

func (x *persistentSequence) Rename (n string) {
  if str.Empty (n) || n == x.name {
    return
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

func equal (as, bs []byte) bool {
  if len (as) != len (bs) { return false }
  for i, a := range (as) {
    if a != bs[i] { return false }
  }
  return true
}

func (x *persistentSequence) e (y *persistentSequence, r Rel) bool {
  if y.name == x.name { return true }
  if x.num != y.num { return false }
  for i := null; i < x.num; i++ {
    x.read (x.buf)
    y.read (y.buf)
    if ! r (x.buf, y.buf) {
      return false
    }
  }
  return true
}

func (x *persistentSequence) Eq (Y Any) bool {
  return x.e (x.imp (Y), Eq)
}

func (x *persistentSequence) cp (y *persistentSequence) {
// alpha Version
  x.file.Clr()
  x.file = internal.New()
  _, err := io.Copy (x.file, y.file)
  if err != nil { Panic ("pseq.cp.bug: io.Copy did not work") }
  x.emptyObject = Clone (y.emptyObject)
  x.Any = Clone (y.Any)
  x.buf = make ([]byte, x.size)
  x.buf1 = make ([]byte, x.size)
  x.ordered = y.ordered
  x.num = y.num
  x.Name (y.name + ".Copy")
}

func (x *persistentSequence) Copy (Y Any) {
  y := x.imp(Y)
  x.cp (y)
}

func (x *persistentSequence) Clone () Any {
  y := new_(Clone(x.emptyObject))
  y.Copy(x)
  return y
}

func (x *persistentSequence) leq (Y Any) bool { // TODO
  y := x.imp (Y)
  if y.name == x.name { return true }
  if x.num != y.num { return false }
  for i := null; i < x.num; i++ {
    x.read (x.buf)
/*
    for x.pos < x.num {
      for {
        Gx1.read (x1.buf)
        if equal (x.buf, x1. buf)
          continue
        } else {
        }
      }
    if ! equal (x.buf, x1.buf) {
      return false
    }
*/
  }
  return true
}

func (x *persistentSequence) Less (Y Object) bool { // TODO
  y := x.imp (Y)
  if y.name == x.name { return false }
  if x.num == y.num { return false }
  return x.leq (Y)
}

func (x *persistentSequence) Num() uint {
  return uint(x.file.Length() / x.size)
//  return uint(x.num)
}

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

func (x *persistentSequence) Ex (a Any) bool {
  x.check (a)
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
    x.pos --
  }
}

func (x *persistentSequence) Seek (n uint) {
  x.pos = uint64(n)
}

func (x *persistentSequence) Jump (forward bool) {
  if forward {
    x.Seek (uint(x.num))
  } else {
    x.Seek (0)
  }
}

func (x *persistentSequence) Offc() bool {
  return x.pos * x.size == x.file.Length()
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

func (x *persistentSequence) Get() Any {
  x.file.Seek (x.pos * x.size)
  if x.file.Position() != x.pos * x.size { Panic1 ("pseq", 10000000 + uint(x.pos)) }
  x.read (x.buf)
  return Clone (Decode (Clone (x.Any), x.buf))
}

func (x *persistentSequence) Put (a Any) {
  x.check (a)
  x.file.Seek (x.pos * x.size)
  x.write (Encode (a))
  x.num = x.file.Length() / x.size
}

func (x *persistentSequence) insert (a Any) {
  if x.pos >= x.num {
    x.pos = x.num
    x.file.Seek (x.file.Length())
    x.write (Encode (a))
    x.pos++
    x.num++
    return
  }
// x.pos < x.num:
  x1 := new_(x.emptyObject).(*persistentSequence)
  x1.Name (x.tmpName)
  x1.Clr()
  x.file.Seek (0)
  if x.pos > 0 {
    for i := null; i < x.pos; i++ {
      x.read (x.buf)
      x1.write (x.buf)
    }
  }
  x1.write (Encode (a))
  if x.pos < x.num {
    for i := x.pos; i < x.num; i++ {
      x.read (x.buf)
      x1.write (x.buf)
    }
  }
  x.pos++
  x.num++
  n := x.num
  p := x.pos
  x.file.Clr()
  x1.file.Rename (x.name)
  x1.file.Fin()
  x1.Fin()
  x.file.Name (x.name)
  x.pos = p
  x.num = n // == x.file.length() / x.size
}

func (x *persistentSequence) insertOrd (a Any) {
  ps := new_(x.emptyObject).(*persistentSequence)
  ps.Name (x.tmpName)
  ps.Clr()
  x.file.Seek (0)
  i := null
  n := x.num
  inserted := false
  p := null
  code := Encode (a)
  for {
    if i == x.num {
      if ! inserted {
        p = i
        ps.write (code)
      }
      break
    }
    x.read (x.buf)
    if ! inserted {
      if Less (code, x.buf) {
        p = i
        ps.write (code)
        inserted = true
      }
      if ! inserted {
        if ! Less (x.buf, code) {
          inserted = true
        }
      }
    }
    ps.write (x.buf)
    i++
  }
  x.file.Clr()
  ps.file.Rename (x.name)
  ps.file.Fin()
  ps.Fin()
  x.file.Name (x.name)
  x.num = x.file.Length() / x.size
  if x.num != n + 1 {
    // noch untersuchen
  }
  x.pos = p + 1
}

func (x *persistentSequence) Ins (a Any) {
  x.check (a)
  if x.ordered {
    x.insertOrd (a)
  } else {
    x.insert (a)
  }
  if x.num != uint64(x.Num()) { Panic ("pseq.Ins: num bug") }
}

func (x *persistentSequence) Del() Any {
  if x.num == 0 || x.pos >= x.num {
    return nil
  }
  n := x.num
  x1 := new_(x.emptyObject).(*persistentSequence)
  x1.Name (x.tmpName)
  x1.Clr()
  x.file.Seek (0)
  var a Any
  for i := null; i < x.num; i++ {
    x.read (x.buf)
    if i == x.pos {
//      a = Decode (x.Any, x.buf)
      a = Decode (Clone (x.Any), x.buf)
    } else {
      x1.write (x.buf)
    }
  }
  if x.pos == x.num - 1 && x.pos > 0 {
    x.pos --
  }
  p := x.pos
  x.file.Clr()
  x1.file.Rename (x.name)
  x1.file.Fin()
  x1.Fin()
  x.file.Name (x.name)
  x.pos = p
  x.num = x.file.Length() / x.size // x.num --
  if x.num + 1 != n {
// errh.Error2 ("what to devil", uint(x.num + 1), "is here loose", uint(n))
  }
  if x.num != uint64(x.Num()) { Panic ("pseq.Del: num bug") }
  return a
}

func (x *persistentSequence) ExPred (p Pred, f bool) bool {
  if x.file.Empty() { return false }
  n := x.file.Length() / x.size
  if n == 0 { return false }
  i := null
  if f {
    i = 0
  } else {
    i = n - 1
  }
  x.file.Seek (i * x.size)
  for {
    x.read (x.buf)
//    if p (Decode (x.emptyObject, x.buf)) {
    if p (Decode (Clone (x.emptyObject), x.buf)) {
      x.pos = i
      return true
    }
    if f {
      if i == n - 1 {
        break
      } else {
        i++
      }
    } else if i == 0 {
      break
    } else {
      i --
    }
  }
  return false
}

func (x *persistentSequence) StepPred (p Pred, f bool) bool {
  n := x.file.Length() / x.size
  if n <= 1 { return false }
  if f && x.pos == n - 1 { return false }
  if ! f && x.pos == 0 { return false }
  i := null
  if x.pos == n {
    if f {
      i = 0
    } else {
      i = n - 1
    }
  } else {
    i = x.pos
    if f {
      i++
    } else {
      i --
    }
  }
  for {
    x.file.Seek (i * x.size)
    x.read (x.buf)
//    if p (Decode (x.emptyObject, x.buf)) {
    if p (Decode (Clone (x.emptyObject), x.buf)) {
      x.pos = i
      break
    }
    if f {
      if i == n - 1 {
        break
      } else {
        i++
      }
    } else {
      if i == 0 {
        break
      } else {
        i --
      }
    }
  }
  return false
}

func (x *persistentSequence) All (p Pred) bool {
  if x.num == 0 { return true }
  x.file.Seek (0)
  for i := null; i < x.num; i++ {
    x.read (x.buf)
//    if ! p (Decode (x.emptyObject, x.buf)) {
    if ! p (Decode (Clone (x.emptyObject), x.buf)) {
      return false
    }
  }
  return true
}

func (x *persistentSequence) Ordered() bool {
  if x.num <= 1 { return true }
  x.file.Seek (0)
  x.read (x.buf)
  for i := one; i < x.num; i++ {
    x.read (x.buf1)
    if Less (x.buf1, x.buf) {
      return false
    }
    copy (x.buf, x.buf1)
    i++
  }
  return true
}

func (x *persistentSequence) Sort() {
// TODO
}

func (x *persistentSequence) Trav (op Op) {
  if x.num != uint64(x.Num()) { Panic ("pseq.Trav: num bug") }
  b := x.file.Length() == 0
  if b { if x.num != 0 || ! x.Empty() { println ("pseq.Trav: oops") } }
  x.file.Seek (0)
  for i := null; i < x.num; i++ {
    x.read (x.buf)
    if uint64(wasRead) < x.size {
      copy (x.buf, Encode (x.emptyObject)) // provisorisch
    }
//    x.Any = Decode (x.emptyObject, x.buf)
    x.Any = Decode (Clone (x.emptyObject), x.buf)
    op (x.Any)
    if ! equal (x.buf, Encode (x.Any)) {
      copy (x.buf, Encode (x.Any))
      x.file.Seek (i * x.size)
      x.write (x.buf)
      x.file.Seek (i * x.size)
    }
  }
  x.file.Fin()
}

func (x *persistentSequence) Filter (Y Iterator, p Pred) {
  y := x.imp (Y)
  if y == nil { return }
  if x.num == 0 { return }
  x.file.Seek (0)
  y.Clr()
  y.pos = 0
  for i := null; i < x.num; i++ {
    x.read (x.buf)
    if p (Decode (Clone (x.emptyObject), x.buf)) {
      y.write (x.buf)
      y.pos++
    }
  }
  y.file.Fin()
  if x.num != uint64(x.Num()) { Panic ("pseq.Filter: x.num bug") }
  if y.num != uint64(y.Num()) { Panic ("pseq.Filter: y.num bug") }
}

func (x *persistentSequence) Cut (Y Iterator, p Pred) {
  y := x.imp (Y)
  if y == nil { return }
  y.Clr()
  if x.name == y.name { return }
  x2 := new_(x.emptyObject).(*persistentSequence)
  x2.Name (x.tmpName)
  x2.Clr()
  x.file.Seek (0)
  x.pos = 0
  for i := null; i < x.num; i++ {
    x.read (x.buf)
    if p (Decode (Clone (x.emptyObject), x.buf)) {
      y.write (x.buf)
      y.pos++
    } else {
      x2.write (x.buf)
      x.pos++
    }
  }
  x.file.Clr()
  x2.file.Rename (x.name)
  x2.file.Fin()
  x2.Fin()
  x.file.Name (x.name)
  y.file.Fin()
  if x.num != uint64(x.Num()) { Panic ("pseq.Cut: x.num bug") }
  if y.num != uint64(y.Num()) { Panic ("pseq.Cut: y.num bug") }
}

func (x *persistentSequence) ClrPred (p Pred) {
  y := new_(x.emptyObject).(*persistentSequence)
  if y == nil { return }
  if x.num == 0 { return }
  x.file.Seek (0)
  n := x.pos
  y.Clr()
  for i := null; i < x.num; i++ {
    x.read (x.buf)
    if p (Decode (Clone (x.emptyObject), x.buf)) {
      if n == i {
        n++
      }
    } else {
      y.write (x.buf)
      y.num++
    }
  }
  y.file.Fin()
  y.Fin()
  y.pos = n
  x.file.Name (x.name)
}

func (x *persistentSequence) Split (Y Iterator) {
  y := x.imp (Y)
  if y == nil { return }
  if x.num == 0 { return }
  y.Clr()
  ps := new_(x.emptyObject).(*persistentSequence)
  ps.Name (x.tmpName)
  ps.Clr()
  x.file.Seek (0)
  if x.pos == 0 {
//    errh.ReportError ("pseq: Split not yet completely implemented") // >>>> alles nach S1
  } else {
    for i := null; i < x.pos; i++ {
      x.read (x.buf)
      ps.write (x.buf)
    }
    if x.pos < x.num {
      for i := one; i <= x.num - x.pos; i++ {
        x.read (x.buf)
        y.write (x.buf)
      }
    }
    y.pos = 0
  }
  x.file.Clr()
  ps.file.Rename (x.name)
  ps.Fin()
  x.file.Name (x.name)
  x.pos = x.num - x.pos - 1
  x.num = x.file.Length() / x.size
  y.num = y.file.Length() / y.size
  y.file.Fin()
}

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
  ps := new_(x.emptyObject).(*persistentSequence)
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

func (x *persistentSequence) Join (Y Iterator) {
  y := x.imp (Y)
  if y == nil { return }
  if x.ordered {
    x.join (y)
  } else {
    x.concatenate (y)
  }
}

/* func init() {
  filenames = seq.New (string)
} */
