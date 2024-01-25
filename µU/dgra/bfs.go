package dgra

// (c) Christian Maurer   v. 231229 - license see µU.go

// >>> Algorithmus von Zhu, Y., Cheung, T.-Y.: A New Distributed Breadth-First-Seach Algorithm
//     Inform. Proc. Letters 25 (1987), 329-333
//
// >>> "visited" used for "echoed"

const (
  label = uint(iota)
  keepon
  stop
  end
  term
)

func (x *distributedGraph) numSendTos() uint {
  s := uint(0)
  for k := uint(0); k < x.n; k++ {
    if x.sendTo[k] {
      s++
    }
  }
  return s
}

func (x *distributedGraph) allSendTosVisited() bool {
  for k := uint(0); k < x.n; k++ {
    if x.sendTo[k] && ! x.visited[k] {
      return false
    }
  }
  return true
}

func (x *distributedGraph) Bfs() {
  x.connect (uint(0))
  defer x.fin()
  m := inf * x.me
  if x.me == x.root {
    x.labeled, x.parent, x.distance = true, x.root, 0
    for i := uint(0); i < x.n; i++ {
      x.child[i] = false
      x.send (i, label + 8 * x.distance + m)
      x.visited[i] = false
      x.sendTo[i] = true
    }
  }
  done := make(chan int, x.n)
  for j := uint(0); j < x.n; j++ {
    go func (i uint) {
      loop:
      for {
        t := x.recv (i).(uint)
        if t % 8 == term {
          break loop
        } else {
          x.chan1 <- t
        }
      }
      done <- 1
    }(j)
  }
  for {
    t := <-x.chan1
    j := x.channel (t / inf)
    t %= inf
    x.distance = t / 8
    switch t % 8 {
    case label:
      if ! x.labeled {
        x.labeled = true
        x.parent = x.nr[j]
        x.distance++
        for k := uint(0); k < x.n; k++ {
          x.sendTo[k] = k != j
        }
        if x.n == 1 { // no neighbours != parent
          x.send (j, end + m)
        } else {
          x.send (j, keepon + m)
        }
      } else {
        if x.parent == x.nr[j] {
          for k := uint(0); k < x.n; k++ {
            if x.sendTo[k] {
              x.send (k, label + 8 * x.distance + m)
              x.visited[k] = false
            }
          }
        } else { // x.parent =! x.nr[j]
          x.send (j, stop + m)
        }
      }
    case keepon:
      x.visited[j] = true
      x.child[j] = true
    case stop:
      if x.nr[j] == x.parent {
        for k := uint(0); k < x.n; k++ {
          if x.child[k] {
            x.send (k, stop + m)
          }
        }
        for k := uint(0); k < x.n; k++ {
          x.send (k, term)
        }
        for k := uint(0); k < x.n; k++ {
          <-done
        }
        return
      } else {
        x.visited[j] = true
        x.sendTo[j] = false
      }
    case end:
      x.child[j] = true
      x.visited[j] = true
      x.sendTo[j] = false
    }
    if x.numSendTos() > 0 {
      if x.allSendTosVisited() {
        if x.me == x.root {
          for k := uint(0); k < x.n; k++ {
            if x.sendTo[k] {
              x.send (k, label + 8 * x.distance + m)
              x.visited[k] = false
            }
          }
        } else {
          x.send (x.channel(x.parent), keepon + m)
        }
      }
    } else { // numSendTos() == 0
      if x.me == x.root {
        for k := uint(0); k < x.n; k++ {
          if x.child[k] {
            x.send (k, stop + m)
          }
        }
        for k := uint(0); k < x.n; k++ {
          x.send (k, term)
        }
        for k := uint(0); k < x.n; k++ {
          <-done
        }
        return
      } else {
        k := x.channel(x.parent)
        x.send (k, end + m)
      }
    }
  }
}
