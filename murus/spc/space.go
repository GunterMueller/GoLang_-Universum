package spc

// (c) Christian Maurer   v. 121204 - license see murus.go

func next (d Direction) Direction {
  return (d + 1) % NDirs
}

func prev (d Direction) Direction {
  return (d + 2) % NDirs
}

func init() {
  for d := D0; d < NDirs; d++ {
//    Origin [d] = 0.0
//    for D:= D0; D < NDirs; D++ { Unit [d][D] = 0.0 }
    Unit [d][d] = 1.0
  }
}
