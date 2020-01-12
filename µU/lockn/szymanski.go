package lockn

// (c) Christian Maurer   v. 190815 - license see µU.go

// >>> Szymanski, B. K.: A Simple Solution to Lamport's Concurrent Programming Problem with Linear Wait.
//     In: Lenfant, J. (ed.): ICS '88, New York. ACM Press (1988) 621-626

import (
  . "µU/obj"
  . "µU/atomic"
)
const (
  outsideCS = uint(iota)
  interested
  waitingForOthers
  inWaitingRoom
  behindWaitingRoom
)
type
  szymanski struct {
                   uint "number of processes"
              flag []uint
                   }

func newSzymanski (n uint) LockerN {
  x := new(szymanski)
  x.uint = uint(n)
  x.flag = make([]uint, n)
  return x
}

func (x *szymanski) allLeqWaitingForOthers() bool {
  for q := uint(0); q < x.uint; q++ {
    w := waitingForOthers
    if w < x.flag[q] {
      return false
    }
  }
  return true
}

func (x *szymanski) exists (i, k uint) bool {
  for j := uint(0); j < x.uint; j++ {
    if x.flag[j] == k {
      return true
    }
  }
  return false
}

func (x *szymanski) allLeqInterested (p uint) bool {
  for q := uint(0); q < p; q++ {
    i := interested
    if i < x.flag[q] {
      return false
    }
  }
  return true
}

func (x *szymanski) allOutsideWaitingRoom (p uint) bool {
  for q := p + 1; q < x.uint; q++ {
    if x.flag[q] == waitingForOthers || x.flag[q] == inWaitingRoom {
      return false
    }
  }
  return true
}

func (x *szymanski) Lock (p uint) {
  Store (&x.flag[p], interested)
  for { // wait until for all j: flag[j] <= waitingForOthers
    if x.allLeqWaitingForOthers() {
      break
    }
    Nothing()
  }
  Store (&x.flag[p], inWaitingRoom)
  if x.exists (p, interested) { // if exists j: flag[j] == interested {
    Store (&x.flag[p], waitingForOthers)
    for { // wait until exists j: flag[j] == behindWaitingRoom }
      if x.exists (p, behindWaitingRoom) {
        break
      }
      Nothing()
    }
  }
  Store (&x.flag[p], behindWaitingRoom)
  for { // wait until for all j > p: flag[j] <= interested ||
        //                           flag[j] == leftWaitingRomm
    if x.allOutsideWaitingRoom (p) {
      break
    }
    Nothing()
  }
  for { // wait until for all j < p: flag[j] <= interested 
    if x.allLeqInterested (p) {
      break
    }
    Nothing()
  }
}

func (x *szymanski) Unlock (p uint) {
  Store (&x.flag[p], outsideCS)
}
