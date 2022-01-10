package seq

// (c) Christian Maurer   v. 220106 - license see µU.go

import (
  "sync"
  . "µU/ker"
  . "µU/obj"
)
type (
  cell struct {
              Any "content of the cell"
         next,
         prev *cell
       }
  sequence struct {
              num,
              pos uint
           anchor,
           actual *cell
          ordered bool
                  }
)
var
  mutex sync.Mutex

func (x *sequence) check (a Any) {
  CheckTypeEq (x.anchor.Any, a)
}

func new_(a Any) Sequence {
  CheckAtomicOrObject(a)
  x := new(sequence)
  x.anchor = new(cell)
  x.anchor.Any = Clone(a)
  x.anchor.next, x.anchor.prev = x.anchor, x.anchor
  x.actual = x.anchor
  return x
}

func (x *sequence) imp (a Any) *sequence {
  y, ok := a.(*sequence)
  if ! ok { TypeNotEqPanic (x, a) }
  x.check (y.anchor.Any)
  return y
}

func checked (a Any) bool {
  return AtomicOrObject (a)
}

func (x *sequence) Empty() bool {
  return x.anchor.next == x.anchor
}

func (x *sequence) Clr() {
  x.anchor.next, x.anchor.prev = x.anchor, x.anchor
  x.actual = x.anchor
  x.num, x.pos = 0, 0
}

func (x *sequence) e (y *sequence, r Rel) bool {
  if x.num != y.num { return false }
  for c, d := x.anchor.next, y.anchor.next; c != x.anchor; c, d = c.next, c.next {
    if ! r (c.Any, d.Any) {
      return false
    }
  }
  return true
}

func (x *sequence) Eq (Y Any) bool {
  return x.e (x.imp (Y), Eq)
}

func (x *sequence) insert (a Any) {
  n := new (cell)
  n.Any = Clone (a)
  n.next, n.prev = x.actual, x.actual.prev
  x.actual.prev.next = n
  x.actual.prev = n
}

func (x *sequence) Copy (Y Any) {
  y := x.imp (Y)
  x.Clr()
  x.anchor.Any = Clone(y.anchor.Any)
  for c := y.anchor.next; c != y.anchor; c = c.next {
    x.insert (c.Any)
  }
  x.num, x.pos = y.num, y.num
}

func (x *sequence) Clone() Any {
  y := new_(Clone(x.anchor.Any))
  y.Copy (x)
  return y
}

func (x *sequence) Less (Y Any) bool {
  y := x.imp (Y)
  if x.num >= y.num { return false }
  c := x.anchor.next
  d := y.anchor.next
  for c != x.anchor {
    for {
      if d == y.anchor {
        return false
      }
      if Eq (c.Any, d.Any) {
        d = d.next
        break
      }
      d = d.next
    }
    c = c.next
  }
  return true
}

func (x *sequence) Codelen() uint {
  n := uint(4) // Codelen (uint32(0))
  for c := x.anchor.next; c != x.anchor; c = c.next {
    n += 4 + Codelen (c.Any)
  }
  return n
}

func (x *sequence) Encode() []byte {
  b := make ([]byte, x.Codelen())
  i, a := uint32(0), uint32(4)
  copy (b[i:a], Encode (uint32(x.num)))
  i += 4
  for l := x.anchor.next; l != x.anchor; l = l.next {
    n := uint32(Codelen (l.Any))
    copy (b[i:i+a], Encode (n))
    i += a
    copy (b[i:i+n], Encode (l.Any))
    i += n
  }
  return b
}

func (x *sequence) Decode (bs []byte) {
  x.Clr()
  i, a := uint32(0), uint32(4)
  x.num = uint(Decode (uint32(0), bs[i:a]).(uint32))
  i += a
  for j := uint32(0); j < uint32(x.num); j++ {
    n := Decode (uint32(0), bs[i:i+a]).(uint32)
    i += a
    x.insert (Decode (Clone (x.anchor.Any), bs[i:i+n]))
    i += n
  }
  return
}

func (x *sequence) Num() uint {
  return x.num
}

func (x *sequence) NumPred (p Pred) uint {
  n := uint(0);
  for c := x.anchor.next; c != x.anchor; c = c.next {
    if p (c.Any) {
      n++
    }
  }
  return n
}

func (x *sequence) Ex (a Any) bool {
  x.check (a)
  p := uint(0)
  for c := x.anchor.next; c != x.anchor; c = c.next {
    if Eq (c.Any, a) {
      x.actual = c
      x.pos = p
      return true
    }
    p++
  }
  return false
}

func (x *sequence) Step (forward bool) {
  if forward {
    if x.actual != x.anchor {
      x.actual = x.actual.next
      x.pos++
    }
  } else if x.actual != x.anchor.next {
    x.actual = x.actual.prev
    x.pos--
  }
}

