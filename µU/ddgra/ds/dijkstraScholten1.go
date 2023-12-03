package dgra

// (c) Christian Maurer   v. 231104 - license see µU.go
//
// >>> Dijkstra-Scholten: Termination Detection for Diffusing Computations
//     Inf. Proc. Letters 11 (1980) 1-4

import (
  . "µU/obj"
  "µU/atomic";
  "µU/time"
  "µU/N"
  "µU/perm"
)

func (x *distributedGraph) lor (s string, i uint) {
  return
  x.log4 (N.String (time.Secnsec()) + " " + s, x.me, "from", x.nr[i],
          "in =", x.C, "out =", x.D)
}

func (x *distributedGraph) los (s string, i uint) {
  return
  x.log4 (N.String (time.Secnsec()) + " " + s, x.me, " to ", x.nr[i],
          "in =", x.C, "out =", x.D)
}

func (x *distributedGraph) DijkstraScholten1 (o Op) {
  x.connect (uint(0)) // iff messages have type uint
  x.Op = o
  for i := uint(0); i < x.n; i++ {
    go x.do1 (i)
  }
  if x.me == x.root {
    p := perm.New (x.n)
    for i := uint(0); i < x.n; i++ {
      j := p.F(i)
      x.ch[j].Send (x.me + inf * message)
      atomic.Inc1 (&x.D)
x.los ("msg", j)
    }
x.log0 (N.String (time.Secnsec()) + N.String (x.me) + " all msgs sent")
    x.Op (x.actVertex)
  }
  <-done
}

func (x *distributedGraph) do1 (i uint) {
  for {
    n := x.ch[i].Recv().(uint)
    content := n % inf // content of received message
    if content != x.nr[i] {
      x.log2 ("content", content, "instead", x.nr[i])
    }
    if n / inf == message {
      if x.Outgoing (i) { x.log ("error: out to", x.nr[i])}
      atomic.Inc1 (&x.C)
x.lor ("msg", i)
      x.mutex.Lock()
      if x.C == 1 {
        x.Op (x.actVertex)
      }
      x.mutex.Unlock()
      x.corn.Ins (i)
      if x.NumNeighboursOut() == 0 {
        x.mutex.Lock()
        if x.C == x.NumNeighboursIn() {
// i.e.    from all incoming neighbours a message received
          for x.C > 0 {
            j := x.corn.Get().(uint)
            x.ch[j].Send (x.me + inf * signal)
            atomic.Dec (&x.C)
x.los ("sig", j)
          }
          x.mutex.Unlock()
          break
        } else { // x.inDefs < x.NumNeighboursIn()
//        messages from incoming neighbours outstanding,
//        i.e. wait for them, so remain in for-loop
          x.mutex.Unlock()
        }
      } else { // x.Outgoing (i) > 0
        x.mutex.Lock()
        if x.C == x.NumNeighboursIn() {
          p := perm.New (x.n)
          for j := uint(0); j < x.n; j++ {
            k := p.F(j)
            if x.Outgoing (k) {
              x.ch[k].Send (x.me + inf * message)
              atomic.Inc1 (&x.D)
x.los ("msg", k)
            }
          }
//        and now wait for the corresponding signals,
//        so remain in for-loop
        } else { // x.inDefs < x.NumNeighboursIn()
//        messages from incoming neighbours outstanding,
//        i.e. wait for them, so remain in for-loop
        }
        x.mutex.Unlock()
      }
    } else { // n/inf == signal
      if x.Incoming (i) { x.log ("error: in from", x.nr[i])}
      atomic.Dec (&x.D)
x.lor ("sig", i)
      if x.me == x.root { // environment
        if x.D == 0 {
x.log (N.String (time.Secnsec()) + " term by sig <", x.nr[i])
          break
        }
      } else { // inner node
        x.mutex.Lock()
        if x.D == 0 {
          for x.C > 0 {
            j := x.corn.Get().(uint)
            x.ch[j].Send (x.me + inf * signal)
            atomic.Dec (&x.C)
x.los ("sig", j)
          }
          x.mutex.Unlock()
          break
        } else { // x.outDefs > 0
          x.mutex.Unlock()
//        must not send any signals, before signals
//        from all outgoing neighbours received,
//        i.e. wait for them, so remain in for-loop
        }
      }
    }
  }
  done <- 0
}
