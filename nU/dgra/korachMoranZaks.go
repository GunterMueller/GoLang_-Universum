package dgra

// (c) Christian Maurer  v. 180819 - license see µU.go
//
// >>> Korach, E., Moran, S., Zaks, S.: Tight Lower and Upper Bounds for
//     Some Distributed Algorithms for a Complete Network of Processors.
//     PODC '84, ACM symposion (1984) 199-207
//
// >>> attempt to port this algorithm to Go XXX but still a lot of work TODO

import (
  . "nU/obj"
  "nU/dgra/status"
  . "nU/dgra/msgkmz"
)
const (
  King = uint(iota)
  Citizen
)
var (
  state uint // King or Citizen
  pik status.Status // phase and identity of the king of the calling process
  unused []bool
  sons = make([]uint, 0)
  father_edge uint // channel-number to parent
  msgchan = make(chan Stream)
)

func (x *distributedGraph) recv() {
  for i := uint(0); i < x.n; i++ {
    go func (j uint) {
      m := x.ch[i].Recv().(Message)
      msgchan <- append (Encode(i), m.Encode()...)
    }(i)
  }
}

func (x *distributedGraph) ssearch (kind byte) {
  // TODO
}

func (x *distributedGraph) king() {
  if state != King { panic("no king") }
  pik = status.New (0, x.me)
  for i := uint(0); i < x.n; i++ {
    unused[i] = true
  }
  m := NewMsg (NKinds, status.New(0, 0))
  for state == King {
    u := x.n
    for i := uint(0); i < x.n; i++ {
      if unused[i] {
        u = i
        break
      }
    }
    if u == x.n { break }
    unused[u] = false
    x.ch[u].Send (NewMsg (Ask, pik))
  label:
    s := <-msgchan
    j := Decode(uint(0), s[:C0]).(uint) // channel number from which the message was received
    m.Decode (s[C0:])
    switch m.Kind() { case Ask:
      if m.Status().Less (pik) {
        goto label
      } else {
        state = Citizen
      }
    case Accept:
      sons = append (sons, j)
      if m.Status().Phase() < pik.Phase() {
        x.ch[j].Send (NewMsg (Update, pik))
      } else { // pik.Phase() <= m.Status().Phase()
        pik.Inc()
        for i := range sons {
          x.ch[i].Send (NewMsg (Update, pik))
        }
      }
    case YourCitizen:
      // do nothing and enter the for-loop agains
    default:
      panic ("should not happen")
    }
  }
  if state == Citizen {
    x.citizen (m)
  } else { // no more unused neighbours
    x.leader = x.me
    m := NewMsg (Leader, pik)
// send m to all vertices in the graph
    for i := uint(0); i < x.n; i++ {
      x.ch[i].Send (m)
// TODO transfer the message by traversing the whole graph
    }
  }
}

func (x *distributedGraph) processAsk (m Message) {
// just received an Ask(s)-message along edge e.
  e := uint(97) // XXX
  p := m.Status()
  if pik.Less (p) {
    x.ch[father_edge].Send (m)
    var m1 Message // ssearch (Update, Accept)
    switch m1.Kind() { case Update:
/*
// just received an Update(s)-message along father_edge.
//  pik.Set (s) // phase of pik increased by >= 1
      m := msgkmz.New (msgkmz.Update, pik)
      for i := range sons {
        x.ch[i].Send (m)
      }
*/
      if m.Status().Eq (pik) { // XXX && e is not a tree edge { // TODO
        x.ch[e].Send (NewMsg(YourCitizen, pik))
      }
    case Accept:
// just received an Accept(s)-message along edge e'.
    }
  }
  if m.Status().Eq (pik) { // XXX && e is not a tree edge { // TODO
    x.ch[e].Send (NewMsg(YourCitizen, pik))
  }
  if m.Status().Less (pik) { // XXX && e is not a tree edge { // TODO
    // Ask-message is ignored
  }
}

func (x *distributedGraph) citizen (m Message) {
// just received an Ask(s)-message along edge e,
// which changed status from King to Citizen.
  e := uint(97) // XXX
  for i := uint(0); i < x.n; i++ {
    if sons[i] == e {
      sons = append (sons[:i], sons[i+1:]...)
      break
    }
  }
  father_edge = e
  x.ch[e].Send (NewMsg (Accept, pik))
//  m = x.ssearch (Update) // XXX
// now m = Update(s) was sent along e
  s := m.Status()
  pik = status.New (s.Phase(), s.Id()) // phase of pik increased by >= 1
  for i := range sons { // send Update to all sons
    x.ch[i].Send (NewMsg (Update, pik))
  }
  for {
//    receive (m) // m will be one of Ask, Update, Accept, Leader
    switch m.Kind() {
    case Ask:
      x.processAsk (m)
    case Update:
// just received an Update(s)-message along father_edge.
      pik = status.New (m.Status().Phase(), m.Status().Id()) // phase of pik increased by >= 1
      m := NewMsg (Update, pik)
      for i := range sons {
        x.ch[i].Send (m)
      }
    case Accept:
      if false { // e' is not a tree edge { // XXX what is e' XXX ?
// just received an Accept(s)-message along edge e which is not a tree edge.
        e := uint(97) // XXX
        sons = append (sons, e)
        m := NewMsg (Update, pik)
        x.ch[e].Send (m)
      } else {
// just received an Accept(s)-message along father_edge.
// This Accept must be a response to an Ask-Message received along edge e.
        e := uint(97) // XXX
        father_edge = e
        sons = append (sons, father_edge)
        var m2 Message //      m2 := x.ssearch (Update)
// just received an Update(s)-message along father_edge.
        s := m2.Status()
        pik = status.New (s.Phase(), s.Id()) // phase of pik increased by >= 1
        m := NewMsg (Update, pik)
        for i := range sons {
          x.ch[i].Send (m)
        }
      }
    }
    if m.Kind() == Leader {
      break
    }
  }
}
