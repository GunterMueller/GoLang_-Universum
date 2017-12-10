package ego

// (c) Christian Maurer   v. 171202 - license see µU.go

import (
  "µU/ker"
  "µU/env"
  "µU/nat"
)

func ego (n uint) uint {
  i, ok := nat.Natural (env.Arg(1))
  if ! ok {
    ker.Panic("falsches Argument")
  }
  if i >= n {
    ker.Panic("zu großes Argument")
  }
  return i
}

func me() uint {
  i, ok := nat.Natural (env.Arg(1))
  if ! ok {
    return uint(1 << 16)
  }
  return i
}
