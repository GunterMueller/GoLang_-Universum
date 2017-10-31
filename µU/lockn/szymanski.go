package lockn

// (c) Christian Maurer   v. 171025 - license see µU.go

// >>> Szymanski, B. K.: A Simple Solution to Lamport's Concurrent Programming Problem with Linear Wait.
//     In: Lenfant, J. (ed.): ICS '88, New York. ACM Press (1988) 621-626

import
  "runtime"
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

func newS (n uint) LockerN {
  x := new(szymanski)
  x.uint = n
  x.flag = make([]uint, n)
  return x
}

func (x *szymanski) allLeqWaitingForOthers (i uint) bool {
  for j := uint(0); j < x.uint; j++ {
    if x.flag[j] > waitingForOthers { return false }
  }
  return true
}

func (x *szymanski) exists (i, k uint) bool {
  for j := uint(0); j < x.uint; j++ {
    if x.flag[j] == k { return true }
  }
  return false
}

func (x *szymanski) allLeqInterested (i uint) bool {
  for j := uint(0); j < i; j++ {
    if x.flag[j] > interested { return false }
  }
  return true
}

func (x *szymanski) allOutsideWaitingRoom (i uint) bool {
  for j := i + 1; j < x.uint; j++ {
    if x.flag[j] == waitingForOthers ||
       x.flag[j] == inWaitingRoom { return false }
  }
  return true
}

func (x *szymanski) Lock (i uint) {
  x.flag[i] = interested
  for { // wait until for all j: flag[j] <= waitingForOthers
    if x.allLeqWaitingForOthers (i) { break }
    runtime.Gosched()
  }
  x.flag[i] = inWaitingRoom
  if x.exists (i, interested) { // if exists j: flag[j] == interested {
    x.flag[i] = waitingForOthers
    for { // wait until exists j: flag[j] == behindWaitingRoom }
      if x.exists (i, behindWaitingRoom) { break }
      runtime.Gosched()
    }
  }
  x.flag[i] = behindWaitingRoom

  for { // wait until for all j > i: flag[j] <= interested ||
        //                           flag[j] = leftWaitingRomm
    if x.allOutsideWaitingRoom (i) { break }
//    runtime.Gosched()
  }

  for { // wait until for all j < i: flag[j] <= interested 
    if x.allLeqInterested (i) { break }
    runtime.Gosched()
  }
}

func (x *szymanski) Unlock (i uint) {
/*
  for { // wait until for all j > i: flag[j] <= interested ||
        //                           flag[j] = leftWaitingRomm
    if x.allOutsideWaitingRoom (i) { break }
//    runtime.Gosched()
  }
*/
  x.flag[i] = outsideCS
}
