package nelect

// (c) murus.org  v. 160101 - license see murus.go

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

func (x *netElection) peterson() int {
  id := int(x.me)
  tid := id
//  phase := uint(1)
//  state := active
  for {
    x.chR.Send (tid)
    ntid := x.chL.Recv().(int)
    x.log (tid, ntid)
    if ntid == id {
      x.chR.Send (-ntid) // announce
      return ntid
    }
    if ntid < 0 {
      x.chR.Send (ntid) // announce
      tid = id
      x.log (tid, ntid)
      return -ntid
    }
    if tid > ntid {
      x.chR.Send (tid)
    } else {
      x.chR.Send (ntid)
    }
    nntid := x.chL.Recv().(int)
    x.log (tid, ntid, nntid)
    if nntid == id {
      x.chR.Send (-nntid) // announce
      return nntid
    }
    if nntid < 0 {
      x.chR.Send (nntid)
      tid = id
      x.log (tid)
      return -nntid
    }
    if ntid >= tid && ntid >= nntid { // false, if tid == maximum
      tid = ntid
      x.log (tid)
//      phase++
    } else {
      break // goto relay (in the first active phase, if tid == maximum)
    }
  }
  x.log (tid)
//  state = relay
  for {
    r := x.chL.Recv().(int)
    if r == id {
      x.chR.Send (-r)
      tid = id
      x.log (tid)
      x.chL.Recv()
      return id
    }
    if r < 0 {
      x.chR.Send (r)
      tid = id
      x.log (tid)
      return -r
    }
    x.chR.Send (r)
  }
}

func (x *netElection) petersonImproved() int {
  id := int(x.me)
  tid := id
  for { // active
//    if x.bool { hint ("active") }
    x.chR.Send (tid)
    ntid := x.chL.Recv().(int)
    if ntid == id { // elected
      x.chR.Send (-ntid) // announce
      return ntid
    }
    if ntid < 0 {
      x.chR.Send (ntid)
      tid = id
      return -ntid
    }
    if tid < ntid {
      break // goto relay
    }
    x.chR.Send (tid)
    ntid = x.chL.Recv().(int)
    if ntid == id { // elected
      x.chR.Send (-ntid) // announce
      return ntid
    }
    if ntid < 0 { // elected
      x.chR.Send (ntid)
      tid = id
      return -ntid
    }
    if ntid >= tid {
      tid = ntid
      x.log (tid)
//      if x.bool { error ("next phase", tid) }
//      x.log (tid )
    } else {
      break // goto relay
    }
  }
  for { // relay
//    if x.bool { hint ("in relay") }
    r := x.chL.Recv().(int)
    if r == id {
      x.chR.Send (-r)
      tid = id
      x.chL.Recv()
      x.log (id) // as r < 0
      return id
    }
    if r < 0 {
      x.chR.Send (r)
      tid = id
      return -r
    }
    x.chR.Send (r)
  }
}
