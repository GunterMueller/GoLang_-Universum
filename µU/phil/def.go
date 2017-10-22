package phil

// (c) Christian Maurer   v. 171019 - license see µU.go

import
  . "µU/lockp"
type
  Philos interface {

  LockerP
}
var
  NPhilos = uint(5)

func NewNaive() LockerP { return new_() }
func NewBounded() LockerP { return newB() }
func NewUnsymmetric() LockerP { return newU() }
func NewSemaphoreUnfair() LockerP { return newSU() }
func NewSemaphoreFair() LockerP { return newSF() }
func NewCriticalSection() LockerP { return newCS() }
func NewMonitor() LockerP { return newM() }
func NewMonitorFair() LockerP { return newMF() }
func NewMonitorUnfair() LockerP { return newMU() }
func NewCondMonitor() LockerP { return newCM() }
func NewChannel() LockerP { return newCh() }
func NewChannelUnsymmetric() LockerP { return newChU() }
