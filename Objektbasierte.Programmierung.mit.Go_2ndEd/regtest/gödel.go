package main

// (c) Christian Maurer   v. 210331 - license see µU.go

import
  . "µU/reg"

func Kopie (a Register) Register {
  if a.Gt0() { goto A }
  return Null()
A: // sei x = a
  a.Dec() // a = x - 1
  k := Kopie (a) // k = x - 1
  a.Inc() // a = x (Wert beim Aufruf)
  k.Inc() // k = x = a
  return k
}

func Plus1 (n Register) Register {
  n1 := Kopie(n)
  n1.Inc()
  return n1
}

func Minus1 (n Register) Register {
  n1 := Kopie(n)
  n1.Dec()
  return n1
}

func Kleinergleich (a, b Register) bool {
  if a.Gt0() { goto A }
  return true // a == 0 <= b
A: // a > 0
  if b.Gt0() { goto B }
  return false // b == 0 < a
B: // a > 0, b > 0
  a1 := Kopie(a)
  b1 := Kopie(b)
  return Kleinergleich (Minus1(a1), Minus1(b1))
}

func Gleich (a, b Register) bool {
  if a.Gt0() { goto A }
// a == 0
  if b.Gt0() { goto F }
  return true // b = 0 = a
A: // a > 0
  if b.Gt0() { goto B }
  return false // b = 0 < a
B: // a > 0, b > 0
  return Gleich (Minus1(a), Minus1(b))
F:
  return false // a = 0 < b
}

func Summe (a, b Register) Register {
  a1 := Kopie(a)
  b1 := Kopie(b)
  if b1.Gt0() { goto A }
  return a1 // b = 0, folglich a + b = a
A:
  return Plus1 (Summe (a1, Minus1 (b1)))
}

func Differenz (a, b Register) Register {
  a1 := Kopie(a)
  b1 := Kopie(b)
  if b1.Gt0() { goto A }
  return a1 // b = 0, folglich a - b = a
A: // b > 0
  if a.Gt0() { goto B }
  return Null() // a = 0 < b, folglich a - b < 0
B: // a > 0, b > 0
  return Differenz (Minus1(a1), Minus1(b1))
}

func IstEins (n Register) bool {
  n1 := Kopie(n)
  if n.Gt0() { goto A }
  return false
A:
  n1.Dec()
  if n1.Gt0() { goto B }
  return true
B:
  return false
}

func IstZwei (n Register) bool {
  n1 := Kopie(n)
  if n.Gt0() { goto A }
  return false
A:
  n1.Dec()
  if n1.Gt0() { goto B }
  return false
B:
  n1.Dec()
  if n1.Gt0() { goto C }
  return true
C:
  return false
}

func Produkt (a, b Register) Register {
  a1 := Kopie(a)
  b1 := Kopie(b)
  if a1.Gt0() { goto A }
  return Null()
A:
  if b1.Gt0() { goto B }
  return Null() // s. oben
B:
  a1.Dec()
  if a1.Gt0() { goto C }
  return b1
C:
  p := Produkt (a1, b1)
  if p.Gt0() { goto D }
  return Null()
D:
  s := Summe (p, b1)
  return s
}

func Div2 (a Register) Register {
  d := Null()
  if a.Gt0() { goto A }
  return d
A:
  a.Dec()
  if a.Gt0() { goto B }
  return d
B:
  a.Dec()
  d.Inc()
  if a.Gt0() { goto A }
  return d
}

func C0 (a, b Register) Register {
  c := Summe (a, b)
  d := Kopie (c)
  d.Inc()
  p := Produkt (c, d)
  f := Div2 (p)
  return Summe (f, a)
}

// Vor.: Es gibt höchstens 2 Register in a.
func C (a Registers) Register {
  if IstZwei (a.Num()) { goto A }
////////////  ^^^^^^^^ XXX XXX XXX XXX XXX
  return C0 (a.Head(), C (a.Tail()))
A:
  return C0 (a.Head(), a.Tail().Head())
}

// Liefert das größte y mit C0(0, y) <= n
func max (n Register) Register {
  n1 := Kopie (n)
  y := Kopie (n)
A:
  c := C0 (Null(), y)
  if Kleinergleich (c, n1) { goto B }
  y.Dec()
  goto A
B:
  return y
}

func D0 (a Register) (Register, Register) {
  return E(a), F(a)
}

func E (n Register) Register {
  n1 := Kopie (n)
  y0 := max(n1)
  return Differenz (n1, C0(Null(), y0))
}

func F (n Register) Register {
  n1 := Kopie(n)
  y0 := max(n1)
  x := Differenz (n1, C0(Null(), y0))
  return Differenz (y0, x)
}

func DD (i, n Register) Register {
  i1 := Kopie(i)
  if i1.Gt0() { goto A }
  return F (n)
A:
  i1.Dec()
  return DD (i1, n)
}

func D1 (i, n Register) Register {
  i1 := Kopie (i)
  if i1.Gt0() { goto A }
  return E (n)
A:
  i1.Dec()
  if i1.Gt0() { goto B }
  return E (F (n))
B:
  return DD (i1, F (n))
}

/*
func D (n Register) Registers {
  list := Empty()
  return list.Cons (D1 (0, n))
  , D1 (1, n), D1 (2, n), ..., D1 (n-1, n))
}
*/

func main() {
  Eins := Null()
  Eins.Inc()
  Zwei := Kopie(Eins)
  Zwei.Inc()
  Drei := Kopie(Zwei)
  Drei.Inc()
  Vier := Kopie(Drei)
  Vier.Inc()
  Fünf := Kopie(Vier)
  Fünf.Inc()
  c := C (New (Zwei, Eins, Drei))
/*
  d := Produkt (Vier, Vier)
  e := Produkt (d, Fünf)
  f := Summe (e, Vier)
  g := Summe (f, Vier)
  h := Summe (g, Vier)
  h.Inc()
  h.Write()
  F(h).Write()
return
  D1 (Null(), h).Write()
  D1 (Eins, h).Write()
  D1 (Zwei, h).Write()
return
*/
  D1 (Null(), c).Write()
  D1 (Eins, c).Write()
  D1 (Zwei, c).Write()
}
