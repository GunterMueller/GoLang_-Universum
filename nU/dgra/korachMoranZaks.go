package dgra

// (c) Christian Maurer  v. 190402 - license see µU.go
//
// >>> Algorithmus von Korach, Moran und Zaks
// >>> Diese Portierung ist noch in Arbeit !

import (
  "strconv"
  "time"
  . "nU/obj"
  "nU/dgra/status"
  . "nU/dgra/msg"
)
const (
  King = byte(iota)
  Citizen
)

func (x *distributedGraph) channelKmz (s Stream, m *Message) uint {
  j := Decode(uint(0), s[:C0()]).(uint) // number of channel on which m was received
  (*m).Decode (s[C0():])
  return j
}

func (x *distributedGraph) logKmz (recv bool, m Message, b byte, i uint) {
  if m.Type() != b { panic(m.String() + " error") }
  ft := "->"; if recv { ft = "<-" }
  println (strconv.FormatInt ((time.Now().UnixNano() % 10000000000), 10),
           x.me, m.String(), ft, x.nr[i])
}

func (x *distributedGraph) korachMoranZaks() {
  x.msgchan = make([]chan Stream, NTypes)
  for i := byte(0); i < NTypes; i++ {
    x.msgchan[i] = make(chan Stream)
  }
  m := NewMsg()
  x.connect(m)
  for i := uint(0); i < x.n; i++ {
    go func (j uint) {
      m  = x.ch[j].Recv().(Message)
      k := m.Type()
      x.msgchan[k] <- append(Encode(j), m.Encode()...)
    }(i)
  }
  x.pik = status.New()
  x.unused = make([]bool, x.n)
  for i := uint(0); i < x.n; i++ {
    x.unused[i] = true
  }
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
x.logKmz (false, m, Ask, u)
    var bs Stream
    label: select {
    case bs = <-x.msgchan[Ask]:
      j = x.channelKmz (bs, &m)
x.logKmz (true, m, Ask, j)
      if m.Status().Less (x.pik) {
        goto label
      } else {
        x.state = Citizen
      }
    case bs = <-x.msgchan[Accept]:
      j = x.channelKmz (bs, &m)
x.logKmz (true, m, Accept, j)
      x.child[j] = true
      if m.Status().Phase() < x.pik.Phase() {
        m.Set (Update, x.pik)
        x.ch[j].Send (m)
x.logKmz (false, m, Update, j)
      } else { // pik.Phase() == m.Status().Phase()
               // as Acc only iff m.Status.Phase() <= x.phik.Phase()
               // => x.me < m.Status().Id()
        x.pik.Inc()
        for i := uint(0); i < x.n; i++ {
          if x.child[i] {
            m.Set (Update, x.pik)
            x.ch[i].Send (m)
x.logKmz (false, m, Update, i)
          }
        }
      }
    case bs = <-x.msgchan[YourCitizen]:
      j = x.channelKmz (bs, &m)
x.logKmz (true, m, YourCitizen, j)
      // do nothing and enter the for-loop agains
    case bs = <-x.msgchan[Update]:
      panic ("kmz Update")
    case bs = <-x.msgchan[Leader]:
      panic ("kmz Leader")
    }
  }
  if x.state == King { // no more unused neighbours
    x.leader = x.me
    m.Set (Leader, x.pik)
    for i := uint(0); i < x.n; i++ {
      x.ch[i].Send (m)
x.logKmz (false, m, Leader, i)
    }
    return
  }
// algorithm for a Citizen:
// received an Ask(s)-message along edge j, which changed status from King to Citizen.
  x.child[j] = false
  x.parentChannel = j
  m.Set (Accept, x.pik)
  x.ch[j].Send (m)
x.logKmz (false, m, Accept, j)
  select {
  case bs := <-x.msgchan[Update]:
    j = x.channelKmz (bs, &m)
x.logKmz (true, m, Update, j)
  }
  x.pik.Copy (m.Status()) // phase of pik increased by >= 1
  m.Set (Update, x.pik)
  for i := uint(0); i < x.n; i++ {
    if x.child[i] {
      x.ch[i].Send (m)
x.logKmz (false, m, Update, i)
    }
  }
  jAsk := x.n
  for {
    select {
    case bs := <-x.msgchan[Ask]:
      j = x.channelKmz (bs, &m)
x.logKmz (true, m, Ask, j)
      jAsk = j
      s := m.Status()
      if x.pik.Less (s) {
        x.ch[x.parentChannel].Send (m)
x.logKmz (false, m, Ask, x.parentChannel)
        for x.pik.Less (s) {
          select {
          case bs := <-x.msgchan[Update]:
            j = x.channelKmz (bs, &m)
x.logKmz (true, m, Update, j)
            if j != x.parentChannel { panic("kmz j != parent") }
            x.pik.Copy (m.Status()) // phase of pik increased by >= 1
            m.Set (Update, x.pik)
            for i := uint(0); i < x.n; i++ {
              if x.child[i] {
                x.ch[i].Send (m)
x.logKmz (false, m, Update, i)
              }
            }
          }
        }
      }
      if s.Eq(x.pik) && j != x.parentChannel { // && j is not a tree edge
        m.Set (YourCitizen, x.pik)
        x.ch[j].Send (m)
x.logKmz (false, m, YourCitizen, j)
      }
      if s.Less (x.pik) && j != x.parentChannel { // && j is not a tree edge
        // Ask-message is ignored
      }
    case bs := <-x.msgchan[Update]:
      j = x.channelKmz (bs, &m)
x.logKmz (true, m, Update, j)
      x.pik.Copy (m.Status) // phase of pik increased by >= 1
      m.Set (Update, x.pik)
      for i := uint(0); i < x.n; i++ {
        if x.child[i] {
          x.ch[i].Send (m)
x.logKmz (false, m, Update, i)
        }
      }
    case bs := <-x.msgchan[Accept]:
      j = x.channelKmz (bs, &m)
x.logKmz (true, m, Accept, j)
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
x.logKmz (false, m, Update, i)
            }
          }
        }
      } else { // j != x.parentChannel <=> j is not a tree edge
        x.child[j] = true
        m.Set (Update, x.pik)
        x.ch[j].Send (m)
x.logKmz (false, m, Update, j)
      }
    case bs := <-x.msgchan[Leader]:
      j = x.channelKmz (bs, &m)
x.logKmz (true, m, Leader, j)
      x.leader = m.Status().Id()
      return
    case <-x.msgchan[YourCitizen]:
      panic("kmz YourCitizen")
    }
  }
}
