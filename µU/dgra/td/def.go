package td

// (c) Christian Maurer   v. 241015 - license see µU.go

const (
  register = iota
  answer
)
type
  TD interface {
    Register (a uint)
    Answer (a uint) uint
    Fin()
  }

func New (h string, port uint16, s bool) TD { return new_(h,port,s) }
