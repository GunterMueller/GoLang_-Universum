package ego

// (c) Christian Maurer   v. 171227 - license see nU.go

import ("strconv"; "nU/env")

func ego (n uint) uint {
  if i, err := strconv.Atoi(env.Arg(1)); err == nil {
    if uint(i) < n {
      return uint(i)
    }
  }
  return uint(1<<16)
}

func me() uint {
  if i, err := strconv.Atoi(env.Arg(1)); err == nil {
    return uint(i)
  }
  return uint(1<<16)
}
