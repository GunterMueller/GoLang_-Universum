package dgra

// (c) Christian Maurer   v. 170803 - license see murus.go

import (
  "murus/ker"
  . "murus/obj"
)
const (
  child = 1027
  noChild = 1099
)

func (x *distributedGraph) bfsx (o Op) {
  x.connect (uint(999))
  defer x.fin()
  x.Op = o
  if x.me == x.root { // root sends the first message
    x.time = x.me
    x.parent = inf + 1 // trick, see below
    for i := uint(0); i < x.n; i++ {
      x.ch[i].Send (x.me)
mutex.Lock(); println ("send", x.me, "to", x.nr[i]); mutex.Unlock()
      x.child[i] = true
    }
  }

  for j := uint(0); j < x.n; j++ {
    go func (i uint) {
      for {
println("wait for", x.nr[i])
        t := x.ch[i].Recv().(uint)
        x.chanuint[i] <- t
        t = <-x.chanuint[i]
        x.ch[i].Send(t)
      }
    }(j)
  }

  for i := uint(0); i < x.n; i++ {
    go func (j uint) {
//      loop:
      for {
        select {
        case got := <-x.chanuint[j]:
          mutex.Lock()
println ("recv", got , "from", x.nr[j])
          switch got {
          case child:
            x.child[j] = true
            x.visited[j] = true
println ("have child", x.nr[j])
          case noChild:
            x.child[j] = false
            x.visited[j] = true
          default:
            if x.parent == inf {
              x.parent = x.nr[j]
println ("have parent", x.nr[j])
              x.chanuint[j] <- child
              x.visited[j] = true
              for k := uint(0); k < x.n; k++ {
                if ! x.visited[k] {
                  x.chanuint[k] <- x.me
println ("send", x.me, "to", x.nr[k])
                  x.visited[k] = true
                }
              }
            } else {
println ("send", noChild, "to", x.nr[j])
              x.chanuint[j] <- noChild
            }
          }
          mutex.Unlock()
//          done <- 0
//          break loop
        default:
        }
        ker.Msleep(1000)
println("select for", x.nr[j])
      }
    }(i)
  }
/*
  for i := uint(0); i < x.n; i++ {
    <-done
  }
  x.Op(x.actVertex)
*/
}
