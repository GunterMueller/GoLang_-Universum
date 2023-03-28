package main

// (c) Christian Maurer   v. 210331 - license see µU.go

import
  . "µU/reg"

func Kopie (a Register) Register {
  if a.Gt0() { goto A }
  return Null()
A:
  a.Dec()
  b := Kopie (a)
  a.Inc()
  b.Inc()
  return b
}

func µ (f RegFunc1) RegFunc {
  return func (b Registers) Register {
           a := Null()
           goto B
         A:
           a.Inc()
         B:
           if f(a, b).Gt0() { goto A }
           return a
         }
}

func r3() Register { d := Null(); d.Inc(); d.Inc(); d.Inc(); return d }

func r7() Register { s := r3(); s.Inc(); s.Inc(); s.Inc(); s.Inc(); return s }

func main() {
}