func (x *sequence) Jump (forward bool) {
  if x.num > 0 {
    if forward {
      x.actual = x.anchor.prev
      x.pos = x.num - 1
    } else {
      x.actual = x.anchor.next
      x.pos = 0
    }
  }
}

func (x *sequence) Eoc (forward bool) bool {
  if x.actual == x.anchor {
    return false
  }
  if forward {
    return (x.actual.next == x.anchor)
  }
  return (x.actual.prev == x.anchor)
}

func (x *sequence) Offc() bool {
  return x.actual == x.anchor
}

func (x *sequence) Pos() uint {
  return x.pos
}

func (x *sequence) Seek (i uint) {
  if i == 0 {
    x.actual = x.anchor.next
    x.pos = 0
    return
  }
  if i >= x.num {
    x.actual = x.anchor
    x.pos = x.num
    return
  }
  if i == 1 {
    x.actual = x.anchor.next.next
    x.pos = 1
    return
  }
  if i + 1 == x.num {
    x.actual = x.anchor.prev
    x.pos = x.num - 1
    return
  }
  for x.pos < i {
    x.actual = x.actual.next
    x.pos++
  }
  for x.pos > i {
    x.actual = x.actual.prev
    x.pos--
  }
}

func (x *sequence) Get() Any {
  if x.actual == x.anchor {
    return nil
  }
  return Clone (x.actual.Any)
}

func (x *sequence) Put (a Any) {
  x.check (a)
  if x.actual == x.anchor {
    x.insert (a)
    x.actual = x.actual.prev
    x.pos = x.num
    x.num++
  } else {
    x.actual.Any = Clone (a)
  }
}

func (x *sequence) Ins (a Any) {
  x.check (a)
  if x.ordered {
    x.actual = x.anchor.next
    x.pos = 0
    for x.actual != x.anchor {
      if Less (x.actual.Any, a) {
        x.actual = x.actual.next
        x.pos++
      } else {
        if Less (a, x.actual.Any) {
          break
        } else { // already there
          return
        }
      }
    }
  }
  x.insert (a)
  x.num++
  x.pos++
}

func (x *sequence) Del() Any {
  if x.actual == x.anchor {
    return nil
  }
  c := x.actual.next
  x.actual.prev.next = c
  c.prev = x.actual.prev
  x.actual = c
  x.num--
  return Clone (x.actual.Any)
}

func (x *sequence) ExPred (p Pred, forward bool) bool {
//  c, i := x.actual, x.pos
  if x.num == 0 {
    return false
  }
  c, i := x.anchor.next, uint(0)
  if ! forward {
    c, i = x.anchor.prev, x.num - 1
  }
  for c != x.anchor {
    if p (c.Any) {
      x.actual, x.pos = c, i
      return true
    }
    if forward {
      c = c.next
      i++
    } else {
      c = c.prev
      i--
    }
/*
    if p (c.Any) {
      x.actual, x.pos = c, i
      return true
    }
*/
  }
  return false
}

func (x *sequence) StepPred (p Pred, forward bool) bool {
  c, i := x.actual, x.pos
  for {
    if forward {
      c = c.next
      i++
    } else {
      c = c.prev
      i--
    }
    if c == x.anchor { break }
    if p (c.Any) {
      x.actual = c
      x.pos = i
      return true
    }
  }
  return false
}

func (x *sequence) All (p Pred) bool {
  for c := x.anchor.next; c != x.anchor; c = c.next {
    if ! p (c.Any) {
      return false
    }
  }
  return true
}

func (x *sequence) Ordered() bool {
  c := x.anchor.next
  if c == x.anchor { return true }
  for c.next != x.anchor {
    if Less (c.Any, c.next.Any) {
      c = c.next
    } else {
      x.ordered = false
      return false
    }
  }
  x.ordered = true
  return true
}

func (x *sequence) Sort() { // quicksort
  if x.num <= 1 { return }
  c := x.anchor.next
  if c == x.anchor { return }
  if c.next == x.anchor { return }
  x.anchor.next = c.next
  c.next.prev = x.anchor
  x.num--
  y := new_(x.anchor.Any).(*sequence)
  d := x.anchor.next
  var e *cell
  for d != x.anchor {
    e = d.next
    if Less (c.Any, d.Any) {
      d.prev.next = d.next
      d.next.prev = d.prev
      d.next, d.prev = y.anchor, y.anchor.prev
      d.prev.next = d
      y.anchor.prev = d
      x.num--
      y.num++
    }
    d = e
  }
  x.Sort()
  y.Sort()
  c.next = y.anchor.next
  y.anchor.next = c
  c.prev = y.anchor
  c.next.prev = c
  y.num++
  x.anchor.prev.next = y.anchor.next
  y.anchor.next.prev = x.anchor.prev
  y.anchor.prev.next = x.anchor
  x.anchor.prev = y.anchor.prev
  x.num += y.num
  x.actual = x.anchor
  x.pos = x.num
}

