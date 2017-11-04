package rw

// (c) Christian Maurer   v. 171101 - license see µU.go

const (
  reader = uint(iota)
  writer
)
const (
  readerIn = uint(iota)
  readerOut
  writerIn
  writerOut
)
const
  inf = 1 << 31 - 1
