package phil

// (c) Christian Maurer   v. 230105 - license see µU.go

import
  "µU/lockn"
type
  Philos interface {

  lockn.LockerN
}
var
  NPhilos = uint(5)

// naive implementation with mutex (deadlock):
func NewNaive() Philos { return new_() }

// naive implementation with semaphores (deadlock):
func NewSemaphore() Philos { return newS() }

// implementation with bounded number of philos:
func NewBounded() Philos { return newB() }

// unsymmetric implementation with mutexes:
func NewUnsymmetric() Philos { return newU() }

// unfair implementation with a semaphore (danger of starvation):
func NewSemaphoreUnfair() Philos { return newSU() }

// fair implementation with a semaphore (due to Dijkstra):
func NewSemaphoreFair() Philos { return newSF() }

// implementation with critital sections:
func NewCriticalSection() Philos { return newCS() }

// implementation with a universal monitor:
func NewMonitor() Philos { return newM() }

// fair implementation with a monitor (due to Dijkstra):
func NewMonitorFair() Philos { return newMF() }

// unfair implementation with a monitor (due to Dijkstra):
func NewMonitorUnfair() Philos { return newMU() }

// implementation with a conditioned monitor:
func NewCondMonitor() Philos { return newCM() }

// implementation with synchronous message passing (due to Ben-Ari):
func NewChannel() Philos { return newCh() }

// unsymmetric implementation with synchronous message passing:
func NewChannelUnsymmetric() Philos { return newChU() }

// implementation with a far monitor (s = name of the server, p = used port,
func NewFarMonitor (h string, p uint16, s bool) Philos { return newFM(h,p,s) }
