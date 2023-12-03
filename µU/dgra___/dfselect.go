package dgra

// (c) Christian Maurer   v. 200728 - license see µU.go

func (x *distributedGraph) dfselect() {
  x.connect (x.leader)
  defer x.fin()
  if x.me == x.root { // root sends the first message
    x.parent = x.root // inf + 1 // trick, see below
//    x.ch[0].Send (x.leader)
    x.send (0, x.leader)
    x.child[0] = true
    x.visited[0] = true
  }
  x.distance, x.diameter = x.n, inf // both variables only used for temporary information
  for i := uint(0); i < x.n; i++ {
    go func (j uint) {
      a := x.ch[j].Recv().(uint)
      v, t := a % inf, a / inf
      mutex.Lock()
      if v > x.me {
        x.leader = v
      }
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
          if u == x.n { // all neighbours visited
            t++
            k = x.channel(x.parent) // send echo back to x.parent
          }
        } else { // parent is already defined
          k = j // send echo back to sender
        }
      } else { // x.visited[j], i.e. echo
        if u == x.n { // all neighbours visited
          t++
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
//      x.ch[k].Send(x.leader + inf * t)
      x.send (k, x.leader + inf * t)
      mutex.Unlock()
      done <- 0
    }(i)
  }
  for i := uint(0); i < x.n; i++ {
    <-done
  }
  if x.me == x.root {
    for i := uint(0); i < x.n; i++ {
      if x.child[i] {
//        x.ch[i].Send(x.leader)
        x.send (i, x.leader)
      }
    }
  } else {
    x.leader = x.ch[x.channel(x.parent)].Recv().(uint)
    for i := uint(0); i < x.n; i++ {
      if x.child[i] {
//        x.ch[i].Send(x.leader)
        x.send (i, x.leader)
      }
    }
  }
}
