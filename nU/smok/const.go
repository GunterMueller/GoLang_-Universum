package smok

// (c) Christian Maurer   v. 171018 - license see nU.go

const (
  agent = uint(iota)
  smokerIn
  smokerOut
)
const (
  paper = uint(iota)
  tobacco
  matches
)

func others (u uint) (uint, uint) {
  return (u + 1) % 3, (u + 2) % 3
}
