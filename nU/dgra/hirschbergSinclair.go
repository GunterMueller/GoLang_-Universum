package dgra

// (c) Christian Maurer   v. 231220 - license see nU.go

import
  "nU/dgra/internal"
const (
  candidate = iota
  lost
  won
)

func (x *distributedGraph) HirschbergSinclair() {
  x.connect (internal.New())
  defer x.fin()
  maxnum := uint(1)
  status := candidate
  replyOk := true
  gotReply := make(chan uint, x.n)
  done := make(chan uint, x.n)
  for i := uint(0); i < x.n; i++ {
    go func (j uint) {
      loop:
      for {
        msg := x.ch[j].Recv().(internal.Message)
        mutex.Lock()
        id, num, maxnum, ok := msg.IdNumsOk()
        switch msg.Kind() {
        case internal.Candidate:
          if id < x.me {
            msg.SetReply (false)
            x.ch[j].Send (msg)
          } else if id > x.me {
            status = lost
            num++
            if num < maxnum {
              msg.SetPass (id, num, maxnum)
              x.ch[1 - j].Send (msg)
            } else { // num >= maxnum
              msg.SetReply (true)
              x.ch[j].Send (msg)
            }
          } else { // id == x.me
            x.leader = x.me
            status = won
            msg.SetLeader (x.me)
            x.ch[1 - j].Send (msg)
          }
        case internal.Reply:
          if id == x.me {
            replyOk = replyOk && ok
            gotReply <- j
          } else { // id != x.me
            x.ch[1 - j].Send (msg)
          }
        case internal.Leader:
          if id == x.me {
            gotReply <- j
            done <- 0
            mutex.Unlock()
            break loop
          } else {
            status = lost
            msg.SetLeader (id)
            x.ch[1 - j].Send (msg)
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
    x.ch[0].Send (msg)
    x.ch[1].Send (msg)
    <-gotReply; <-gotReply
    if ! replyOk {
      status = lost
    }
    if status != candidate { break }
    maxnum *= 2
  }
  <-done; <-done
}
