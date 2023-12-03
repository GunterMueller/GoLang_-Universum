package dgra

// (c) Christian Maurer   v. 231105 - license see µU.go
//
// >>> Dijkstra-Scholten: Termination Detection for Diffusing Computations
//     Inf. Proc. Letters 11 (1980) 1-4

import (
//  "µU/ker"
  . "µU/obj"
  "µU/time"
  "µU/ego"
  "µU/rand"
  "µU/atomic";
  "µU/perm"
//  "µU/errh"
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
  me := ego.Me()
  n := uint(len(x.Esel()[me]))
  for i := uint(0); i < n; i++ {
    e := x.Esel()[me][i]
    println (e, " ")
/*/
  0 -> 1, 0 -> 3
  1 -> 4
  2 -> 1
  3 -> 4, 3 -> 6, 3 -> 7
  4 -> 5, 4 -> 6
  5 -> 3, 5 -> 8
  6 -> 7
  7 -> 8
  8 -> 2
/*/
//    x.connectN (uint(0), true) // the messages have type uint
  }
  for i := uint(0); i < x.n; i++ {
    go x.do (i)
  }
  if x.me == x.root {
    p := perm.New (x.n)
    for i := uint(0); i < x.n; i++ {
      x.ch[p.F(i)].Send (x.me + inf * message)
      atomic.Inc1 (&x.D)
// println ("sent ", x.me + inf * message)
      sleep()
    }
    x.Op (x.actVertex)
  }
  <-done
}

func (x *distributedGraph) do (i uint) {
  for {
    n := x.ch[i].Recv().(uint)
// println ("rcvd ", n)
    if n / inf == message {
      atomic.Inc1 (&x.C)
      x.mutex.Lock()
      if x.C == 1 { // perform x.Op only once
        x.Op (x.actVertex)
      }
      x.mutex.Unlock()
      x.corn.Ins (i)
      if x.NumNeighboursOut() == 0 {
        x.mutex.Lock()
        if x.C == x.NumNeighboursIn() { // received a message from all predecessors
          for x.C > 0 {
            j := x.corn.Get().(uint)
            x.ch[j].Send (x.me + inf * signal)
            atomic.Dec (&x.C)
// println ("sent ", x.me + inf * signal)
            sleep()
          }
          x.mutex.Unlock()
          break
        } else { // x.C < x.NumNeighboursIn(); wait for outstanding messages
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
// println ("sent ", x.me + inf * message)
              atomic.Inc1 (&x.D)
              sleep()
            }
          }
//        wait for the corresponding signals
        } else { // x.C < x.NumNeighboursIn()
//        wait for outstanding messages
        }
        x.mutex.Unlock()
      }
    } else { // n / inf == signal
      atomic.Dec (&x.D)
      if x.me == x.root {
        if x.D == 0 {
          break
        }
      } else {
        x.mutex.Lock()
        if x.D == 0 {
          for x.C > 0 {
            j := x.corn.Get().(uint)
            x.ch[j].Send (x.me + inf * signal)
            atomic.Dec (&x.C)
// println ("sent ", x.me + inf * signal)
            sleep()
          }
          x.mutex.Unlock()
          break
        } else { // x.D > 0
          x.mutex.Unlock()
//        to keep the invariant C = 0 => D = 0 do not send any signals, before D = 0,
//        i.e. signals from all successors received, so wait for those
        }
      }
    }
  }
  done <- 0
}
