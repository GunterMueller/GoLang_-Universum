package main

// (c) Christian Maurer   v. 210331 - license see µU.go

import
  . "µU/reg"

func kopie (a Register) Register {
  if a.Gt0() { goto A }
  return Null()
A:
  a.Dec()
  b := kopie (a)
  a.Inc()
  b.Inc()
  return b
}

func eins() Register {
  e := Null()
  e.Inc()
  return e
}

func succ (a Register) Register {
  b := kopie (a)
  b.Inc()
  return b
}

func summe (a, b Register) Register {
  if a.Gt0() { goto A }
  return kopie(b)
A:
  c := kopie (a)
  c.Dec()
  return succ (summe (c, b))
}

func produkt (a, b Register) Register {
  if a.Gt0() { goto A }
  return Null()
A:
  c := kopie(a)
  c.Dec()
  return summe (produkt(c, b), b)
}

func r2() Register { z := Null(); z.Inc(); z.Inc(); return z }
func r4() Register { v := r2(); v.Inc(); v.Inc(); return v }
func r7() Register { s := r4(); s.Inc(); s.Inc(); s.Inc(); return s }
func r10() Register { z := r7(); z.Inc(); z.Inc(); z.Inc(); return z }
func r16() Register { s := r10(); s.Inc(); s.Inc(); s.Inc(); s.Inc(); s.Inc(); s.Inc(); return s }

func main() {
  a, b := r7(), r16()
  a.Write()
  b.Write()
  summe (a, b).Write()
  produkt (a, b).Write()
}
