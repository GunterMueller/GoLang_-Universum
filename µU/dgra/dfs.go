package dgra

// (c) Christian Maurer   v. 241004 - license see µU.go

func (x *distributedGraph) Dfs() {
  x.connect (x.time) // netchannels have the type uint
  defer x.fin()
  if x.me == x.root {
    x.parent = x.root
    x.time = 0
    x.send (0, x.time) // root sends first
    x.child[0] = true
    x.visited[0] = true
  }
  x.distance, x.diameter = x.n, inf
  if x.demo { x.Write() }
  for i := uint(0); i < x.n; i++ {
    go func (j uint) {
      t := x.recv (j).(uint)
      mutex.Lock()
      if x.distance == j && x.diameter == t { // t unchanged back on
                                              // the j-th netchannel,
        x.child[j] = false // thus x.nr[j] is no child of x.me
      }
      u := x.next(j) // == x.n iff all netchannels != j
                     // are already marked
      k := u // channel for the next send operation
      if x.visited[j] { // i.e. echo
        if u == x.n { // all netchannels marked
          t++
          x.time1 = t
          if x.me == x.root { // root must not send an echo any more
            mutex.Unlock()
            done <- 0
            return
          }
          k = x.channel(x.parent) // echo x.time1 back to sender
        } else {
          // k == u < x.n, t unchange as probe ahead to x.nr[u]
        }
      } else { // ! x.visited[j], i.e. probe
        x.visited[j] = true
        if x.parent < inf { // father already defined
          k = j // t unchanged as echo back to sender
        } else { // x.parent == inf, i.e. father yet
                 // undefined (not for root!)
          x.parent = x.nr[j]
          t++
          x.time = t // if u < x.n probe x.time ahead to x.nr[u]
          if u == x.n { // all netchannels marked
            t++
            x.time1 = t
            k = j // echo x.time1 back to sender (father)
          }
        }
      }
      x.visited[k] = true
      if k == u {
        x.distance, x.diameter = k, t // save k, t for the above test
        x.child[k] = true // tentatively
      }
      x.send (k, t) // for k == u probe, otherwise echo
      if x.demo { x.Write() }
      mutex.Unlock()
      done <- 0
    }(i)
  }
  for i := uint(0); i < x.n; i++ { // wait for all processes
    <-done                         // to be completed
  }
}
