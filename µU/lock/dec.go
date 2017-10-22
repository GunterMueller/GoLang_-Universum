package lock

// (c) Christian Maurer   v. 171013 - license see µU.go

type
  dec struct {
             int32
             }

func newDEC() Locker {
  return &dec { int32: 1 }
}

func (x *dec) Lock() {
  for DecrementInt32 (&x.int32) {
    null()
  }
}

func (x *dec) Unlock() {
  x.int32 = 1
}
