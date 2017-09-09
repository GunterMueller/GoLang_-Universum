package lock

// (c) Christian Maurer   v. 161216 - license see murus.go

type
  dec struct {
           n int32
             }

func newDEC() Locker {
  return &dec { n: int32(1) }
}

func (L *dec) Lock() {
  for DecrementInt32 (&L.n) {
    null()
  }
}

func (L *dec) Unlock() {
  L.n = 1
}
