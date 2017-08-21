package dgra

// (c) murus.org  v. 170802 - license see murus.go

import (
  "murus/ker"
  . "murus/obj"
)

func (x *distributedGraph) dfsx (o Op) {
  x.connect (x.time)
  defer x.fin()
  x.Op = o
  if x.me == x.root { // root sends the first message
    x.time = 0
    x.parent = inf + 1 // trick, see below
    x.ch[0].Send (x.time)
mutex.Lock(); println ("send 0 to", x.nr[0]); mutex.Unlock()
    x.child[0] = true
    x.visited[0] = true
  }
  x.distance, x.diameter = x.n, inf // both variables only used for temporary information

  for j := uint(0); j < x.n; j++ {
    go func (i uint) {
      for {
        t := x.ch[i].Recv().(uint)
        x.chanuint[i] <- t
        t = <-x.chanuint[i]
        x.ch[i].Send(t)
      }
    }(j)
  }

  for i := uint(0); i < x.n; i++ {
    go func (j uint) {
      loop:
      for {
        select {
        case t := <-x.chanuint[j]:
          mutex.Lock()
println ("recv", t, "from", x.nr[j])
          if x.distance == j && x.diameter == t { // unchanged t back from this channel
            x.child[j] = false // so x.nr[j] is no child of x.me
          }
          u := x.next (j) // == x.n, iff all neighbours != j are visited
          k := u
          if ! x.visited[j] { // probe
            x.visited[j] = true
            if x.parent == inf { // parent is undefined (not for root - see trick)
              x.parent = x.nr[j]
              t++
              x.time = t
println("time =", t)
              if u == x.n { // all neighbours visited
                t++
                x.time1 = t
println("time1 =", t)
                k = x.channel(x.parent) // send echo back to x.parent
              }
            } else { // parent is already defined
              k = j // send echo back to sender
            }
          } else { // x.visited[j], i.e. echo
            if u == x.n { // all neighbours visited
              t++
              x.time1 = t
println("time1 =", t)
              if x.me == x.root { // root must not reply any more
                done <- 0
                return
              }
              k = x.channel(x.parent) // send echo back to x.parent
            }
          }
          x.visited[k] = true
          if k == u { // send probe
            x.distance, x.diameter = k, t
            x.child[k] = true // temptative
          }
println ("send", t, "to", x.nr[k])
          x.chanuint[k] <- t
          mutex.Unlock()
          done <- 0
          break loop
        default:
        }
        ker.Msleep(3000)
      }
    }(i)
  }
  for i := uint(0); i < x.n; i++ {
    <-done
  }
  x.Op(x.actVertex)
}
