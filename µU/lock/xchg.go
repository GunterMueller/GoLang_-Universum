package lock

// (c) Christian Maurer   v. 161216 - license see µU.go

type
  xchg struct {
              uint32 "0 or 1; initially 1"
              }

func newXCHG() Locker {
  return &xchg { 1 }
}

func (x *xchg) Lock() {
  local := uint32(0)
  for ExchangeUint32 (&x.uint32, local) == 0 {
    null()
  }
}

func (x *xchg) Unlock() {
//  local := uint32(1); local = ExchangeUint32 (&x.uint32, local)
  x.uint32 = 1
}
