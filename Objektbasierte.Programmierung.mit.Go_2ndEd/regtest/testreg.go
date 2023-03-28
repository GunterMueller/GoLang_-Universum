package main

// (c) Christian Maurer   v. 210331 - license see µU.go

import
  . "µU/reg"

// Liefert ein Register, dessen Wert um 1 größer ist
// als der Wert von a.
func plus1 (a Register) Register {
  x := Kopie(a)
  x.Inc()
  return x
}

func plus2 (a Register) Register {
  x := Kopie(a)
  x.Inc()
  x.Inc()
  return x
}

func minus1 (a Register) Register {
  x := Kopie(a)
  x.Dec()
  return x
}

/*
func produkt (a, b Register) Register {
  p := Null()
  if a.Gt0() { goto A }
  return p // geht auch goto B, aber dann muss x:= Null() vorne stehen !
A:
  if b.Gt0() { goto B }
  return p // s. oben
B:
  x := Null()
C:
  Add (p, a)
  b.Dec()
  x.Inc()
  if b.Gt0() { goto C }
D:
  b.Inc()
  x.Dec()
  if x.Gt0() { goto D }
  return p
}
*/

func Kopie (a Register) Register {
  k := Null() // Register zur Aufnahme der Kopie+
  if a.Gt0() { goto A }
  return k // a = 0, k = 0
A: // sei x = Wert von a beim Aufruf
  h := Null() // Hilfsregister zum Protokollieren
              // der Anzahl der a.Dec()-Anweisungen
  // h = k = 0, a + h = a = x
B:
  a.Dec()
  k.Inc()
  h.Inc() // a + h = x, h = k > 0
  if a.Gt0() { goto B }
  // a = 0, folglich a + h = x = h, deshalb k = x,
  // aber wegen x = h > 0 gilt a < a + h = x
  // bei der Rückgabe von k muss aber a = x gelten,
  // deshalb a wieder sooft erhöhen, wie in h protolliert ist:
C:
  a.Inc()
  h.Dec() // a + h = x
  if h.Gt0() { goto C }
  // h = 0, folglich a = x
  return k // k = a
}

func KopieRek (a Register) Register {
  if a.Gt0() { goto A }
  return Null()
A: // sei x = a
  a.Dec() // a = x - 1
  k := KopieRek (a) // k = x - 1
  a.Inc() // a = x (Wert beim Aufruf)
  k.Inc() // k = x = a
  return k
}

func gleich (a, b Register) bool {
  x := Kopie (a) // x = a
  y := Kopie (b) // y = b, folglich a + y = b + x
A:
  if x.Gt0() { goto B }
// x == 0
  if y.Gt0() { goto F } // x = 0 < y
  return true // y = 0 = x, folglich a = a + y = b + x = b
B: // x > 0
  if y.Gt0() { goto C } // x > 0, y > 0
  return false // x > 0 = y, folglich a + y = a = b + x > b
C: // x > 0, y > 0
  x.Dec()
  y.Dec() // a + y = b + x
  goto A
F:
  return false // x = 0 < y, folglich a < a + y == b + x == b
}

func gleichRek (a, b Register) bool {
  if a.Gt0() { goto A }
// a == 0
  if b.Gt0() { goto F }
  return true // b = 0 = a
A: // a > 0
  if b.Gt0() { goto B }
  return false // b = 0 < a
B: // a > 0, b > 0
  return gleichRek (minus1(a), minus1(b))
F:
  return false // a = 0 < b
}

func kleinergleich (a, b Register) bool {
  if a.Gt0() { goto A }
  return true // a == 0 <= b
A: // a > 0
  if b.Gt0() { goto B }
  return false // b == 0 < a
B: // a > 0, b > 0
  return kleinergleich (minus1(a), minus1(b))
}

func differenz (a, b Register) Register {
  if b.Gt0() { goto A }
  return Kopie(a) // b = 0, folglich a - b = a
A: // b > 0
  if a.Gt0() { goto B }
  return Null() // a = 0 < b, folglich a - b < 0
B: // a > 0, b > 0
  return differenz (minus1(a), minus1(b))
}

