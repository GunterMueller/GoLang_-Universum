package phil

// (c) Christian Maurer   v. 170627 - license see µu.go

import
  "sync"
type state byte; const (
  satisfied = state(iota)
  hungry
  starving
  hasRightFork
  hasLeftFork
  dining
  nStates
)
var (
  status [NPhilos]state
  mutex sync.Mutex
)

func left (i uint) uint {
  return (i + NPhilos - 1) % NPhilos
}

func right (i uint) uint {
  return (i + 1) % NPhilos
}

func changeStatus (i uint, s state) {
  mutex.Lock()
  status[i] = s
  mutex.Unlock()
}
