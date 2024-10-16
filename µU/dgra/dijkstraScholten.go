package dgra

// (c) Christian Maurer   v. 240113 - license see µU.go

import (
  "os"
  . "µU/obj"
  "µU/errh"
  "µU/rand"
  "µU/fmon"
)
const (
  message = uint(iota)
  signal
)
const
  Max = uint(1000)
var
  nSigs = make([]uint, 0)

/*/
func (x *distributedGraph) c (s string, n uint) {
  errh.Error (s, n); return
  errh.Error3 (s, n, " C", x.C[x.me], " D", x.D[x.me])
  errh.Hint2 ("C", x.C[x.me], " D", x.D[x.me])
}
/*/

func (x *distributedGraph) ds (a any, i uint) any {
  errh.Error0 ("waiting")
  var m uint
  x.awaitAllMonitors()
  n := a.(uint)
  j := x.channel (n)
  if n >= Max {
    for k := uint(0); k < x.n; k++ {
      if x.Outgoing (k) {
        x.mon[k].F (Max + x.me, message)
      }
    }
    errh.Error0 ("terminated")
    os.Exit (0)
  }
  if i == message {
    errh.DelHint()
    x.C[x.me]++
    x.corn[x.me].Ins (n)
// x.c ("recv msg from", n)
    if x.NumNeighboursOut() == 0 {
      x.C[x.me]--
// x.c ("E send sig to", x.nr[j])
      x.mon[j].F (x.me, signal)
    } else {
      if true { // rand.Natural (2) == 0 {
        p := make([]uint, 0)
        for a := uint(0); a < x.n; a++ {
          if x.Outgoing (a) { p = append (p, a) }
        }
        if len(p) == 0 {
          m = 0
        } else { 
          k := rand.Natural (uint(len(p))); m = p[k]
        }
        x.D[x.me]++
// x.c ("A send msg to", x.nr[m])
        x.mon[m].F (x.me, message)
      } else { // k == 1
        x.C[x.me]--
        m = x.corn[x.me].Get().(uint)
// x.c ("C send sig to", x.nr[m])
        x.mon[m].F (x.me, signal)
      }
    }
  } else { // i == signal
    errh.DelHint()
    x.D[x.me]--
// x.c ("recv sig from", n)
    if x.D[x.me] == 0 {
      if x.me == x.root {
        nSigs = append (nSigs, n)
        if uint(len(nSigs)) == x.n {
          errh.Error0 ("done"); done <- 0
        }
      }
    } else { // x.D[x.me] > 0
      if true { // rand.Natural (2) == 0 {
        p := make([]uint, 0)
        for a := uint(0); a < x.n; a++ {
          if x.Outgoing (a) { p = append (p, a) }
        }
        if len(p) == 0 {
          m = 0
        } else { 
          k := rand.Natural (uint(len(p))); m = p[k]
        }
        x.D[x.me]++
// x.c ("B send msg to", x.nr[m])
        x.mon[m].F (x.me, message)
      } else { // k == 1
        x.C[x.me]--
        m = x.corn[x.me].Get().(uint)
// x.c ("D send sig to", x.nr[m])
        x.mon[m].F (x.me, signal)
      }
    }
  }
  return x.me
}

func (x *distributedGraph) DijkstraScholten (o Op) {
  go func() {fmon.New (uint(0), 2, x.ds, AllTrueSp, x.actHost, p0 + uint16(2 * x.me), true)}()
  for i := uint(0); i < x.n; i++ {
    x.mon[i] = fmon.New (uint(0), 2, x.ds, AllTrueSp, x.actHost, p0 + uint16(2 * x.nr[i]), false)
  }
  defer x.finMon()
  x.awaitAllMonitors()
  x.Op = o
  if x.me == x.root {
    x.C[x.me] = 0
    for i := uint(0); i < x.n; i++ {
      x.D[x.me]++
// x.c ("send msg to", x.nr[i])
      x.mon[i].F (x.root, message)
    }
  }
  <-done
  if x.me == x.root {
    for i := uint(0); i < x.n; i++ {
      x.mon[i].F (Max + x.me, message)
    }
  }
  errh.Error0 ("terminated")
}
