package seq

// (c) Christian Maurer   v. 170501 - license see µu.go

import
  . "µu/obj"
type (
  cell struct {
              Any
         next,
         prev *cell
       }
  sequence struct {
              num,
              pos uint
   anchor, actual *cell
          ordered bool
                  }
)

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

func (x *sequence) remove() {
  a := x.actual.next
  x.actual.prev.next = a
  a.prev = x.actual.prev
  x.actual.prev, x.actual.next = nil, nil
  x.actual = a
}

func (x *sequence) Clr() {
  x.actual = x.anchor.next
  for x.actual != x.anchor {
    x.remove()
  }
  x.num, x.pos = 0, 0
}

func (x *sequence) ins (a Any) {
  c := new (cell)
  c.Any = Clone (a)
  c.next, c.prev = x.actual, x.actual.prev
  x.actual.prev.next = c
  x.actual.prev = c
}

func (x *sequence) Copy (Y Any) {
  y := x.imp (Y)
  x.Clr()
  x.anchor.Any = Clone(y.anchor.Any)
  for l := y.anchor.next; l != y.anchor; l = l.next {
    x.ins (l.Any)
  }
  x.num, x.pos = y.num, y.num
}

func (x *sequence) Clone() Any {
  y := new_(Clone(x.anchor.Any))
  y.Copy (x)
  return y
}

func (x *sequence) e (y *sequence, r Rel) bool {
  if x.num != y.num { return false }
  for l, l1 := x.anchor.next, y.anchor.next; l != x.anchor; l, l1 = l.next, l1.next {
    if ! r (l.Any, l1.Any) {
      return false
    }
  }
  return true
}

func (x *sequence) Eq (Y Any) bool {
  return x.e (x.imp (Y), Eq)
}

func (x *sequence) Less (Y Any) bool {
  y := x.imp (Y)
  if x.num >= y.num { return false }
  l := x.anchor.next
  l1 := y.anchor.next
  for l != x.anchor {
    for {
      if l1 == y.anchor {
        return false
      }
      if Eq (l.Any, l1.Any) {
        l1 = l1.next
        break
      }
      l1 = l1.next
    }
    l = l.next
  }
  return true
}

func (x *sequence) Num() uint {
  return x.num
}

func (x *sequence) NumPred (p Pred) uint {
  n := uint(0);
  for l := x.anchor.next; l != x.anchor; l = l.next {
    if p (l.Any) {
      n++
    }
  }
  return n
}

func (x *sequence) Ex (a Any) bool {
  x.check (a)
  p := uint(0)
  for l := x.anchor.next; l != x.anchor; l = l.next {
    if Eq (l.Any, a) {
      x.actual = l
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
      x.pos ++
    }
  } else if x.actual != x.anchor.next {
    x.actual = x.actual.prev
    x.pos --
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
    x.pos ++
  }
  for x.pos > i {
    x.actual = x.actual.prev
    x.pos --
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
    x.ins (a)
    x.actual = x.actual.prev
    x.pos = x.num
    x.num ++
  } else {
    x.actual.Any = Clone (a)
  }
}

func (x *sequence) insert (a Any) {
  x.ins (a)
  x.num ++
  x.pos ++
}

