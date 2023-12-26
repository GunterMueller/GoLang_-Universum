package dgra

// (c) Christian Maurer   v. 231105 - license see µU.go
//
// >>> Dijkstra-Scholten: Termination Detection for Diffusing Computations
//     Inf. Proc. Letters 11 (1980) 1-4

import (
//  "µU/ker"
  . "µU/obj"
  "µU/time"
//  "µU/ego"
  "µU/rand"
  "µU/atomic";
//  "µU/errh"
  "µU/vtx"
//  "µU/edg"
//  "µU/host"
  "µU/nchan"
)
const (
  message = uint(iota)
  signal
)

func sleep() {
  time.Msleep (rand.Natural (500))
}

func (x *distributedGraph) DijkstraScholten (o Op) {
  x.Op = o
  x.Trav (func (a any) { v := a.(vtx.Vertex)
                         x.Ex (v)
                         n := v.Val()
//                         ch := make ([]nchan.NetChannel, 0)
                         for i := uint(0); i < x.NumNeighboursIn(); i++ {
                           nb := x.NeighbourIn(i).(vtx.Vertex)
                           if nb != nil {
                             nn := nb.(vtx.Vertex).Val()
                             println (nn, "->", n)
                             x.ch = append (x.ch, nchan.NewN (uint(0), "jupiter",
                                                              nchan.Port (uint(0), n, nn, 1),
                                                              true))
                           }
                         }
                       })
//    x.connectN (uint(0), true) // the messages have type uint
/*/
  }
  for i := uint(0); i < x.n; i++ {
    go x.do (i)
  }
  if x.me == x.root {
    for i := uint(0); i < x.n; i++ {
      x.ch[i].Send (x.me + inf * message)
      atomic.Inc1 (&x.D[i])
// println ("sent ", x.me + inf * message)
      sleep()
    }
    x.Op (x.actVertex)
  }
  <-done
/*/
}

func (x *distributedGraph) do (i uint) {
  for {
    n := x.ch[i].Recv().(uint)
// println ("rcvd ", n)
    if n / inf == message {
      atomic.Inc1 (&x.C[i])
      x.mutex.Lock()
      if x.C[i] == 1 { // perform x.Op only once
        x.Op (x.actVertex)
      }
      x.mutex.Unlock()
      x.corn.Ins (i)
      if x.NumNeighboursOut() == 0 {
        x.mutex.Lock()
        if x.C[i] == x.NumNeighboursIn() { // received a message from all predecessors
          for x.C[i] > 0 {
            j := x.corn.Get().(uint)
            x.ch[j].Send (x.me + inf * signal)
            atomic.Dec (&x.C[i])
// println ("sent ", x.me + inf * signal)
            sleep()
          }
          x.mutex.Unlock()
          break
        } else { // x.C[i] < x.NumNeighboursIn(); wait for outstanding messages
          x.mutex.Unlock()
        }
      } else { // x.Outgoing (i) > 0
        x.mutex.Lock()
        if x.C[i] == x.NumNeighboursIn() {
          for j := uint(0); j < x.n; j++ {
            if x.Outgoing (j) {
              x.ch[j].Send (x.me + inf * message)
// println ("sent ", x.me + inf * message)
              atomic.Inc1 (&x.D[i])
              sleep()
            }
          }
//        wait for the corresponding signals
        } else { // x.C[i] < x.NumNeighboursIn()
//        wait for outstanding messages
        }
        x.mutex.Unlock()
      }
    } else { // n / inf == signal
      atomic.Dec (&x.D[i])
      if x.me == x.root {
        if x.D[i] == 0 {
          break
        }
      } else {
        x.mutex.Lock()
        if x.D[i] == 0 {
          for x.C[i] > 0 {
            j := x.corn.Get().(uint)
            x.ch[j].Send (x.me + inf * signal)
            atomic.Dec (&x.C[i])
// println ("sent ", x.me + inf * signal)
            sleep()
          }
          x.mutex.Unlock()
          break
        } else { // x.D[i] > 0
          x.mutex.Unlock()
//        to keep the invariant C = 0 => D = 0 do not send any signals, before D = 0,
//        i.e. signals from all successors received, so wait for those
        }
      }
    }
  }
  done <- 0
}
