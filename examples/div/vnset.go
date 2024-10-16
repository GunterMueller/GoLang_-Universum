package main

import
  . "µU/vnset"

func main() {
  const n = uint(10)
  var set[n]VonNeumannSet
  for i := uint(0); i < n; i++ { set[i] = Ordinal(i) }
  println(set[9].BigIntersection().String())
  a := EmptySet(); a.Defined("{0,1,2,{4,3},3,4,5,6,7,  9}")
  println(a.String())
  bs := a.Encode()
  a = EmptySet()
  a.Decode(bs)
  println(a.String())
  b := EmptySet(); b.Defined("{  1,2,3,4,  6,7,8,9}")
  c := EmptySet(); c.Defined("{0,1,2,3,4,5,  7,8,9}")
  d := EmptySet(); d.Defined("{0,1,  3,  5,6,7,8  }")
  e := EmptySet(); e.Defined("{0,1,2,3,4,5,  7,  9}")
//                               1   3       7
  println (a.String())
  println (b.String())
  println (c.String())
  println (d.String())
  println (e.String())
  x := SetOf (a, b, c, d, e)
  println(x.String()); println()
  println(x.BigIntersection().String()) // exspected: {1,3,7}
  println()
  a = set[3].Powerset()
//  if a.Founded() { println (a.String(), "founded") } else { println (a.String(), "not founded") }
  if a.Transitive() { println (a.String(), "transitive") } else { println (a.String(), "not transitive") }
  a = SetOf(set[1], set[3], set[5])
//  if a.Founded() { println (a.String(), "founded") } else { println (a.String(), "not founded") }
  if a.Transitive() { println (a.String(), "transitive") } else { println (a.String(), "not transitive") }
  if set[4].Transitive() { println ("4 transitive") }
//  if s[2].Powerset().Founded() { println ("P(2) founded") }
  x = EmptySet()
  var s string
  s = "{0,{0},{0,{0}}}"
  s = "{{{{{0},{{{0}}}}}}}"
  s = "{0,{1}}"
  s = "{1,{2},{3,{4}}}"
  s = "{{0},{0,{0}},{{0},{0,{0}},{0,{{0},{{{0}}}}}}}"
  s = "{3, 5, 6, 9}"
  if x.Defined(s) {
    println(s)
    println(x.String())
  } else {
    println ("Mist")
  }
//  if Ordinal(3).Founded() { println ("3 is founded")}
  if Ordinal(3).Transitive() { println ("3 is transitive")}
//  if Ordinal(3).Powerset().Founded() { println ("P(3) is founded")}
  if Ordinal(3).Powerset().Transitive() { println ("P(3) is transitive")}
  set[0] = EmptySet()
  for i := uint(1); i < n; i++ {
    set[i] = set[i-1].Succ()
//    if ! set[i].Eq (Ord(i)) { println ("oops", i) }
    println (i, "=", set[i].String())
  }
  if ! set[0].Element (set[0]) { println("0 nicht in 0") } else { println("FALSCH: 0 in 0") }
  if  set[0].Element (set[1]) { println("0 in 1") } else { println("FALSCH: 0 nicht in 1") }
  if  set[0].Element (set[9]) { println("0 in 9") } else { println("FALSCH: 0 nicht in 9") }
  if ! set[1].Element (set[0]) { println("1 nicht in 0") } else { println("FALSCH: 1 in 0") }
  if !set[1].Element (set[1]) { println("1 nicht in 1") } else { println("FALSCH: 1 in 1") }
  if  set[1].Element (set[5]) { println("1 in 5") } else { println("FALSCH: 1 nicht in 5") }
  if  set[1].Element (set[9]) { println("1 in 9") } else { println("FALSCH: 1 nicht in 9") }
  if ! set[5].Element (set[5]) { println("5 nicht in 5") } else { println("FALSCH: 5 in 5") }
  if  set[5].Element (set[9]) { println("5 in 9") } else { println("FALSCH: 5 nicht in 9") }
  if  set[5].Subset (set[9]) { println("5 ist Teilmenge von 9.") } else { println("FALSCH: 5 ist keine Teilmenge von 9.") }
  if  set[5].Less (set[9]) { println("5 ist echte Teilmenge von 9.") } else { println("FALSCH: 5 ist keine echte Teilmenge von 9.") }
  if  set[9].Subset (set[9]) { println("9 ist Teilmenge von 9.") } else { println("FALSCH: 9 ist keine Teilmenge von 9.") }
  if !set[9].Less (set[9]) { println("9 ist keine echte Teilmenge von 9.") } else { println("FALSCH: 9 ist echte Teilmenge von 9.") }
  if  set[9].Transitive() { println("9 ist transitiv.") } else { println("FALSCH: 9 ist nicht transitiv.") }
  if  set[8].BigUnion().Eq (set[7]) { println ("BigUnion(8) gleich 7.") } else { println ("FALSCH: Union(8) nicht gleich 7.") }
  if set[5].Union (set[9]).Eq(set[9]) { println ("5 vereinigt mit 9 gleich 9.") } else { println ("FALSCH: 5 vereinigt mit 9 ist nicht gleich 9.") }
  if set[9].Union (set[0]).Eq(set[9]) { println ("9 vereinigt mit 0 gleich 9.") } else { println ("FALSCH: 9 vereinigt mit 0 ist nicht gleich 9.") }
  println (set[3].String())
  println (set[3].Powerset().String())
}
