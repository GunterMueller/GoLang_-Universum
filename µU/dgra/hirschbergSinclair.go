package dgra

// (c) Christian Maurer   v. 231229 - license see µU.go
//
// >>> D.S. Hirschberg, J. B. Sinclair: Decentralized Extrema-Finding in
//     Circular Configurations of Processes. CACM 23 (1980), 627 - 628

import (
  "sync"
  "µU/ker"
  "µU/dgra/internal"
)
const (
  candidate = iota
  lost
  won
)

func (x *distributedGraph) HirschbergSinclair() {
  x.connect (internal.New())
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
        msg := x.recv (j).(internal.Message)
        mutex.Lock()
        id, num, maxnum, ok := msg.IdNumsOk()
        switch msg.Kind() {
        case internal.Candidate:
          if id < x.me {
            msg.SetReply (false)
            x.send (j, msg)
          } else if id > x.me {
            status = lost
            num++
            if num < maxnum {
              if j > 1 { ker.Panic1 ("1 < j ==", j) }
              msg.SetPass (id, num, maxnum)
              x.send (1 - j, msg)
            } else { // num >= maxnum
              msg.SetReply (true)
              x.send (j, msg)
            }
          } else { // id == x.me
            x.leader = x.me
            status = won
            msg.SetLeader (x.me)
            x.send (1 - j, msg)
          }
        case internal.Reply:
          if id == x.me {
            replyOk = replyOk && ok
            gotReply <- j
          } else { // id != x.me
            x.send (1 - j, msg) // pass msg
          }
        case internal.Leader:
          if id == x.me {
            gotReply <- j // trick to force termination
            done <- 0
            mutex.Unlock()
            break loop
          } else { // id != x.me
            status = lost
            msg.SetLeader (id)
            x.send (1 - j, msg)
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
    msg := internal.New()
    msg.SetPass (x.me, 0, maxnum)
    x.send (0, msg)
    x.send (1, msg)
    <-gotReply; <-gotReply // await 2 responses
    if ! replyOk {
      status = lost
    }
    if status != candidate {
      break
    }
    maxnum *= 2
  }
  <-done; <-done
}
