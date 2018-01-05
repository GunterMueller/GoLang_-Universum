package dgra

// (c) Christian Maurer   v. 171125 - license see nU.go

func (x *distributedGraph) peterson() {
  x.connect(uint(0))
  defer x.fin()
  out, in := uint(0), uint(1)
  if x.Graph.Outgoing(1) { in, out = out, in }
  tid := x.me
  for {
println ("send", tid)
    x.ch[out].Send (tid)
    ntid := x.ch[in].Recv().(uint)
println ("recv", ntid)
    if ntid == x.me {
      x.leader = x.me
println ("send", x.leader + inf)
      x.ch[out].Send (x.leader + inf)
      return
    }
    if ntid >= inf {
      x.leader = ntid - inf
println ("send", ntid)
      x.ch[out].Send (ntid)
//      tid = x.me
      return
    }
    if tid > ntid {
println ("send", tid)
      x.ch[out].Send (tid)
    } else {
println ("send", ntid)
      x.ch[out].Send (ntid)
    }
    nntid := x.ch[in].Recv().(uint)
    if nntid == x.me {
      x.leader = x.me
println ("send", x.leader + inf)
      x.ch[out].Send (x.leader + inf)
      return
    }
    if nntid >= inf {
      x.leader = nntid - inf
println ("send", nntid)
      x.ch[out].Send (nntid)
      return
    }
    if ntid >= tid && ntid >= nntid {
      tid = ntid
    } else {
      break
    }
  }
  for {
    n := x.ch[in].Recv().(uint)
    if n == x.me {
      x.leader = x.me
      x.ch[out].Send (x.leader + inf)
      affe := x.ch[in].Recv().(uint)
println ("recv", affe - inf)
      return
    }
    if n >= inf {
      x.leader = n - inf
println ("send", n)
      x.ch[out].Send (n)
      return
    }
println ("send", n)
    x.ch[out].Send (n)
  }
}
