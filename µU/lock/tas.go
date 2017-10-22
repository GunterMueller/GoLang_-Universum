package lock

// (c) Christian Maurer   v. 171013 - license see µU.go

type
  tas struct {
             bool "locked"
             }

func newTAS() Locker {
  return new(tas)
}

func (x *tas) Lock() {
  for TestAndSet (&x.bool) {
    null()
  }
}

func (x *tas) Unlock() {
  x.bool = false
}
