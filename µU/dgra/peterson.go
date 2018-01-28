package dgra

// (c) Christian Maurer   v. 170423 - license see µU.go

// >>> Gary L. Peterson: An O(n log n) Unidirectional Algorithm
//     for the Circular Extrema Problem. ACM TOPLAS (1982), 758-762

func (x *distributedGraph) peterson() {
  x.connect(uint(0))
  defer x.fin()
  out, in := uint(0), uint(1)
  if x.Graph.Outgoing(1) { in, out = out, in }
  tid := x.me
  for {
    x.ch[out].Send (tid)
    ntid := x.ch[in].Recv().(uint)
    if ntid == x.me {
      x.leader = x.me
      x.ch[out].Send (x.leader + inf)
      return
    }
    if ntid >= inf {
      x.leader = ntid - inf
      x.ch[out].Send (ntid)
      return
    }
    if tid > ntid {
      x.ch[out].Send (tid)
    } else {
      x.ch[out].Send (ntid)
    }
    nntid := x.ch[in].Recv().(uint)
    if nntid == x.me {
      x.leader = x.me
      x.ch[out].Send (x.leader + inf)
      return
    }
    if nntid >= inf {
      x.leader = nntid - inf
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
      x.ch[in].Recv()
      return
    }
    if n >= inf {
      x.leader = n - inf
      x.ch[out].Send (n)
      return
    }
    x.ch[out].Send (n)
  }
}
