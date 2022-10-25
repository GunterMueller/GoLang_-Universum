package seq

// (c) Christian Maurer   v. 221021 - license see µU.go

import (
  "sync"
  . "µU/ker"
  . "µU/obj"
)
type (
  cell struct {
              any "content of the cell"
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

func (x *sequence) check (a any) {
  CheckTypeEq (x.anchor.any, a)
}

func new_(a any) Sequence {
  CheckAtomicOrObject(a)
  x := new(sequence)
  x.anchor = new(cell)
  x.anchor.any = Clone(a)
  x.anchor.next, x.anchor.prev = x.anchor, x.anchor
  x.actual = x.anchor
  return x
}

func (x *sequence) imp (a any) *sequence {
  y, ok := a.(*sequence)
  if ! ok { TypeNotEqPanic (x, a) }
  x.check (y.anchor.any)
  return y
}

func checked (a any) bool {
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
    if ! r (c.any, d.any) {
      return false
    }
  }
  return true
}

func (x *sequence) Eq (Y any) bool {
  return x.e (x.imp (Y), Eq)
}

func (x *sequence) insert (a any) {
  n := new (cell)
  n.any = Clone (a)
  n.next, n.prev = x.actual, x.actual.prev
  x.actual.prev.next = n
  x.actual.prev = n
}

func (x *sequence) Copy (Y any) {
  y := x.imp (Y)
  x.Clr()
  x.anchor.any = Clone(y.anchor.any)
  for c := y.anchor.next; c != y.anchor; c = c.next {
    x.insert (c.any)
  }
  x.num, x.pos = y.num, y.num
}

func (x *sequence) Clone() any {
  y := new_(Clone(x.anchor.any))
  y.Copy (x)
  return y
}

func (x *sequence) Less (Y any) bool {
  y := x.imp (Y)
  if x.num >= y.num { return false }
  c := x.anchor.next
  d := y.anchor.next
  for c != x.anchor {
    for {
      if d == y.anchor {
        return false
      }
      if Eq (c.any, d.any) {
        d = d.next
        break
      }
      d = d.next
    }
    c = c.next
  }
  return true
}

func (x *sequence) Leq (Y any) bool {
  return x.Less (Y) || x.Eq (Y)
}

func (x *sequence) Codelen() uint {
  n := uint(4) // Codelen (uint32(0))
  for c := x.anchor.next; c != x.anchor; c = c.next {
    n += 4 + Codelen (c.any)
  }
  return n
}

func (x *sequence) Encode() []byte {
  b := make ([]byte, x.Codelen())
  i, a := uint32(0), uint32(4)
  copy (b[i:a], Encode (uint32(x.num)))
  i += 4
  for l := x.anchor.next; l != x.anchor; l = l.next {
    n := uint32(Codelen (l.any))
    copy (b[i:i+a], Encode (n))
    i += a
    copy (b[i:i+n], Encode (l.any))
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
    x.insert (Decode (Clone (x.anchor.any), bs[i:i+n]))
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
    if p (c.any) {
      n++
    }
  }
  return n
}

func (x *sequence) Ex (a any) bool {
  x.check (a)
  p := uint(0)
  for c := x.anchor.next; c != x.anchor; c = c.next {
    if Eq (c.any, a) {
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

func (x *sequence) Get() any {
  if x.actual == x.anchor {
    return nil
  }
  return Clone (x.actual.any)
}

func (x *sequence) Put (a any) {
  x.check (a)
  if x.actual == x.anchor {
    x.insert (a)
    x.actual = x.actual.prev
    x.pos = x.num
    x.num++
  } else {
    x.actual.any = Clone (a)
  }
}

func (x *sequence) Ins (a any) {
  x.check (a)
  if x.ordered {
    x.actual = x.anchor.next
    x.pos = 0
    for x.actual != x.anchor {
      if Less (x.actual.any, a) {
        x.actual = x.actual.next
        x.pos++
      } else {
        if Less (a, x.actual.any) {
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

func (x *sequence) Del() any {
  if x.actual == x.anchor {
    return nil
  }
  c := x.actual.next
  x.actual.prev.next = c
  c.prev = x.actual.prev
  x.actual = c
  x.num--
  return Clone (x.actual.any)
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
    if p (c.any) {
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
    if p (c.any) {
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
    if p (c.any) {
      x.actual = c
      x.pos = i
      return true
    }
  }
  return false
}

func (x *sequence) All (p Pred) bool {
  for c := x.anchor.next; c != x.anchor; c = c.next {
    if ! p (c.any) {
      return false
    }
  }
  return true
}

func (x *sequence) Ordered() bool {
  c := x.anchor.next
  if c == x.anchor { return true }
  for c.next != x.anchor {
    if Less (c.any, c.next.any) {
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
  y := new_(x.anchor.any).(*sequence)
  d := x.anchor.next
  var e *cell
  for d != x.anchor {
    e = d.next
    if Less (c.any, d.any) {
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

func (x *sequence) ExGeq (a any) bool {
  if ! x.ordered { Panic ("x is not ordered") }
  x.check (a)
  p := uint(0)
  for c := x.anchor.next; c != x.anchor; c, p = c.next, p + 1 {
    if Less (a, c.any) {
      x.actual = c
      x.pos = p
      return true
    }
  }
  return false
}

func (x *sequence) Trav (op Op) {
  for c := x.anchor.next; c != x.anchor; c = c.next {
    op (c.any)
  }
}

func (x *sequence) Filter (Y Collector, p Pred) {
  y := x.imp (Y)
  if y == x { return }
  y.Clr()
  for c := x.anchor.next; c != x.anchor; c = c.next {
    if p (c.any) {
      y.insert (c.any)
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
    if p (c.any) {
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
    if p (c.any) {
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
  if ! Eq (x.anchor.any, y.anchor.any) { return }
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
    if Less (y.actual.any, c.any) {
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
