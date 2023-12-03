package dgra

// (c) Christian Maurer   v. 220816 - license see µU.go
//
// >>> D.S. Hirschberg, J. B. Sinclair: Decentralized Extrema-Finding in
//     Circular Configuations of Processes. CACM 23 (1980), 627 - 628

import (
  "sync"
  "µU/dgra/internal"
)
const (
  candidate = iota
  lost
  won
)

func (x *distributedGraph) hirschbergSinclair() {
  x.connect (msg.New())
  defer x.fin()
  status := candidate
  maxnum := uint(1)
  replyOk := true
  var mutex sync.Mutex
  gotReply := make(chan uint, x.n)
  done := make(chan uint, x.n)
  for i := uint(0); i < x.n; i++ { // listen on all networkchannels
    go func (j uint) {
      loop:
      for {
        message := x.ch[j].Recv().(msg.Message)
        mutex.Lock()
        id, num, maxnum, ok := message.IdNumsOk()
        switch message.Kind() {
        case msg.Candidate:
          if id < x.me {
            message.SetReply (false)
//            x.ch[j].Send (message)
            x.send (j, message)
          } else if id > x.me {
            status = lost
            num++
            if num < maxnum {
              message.SetPass (id, num, maxnum)
//              x.ch[1 - j].Send (message)
              x.send (1 - j, message)
            } else { // num >= maxnum
              message.SetReply (true)
//              x.ch[j].Send (msg)
              x.send (j, message)
            }
          } else { // id == x.me
            x.leader = x.me
            status = won
            message.SetLeader (x.me)
//            x.ch[1 - j].Send (message)
            x.send (1 - j, message)
          }
        case msg.Reply:
          if id == x.me {
            replyOk = replyOk && ok
            gotReply <- j
          } else { // id != x.me
//            x.ch[1 - j].Send (message) // pass msg
            x.send (1 - j, message) // pass msg
          }
        case msg.Leader:
          if id == x.me {
            gotReply <- j // trick to force termination
            done <- 0
            mutex.Unlock()
            break loop
          } else {
            status = lost
            message.SetLeader (id)
//            x.ch[1 - j].Send (message)
            x.send (1 - j, message)
            x.leader = id
            done <- 0
            mutex.Unlock()
            break loop
          }
        }
        mutex.Unlock()
      }
    }(i)
  }
  for status == candidate {
    replyOk = true
    message := msg.New()
    message.SetPass (x.me, 0, maxnum)
//    x.ch[0].Send (message)
    x.send (0, message)
//    x.ch[1].Send (message)
    x.send (1, message)
    <-gotReply; <-gotReply // await 2 responses
    if ! replyOk {
      status = lost
    }
    if status != candidate { break }
    maxnum *= 2
  }
  <-done; <-done
}
