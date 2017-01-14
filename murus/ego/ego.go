package ego

// (c) murus.org  v. 151110 - license see murus.go

import (
  "murus/ker"; "murus/env"
  "murus/nat"
)

func ego (n uint) uint {
  e, ok := nat.Natural (env.Par(1))
  if ! ok {
    ker.Panic ("falsches Argument")
  }
  if e >= n {
    ker.Panic ("zu großes Argument")
  }
  return e
}
