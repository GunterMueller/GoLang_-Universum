package dgra

// (c) Christian Maurer  v. 180902 - license see µU.go
//
// >>> Korach, E., Moran, S., Zaks, S.: Tight Lower and Upper Bounds for
//     Some Distributed Algorithms for a Complete Network of Processors.
//     PODC '84, ACM symposion (1984) 199-207
//
// >>> Diese Portierung ist vermutlich noch nicht korrekt.

import (
  . "µU/obj"
  "µU/dgra/status"
  . "µU/dgra/msg"
)
const (
  King = byte(iota)
  Citizen
)

func (x *distributedGraph) channelKmz (s Stream, m *Message) uint {
  j := Decode(uint(0), s[:C0()]).(uint) // number of channel on
  (*m).Decode (s[C0():])                // which m was received
  return j
}

func (x *distributedGraph) korachMoranZaks() {
  for i := byte(0); i < NTypes; i++ {
    x.msgchan[i] = make(chan Stream)
  }
  for i := uint(0); i < x.n; i++ {
    go func (j uint) {
      m := x.ch[j].Recv().(Message)
      k := m.Type()
      x.msgchan[k] <- append(Encode(j), m.Encode()...)
    }(i)
  }
  x.pik = status.New()
  x.pik.Set (-1, x.me)
  if x.me == x.root { x.pik.Inc() }
  for i := uint(0); i < x.n; i++ {
    x.unused[i] = true
  }
  m := NewMsg()
  var j uint
// algorithm for a King:
  for x.state == King {
    u := x.n
    for i := uint(0); i < x.n; i++ {
      if x.unused[i] {
        u = i
        break
      }
    }
    if u == x.n { break }
    x.unused[u] = false
    m.Set (Ask, x.pik)
    x.ch[u].Send (m)
    label: select {
    case bs := <-x.msgchan[Ask]:
      j = x.channelKmz (bs, &m)
      if m.Status().Less (x.pik) {
        goto label
      } else {
        x.state = Citizen
      }
    case bs := <-x.msgchan[Accept]:
      j = x.channelKmz (bs, &m)
      x.child[j] = true
      if m.Status().Phase() < x.pik.Phase() {
        m.Set (Update, x.pik)
        x.ch[j].Send (m)
      } else { // pik.Phase() == m.Status().Phase()
               // as Acc only iff m.Status.Phase() <= x.phik.Phase()
               // => x.me < m.Status().Id()
        x.pik.Inc()
        for i := uint(0); i < x.n; i++ {
          if x.child[i] {
            m.Set (Update, x.pik)
            x.ch[i].Send (m)
          }
        }
      }
    case <-x.msgchan[YourCitizen]:
      // do nothing and enter the for-loop agains
    default:
      panic ("kmz 1")
    }
  }
  if x.state == King { // no more unused neighbours
    x.leader = x.me
    m.Set (Leader, x.pik)
    for i := uint(0); i < x.n; i++ {
      x.ch[i].Send (m)
    }
    return
  }
// algorithm for a Citizen:
// received an Ask(s)-message along edge j,
// which changed status from King to Citizen.
  x.child[j] = false
  x.parentChannel = j
  m.Set (Accept, x.pik)
  x.ch[j].Send (m)
  select {
  case bs := <-x.msgchan[Update]:
    j = x.channelKmz (bs, &m)
  }
  x.pik.Copy (m.Status()) // phase of pik increased by >= 1
  m.Set (Update, x.pik)
  for i := uint(0); i < x.n; i++ {
    if x.child[i] {
      x.ch[i].Send (m)
    }
  }
  jAsk := x.n
  for {
    select {
    case bs := <-x.msgchan[Ask]:
      j = x.channelKmz (bs, &m)
      jAsk = j
      s := m.Status()
      if x.pik.Less (s) {
        x.ch[x.parentChannel].Send (m)
        for x.pik.Less (s) {
          select {
          case bs := <-x.msgchan[Update]:
            j = x.channelKmz (bs, &m)
            if j != x.parentChannel { panic("kmz 2") }
            x.pik.Copy (m.Status()) // phase of pik increased by >= 1
            m.Set (Update, x.pik)
            for i := uint(0); i < x.n; i++ {
              if x.child[i] {
                x.ch[i].Send (m)
              }
            }
          }
        }
      }
      if s.Eq(x.pik) && j != x.parentChannel { // && j is not a tree edge
        m.Set (YourCitizen, x.pik)
        x.ch[j].Send (m)
      }
      if s.Less (x.pik) && j != x.parentChannel { // && j is not a tree edge
        // Ask-message is ignored
      }
    case bs := <-x.msgchan[Update]:
      j = x.channelKmz (bs, &m)
      x.pik.Copy (m.Status) // phase of pik increased by >= 1
      m.Set (Update, x.pik)
      for i := uint(0); i < x.n; i++ {
        if x.child[i] {
          x.ch[i].Send (m)
        }
      }
    case bs := <-x.msgchan[Accept]:
      j = x.channelKmz (bs, &m)
      if j == x.parentChannel {
// received an Accept(s)-message on parentChannel
// this Accept must be a response to an Ask-Message received on channel jAsk
        x.child[j] = true
        x.parentChannel = jAsk
        select {
        case bs := <-x.msgchan[Update]:
          j = x.channelKmz (bs, &m)
          x.pik.Copy (m.Status) // phase of pik increased by >= 1
          m.Set (Update, x.pik)
          for i := uint(0); i < x.n; i++ {
            if x.child[i] {
              x.ch[i].Send (m)
            }
          }
        }
      } else { // j != x.parentChannel <=> j is not a tree edge
        x.child[j] = true
        m.Set (Update, x.pik)
        x.ch[j].Send (m)
      }
    case bs := <-x.msgchan[Leader]:
      j = x.channelKmz (bs, &m)
      x.leader = m.Status().Id()
      return
    case <-x.msgchan[YourCitizen]:
      panic("kmz 3")
    }
  }
}
