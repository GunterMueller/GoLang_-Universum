package dgra

// (c) Christian Maurer   v. 231217 - license see µU.go

// >>> Gary L. Peterson: An O(n log n) Unidirectional Algorithm
//     for the Circular Extrema Problem. ACM TOPLAS (1982), 758-762

func (x *distributedGraph) Peterson() uint {
  x.connect(uint(0))
  defer x.fin()
  out, in := uint(0), uint(1)
  if x.Outgoing(1) { in, out = out, in }
  tid := x.me
  for {
    x.ch[out].Send (tid); println ("sent tid ==", tid)
    ntid := x.ch[in].Recv().(uint); println ("recv ntid ==", ntid)
    if ntid == x.me {
      x.leader = x.me
      x.ch[out].Send (x.leader + inf)
      return x.leader
    }
    if ntid >= inf {
      x.leader = ntid - inf
      x.ch[out].Send (ntid); println ("sent tid ==", tid)
      return x.leader
    }
    if tid > ntid {
      x.ch[out].Send (tid); println ("sent tid ==", tid)
    } else {
      x.ch[out].Send (ntid)
    }
    nntid := x.ch[in].Recv().(uint); println ("recv nntid ==", nntid)
    if nntid == x.me {
      x.leader = x.me
      x.ch[out].Send (x.leader + inf); println ("sent x.leader + inf ==", x.leader + inf)
      return x.leader
    }
    if nntid >= inf {
      x.leader = nntid - inf
      x.ch[out].Send (nntid); println ("sent nntid == ", nntid)
      return x.leader
    }
    if ntid >= tid && ntid >= nntid {
      tid = ntid
    } else {
      break
    }
  }
  for {
    n := x.ch[in].Recv().(uint); println ("recv n ==", n)
    if n == x.me {
      x.leader = x.me
      x.ch[out].Send (x.leader + inf); println ("sent", x.leader + inf)
      r := x.ch[in].Recv().(uint); println ("recv r ==", r)
      return x.leader
    }
    if n >= inf {
      x.leader = n - inf
      x.ch[out].Send (n); print ("sent n = =", n)
      return x.leader
    }
    x.ch[out].Send (n); print ("sent n == ", n)
  }
}
