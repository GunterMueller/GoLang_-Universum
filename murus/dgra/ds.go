package dgra

// (c) Christian Maurer   v. 170506 - license see murus.go

import
  . "murus/obj"

func (x *distributedGraph) ds (o Op) {
  x.connect (x.me)
  defer x.fin()
  x.Op = o
  x.parent = inf
  if x.me == x.root { // root sends the first message
    x.parent = inf + 1 // trick, see below
    x.visited[0] = true
    x.child[0] = true
    x.ch[0].Send (x.me)
  }
  for i := uint(0); i < x.n; i++ {
    go func (j uint) {
      r := x.ch[j].Recv().(uint)
      sender := r % 256
      if r / 256 == inf {
        if x.child[j] {
          x.child[j] = false
        }
      } else {
        if sender != x.nr[j] { x.log2("sender =", sender, "!= nr[j]", x.nr[j]) }
        if x.parent == inf { // not for root
          x.parent = sender
          for i := uint(0); i < x.n; i++ {
            if i != j && i != x.parent {
              x.child[i] = true
              x.ch[i].Send(x.me)
            }
          }
        } else {
          x.ch[i].Send(x.me + 256 * inf)
        }
      }
      done <- 0
    }(i)
  }
  for i := uint(0); i < x.n; i++ {
    <-done
  }
  x.Op (x.actVertex)
}