func (x *sequence) ExGeq (a Any) bool {
  if ! x.ordered { Panic ("x is not ordered") }
  x.check (a)
  p := uint(0)
  for c := x.anchor.next; c != x.anchor; c, p = c.next, p + 1 {
    if Less (a, c.Any) {
      x.actual = c
      x.pos = p
      return true
    }
  }
  return false
}

func (x *sequence) Trav (op Op) {
  for c := x.anchor.next; c != x.anchor; c = c.next {
    op (c.Any)
  }
}

func (x *sequence) Filter (Y Collector, p Pred) {
  y := x.imp (Y)
  if y == x { return }
  y.Clr()
  for c := x.anchor.next; c != x.anchor; c = c.next {
    if p (c.Any) {
      y.insert (c.Any)
      y.num++
    }
  }
  y.pos = x.num
}

func (x *sequence) Cut (Y Collector, p Pred) {
  y := x.imp (Y)
  if y == x { return }
  y.Clr()
  c := x.anchor.next
  var d *cell
  for c != x.anchor {
    d = c.next
    if p (c.Any) {
      c.prev.next = c.next
      c.next.prev = c.prev
      c.next, c.prev = y.anchor, y.anchor.prev
      c.prev.next = c
      y.anchor.prev = c
      x.num--
      y.num++
    }
    c = d
  }
  x.actual = x.anchor
  x.pos = x.num
  y.actual = y.anchor
  y.pos = y.num
}

func (x *sequence) ClrPred (p Pred) {
  c := x.anchor.next
  for c != x.anchor {
    d := c.next
    if p (c.Any) {
      c.prev.next = c.next
      c.next.prev = c.prev
      c.prev, c.next = nil, nil
      if x.actual == c {
        x.actual = d
        x.pos++
      }
      x.num--
    } else {
    }
    c = d
  }
}

func (x *sequence) concatenate (y *sequence) {
  if y.Empty() { return }
  if ! Eq (x.anchor.Any, y.anchor.Any) { return }
  x.anchor.prev.next = y.anchor.next
  y.anchor.next.prev = x.anchor.prev
  y.anchor.prev.next = x.anchor
  x.anchor.prev = y.anchor.prev
  x.num += y.num
  if x.actual == x.anchor {
    x.actual = y.anchor.next
  }
  y.Clr()
}

func (x *sequence) join (y *sequence) {
  if y.anchor == y.anchor.next {
    x.actual = x.anchor
    x.pos = x.num
    return
  }
  c := x.anchor.next
  y.actual = y.anchor.next
  for {
    if y.actual == y.anchor { break }
    if c == x.anchor { break }
    if Less (y.actual.Any, c.Any) {
      y.anchor.next = y.actual.next
      y.actual.prev = c.prev
      c.prev.next = y.actual
      c.prev = y.actual
      y.actual.next = c
      y.actual = y.anchor.next
    } else {
      c = c.next
    }
  }
  if y.actual != y.anchor {
    c = c.prev // == x.anchor.prev
    c.next = y.actual
    y.actual.prev = c
    x.anchor.prev = y.anchor.prev
    y.anchor.prev.next = x.anchor
  }
  x.actual = x.anchor
  x.num += y.num
  x.pos = x.num
  y.anchor.next, y.anchor.prev = y.anchor, y.anchor
  y.actual = y.anchor
  y.num, y.pos = 0, 0
}

func (x *sequence) Join (Y Collector) {
  y := x.imp (Y)
  if x.ordered {
    x.join (y)
  } else {
    x.concatenate (y)
  }
}

func (x *sequence) Reverse() {
  if x.ordered { return }
  c := x.anchor
  d := c.next
  for d != x.anchor {
    d = c.next
    c.next = c.prev
    c.prev = d
    c = d
  }
}

func (x *sequence) Rotate (forward bool) {
  if x.ordered { return }
  if x.anchor.next == x.anchor || x.anchor.next == x.anchor.prev {
    return
  }
  if forward {
    c := x.anchor.prev
    c.prev.next = x.anchor
    x.anchor.prev = c.prev
    c.prev = x.anchor
    c.next = x.anchor.next
    x.anchor.next = c
    c.next.prev = c
  } else {
    c := x.anchor.next
    c.next.prev = x.anchor
    x.anchor.next = c.next
    c.next, c.prev = x.anchor, x.anchor.prev
    x.anchor.prev = c
    c.prev.next = c
  }
}
