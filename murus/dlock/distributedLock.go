package dlock

// (c) murus.org  v. 161224 - license see murus.go
//
/* >>> Distributed mutual exclusion due to
       Ricart, G., Agrawala, A. K.:
         An Optimal Algorithm for Mutual Exclusion in Computer Networks
         Commun. ACM 24 (1981), 9-17, 581 */

import (
  "sync"
  "murus/host"
  "murus/nchan"
)
const
  ok = uint(1)
type
  distributedLock struct {
                         uint "number of hosts involved"
                    host []host.Host
                      me uint
                 request,
                   reply []nchan.NetChannel
         mutex, critSect sync.Mutex
           time, ownTime,
                nReplies uint
              requesting bool
                deferred []bool
                         }

func pp (n, x, y uint) uint16 { // 0 <= result < n * (n - 1) / 2
  c := uint16(0)
  for i := uint(0); i < n; i++ {
    for j := i+1; j < n; j++ {
      if i == x && j == y {
        return c
      }
      c++
    }
  }
  return uint16(n * (n - 1) / 2)
}

func qq (n uint, k uint16) (uint, uint) { // inverse of pp
  c := uint(0)
  for i := uint(0); i < n; i++ {
    for j := i+1; j < n; j++ {
      if c == uint(k) {
        return i, j
      }
      c++
    }
  }
  return n, n // nonsense, just to shut the mouth of the compiler
}

func new_(me uint, hs []host.Host, port uint16) DistributedLock {
  n := uint(len(hs))
  if n < 2 || me >= n { return nil }
  x := new (distributedLock)
  x.uint = n
  x.deferred = make ([]bool, x.uint)
  x.host = make ([]host.Host, x.uint)
  for i := uint(0); i < x.uint; i++ {
    x.host[i] = hs[i].Clone().(host.Host)
  }
  x.me = me
  x.critSect.Lock()
  x.request, x.reply = make([]nchan.NetChannel, x.uint), make([]nchan.NetChannel, x.uint)
  n2 := uint16(n * (n - 1) / 2)
  for i := uint(0); i < x.uint; i++ {
    if i != x.me {
//      p := port + uint16(a * x.uint + b)
      p := pp(x.uint, x.me, i); if i < x.me { p = pp(x.uint, i, x.me) } // 0 <= p < n2
      p += port
      x.request[i] = nchan.New (x.ownTime, x.me, i, x.host[i], p)
      x.reply[i] = nchan.New (x.ownTime, x.me, i, x.host[i], p + uint16(n2))
    }
  }
  for i := uint(0); i < x.uint; i++ {
    if i != x.me {
      go func (n uint) { // bookkeeping of requests
        for {
          otherTime := x.request[n].Recv().(uint)
          x.mutex.Lock()
          if otherTime > x.time {
            x.time = otherTime
          }
          if x.requesting && (x.ownTime < otherTime || (x.ownTime == otherTime && x.me < n)) {
            x.deferred[n] = true
          } else {
            x.reply[n].Send(ok)
          }
          x.mutex.Unlock()
        }
      }(i)
      go func (n uint) { // bookkeeping of ok-replies
        for {
          _ = x.reply[n].Recv().(uint)
          x.mutex.Lock()
          x.nReplies++
          if x.nReplies == x.uint - 1 {
            x.critSect.Unlock()
          }
          x.mutex.Unlock()
        }
      }(i)
    }
  }
  return x
}

func (x *distributedLock) Lock() {
  x.mutex.Lock()
  x.requesting = true
  x.ownTime = x.time + 1
  x.nReplies = 0
  x.mutex.Unlock()
  for j := uint(0); j < x.uint; j++ {
    if uint(j) != x.me {
      x.request[j].Send (x.ownTime)
    }
  }
  x.critSect.Lock()
}

func (x *distributedLock) Unlock() {
  x.mutex.Lock()
  x.requesting = false
  x.mutex.Unlock()
  for j := uint(0); j < x.uint; j++ {
    if x.deferred[j] {
      x.deferred[j] = false
      x.reply[j].Send (ok)
    }
  }
}
