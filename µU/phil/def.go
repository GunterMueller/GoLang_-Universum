package phil

// (c) Christian Maurer   v. 190321 - license see µU.go

import
  . "µU/lockn"
type
  Philos interface {

  LockerN
}
var
  NPhilos = uint(5)

func NewNaive() Philos { return new_() }
func NewSemaphore() Philos { return newS() }
func NewBounded() Philos { return newB() }
func NewUnsymmetric() Philos { return newU() }
func NewSemaphoreUnfair() Philos { return newSU() }
func NewSemaphoreFair() Philos { return newSF() }
func NewCriticalSection() Philos { return newCS() }
func NewMonitor() Philos { return newM() }
func NewMonitorFair() Philos { return newMF() }
func NewMonitorUnfair() Philos { return newMU() }
func NewCondMonitor() Philos { return newCM() }
func NewChannel() Philos { return newCh() }
func NewChannelUnsymmetric() Philos { return newChU() }
func NewFarMonitor (h string, p uint16, s bool) Philos { return newFM(h,p,s) }
func NewGuardedSelect() Philos { return newGS() }
