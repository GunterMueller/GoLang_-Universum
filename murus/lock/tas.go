package lock

// (c) murus.org  v. 161216 - license see murus.go

type
  tas struct {
      locked bool
             }

func newTAS() Locker {
  return new (tas)
}

func (L *tas) Lock() {
  for TestAndSet (&L.locked) {
    null()
  }
}

func (L *tas) Unlock() {
  L.locked = false
}
