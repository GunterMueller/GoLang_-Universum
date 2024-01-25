package dgra

// (c) Christian Maurer   v. 231220 - license see nU.go

import (
  . "nU/obj"
  "nU/vtx"
  "nU/fmon"
)

func (x *distributedGraph) r1 (a any, i uint) any {
  x.awaitAllMonitors()
  bs := a.(Stream)
  s := Decode(uint(0), bs[:C0]).(uint)
  j := x.channel(s)
  switch i {
  case visit:
    x.visited[j] = true
  case discover:
    for k := uint(0); k < x.n; k++ {
      if k != j {
        x.mon[k].F(Encode(x.me), visit)
      }
    }
    x.cycle = x.decodedGraph(bs[C0:])
    x.cycle.Ins (x.actVertex)
    a0, a1 := x.cycle.Get2()
    e := x.edge (a0.(vtx.Vertex), a1.(vtx.Vertex)) // colokal == lokal
    x.cycle.Edge (e)
    x.cycle.Write()
    pause()
    for k := uint(0); k < x.n; k++ {
      if k != j && ! x.visited[k] {
        x.visited[k] = true
        x.child[k] = true
        bs = append(Encode(x.me), x.cycle.Encode()...)
        bs = x.mon[k].F(bs, discover).(Stream)
        x.cycle = x.decodedGraph(bs[C0:])
        x.cycle.Write()
        pause()
      }
    }
    return append(Encode(x.me), x.cycle.Encode()...)
  case distribute:
    x.cycle = x.decodedGraph(bs)
    x.cycle.Write()
    pause()
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
        x.mon[k].F(bs, distribute)
      }
    }
    done <- 0
  }
  return nil
}

func (x *distributedGraph) Ring1() {
  go func() {
    fmon.New (nil, 3, x.r1, AllTrueSp,
              x.actHost, p0 + uint16(3 * x.me), true)
  }()
  for i := uint(0); i < x.n; i++ {
    x.mon[i] = fmon.New (nil, 3, x.r1, AllTrueSp,
                         x.host[i], p0 + uint16(3 * x.nr[i]), false)
  }
  defer x.finMon()
  x.awaitAllMonitors()
  x.cycle.Clr()
  if x.me == x.root {
    x.cycle.Ins (x.actVertex)
    x.cycle.Mark (x.actVertex)
    x.cycle.Write()
    pause()
    for k := uint(0); k < x.n; k++ {
      x.mon[k].F(Encode(x.me), visit)
    }
    for k := uint(0); k < x.n; k++ {
      if ! x.visited[k] {
        x.visited[k] = true
        x.child[k] = true
        x.cycle.Ex (x.actVertex)
        bs := append(Encode(x.me), x.cycle.Encode()...)
        bs = x.mon[k].F(bs, discover).(Stream)
        x.cycle = x.decodedGraph(bs[C0:])
        x.cycle.Write()

      }
    }
    v := x.cycle.Get().(vtx.Vertex)
    x.cycle.Ex2 (v, x.actVertex)
    x.cycle.Edge (x.edge (v, x.actVertex))
    x.cycle.Write()
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
        x.mon[k].F(x.cycle, distribute)
      }
    }
  } else {
    <-done
  }
}
