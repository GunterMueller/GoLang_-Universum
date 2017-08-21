package dgra

// (c) murus.org  v. 170403 - license see murus.go
//
// >>> D.S. Hirschberg, J. B. Sinclair: Decentralized Extrema-Finding in
//     Circular Configuations of Processes. CACM 23 (1980), 627 - 628

import
  "murus/dgra/internal"
const (
  candidate = iota
  lost
  won
)

func (x *distributedGraph) hirschbergSinclair() {
  x.connect (internal.New())
  defer x.fin()
  maxnum := uint(1)
  status := candidate
  replyOk := true
  gotReply := make(chan uint, x.n)
  done := make(chan uint, x.n)
  for i := uint(0); i < x.n; i++ { // listen to both sides
    go func (j uint) {
      loop:
      for {
        msg := x.ch[j].Recv().(internal.Message)
        mutex.Lock()
        typ := msg.Typ()
        val, num, maxnum, ok := msg.Content()
        switch typ {
        case internal.Candidate:
          if val < x.me {
            msg.Reply (false); x.ch[j].Send(msg)
          } else if val > x.me {
            status = lost
            num++
            if num < maxnum {
              msg.PassCandidate (val, num, maxnum); x.ch[1 - j].Send(msg)
            } else { // nm[j] >= maxnum
              msg.Reply (true); x.ch[j].Send(msg)
            }
          } else { // val[j] == x.net.Me()
            x.leader = x.me
            status = won
            msg.Define (internal.Leader, x.me); x.ch[1 - j].Send(msg)
          }
        case internal.Reply:
          if val == x.me {
            replyOk = replyOk && ok
            gotReply <- j
          } else { // val[j] != x.net.Me()
            x.ch[1 - j].Send(msg) // pass msg
          }
        case internal.Leader:
          if val == x.me {
            gotReply <- j // trick to force termination
            done <- 0
            mutex.Unlock()
            break loop
          } else {
            status = lost
            msg.Define (internal.Leader, val); x.ch[1 - j].Send(msg)
            x.leader = val
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
    msg.PassCandidate (x.me, 0, maxnum)
    x.ch[0].Send(msg)
    x.ch[1].Send(msg)
    <-gotReply; <-gotReply // await 2 respomses
    if ! replyOk {
      status = lost
    }
    if status != candidate { break }
    maxnum *= 2
  }
  <-done; <-done
}
