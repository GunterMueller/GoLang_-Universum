package pqu

// (c) Christian Maurer   v. 220420 - license see µU.go

import
  . "µU/obj"
type
  prioQueue struct {
              heap []any // heap[0] = pattern object
                   }

func new_(a any) PrioQueue {
  if a == nil { return nil }
  CheckAtomicOrObject (a)
  x := new(prioQueue)
  x.heap = make([]any, 1)
  x.heap[0] = Clone(a)
  return x
}

func (x *prioQueue) Empty() bool {
  return len(x.heap) == 1
}

func (x *prioQueue) Num() uint {
  return uint(len(x.heap)) - 1
}

func (x *prioQueue) sift (i uint) {
  m, n, k := i, 2 * i, 2 * i + 1
  if n <= x.Num() && Less (x.heap[i], x.heap[n]) {
    m = n
  }
  if k <= x.Num() && Less (x.heap[m], x.heap[k]) {
    m = k
  }
  if m != i {
    x.heap[i], x.heap[m] = x.heap[m], x.heap[i]
    x.sift (m)
  }
}

func (x *prioQueue) Ins (a any) {
  CheckTypeEq (a, x.heap[0])
  x.heap = append (x.heap, Clone(a))
  n := uint(len(x.heap))
  i := n - 1
  for i > 1 && Less (x.heap[i/2], a) {
    x.heap[i] = x.heap[i/2]
    i /= 2
  }
  x.heap[i] = Clone(a)
}

func (x *prioQueue) Get() any {
  if x.Empty() { return nil }
  a := x.heap[1] // 22
  if x.Num() == 1 {
    x.heap = x.heap[:1]
    return a
  }
  x.heap[1] = x.heap[x.Num()]
  n := uint(len(x.heap))
  x.heap = x.heap[:n-1]
  x.sift (1)
  return a
}
