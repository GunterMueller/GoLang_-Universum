package ego

// (c) Christian Maurer   v. 221213 - license see µU.go

import (
  "µU/ker"
  "µU/env"
  "µU/N"
)

func ego (k uint) uint {
  i, ok := N.Natural (env.Arg(1))
  if ! ok {
    ker.Panic("falsches Argument")
  }
  if i >= k {
    ker.Panic("zu großes Argument")
  }
  return i
}

func me() uint {
  i, ok := N.Natural (env.Arg(1))
  if ! ok {
    return uint(1 << 16)
  }
  return i
}
