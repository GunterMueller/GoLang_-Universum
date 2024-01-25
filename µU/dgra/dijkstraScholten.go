package dgra

// (c) Christian Maurer   v. 240104 - license see µU.go
//
// >>> under development, not yet correct:
//     works for dgra.G4ds. but not for dgra.G8ds

import (
  "os"
  . "µU/obj"
  "µU/time"
  "µU/fmon"
  "µU/errh"
)
const (
  message = uint(iota)
  signal
)
const
  M = uint(1000)
var
  nSigs = make([]uint, 0)

func (x *distributedGraph) c0 (t string, a uint) {
  errh.Error (t, a); return
  _, _, s := time.ActTime()
  println (s, t, a); return
  errh.Error2 ("", s, t, a)
}

func (x *distributedGraph) c (t string, a uint) {
//  time.Sleep (10)
  x.c0 (t, a)
}

func (x *distributedGraph) d (a any, i uint) any {
  x.awaitAllMonitors()
  n := a.(uint)
  j := x.channel(n)
  switch i {
  case message:
    if n >= M {
x.c0 ("recv M from", n % M)
      for k := uint(0); k < x.n; k++ {
        if x.Outgoing (k) {
x.c ("send M to", x.nr[k])
          x.mon[k].F (M + x.me, message)
        }
      }
      errh.Error0 ("terminated")
      os.Exit (0)
    }
//    x.Op (x.me)
x.c0 ("recv msg from", n % M)
    x.C[x.me]++
    x.corn[x.me].Ins (n)
    x.visited[j] = true
    for k := uint(0); k < x.n; k++ {
      if x.Outgoing (k) && x.visited[k] {
        if x.corn[x.me].Empty() { panic ("affe") }
        x.C[x.me]--
        m := x.corn[x.me].Get().(uint)
x.c ("send sig to", x.nr[m])
        x.mon[m].F (x.me, signal)
      } else {
        x.visited[k] = true
        if x.Outgoing (k) {
          x.D[x.me]++
x.c ("send msg to", x.nr[k])
          x.mon[k].F (x.me, message)
        }
      }
    }
    if x.NumNeighboursOut() == 0 {
      x.C[x.me]--
//         x.corn[x.me].Get().(uint) == 2
      m := uint(1); if x.Incoming (0) { m = 0 } // XXX
x.c ("send sig to", x.nr[m])
      x.mon[m].F (x.me, signal)
    }
    if x.D[x.me] == 0 {
      if x.C[x.me] > 0 {
        if ! x.corn[x.me].Empty() {
          m := x.corn[x.me].Get().(uint)
x.c ("send sig to", x.nr[m])
          x.mon[m].F (x.me, signal)
        } 
      } 
    } 
  case signal:
x.c0 ("recv sig from", n)
    x.D[x.me]--
    if x.D[x.me] == 0 {
      if x.me == x.root {
        nSigs = append (nSigs, n)
        if uint(len(nSigs)) == x.n {
          errh.Error0 ("done")
          done <- 0
        }
      }
    } else {
      if ! x.corn[x.me].Empty() {
        x.C[x.me]--
        m := x.corn[x.me].Get().(uint)
x.c0 ("send sig to", x.nr[m])
        x.mon[m].F (x.me, signal)
      }
    }
  }
  return x.me
}

func (x *distributedGraph) DijkstraScholten (o Op) {
  x.visited[x.root] = true
  go func() {fmon.New (uint(0), 2, x.d, AllTrueSp, x.actHost, p0 + uint16(2 * x.me), true)}()
  for i := uint(0); i < x.n; i++ {
    x.mon[i] = fmon.New (uint(0), 2, x.d, AllTrueSp, x.actHost, p0 + uint16(2 * x.nr[i]), false)
  }
  defer x.finMon()
  x.awaitAllMonitors()
  x.Op = o
  if x.me == x.root {
    for i := uint(0); i < x.n; i++ {
      x.visited[i] = true
      x.D[x.me]++
x.c0 ("send msg to", x.nr[i])
      x.mon[i].F (x.root, message)
    }
  }
  <-done
  if x.me == x.root {
    for i := uint(0); i < x.n; i++ {
      x.mon[i].F (M + x.me, message)
    }
  }
  errh.Error0 ("terminated")
}
