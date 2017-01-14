package nelect

// (c) murus.org  v.170113 - license see murus.go
//
// >>> D.S. Hirschberg, J. B. Sinclair: Decentralized Extrema-Finding in
//     Circular Configuations of Processes. CACM 23 (1980), 627 - 628
//
// XXX Yet a bug in the implementation: Clean termination is missing !
//     At the moment just a horrible workaround by forcing termination via os.Exit().

import (
  "time"
//  "os"
//  "murus/errh"
  . "murus/nelect/internal"
  "murus/nchan"
)
type status byte; const (
  candidate = status(iota)
  lost
  won
)

func (x *netElection) hirschbergSinclair() uint {
  ch := make([]nchan.NetChannel, 2)
  m := make([]Message, 2)
  for i := uint(0); i < 2; i++ { m[i] = NewMsg() }
  value := make([]uint, 2)
  lg := make([]uint, 2)
  lgmax := make([]uint, 2)
  ok := make([]bool, 2)
  mT := make([]Type, 2)
  side := make([]string, 2)
  ch[0], ch[1] = x.chR, x.chL
  winner := x.uint
  status := candidate
  respOk := make([]bool, 2)
  ended := make([]bool, 2)
//  for i:= 0; i < 2; i++ { respOk[i] = true }
  gotResp := make(chan uint, 2)
  for i := uint(0); i < 2; i++ { // listen on both sides
    go func (j uint) {
      loop:
      for {
        side[j] = "right"; if j == 1 { side[j] = "left" }
        m[j] = ch[j].Recv().(Message)
        mT[j], value[j], lg[j], lgmax[j], ok[j] = m[j].Content()
        switch mT[j] {
        case Candidate:
          if value[j] < x.me {
            m[j].Reply (false); ch[j].Send(m[j])
          } else if value[j] > x.me {
            status = lost
            lg[j]++
            if lg[j] < lgmax[j] {
              m[j].PassCandidate (value[j], lg[j], lgmax[j]); ch[1-j].Send(m[j])
            } else { // lg[j] >= lgmax[j]
              m[j].Reply (true); ch[j].Send(m[j])
            }
          } else { // value[j] == x.me
            status = won
//////////////////////////////////////////////////////
            m[j].PassWon (x.me); ch[1-j].Send(m[j]) //
//////////////////////////////////////////////////////
            println("from " + side[j] + ": I won", x.me)
            winner = x.me
//            errh.Error("I am the winner", winner)
//            os.Exit(int(winner)) // horrible workaround
            break loop
          }
        case Reply:
          if value[j] == x.me {
            respOk[j] = respOk[j] && ok[j]
            gotResp <- j
          } else { // value[j] != x.me
            m[j].Pass(); ch[1-j].Send(m[j])
          }
        case Won:
          if value[j] == x.me {
            panic("unreachable") // XXX  Why the hell ?
          } else {
            status = lost
//////////////////////////////////////////////////////////
            m[j].PassWon (value[j]); ch[1-j].Send(m[j]) //
//////////////////////////////////////////////////////////
            println("from " + side[j] + ": winner is", value[j])
            winner = value[j]
//            errh.Error("the winner is", winner)
//            os.Exit(int(winner)) // horrible workaround
            break loop
          }
        }
      }
      ended[j] = true
//      println("ended " + side[j])
    }(i)
  }
  maxnum := uint(1)
//  for winner == x.uint {
//  for ! (ended[0] && ended[1]) {
//  for {
  for status == candidate {
// if ended[0] { break }
    if winner != x.uint { break }
    msg := NewMsg()
    respOk[0], respOk[1] = true, true
    msg.PassCandidate (x.me, 0, maxnum)
//    if winner != x.uint { break }
    ch[0].Send(msg)
    ch[1].Send(msg)
//    if winner != x.uint { break }
    <-gotResp; <-gotResp // await 2 respomses
    if ! (respOk[0] && respOk[1]) {
      status = lost
    }
    maxnum *= 2
//    println("next round", maxnum)
  }
time.Sleep(2 * 1e9)
println("Ende erreicht")
  return winner
}
