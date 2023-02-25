package barb

// (c) Christian Maurer   v. 230105 - license see µU.go

type
  Barber interface {

  Customer()
  Barber()
}

// implementation with a semaphore:
func NewSem() Barber { return newS() }

// implementation with a semaphore by direct handover of the mutual exclusion:
func NewDir() Barber { return newD() }

// implementation with a critical sections:
func NewCS() Barber { return newCS() }

// implementation with a universal monitor:
func NewMon() Barber { return newM() }

// implementation with a conditioned monitor:
func NewCondMon() Barber { return newCM() }

// implementation with conditions in package "sync" (due to Gregory Andrews):
func NewAndrews() Barber { return newA() }
