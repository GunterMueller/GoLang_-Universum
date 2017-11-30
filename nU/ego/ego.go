package ego

// (c) Christian Maurer   v. 170424 - license see nU.go

import "nU/env"

/*
func ego (n uint) uint {
  i, ok := nat.Natural (env.Par(1))
  if ! ok {
    panic("falsches Argument")
  }
  if i >= n {
    panic("zu großes Argument")
  }
  return i
}
*/

func me() uint {
  return uint(env.Par1() - '0')
}
