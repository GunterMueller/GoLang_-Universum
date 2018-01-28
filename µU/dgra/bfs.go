package dgra

// (c) Christian Maurer   v. 171226 - license see µU.go

// >>> Algorithmus von Zhu, Y., Cheung, T.-Y.: A New Distributed Breadth-First-Seach Algorithm
//     Inform. Proc. Letters 25 (1987), 329-333
//
// >>> "visited" used for "echoed"

import
  . "µU/obj"
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

func (x *distributedGraph) bfs (o Op) {
  x.connect (uint(0))
  defer x.fin()
  x.Op = o
  m := inf * x.me
  if x.me == x.root {
    x.labeled, x.parent, x.distance = true, x.root, 0
    for i := uint(0); i < x.n; i++ {
      x.child[i] = false
      x.ch[i].Send (label + 8 * x.distance + m)
// x.log("label to", x.nr[i])
      x.visited[i] = false
      x.sendTo[i] = true
    }
  }
  done := make(chan int, x.n)
  for j := uint(0); j < x.n; j++ {
    go func (i uint) {
      loop:
      for {
        t := x.ch[i].Recv().(uint)
// // x.log("recv from", x.nr[i])
        if t % 8 == term {
// if x.me == x.root { x.log2(x.me, "recv term from", x.nr[i]) }
          break loop
        } else {
          x.chan1 <- t
        }
      }
      done <- 1
// x.log2("sent done", i, "/", x.nr[i])
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
if x.numSendTos() > 0 { panic ("oops") }
          x.ch[j].Send (end + m)
// x.log("end to", x.nr[j])
        } else {
          x.ch[j].Send (keepon + m)
// x.log("keepon to", x.nr[j])
        }
      } else {
        if x.parent == x.nr[j] {
          for k := uint(0); k < x.n; k++ {
            if x.sendTo[k] {
              x.ch[k].Send (label + 8 * x.distance + m)
              x.visited[k] = false
            }
          }
        } else { // x.parent =! x.nr[j]
          x.ch[j].Send (stop + m)
// x.log("stop to", x.nr[j])
        }
      }
    case keepon:
      x.visited[j] = true
      x.child[j] = true
    case stop:
      if x.nr[j] == x.parent {
        for k := uint(0); k < x.n; k++ {
          if x.child[k] {
            x.ch[k].Send (stop + m)
// x.log("stop to", x.nr[k])
          }
        }
        x.Op (x.actVertex)
        for k := uint(0); k < x.n; k++ {
          x.ch[k].Send (term)
// if x.nr[k] == x.root { x.log2(x.me, "sent term to", x.nr[k]) }
// x.log("term to", x.nr[k])
        }
        for k := uint(0); k < x.n; k++ {
// x.logo("got done", k)
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
              x.ch[k].Send (label + 8 * x.distance + m)
// x.log("label to", x.nr[k])
              x.visited[k] = false
            }
          }
        } else {
          x.ch[x.channel(x.parent)].Send (keepon + m)
// x.log("keepon to", x.nr[k])
        }
      }
    } else { // numSendTos() == 0
      if x.me == x.root {
        for k := uint(0); k < x.n; k++ {
          if x.child[k] {
            x.ch[k].Send (stop + m)
// x.log("stop to", x.nr[k])
          }
        }
        x.Op (x.actVertex)
        for k := uint(0); k < x.n; k++ {
          x.ch[k].Send (term)
// x.log2("sent term to", x.nr[k], "on", k)
// x.log("term to", x.nr[k])
        }
        for k := uint(0); k < x.n; k++ {
// x.log("got done", k)
          <-done
        }
        return
      } else {
        k := x.channel(x.parent); x.ch[k].Send (end + m)
// x.log("end to", x.nr[k])
      }
    }
  }
}
