package qu

// (c) Christian Maurer   v. 220801 - license see nU.go

import
  . "nU/obj"
type
  queue struct {
               any "Musterobjekt"
            qs []any
               }

func new_(a any) Queue {
  x := new(queue)
  x.any = Clone(a)
  x.qs = make([]any, 0)
  return x
}

func (x *queue) Empty() bool {
  return len(x.qs) == 0
}

func (x *queue) Num() int {
  return len(x.qs)
}

func (x *queue) Enqueue (a any) {
  x.qs = append(x.qs, a)
}

func (x *queue) Dequeue() any {
  if x.Empty() {
    return x.any
  }
  a := x.qs[0]
  x.qs = x.qs[1:]
  return a
}
