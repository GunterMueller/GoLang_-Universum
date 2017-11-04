package phil

// (c) Christian Maurer   v. 171019 - license see µU.go

import
  . "µU/lockn"
type
  Philos interface {

  LockerN
}
var
  NPhilos = uint(5)

func NewNaive() LockerN { return new_() }
func NewSemaphore() LockerN { return newS() }
func NewBounded() LockerN { return newB() }
func NewUnsymmetric() LockerN { return newU() }
func NewSemaphoreUnfair() LockerN { return newSU() }
func NewSemaphoreFair() LockerN { return newSF() }
func NewCriticalSection() LockerN { return newCS() }
func NewMonitor() LockerN { return newM() }
func NewMonitorFair() LockerN { return newMF() }
func NewMonitorUnfair() LockerN { return newMU() }
func NewCondMonitor() LockerN { return newCM() }
func NewChannel() LockerN { return newCh() }
func NewChannelUnsymmetric() LockerN { return newChU() }

func NewChannel1() LockerN { return newCh1() }
