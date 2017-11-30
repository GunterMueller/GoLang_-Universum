package dlock

// (c) Christian Maurer   v. 171125 - license see µU.go

import
  . "µU/lock"
type
  DistributedLock interface {

  Locker
}

// Pre: h is the slice of all hostnames involved,
//      me is the identity of the calling process (me < len(h)).
//      The ports p..p+n^2 are not used by a network service.
// Returns a new distributed Lock.
func New (me uint, h []string, p uint16) DistributedLock { return new_(me,h,p) }
