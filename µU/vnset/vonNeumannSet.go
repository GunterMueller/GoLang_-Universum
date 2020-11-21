package vnset

// (c) Christian Maurer   v. 201014 - license see µU.go

import (
  . "µU/obj"
  "µU/z"
  "µU/str"
)
type
  set struct {
        elem []*set
             }
const
  nn = 10
var
  smallord [nn]*set

func init() {
  for i := uint(0); i < nn; i++ {
    smallord[i] = ord(i)
  }
}

func emptySet() VonNeumannSet {
  x := new(set)
  x.elem = make([]*set, 0)
  return x
}

func (x *set) imp (Y Any) *set {
  y, ok := Y.(*set)
  if ! ok { TypeNotEqPanic(x, Y) }
  return y
}

func (x *set) Subset (Y VonNeumannSet) bool {
  if x.Empty() { return true }
  y := x.imp(Y)
  for _, e := range x.elem {
    if ! e.Element(y) { return false }
  }
  return true
}

func (x *set) Eq (Y Any) bool {
  y := x.imp(Y)
  return x.Subset(y) && y.Subset(x)
}

func (x *set) Less (Y Any) bool {
  y := x.imp(Y)
  if x.Num() == y.Num() { return false }
  return x.Subset (y)
}

func (x *set) Copy (Y Any) {
  y := x.imp(Y)
  n := len(y.elem)
  x.elem = make([]*set, n)
  for i, e := range y.elem {
    x.elem[i] = Clone(e).(*set)
  }
}

func (x *set) Clone() Any {
  y := EmptySet()
  y.Copy(x)
  return y
}

func (x *set) Empty() bool {
  return len(x.elem) == 0
}

func (x *set) Clr() {
  x.elem = make ([]*set, 0)
}

func (x *set) Codelen() uint {
  return uint(len(x.String()))
}

func (x *set) Encode() Stream {
  return Stream(x.String())
}

func (x *set) Decode (bs Stream) {
  x.Defined(string(bs))
}

func (x *set) String() string {
  for i := uint(0); i < 10; i++ {
    if x.Eq(smallord[i]) {
      return string(byte(i + '0'))
    }
  }
  if x.Empty() { return "0" }
  s := "{"
  for i, e := range x.elem {
    t := e.String()
    switch t {
    case "{0}": t = "1"
    case "{0,1}": t = "2"
    case "{0,1,2}": t = "3"
    case "{0,1,2,3}": t = "4"
    case "{0,1,2,3,4}": t = "5"
    case "{0,1,2,3,4,5}": t = "6"
    case "{0,1,2,3,4,5,6}": t = "7"
    case "{0,1,2,3,4,5,6,7}": t = "8"
    case "{0,1,2,3,4,5,6,7,8}": t = "9"
    case "{0,1,2,3,4,5,6,7,8,9}": t = "10"
    }
    s += t
    if i + 1 < len(x.elem) { s += "," }
  }
  return s + "}"
}

func (x *set) Defined (s string) bool {
  str.OffBytes (&s, ' ')
  if len(s) == 1 {
    if z.IsDigit (s[0]) {
      *x = *smallord[uint(s[0]) - uint('0')]
      return true
    }
  }
  ss := str.SplitBrackets (s, ',', '{') // geht schief bei Komma zuviel
  if ss == nil {
    return false
  }
  for _, t := range ss {
    if t == "" {
      return false
    } else {
      y := emptySet().(*set)
      if y.Defined (t) {
        x.elem = append (x.elem, y)
      } else {
        return false
      }
    }
  }
  return true
}

func (x *set) Element (Y VonNeumannSet) bool {
  y := x.imp(Y)
  if y.Empty() { return false }
  if x.Empty() {
    for _, e := range y.elem {
      if e.Empty() { return true }
    }
    return false
  }
  for _, e := range y.elem {
    if x.Eq (e) { return true }
  }
  return false
}

func (x *set) Num() uint {
  return uint(len(x.elem))
}

func (x *set) Singleton() VonNeumannSet {
  y := EmptySet().(*set)
  y.elem = []*set { x.Clone().(*set) }
  return y
}

func (x *set) Doubleton (Y VonNeumannSet) VonNeumannSet {
  z := x.Singleton().(*set)
  if ! x.Eq(Y) {
    z.elem = append(z.elem, Y.Clone().(*set))
  }
  return z
}

