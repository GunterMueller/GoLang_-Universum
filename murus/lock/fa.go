package lock

// (c) murus.org  v. 161216 - license see murus.go

type
  fa struct {
          n uint32
            }

func newFA() Locker {
  return new (fa)
}

func (L *fa) Lock() {
  for FetchAndAddUint32 (&L.n, uint32(1)) != 0 {
    null()
  }
}

func (L *fa) Unlock() {
  L.n = uint32(0)
}
