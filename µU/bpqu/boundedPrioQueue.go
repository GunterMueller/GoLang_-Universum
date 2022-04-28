package bpqu

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  "sync"
  . "µU/obj"
)
type
  boundedPrioQueue struct {
                     heap []any // heap[0] to save the type
                 cap, num uint
                          }
var
  mutex sync.Mutex

func new_(a any, m uint) BoundedPrioQueue {
  CheckAtomicOrObject (a)
  if m == 0 { return nil }
  x := new(boundedPrioQueue)
  x.heap = make([]any, m + 1)
  x.heap[0] = Clone(a)
  x.cap = m
  return x
}

func (x *boundedPrioQueue) Clr() {
  x.num = 0
}

func (x *boundedPrioQueue) Empty() bool {
  return x.num == 0
}

func (x *boundedPrioQueue) Num() uint {
  return x.num
}

func (x *boundedPrioQueue) Full() bool {
  return x.num == x.cap
}

// lift heap [i] as far as necessary to restore the heap invariant heap [i] <= heap [j]
// for all i:= 1, ..., (x.num - 1) / 2, j == 2 * i and j == 2 * i + 1
func (x *boundedPrioQueue) lift() {
  mutex.Lock()
  i := x.num
  for {
    if i == 1 {
      break
    }
    j := i / 2 // index above i
    if Less (x.heap [j], x.heap [i]) {
      break // i < x.num, above i heap invariant is ok
    } else {
      x.heap[i], x.heap[j] = x.heap[j], x.heap[i]
    }
    i = j // continue above
  }
  mutex.Unlock()
}

func (x *boundedPrioQueue) Ins (a any) {
  if x.num == x.cap { return } // x full
  CheckTypeEq (a, x.heap[0])
  x.num++
  x.heap [x.num] = Clone (a)
  go x.lift ()
}

// sift heap[1] as far as necessary to restore the heap invariant
func (x *boundedPrioQueue) sift() {
  mutex.Lock()
  i := uint(1)
  for {
    if i > x.num / 2 {
      break // nothing more under i
    }
    j := 2 * i // left under i
    if j < x.num && ! Less (x.heap[j], x.heap[j + 1]) {
      j++ // right under i
    }
    if Less (x.heap [i], x.heap [j]) {
      break
    } else {
      x.heap[i], x.heap[j] = x.heap[j], x.heap[i]
      i = j
    }
  }
  mutex.Unlock()
}

func (x *boundedPrioQueue) Get() any {
  if x.num == 0 { return nil }
  return x.heap[1]
}

func (x *boundedPrioQueue) Del() any {
  if x.num == 0 { return nil }
  a := x.heap[1]
  x.heap[1] = x.heap[x.num]
  x.num--
  go x.sift()
  return Clone (a)
}
