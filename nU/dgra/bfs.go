package dgra

// (c) Christian Maurer   v. 171125 - license see nU.go

import . "nU/obj"

const (LABEL = uint(iota); KEEPON; STOP; END; TERM)
const high = 1<<16

func (x *distributedGraph) sendTos() uint {
  v := uint(0)
  for k := uint(0); k < x.n; k++ {
    if x.sendTo[k] { v++ }
  }
  return v
}

func (x *distributedGraph) allSendTosEchoed() bool {
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
  h := high * x.me
  if x.me == x.root {
    x.labeled, x.parent, x.time = true, x.me, 0
    for i := uint(0); i < x.n; i++ {
      x.sendTo[i], x.child[i] = true, false
      x.ch[i].Send (LABEL + 8 * 0 + h)
      x.visited[i] = false
    }
  }
  done := make(chan int, x.n)
  for j := uint(0); j < x.n; j++ {
    go func (i uint) {
      loop:
      for {
        t := x.ch[i].Recv().(uint)
        if t % 8 == TERM {
          break loop
        } else {
          x.chan1 <- t
        }
      }
      done <- 1
    }(j)
  }
  for {
    t := <-x.chan1; j := x.channel (t / high)
    t = t % high; x.distance = t / 8
    msg := t % 8
    switch msg {
    case LABEL:
      if ! x.labeled {
        x.labeled, x.parent, x.time = true, x.nr[j], x.distance + 1
        for k := uint(0); k < x.n; k++ {
          x.sendTo[k] = k != j
        }
        if x.sendTos() == 0 {
          x.ch[j].Send (END + h)
        } else {
          x.ch[j].Send (KEEPON + h)
        }
      } else {
        if x.parent == x.nr[j] {
          for k := uint(0); k < x.n; k++ {
            if x.sendTo[k] {
              x.ch[k].Send (LABEL + 8 * x.time + h)
              x.visited[k] = false
            }
          }
        } else { // x.parent =! x.nr[j]
          x.ch[j].Send (STOP + h)
        }
      }
    case KEEPON:
      x.visited[j] = true
      x.child[j] = true
    case STOP:
      if x.nr[j] == x.parent {
        for k := uint(0); k < x.n; k++ {
          if x.child[k] {
            x.ch[k].Send (STOP + h)
          }
        }
        x.Op (x.me)
        for k := uint(0); k < x.n; k++ {
          x.ch[k].Send(TERM)
        }
        for k := uint(0); k < x.n; k++ {
          <-done
        }
        return
      } else {
        x.visited[j] = true
        x.sendTo[j] = false
      }
    case END:
      x.visited[j] = true
      x.child[j], x.sendTo[j] = true, false
    case TERM:
      // ignore
    }
    if x.sendTos() == 0 {
      if x.parent == x.me {
        for k := uint(0); k < x.n; k++ {
          if x.child[k] {
            x.ch[k].Send (STOP + h)
          }
        }
        x.Op (x.me)
        for k := uint(0); k < x.n; k++ {
          x.ch[k].Send(TERM)
        }
        for k := uint(0); k < x.n; k++ {
          <-done
        }
        return
      } else {
        k := x.channel(x.parent); x.ch[k].Send (END + h)
      }
    } else {
      if x.allSendTosEchoed() {
        if x.parent == x.me {
          for k := uint(0); k < x.n; k++ {
            if x.sendTo[k] {
              x.ch[k].Send (LABEL + 8 * x.time + h)
              x.visited[k] = false
            }
          }
        } else {
          k := x.channel(x.parent); x.ch[k].Send (KEEPON + h)
        }
      }
    }
  }
}
