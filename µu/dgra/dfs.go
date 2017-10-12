package dgra

// (c) Christian Maurer   v. 170504 - license see µu.go

import
  . "µu/obj"

func (x *distributedGraph) dfs (o Op) {
  x.connect (x.time)
  defer x.fin()
  x.Op = o
  if x.me == x.root { // root sends the first message
    x.time = 0
    x.parent = inf + 1 // trick, see below
    x.ch[0].Send (x.time)
 println ("send", x.time, "to", x.nr[0])
    x.child[0] = true
    x.visited[0] = true
  }
  x.distance, x.diameter = x.n, inf // both variables only used for temporary information
  for i := uint(0); i < x.n; i++ {
    go func (j uint) {
      t := x.ch[j].Recv().(uint)
 println ("recv", t, "from", x.nr[j])
      mutex.Lock()
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
          if u == x.n { // all neighbours visited
            t++
            x.time1 = t
            k = x.channel(x.parent) // send echo back to x.parent
          }
        } else { // parent is already defined
          k = j // send echo back to sender
        }
      } else { // x.visited[j], i.e. echo
        if u == x.n { // all neighbours visited
          t++
          x.time1 = t
          if x.me == x.root { // root must not reply any more
            mutex.Unlock()
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
      x.ch[k].Send(t)
      mutex.Unlock()
      done <- 0
    }(i)
  }
  for i := uint(0); i < x.n; i++ {
    <-done
  }
  x.Op(x.actVertex)
}
