package qu

// (c) Christian Maurer   v. 170620 - license see µU.go

import
  . "µU/obj"
type
  queue struct {
               Any "to maintain the type of objects in the queue"
             s []Any
               }

func new_(a Any) Queue {
  x := new(queue)
  x.Any = Clone(a)
  x.s = make([]Any, 0)
  return x
}

func (x *queue) Empty() bool {
  return len(x.s) == 0
}

func (x *queue) Num() uint {
  return uint(len(x.s))
}

func (x *queue) Ins (a Any) {
  CheckTypeEq (a, x.Any)
  x.s = append (x.s, a)
}

func (x *queue) Get() Any {
  if x.Empty() { return nil }
  a := x.s[0]
  x.s = x.s[1:]
  return a
}
