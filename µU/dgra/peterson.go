package dgra

// (c) Christian Maurer   v. 170423 - license see µU.go

// >>> Gary L. Peterson: An O(n log n) Unidirectional Algorithm
//     for the Circular Extrema Problem. ACM TOPLAS (1982), 758-762

/*
Es können nie zwei im Ring hintereinander liegende akive Prozesse in die nächste Phase gelangen.

id     n-1 -----> n -----> n+1 
tid     a         b         _
ntid    _         a         b
nntid   _      max(a,_)  max(b,a)

Denn für Prozesse $n$ und $n+1$ mit den temporären Identitäten $id_n = a$ und $id_{n+1} = b\neq a$
hätte man in diesem Fall $a\geq b$ und $b\geq a$, also -- im Widerspruch zu $a\neq b$ -- $a=b$.
*/

func (x *distributedGraph) peterson() {
  x.connect(uint(0))
  defer x.fin()
  out, in := uint(0), uint(1); if x.Graph.Outgoing(1) { in, out = out, in }
  m, tid := inf, x.me
  for {
    x.ch[out].Send (tid)
    ntid := x.ch[in].Recv().(uint)
    if ntid == x.me {
      x.ch[out].Send (ntid + m)
      x.leader = x.me
      return
    }
    if ntid >= m {
      x.ch[out].Send (ntid)
      tid = x.me
      x.leader = ntid - m
      return
    }
    if tid > ntid {
      x.ch[out].Send (tid)
    } else {
      x.ch[out].Send (ntid)
    }
    nntid := x.ch[in].Recv().(uint)
    if nntid == x.me {
      x.ch[out].Send (nntid + m)
      x.leader = x.me
      return
    }
    if nntid >= m {
      x.ch[out].Send (nntid)
      x.leader = nntid - m
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
      x.ch[out].Send (n + m)
      x.ch[in].Recv()
      x.leader = x.me
      return
    }
    if n >= m {
      x.ch[out].Send (n)
      x.leader = n - m
      return
    }
    x.ch[out].Send (n)
  }
}
