package msghs

// (c) Christian Maurer   v. 171217 - license see µU.go

import
  . "µU/obj"
const (
  Candidate = byte(iota)
  Reply
  Leader
)
type
  Message interface {

  Equaler
  Coder

  Type() byte
  Val() uint
  Num() uint
  Maxnum() uint
  Ok() bool

  Pass (t byte, v, n, m uint)
  Reply (t bool)
  Define (t byte, v uint)
}

func New() Message { return new_() }
