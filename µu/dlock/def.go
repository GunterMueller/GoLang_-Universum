package dlock

// (c) Christian Maurer   v. 161222 - license see µu.go

import (
  . "µu/lock"
  "µu/host"
)
type
  DistributedLock interface {

  Locker
}

// Pre: h is the slice of all hosts involved,
//      me is the identity of the calling process (me < len(h)).
//      The ports p..p+n^2 are not used by a network service.
// Returns a new distributed Lock.
func New (me uint, h []host.Host, p uint16) DistributedLock { return new_(me,h,p) }
