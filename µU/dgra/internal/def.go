package internal

// (c) Christian Maurer   v. 171219 - license see µU.go

import
  . "µU/obj"
const (
  Candidate = byte(iota) // Type of message
  Reply
  Leader
)
type
  Message interface { // quintuple (Type, id, num, maxnum, ok)

  Equaler
  Coder

// Returns the type of the message.
  Type() byte

// Returns the quadruple (id, num, maxnum, ok) of the message.
  IdNumsOk() (uint, uint, uint, bool)

// The message consists of type Candidate,
// id i, num n, maxnum n and undefined ok.
// ok is not changed.
  SetPass (i, n, m uint)

// The message consists of type Reply and ok b,
// the other components are unchanged.
  SetReply (b bool)

// The message consists of type Leader and id i,
// the other components are unchanged.
  SetLeader (i uint)
}

// Returns an new message, consisting
// of zero values in all components.
func New() Message { return new_() }
