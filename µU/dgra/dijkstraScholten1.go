package dgra

// (c) Christian Maurer   v. 200728 - license see µU.go
//
// >>> Dijkstra-Scholten: Termination Detection for Diffusing Computations
//     Inf. Proc. Letters 11 (1980) 1-4

import (
  "reflect"
  . "µU/obj"
  . "µU/atomic";
  "µU/perm"
  "µU/errh"
)

func (x *distributedGraph) DijkstraScholten1 (o Op) {
  x.connect (nil)
  x.Op = o
  x.tree.Clr()
  x.tree.Ins (x.actVertex)
  x.tree.Write()
  for i := uint(0); i < x.n; i++ {
    go x.do1 (i)
  }
  if x.me == x.root {
    bs := append (Encode(message), x.tree.Encode()...)
    p := perm.New (x.n)
    for i := uint(0); i < x.n; i++ {
      j := p.F(i)
//      x.ch[j].Send (bs)
      x.send (j, bs)
      Inc (&x.D)
    }
    x.Op (x.actVertex)
  }
  <-done
}

func (x *distributedGraph) do1 (i uint) {
  for {
errh.Error ("waiting for msg from", x.nr[i])
    a := x.ch[i].Recv()
    t := reflect.TypeOf (a)
    if t == nil { errh.Error0 ("nil"); break }
    bs := a.(Stream)
    ms := Decode(uint(0), bs[:c0])
    if ms == message {
      x.tree = x.decodedGraph(bs[c0:])
      x.tree.Write()
      Inc (&x.C)
      x.Mutex.Lock()
      if x.C == 1 { // insert actVertex and perform x.Op only once
        x.tree.Ex (x.nb[i])
        x.tree.Ins (x.actVertex)
        x.tree.Edge (x.directedEdge(x.nb[i], x.actVertex))
        x.tree.Write()
        x.Op (x.actVertex)
      }
      x.Mutex.Unlock()
      x.corn.Ins (i)
      if x.NumNeighboursOut() == 0 {
        x.Mutex.Lock()
        if x.C == x.NumNeighboursIn() {
//        received a message from all predecessors
          bs = append (Encode(signal), x.tree.Encode()...)
          for x.C > 0 {
            j := x.corn.Get().(uint)
//            x.ch[j].Send (bs)
            x.send (j, bs)
            Dec (&x.C)
          }
          x.Mutex.Unlock()
          break
        } else { // x.C < x.NumNeighboursIn()
          x.Mutex.Unlock()
//        wait for the outstanding messages
        }
      } else { // x.Outgoing (i) > 0
        x.Mutex.Lock()
        if x.C == x.NumNeighboursIn() {
          bs = append (Encode(message), x.tree.Encode()...)
          p := perm.New (x.n)
          for j := uint(0); j < x.n; j++ {
            k := p.F(j)
            if x.Outgoing (k) {
//              x.ch[k].Send (bs)
              x.send (k, bs)
              Inc (&x.D)
            }
          }
//        message from all predecessors received,
//        wait for the corresponding signals
        } else { // x.C < x.NumNeighboursIn()
//        wait for the outstanding messages
        }
        x.Mutex.Unlock()
      }
    } else { // ms == signal
      x.tree.Add (x.decodedGraph(bs[c0:]))
      x.tree.Write()
      Dec (&x.D)
      if x.me == x.root { // environment
        if x.D == 0 {
          break
        }
      } else { // inner node
        x.Mutex.Lock()
        if x.D == 0 {
          bs = append (Encode(signal), x.tree.Encode()...)
          for x.C > 0 {
            j := x.corn.Get().(uint)
//            x.ch[j].Send (bs)
            x.send (j, bs)
            Dec (&x.C)
          }
          x.Mutex.Unlock()
          break
        } else { // x.D > 0
          x.Mutex.Unlock()
//        to keep the invariant C = 0 => D = 0
//        do not send any signals, before D = 0,
//        i.e. signals from all successors received,
//        so wait for those
        }
      }
    }
  }
  done <- 0
}
