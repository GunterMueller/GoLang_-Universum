package asem

// (c) Christian Maurer   v. 171124 - license see nU.go

type AddSemaphore interface {

  P (n uint)
  V (n uint)
}

func New (n uint) AddSemaphore { return new_(n) }

func NewNaive (n uint) AddSemaphore { return newNaive (n) }
