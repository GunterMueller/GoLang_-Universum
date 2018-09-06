package internal

// (c) Christian Maurer   v. 180819 - license see µU.go

import
  . "µU/obj"
const (
  Candidate = byte(iota) // types of messages
  Reply
  Leader
)
type
  Message interface { // quintuple (type, id, num, maxnum, ok)

  Equaler
  Coder

// Returns the type of x.
  Type() byte

// Returns the quadruple (id, num, maxnum, ok) of x.
  IdNumsOk() (uint, uint, uint, bool)

// x consists of type Candidate, id i, num n and maxnum n.
// ok of x is not changed.
  SetPass (i, n, m uint)

// x consists of type Reply and ok b,
// the other components of x are unchanged.
  SetReply (b bool)

// x consists of type Leader and id i,
// the other components of x are unchanged.
  SetLeader (i uint)
}

// Returns an new message, consisting
// of zero values in all components.
func New() Message { return new_() }
