package obj

// (c) Christian Maurer   v. 211104 - license see µU.go

func sum (as []Adder) Adder {
  a0 := as[0]
  for i, a := range as {
    if i > 0 {
      a0.Add (a)
    }
  }
  return a0
}
