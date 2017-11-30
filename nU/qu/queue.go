package qu

// (c) Christian Maurer   v. 171125 - license see nU.go

import . "nU/obj"

type queue struct {
  Any "Musterobjekt"
  qs []Any
}

func new_(a Any) Queue {
  x := new(queue)
  x.Any = Clone(a)
  x.qs = make([]Any, 0)
  return x
}

func (x *queue) Empty() bool {
  return len(x.qs) == 0
}

func (x *queue) Num() int {
  return len(x.qs)
}

func (x *queue) Enqueue (a Any) {
  x.qs = append(x.qs, a)
}

func (x *queue) Dequeue() Any {
  if x.Empty() {
    return x.Any
  }
  a := x.qs[0]
  x.qs = x.qs[1:]
  return a
}