func summe (a, b Register) Register {
  s := Kopie (a)
  x := Kopie (b) // s + x == a + b
  if x.Gt0() { goto A }
  return s // x == 0, s == a
A:
  s.Inc()
	x.Dec() // s + 1 + x - 1 == s + x == a + b
  if x.Gt0() { goto A }
  return s // x == 0, folglich s + x == s == a + b 
}

func summeRek (a, b Register) Register {
  if b.Gt0() { goto A }
  return Kopie(a) // b = 0, folglich a + b = a
A:
  return plus1 (summeRek (a, minus1 (b)))
}

func produkt (a, b Register) Register {
  if a.Gt0() { goto A }
  return Null() // a = 0, folglich a * b = 0
A:
  return summe (produkt (minus1 (a), b), b)
}

func div2 (a Register) Register {
// TODO
  d := Null()
  if a.Gt0() { goto A }
  goto D
A:
/*
  a.Dec()
  if a.Gt0() { goto C }
  return b
C:
  d.Inc()
  a.Dec()
  if a.Gt0() { goto B }
*/
D:
  return d
}

// Liefert 0, falls a den Wert 0 hat, ansonsten die Summe
// der ersten n natürlichen Zahlen, wobei n der Wert von ist.
func Gauß (a Register) Register {
  g := Null()
  if a.Gt0() { goto A }
  return g // g = a = 0
A:
  x := Kopie(a)
  i := Null()
  k := Null()
B:
  x.Dec()
  i.Inc()
C:
  g.Inc()
  k.Inc()
  i.Dec()
  if i.Gt0() { goto C }
D:
  k.Dec()
  i.Inc()
  if k.Gt0() { goto D }
  if x.Gt0() { goto B }
  return g
}

func Gaußrek (a Register) Register {
  if a.Gt0() { goto A }
  return Null()
A:
  return summe (a, Gaußrek(minus1(a)))
}

func eq0 (a Register) bool {
  if a.Gt0() { goto A }
  return false
A:
  return true
}

func µ (f RegFunc1) RegFunc {
  return func (a Registers) Register {
           n := Null()
           goto B
         A:
           n.Inc()
         B:
           if eq0( f (n, a) ) { goto A }
           return n
         }
}

func zwei() Register {
  zwei := Null()
  zwei.Inc()
  zwei.Inc()
  return zwei
}

func drei() Register {
  drei := zwei()
  drei.Inc()
  return drei
}

func sechs() Register {
  return produkt (zwei(), drei())
}

func folgensumme (a Registers) Register {
  if a.NotEmpty() { goto A }
  return Null()
A:
  return summe (a.Head(), folgensumme (a.Tail()))
}

func Ackermann (a, b Register) Register {
  if a.Gt0() { goto A }
  return plus1(b) // a = 0
A:
  if b.Gt0() { goto B }
  return Ackermann (minus1 (a), plus1 (Null())) // a > 0, b = 0
B:
  return Ackermann (minus1 (a), Ackermann (a, minus1 (b))) // a > 0, b > 0
}

func main() {
  x := drei()
//  x.Write()
  y := produkt (zwei(), x)
//  x.Write()
//  y.Write()
  z := Kopie (y)
  z.Dec()
//  z.Write()
//  differenz (z, x).Write()
//  summe (x, x).Write()
//  drei().Write()
  vier := plus1(drei())
//  vier.Write()
//  acht := produkt(zwei(), vier)
//  acht.Write()
  vierundzwanzig := produkt (drei(), produkt(zwei(), vier))
  vierundzwanzig.Write()
// 
//  sechs().Write()
  summe (vierundzwanzig, sechs()).Write()
//  drei().Write()
//  Gaußrek(Zwanzig()).Write()
/*
  x := hundert92()
  KopieRek (x).Write()
  x.Write()
*/
//  Eins := plus1(Null())
  fünf := plus1(vier)
//  fünf.Write()
  folgensumme (New (produkt (sechs(), fünf), vier, sechs(),
               summe(drei(), sechs()), zwei())).Write() // 51
/*
  x := Sequence (produkt (sechs(), Fuenf), vier, sechs(), summe(drei(), sechs()), zwei())
  folgensumme (x).Write()
  folgensumme (x.Tail().Cons(vier)).Write()
  Ackermann (drei(), zwei()).Write()
*/
}
