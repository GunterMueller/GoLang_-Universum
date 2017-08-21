package internal

// (c) murus.org  v. 170424 - license see murus.go

import (
  "murus/ker"
  . "murus/obj"
  "murus/col"
  "murus/scr"
  "murus/text"
)
const
  max = 2 * N
type
  page struct {
       object Object
          len,       // Codelen of a content
          num uint32 // number of nonempty objects on the page
          pos [max+2]uint32
      content [max+1]Object
              }

func new_(a Any) Page {
  o := object (a)
  x := new (page)
  x.object = o.Clone().(Object)
  x.len = uint32(o.Codelen())
  for i := 0; i <= max; i++ {
    x.content[i] = o.Clone().(Object)
  }
  return x
}

func imp (X Any) *page {
  x, ok := X.(*page)
  if ! ok { TypeNotEqPanic (x, X) }
  return x
}

func object (a Any) Object {
  o, ok := a.(Object)
  if ! ok { TypeNotEqPanic (o, a) }
  return o
}

func (x *page) Empty() bool {
  if x.num > 0 { return false }
  for i := 0; i < max; i++ {
    if x.pos[i] > 0 { return false }
    if ! x.content[i].Empty() { return false }
  }
  if x.pos[max] > 0 { return false }
  return true
}

func (x *page) Eq (Y Any) bool {
  y := imp (Y)
  if x.num != y.num { return false }
  for i := 0; i < max; i++ {
    if x.pos[i] != y.pos[i] { return false }
    if ! x.content[i].Eq (y.content[i]) { return false }
  }
  if x.pos[max] != y.pos[max] { return false }
  return true
}

func (x *page) Less (Y Any) bool {
  return false
}

func (x *page) Copy (Y Any) {
  y := imp (Y)
  x.num = y.num
  for i := 0; i < max; i++ {
    x.pos[i] = y.pos[i]
    x.content[i].Copy (y.content[i])
  }
  x.pos[max] = y.pos[max]
}

func (x *page) Clone() Any {
  y := New (x.object)
  y.Copy (x)
  return y
}

func (x *page) Clr() {
  x.num = 0
  for i := uint(0); i < max; i++ {
    x.pos[i] = 0
    x.content[i].Clr()
  }
  x.pos[max] = 0
}

const
  cluint32 = 4 // Codelen(uint32(0))

func (x *page) Codelen() uint {
  return cluint32 + max * (cluint32 + uint(x.len)) + cluint32
}

func (x *page) Encode() []byte {
  bs := make ([]byte, x.Codelen())
  j := 0
  a := cluint32
  copy (bs[j:j+a], Encode (x.num))
  j += a
  for i := 0; i < max; i++ {
    a = cluint32
    copy (bs[j:j+a], Encode (x.pos[i]))
    j += a
    a = int(x.len)
    copy (bs[j:j+a], x.content[i].Encode())
    j += a
  }
  a = cluint32
  copy (bs[j:j+a], Encode (x.pos[max]))
  return bs
}

func (x *page) Decode (bs []byte) {
  j := 0
  a := cluint32
  x.num = Decode (uint32(0), bs[j:j+a]).(uint32)
  j += a
  for i := 0; i < max; i++ {
    a = cluint32
    x.pos[i] = Decode (uint32(0), bs[j:j+a]).(uint32)
    j += a
    a = int(x.len)
    x.content[i].Decode (bs[j:j+a])
    j += a
  }
  a = cluint32
  x.pos[max] = Decode (uint32(0), bs[j:j+a]).(uint32)
}

func (x *page) PutNum (n uint) {
  if n > max { ker.Oops() }
  x.num = uint32(n)
}

func (x *page) GetNum() uint {
  return uint(x.num)
}

func (x *page) PutPos (p, n uint) {
  if p > max + 1 { ker.Oops() }
  x.pos[p] = uint32(n)
}

func (x *page) GetPos (p uint) uint {
  if p > max + 1 { ker.Oops() }
  return uint(x.pos[p])
}

func (x *page) Put (p uint, o Object) {
  if p > max + 1 { ker.Oops() }
  x.content[p] = o.Clone().(Object)
}

func (x *page) Get (p uint) Object {
  if p > max + 1 { ker.Oops() }
  return x.content[p].Clone().(Object)
}

func (x *page) Oper (p uint, op Op) {
  op (x.content[p])
}

func (x *page) Ins (o Object, p, n uint) {
  if p < uint(x.num) {
    for i := uint(x.num); i >= p + 1; i-- {
      x.pos[i + 1] = x.pos[i]
      x.content[i] = x.content[i - 1]
    }
  }
  x.content[p] = o
  x.pos[p + 1] = uint32(n)
  x.num ++
  if x.num < max {
    for i := x.num; i < max; i++ {
      x.content[i] = x.object
      x.pos[i + 1] = 0
    }
  }
}

func (x *page) IncNum() {
  x.num ++
}

func (x *page) DecNum() {
  if x.num == 0 { ker.Oops() }
  x.num --
}

func (x *page) RotLeft() {
  for i := uint32(1); i < x.num; i++ {
    x.content[i - 1] = x.content[i]
    x.pos[i - 1] = x.pos[i]
  }
  x.content[x.num - 1] = x.object
  x.pos[x.num - 1] = x.pos[x.num]
  x.pos[x.num] = 0
  x.num --
}

func (x *page) RotRight() {
  x.pos[x.num + 1] = x.pos[x.num]
//  for i := x.num - 1; i >= 0; i-- { // does not work, because for uint: 0-- == 2^32 - 1  !
  i := x.num - 1
  for {
    x.content[i + 1], x.pos[i + 1] = x.content[i], x.pos[i]
    if i == 0 {
      break
    }
    i--
  }
}

func (x *page) Join (p uint) {
  if p < uint(x.num) {
    for i := p; i < uint(x.num); i++ {
      x.content[i - 1] = x.content[i]
      x.pos[i] = x.pos[i + 1]
    }
  }
  x.content[x.num - 1] = x.object
  x.pos[x.num] = 0
  x.num --
}

func (x *page) Del (p uint) {
  if p + 1 < uint(x.num) {
    for i := p + 1; i < uint(x.num); i++ {
      x.content[i - 1] = x.content[i]
      x.pos[i] = x.pos[i + 1]
    }
  }
  x.content[x.num - 1] = x.object
  x.pos[x.num] = 0
}

func (x *page) ClrLast() {
  x.content[x.num - 1] = x.object
  x.pos[x.num - 1] = x.pos[x.num] // ?
  x.pos[x.num] = 0
  x.num --
}

func (x *page) Write (l, c uint) {
  scr.Colours (col.White, col.Blue)
  scr.WriteNat (uint(x.num), l, c)
  c += 4
  for i := uint(0); i < max; i++ {
    scr.Colours (col.Yellow, col.Red)
    scr.WriteNat (uint(x.pos[i]), l, c)
    c += 4
    scr.Colours (col.White, col.Blue)
    scr.Write (x.content[i].(text.Text).String(), l, c)
    c += 10
  }
  scr.Colours (col.Yellow, col.Red)
  scr.WriteNat (uint(x.pos[max]), l, c)
}
