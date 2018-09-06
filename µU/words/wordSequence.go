package words

// (c) Christian Maurer   v. 180815 - license see µU.go

import (
  . "µU/obj"
  "µU/ker"
//  "µU/nat"
  "µU/kbd"
  "µU/col"
  "µU/scr"
//  "µU/errh"
  "µU/text"
//  "µU/sort"
//  "µU/set"
)
const (
  maxL = 16
  maxK = 16
)
type
  wordSequence struct {
          num, length uint
                 word []text.Text
                      }

func (x *wordSequence) imp (Y Any) *wordSequence {
  y, ok := Y.(*wordSequence)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func new_(n, l uint) WordSequence {
  if n == 0 || l == 0 { ker.Shit() }
  if n > maxK { n = maxK }
  if l > maxL { l = maxL }
  x := new(wordSequence)
  x.num, x.length = n, l
  x.word = make ([]text.Text, n)
  for i := uint(0); i < x.num; i++ {
    x.word[i] = text.New (x.length)
  }
  return x
}

func (x *wordSequence) Empty() bool {
  for i := uint(0); i < x.num; i++ {
    if ! x.word[i].Empty () {
      return false
    }
  }
  return true
}

func (x *wordSequence) Clr () {
  for i := uint(0); i < x.num; i++ {
    x.word[i].Clr ()
  }
}

func (x *wordSequence) Copy (Y Any) {
  y := x.imp (Y)
  x.num = y.num
  for i := uint(0); i < x.num; i++ {
    x.word[i].Copy (y.word[i])
  }
}

func (x *wordSequence) Eq (Y Any) bool {
  y := x.imp (Y)
  if x.num != y.num { return false }
  for i := uint(0); i < x.num; i++ {
    if x.word[i].Eq (y.word[i]) {
      return false
    }
  }
  return true
}

func (x *wordSequence) Clone() Any {
  y := New (x.num, x.length)
  y.Copy (x)
  return y
}

func (x *wordSequence) EquivSub (Y Object) bool {
  y := x.imp (Y)
  for i := uint(0); i < y.num; i++ {
    e := false
    for j := uint(0); j < y.num; j++ {
      _, e1 := x.word[i].EquivSub (y.word[j])
      e = e || e1
    }
    if ! e {
      return false
    }
  }
  return true
}

func (x *wordSequence) Less (Y Any) bool {
  if x.num == 0 { return false }
  y := x.imp (Y)
  if x.num == y.num {
    for i := uint(0); i < x.num; i++ {
      if y.word[i].Less (x.word[i]) {
        return false
      }
      if x.word[i].Less (y.word[i]) {
        return true
      }
    }
    return false
  }
  return x.num < y.num
}

func (x *wordSequence) Num (Y Object) uint {
  y := x.imp (Y)
  a := uint(0)
  i, j := uint(0), uint(0)
  for i < y.num && j < y.num {
    if x.word[i].Less (y.word[j]) {
      i ++
    }
    if y.word[j].Less (x.word[i]) {
      j ++
    }
    if x.word[i].Eq (y.word[j]) {
      i ++
      j ++
      a ++
    }
  }
  return a
}

func (x *wordSequence) Colours (f, b col.Colour) {
  for i := uint(0); i < x.num; i++ {
    x.word[i].Colours (f, b)
  }
}

func (x *wordSequence) Height() uint {
  return x.num / (scr.NColumns () / (x.length + 1))
}

func (x *wordSequence) Write (l, c uint) {
  n := scr.NColumns () / (x.length + 1)
  for i := uint(0); i < x.num; i++ {
    x.word[i].Write (l + i / n, c + (i % n) * (x.length + 1))
  }
}

func (x *wordSequence) Print (l, c uint) {
  n := scr.NColumns () / (x.length + 1)
  for i := uint(0); i < x.num; i++ {
    x.word[i].Print (l + i / n, c + (i % n) * (x.length + 1))
  }
}

func (x *wordSequence) Edit (l, c uint) {
  x.Write (l, c)
  n := scr.NColumns() / (x.length + 1)
  i := uint(0)
  loop: for {
    x.word[i].Edit (l + i / n, c + (i % n) * (x.length + 1))
    comm, d := kbd.LastCommand()
    switch comm { case kbd.Esc:
      break loop
    case kbd.Up:
      if d > 0 { break loop }
      if i > 0 {
        i --
      }
    case kbd.Enter, kbd.Down:
      if d > 0 { break loop }
      if i + 1 < x.num {
        i ++
      } else {
        break loop
      }
    case kbd.Tab:
      if d == 0 {
        if i + 1 < x.num {
          i ++
        } else {
          i = 0
        }
      } else {
        if i > 0 {
          i --
        } else {
          i = x.num - 1
        }
      }
    }
  }
  x.Write (l, c)
}

func (x *wordSequence) Codelen() uint {
  return 4 + x.num * x.length
}

func (x *wordSequence) Encode() []byte {
  b := make(Stream, x.Codelen())
  i, a := uint(0), uint(4)
  copy (b[i:i+a], Encode (x.num))
  i += a
  a = x.length
  for i := uint(0); i < x.num; i++ {
    copy (b[i:i+a], x.word[i].Encode ())
    i += a
  }
  return b
}

func (x *wordSequence) Decode (b []byte) {
  i, a := uint(0), uint(4)
  x.num = Decode (uint(0), b[i:i+a]).(uint)
  i += a
  a = x.length
  for i := uint(0); i < x.num; i++ {
    x.word[i].Decode (b[i:i+a])
    i += a
  }
}