func (x *sequence) Ins (a Any) {
  x.check (a)
  if x.ordered {
    x.actual = x.anchor.next
    x.pos = 0
    for x.actual != x.anchor {
      if Less (x.actual.Any, a) {
        x.actual = x.actual.next
        x.pos ++
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
}

func (x *sequence) InsRel (a Any, r Rel) {
  x.check (a)
  x.actual, x.pos = x.anchor.next, 0
  for x.actual != x.anchor {
    if r (x.actual.Any, a) {
      x.actual = x.actual.next
      x.pos ++
    } else {
      break
    }
  }
  x.insert (a)
}

func (x *sequence) Del() Any {
  if x.actual == x.anchor {
    return nil
  }
  defer x.remove()
  x.num--
  return Clone (x.actual.Any)
}

func (x *sequence) ExPred (p Pred, forward bool) bool {
//  s, i := x.actual, x.pos
  if x.num == 0 {
    return false
  }
  s, i := x.anchor.next, uint(0)
  if ! forward {
    s, i = x.anchor.prev, x.num - 1
  }
  for s != x.anchor {
    if p (s.Any) {
      x.actual, x.pos = s, i
      return true
    }
    if forward {
      s = s.next
      i++
    } else {
      s = s.prev
      i--
    }
/*
    if p (s.Any) {
      x.actual, x.pos = s, i
      return true
    }
*/
  }
  return false
}

func (x *sequence) StepPred (p Pred, forward bool) bool {
  s, i := x.actual, x.pos
  for {
    if forward {
      s = s.next
      i++
    } else {
      s = s.prev
      i--
    }
    if s == x.anchor { break }
    if p (s.Any) {
      x.actual = s
      x.pos = i
      return true
    }
  }
  return false
}

func (x *sequence) All (p Pred) bool {
  for l := x.anchor.next; l != x.anchor; l = l.next {
    if ! p (l.Any) {
      return false
    }
  }
  return true
}

func (x *sequence) Ordered() bool {
  l := x.anchor.next
  if l == x.anchor { return true }
  for l.next != x.anchor {
    if Less (l.Any, l.next.Any) {
      l = l.next
    } else {
      return false
    }
  }
  return true
}

func (x *sequence) Sort() {
  x.ordered = true
  if x.Ordered() { return }
  l := x.anchor.next
  if l == x.anchor { return }
  if l.next == x.anchor { return }
  x.anchor.next = l.next
  l.next.prev = x.anchor
  x.num --
  var y *sequence
  y = new_(x.anchor.Any).(*sequence)
  l1 := x.anchor.next
  var l2 *cell
  for l1 != x.anchor {
    l2 = l1.next
    if Less (l.Any, l1.Any) {
      l1.prev.next = l1.next
      l1.next.prev = l1.prev
      l1.next, l1.prev = y.anchor, y.anchor.prev
      l1.prev.next = l1
      y.anchor.prev = l1
      x.num --
      y.num ++
    }
    l1 = l2
  }
  x.Sort()
  y.Sort()
  l.next = y.anchor.next
  y.anchor.next = l
  l.prev = y.anchor
  l.next.prev = l
  y.num ++
  x.concatenate (y)
  x.actual = x.anchor
  x.pos = x.num
}

func (x *sequence) ExGeq (a Any) bool {
  if ! x.ordered {
    return false // TODO Panic ?
  }
  x.check (a)
  p := uint(0)
  for l := x.anchor.next; l != x.anchor; l, p = l.next, p + 1 {
    if Less (a, l.Any) {
      x.actual = l
      x.pos = p
      return true
    }
  }
  return false
}

func (x *sequence) Trav (op Op) {
  for l := x.anchor.next; l != x.anchor; l = l.next {
    op (l.Any)
  }
}

func (x *sequence) Filter (Y Iterator, p Pred) {
  y := x.imp (Y)
  if y == x { return }
  y.Clr()
  for l := x.anchor.next; l != x.anchor; l = l.next {
    if p (l.Any) {
      y.ins (l.Any)
      y.num ++
    }
  }
  y.pos = x.num
}

func (x *sequence) Split (Y Iterator) {
  y := x.imp (Y)
  if y == x { return }
  y.Clr()
  if x.actual == x.anchor { return }
  y.anchor.next, y.anchor.prev = x.actual, x.anchor.prev
  x.anchor.prev.next = y.anchor
  x.anchor.prev = x.actual.prev
  x.actual.prev.next = x.anchor
  x.actual.prev = y.anchor
  x.actual = x.anchor
  y.actual = y.anchor.next
  y.num = x.num - x.pos
  x.num = x.pos
  x.pos = x.num
}

func (x *sequence) Cut (Y Iterator, p Pred) {
  y := x.imp (Y)
  if y == x { return }
  y.Clr()
  l := x.anchor.next
  var l1 *cell
  for l != x.anchor {
    l1 = l.next
    if p (l.Any) {
      l.prev.next = l.next
      l.next.prev = l.prev
      l.next, l.prev = y.anchor, y.anchor.prev
      l.prev.next = l
      y.anchor.prev = l
      x.num --
      y.num ++
    }
    l = l1
  }
  x.actual = x.anchor
  x.pos = x.num
  y.actual = y.anchor
  y.pos = y.num
}

func (x *sequence) ClrPred (p Pred) {
  l := x.anchor.next
  for l != x.anchor {
    a := l
    l := l.next
    if p (a.Any) {
      a.prev.next = a.next
      a.next.prev = a.prev
      a.prev, a.next = nil, nil
      if x.actual == a {
        x.actual = l
        x.pos ++
      }
      x.num --
    }
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
  l := x.anchor.next
  y.actual = y.anchor.next
  for {
    if y.actual == y.anchor { break }
    if l == x.anchor { break }
    if Less (y.actual.Any, l.Any) {
      y.anchor.next = y.actual.next
      y.actual.prev = l.prev
      l.prev.next = y.actual
      l.prev = y.actual
      y.actual.next = l
      y.actual = y.anchor.next
    } else {
      l = l.next
    }
  }
  if y.actual != y.anchor {
    l = l.prev // == x.anchor.prev
    l.next = y.actual
    y.actual.prev = l
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

func (x *sequence) Join (Y Iterator) {
  y := x.imp (Y)
  if x.ordered {
    x.join (y)
  } else {
    x.concatenate (y)
  }
}

// Not documented - destroys the order, if x is ordered !!!
func (x *sequence) Reverse() {
//  if x.ordered { return }
  l := x.anchor
  l1 := l.next
  for l1 != x.anchor {
    l1 = l.next
    l.next = l.prev
    l.prev = l1
    l = l1
  }
}

// Not documented - destroys the order, if x is ordered !!!
func (x *sequence) Rotate (forward bool) {
//  if x.ordered { return }
  if x.anchor.next == x.anchor || x.anchor.next == x.anchor.prev {
    return
  }
  if forward {
    l := x.anchor.prev
    l.prev.next = x.anchor
    x.anchor.prev = l.prev
    l.prev = x.anchor
    l.next = x.anchor.next
    x.anchor.next = l
    l.next.prev = l
  } else {
    l := x.anchor.next
    l.next.prev = x.anchor
    x.anchor.next = l.next
    l.next, l.prev = x.anchor, x.anchor.prev
    x.anchor.prev = l
    l.prev.next = l
  }
}

func (x *sequence) MinCodelen() uint {
  return Codelen (x.num)
}

func (x *sequence) Codelen() uint {
  n := uint(4) // Codelen (uint32(0))
  for l := x.anchor.next; l != x.anchor; l = l.next {
    n += 4 + Codelen (l.Any)
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
    x.ins (Decode (Clone (x.anchor.Any), bs[i:i+n]))
    i += n
  }
  return
}

func (x *sequence) Slice() []Any {
  a := make ([]Any, x.Num())
  for i, l := 0, x.anchor.next; l != x.anchor; i, l = i+1, l.next {
    a[i] = Clone (l.Any)
  }
  return a
}

func (x *sequence) Deslice (b []Any) {
  x.Clr()
  for _, a := range b {
    x.Ins (a)
  }
}
