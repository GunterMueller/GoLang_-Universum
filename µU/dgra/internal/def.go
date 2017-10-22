package internal

// (c) Christian Maurer   v. 170209 - license see µU.go

import
  . "µU/obj"
type
  MsgType byte; const (
  Candidate = MsgType(iota) // for HirschberSinclair
  Reply
  Leader
)
type
  Message interface {

  Equaler; Coder //  Object

  Typ() MsgType
  Content() (uint, uint, uint, bool)
  Val() uint
  PassCandidate (i, r, d uint)
  Reply (t bool)
  Define (t MsgType, v uint)
}

func New() Message { return new_() }
