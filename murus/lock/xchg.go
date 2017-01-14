package lock

// (c) murus.org  v. 161216 - license see murus.go

type
  xchg struct {
            n uint32
              }

func newXCHG() Locker {
  return &xchg { 1 }
}

func (L *xchg) Lock() {
  local:= uint32(0)
  for ExchangeUint32 (&L.n, local) == 0 {
    null()
  }
}

func (L *xchg) Unlock() {
//  local:= uint32(1); local = ExchangeUint32 (&L.n, local)
  L.n = 1
}