func (x *set) Zero() bool {
  return x.Empty()
}

func (x *set) Sum (Y, Z Adder) {
  x.Copy (Y)
  x.Add (Z)
}

func (x *set) Add (Y ...Adder) {
  for _, z := range Y {
    y := x.imp(z)
    for _, e := range y.elem {
      if ! e.Element (x) {
        x.elem = append(x.elem, e)
      }
    }
  }
}

func (x *set) Sub (Y ...Adder) {
  z := EmptySet()
  z.Add(Y...)
  m := EmptySet().(*set)
  for _, a := range x.elem {
    if ! a.Element(z) {
      m.elem = append(m.elem, a)
    }
  }
  *x = *m
}

func (x *set) Union (Y VonNeumannSet) VonNeumannSet {
  x.Add(Y)
  return x
}

func (x *set) BigUnion() VonNeumannSet {
  y := EmptySet().(*set)
  for _, e := range x.elem {
    y.Add(e)
  }
  return y
}

func (x *set) Intersection (Y VonNeumannSet) VonNeumannSet {
  y := x.imp(Y)
  z := new(set)
  z.elem = make([]*set, 0)
  if x.Empty() || y.Empty() {
    return z
  }
  for _, e := range x.elem {
    if e.Element(y) {
      z.elem = append(z.elem, e)
    }
  }
  return z
}

func (x *set) BigIntersection() VonNeumannSet {
  z := EmptySet().(*set)
  n := len(x.elem)
  if n == 0 { return z }
  if x.elem[0].Empty() { return z }
  for _, a := range x.elem[0].elem {
    ok := true
    for i := 1; i < n; i++ {
      if !a.Element(x.elem[i]) {
        ok = false
        break
      }
    }
    if ok {
      z.elem = append(z.elem, a)
    }
  }
  return z
}

func (x *set) Succ() VonNeumannSet {
  y := x.Clone().(*set)
  y.Add (y.Singleton())
  return y
}

func (x *set) Powerset() VonNeumannSet {
  if x.Empty() {
    return x.Singleton()
  }
  n := len(x.elem)
  a := x.elem[n-1].Singleton()
  y := new(set)
  y.elem = x.elem[:n-1]
  y = y.Powerset().(*set)
  n = len(y.elem)
  z := new(set)
  z.elem = make([]*set, n)
  for i, e := range y.elem {
    z.elem[i] = e.Union(a).(*set)
  }
  return y.Union(z)
}

func (x *set) Comprehension (p Predicate) VonNeumannSet {
  y := EmptySet().(*set)
  for _, e := range x.elem {
    if p(e) {
      y.elem = append(y.elem, e)
    }
  }
  return y
}

func (x *set) Transitive() bool {
  if x.Empty() { return true }
  for _, e := range x.elem {
    if ! e.Subset (x) { return false }
  }
  return true
}

func (x *set) KuratowskiPair (Y VonNeumannSet) VonNeumannSet {
  y := x.imp(Y)
  z := emptySet().(*set)
  z.elem = make([]*set, 2)
  z.elem[0], z.elem[1] = x.Singleton().(*set), x.Doubleton(y).(*set)
  return z
}

func (x *set) Pair (Y VonNeumannSet) VonNeumannSet {
  y := x.imp(Y)
  u := emptySet().(*set)
  u.elem = make([]*set, x.Num())
  for i, a := range x.elem {
    u.elem[i] = smallord[0].Pair(a).(*set)
  }
  v := emptySet().(*set)
  v.elem = make([]*set, y.Num())
  for i, b := range y.elem {
    v.elem[i] = smallord[1].Pair(b).(*set)
  }
  return u.Union(v).(*set)
}

func ord (i uint) *set {
  if i == 0 { return emptySet().(*set) }
  return ord(i-1).Succ().(*set)
}

func ordinal (i uint) VonNeumannSet {
  if i < nn {
    return smallord[i]
  }
  return ord(i)
}

func setOf (X ...VonNeumannSet) VonNeumannSet {
  x := emptySet().(*set)
  x.elem = make([]*set, len(X))
  for i, e:= range X {
    x.elem[i] = e.(*set)
  }
  return x
}
