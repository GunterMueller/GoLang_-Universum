package asem

// (c) Christian Maurer   v. 170410 - license see µU.go

type
  AddSemaphore interface { // Specs: Buy my book and read the section on additive semaphores.

  P (n uint)
  V (n uint)
}

func New (n uint) AddSemaphore { return new_(n) }

func NewNaive (n uint) AddSemaphore { return newNaive (n) }
