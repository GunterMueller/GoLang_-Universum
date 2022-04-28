package f

// (c) Christian Maurer   v. 220420 - license see µU.go

import
  . "µU/obj"

func pow (a, b uint) uint {
  if b == 0 { return 1 }
  return a * pow (a, b - 1)
}

func f (a any, i uint) any {
  p := Decode(UintStream{0, 0}, a.(Stream)).(UintStream)
  switch i {
  case Plus:
    return p[0] + p[1]
  case Times:
    return p[0] * p[1]
  case ToThe:
    return pow (p[0], p[1])
  }
  return uint(0)
}
