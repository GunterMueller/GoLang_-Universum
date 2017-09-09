package ego

// (c) Christian Maurer   v. 170424 - license see murus.go

import (
  "murus/ker"
  "murus/env"
  "murus/nat"
)

func ego (n uint) uint {
  i, ok := nat.Natural (env.Par(1))
  if ! ok {
    ker.Panic("falsches Argument")
  }
  if i >= n {
    ker.Panic("zu großes Argument")
  }
  return i
}

func me() uint {
  i, ok := nat.Natural (env.Par(1))
  if ! ok {
    return uint(1 << 16)
  }
  return i
}
