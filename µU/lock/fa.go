package lock

// (c) Christian Maurer   v. 171013 - license see µU.go

type
  fa struct {
            uint32
            }

func newFA() Locker {
  return new (fa)
}

func (x *fa) Lock() {
  for FetchAndAddUint32 (&x.uint32, 1) != 0 {
    null()
  }
}

func (x *fa) Unlock() {
  x.uint32 = 0
}
