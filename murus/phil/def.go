package phil

// (c) murus.org  v. 170627 - license see murus.go

import
  . "murus/lockp"
const
  NPhilos = 5
type
  Philos interface {

  LockerP
}

func NewNaive() LockerP { return new_() }
func NewBounded() LockerP { return newB() }
func NewUnsymmetric() LockerP { return newU() }
func NewSemaphoreUnfair() LockerP { return newSU() }
func NewSemaphoreFair() LockerP { return newSF() }
func NewCriticalSection() LockerP { return newCS() }
func NewCriticalSectionAging() LockerP { return newCSA() }
func NewMonitor() LockerP { return newM() }
func NewMonitorFair() LockerP { return newMF() }
func NewMonitorUnfair() LockerP { return newMU() }
func NewCondMonitor() LockerP { return newCM() }
func NewChannel() LockerP { return newCh() }
func NewChannelUnsymmetric() LockerP { return newChU() }
