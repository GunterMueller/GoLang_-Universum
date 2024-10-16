package dgra

// (c) Christian Maurer   v. 241007 - license see µU.go

// >>> Gary L. Peterson: An O(n log n) Unidirectional Algorithm
//     for the Circular Extrema Problem. ACM TOPLAS (1982), 758-762

func (x *distributedGraph) Peterson() {
  x.connect (uint(0))
  defer x.fin()
  in, out := uint(0), uint(1)
  if x.Graph.Outgoing(1) { in, out = out, in }
  tid := x.me
  for { // active
    x.send (out, tid)
    ntid := x.recv (in).(uint)
    if ntid == x.me {
      x.send (out, ntid + inf)
      x.leader = x.me
      return
    }
    if ntid >= inf {
      x.leader = ntid - inf
      x.send (out, ntid)
      return
    }
    if tid > ntid { // tid < inf
      x.send (out, tid)
    } else {
      x.send (out, ntid)
    }
    nntid := x.recv (in).(uint)
    if nntid == x.me {
      x.send (out, nntid + inf)
      x.leader = x.me
      return
    }
    if nntid >= inf {
      x.leader = nntid - inf
      x.send (out, nntid)
      return
    }
    if ntid >= tid && ntid >= nntid { // ntid < inf
      tid = ntid
    } else {
      break
    }
  }
  for { // relay
    n := x.recv (in).(uint)
    if n == x.me {
      x.leader = x.me
      x.send (out, x.leader + inf)
      x.recv (in)
      return
    }
    if n >= inf {
      x.leader = n - inf
      x.send (out, n)
      return
    }
    x.send (out, n)
  }
}
