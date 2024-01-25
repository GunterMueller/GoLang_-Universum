package dlock

// (c) Christian Maurer   v. 171202 - license see nU.go

import (
  "sync"
  "nU/nchan"
)
const
  ok = uint(1)
type
  distributedLock struct {
                         uint "Anzahl der beteiligten Hosts"
                      me uint // Identität des aufrufenden Prozesses
                    host []string
          request, reply []nchan.NetChannel // 0 für Reqests, 1 für Replies
             time, time1, // eigene und andere Zeit
                nReplies uint
              requesting bool
                deferred []bool
         mutex, critSect sync.Mutex
                         }

func new_(me uint, hs []string, p uint16) DistributedLock {
  n := uint(len(hs))
  if n < 2 || me >= n { return nil }
  x := new (distributedLock)
  x.uint = n
  x.deferred = make ([]bool, x.uint)
  x.host = make ([]string, x.uint)
  for i := uint(0); i < x.uint; i++ {
    x.host[i] = hs[i]
  }
  x.me = me
  x.critSect.Lock()
  x.request = make([]nchan.NetChannel, x.uint)
  x.reply = make([]nchan.NetChannel, x.uint)
  p += nchan.Port0
  for k := uint(0); k < x.uint; k++ {
    i, j := x.me, k
    if k != x.me {
      if k < x.me {
        i, j = j, i
      }
      x.request[k] = nchan.New (x.time, x.me, k, x.host[k], nchan.Port (x.uint, i, j, 0) + p)
      x.reply[k] = nchan.New (x.time, x.me, k, x.host[k], nchan.Port (x.uint, i, j, 1) + p)
    }
  }
  for i := uint(0); i < x.uint; i++ {
    if i != x.me {
      go func (n uint) { // bookkeeping of request channels
        for {
          otherTime := x.request[n].Recv().(uint)
          x.mutex.Lock()
          if otherTime > x.time {
            x.time = otherTime
          }
          if x.requesting && (x.time1 < otherTime || (x.time1 == otherTime && x.me < n)) {
            x.deferred[n] = true
          } else {
            x.reply[n].Send(ok)
          }
          x.mutex.Unlock()
        }
      }(i)
      go func (n uint) { // bookkeeping of ok-replies
        for {
          x.reply[n].Recv()
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
  x.time1 = x.time + 1
  x.nReplies = 0
  x.mutex.Unlock()
  for j := uint(0); j < x.uint; j++ {
    if uint(j) != x.me {
      x.request[j].Send (x.time1)
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
