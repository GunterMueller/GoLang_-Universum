package rw

// (c) murus.org  v. 170411 - license see murus.go
//
//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 75 ff., 85 ff., 137 ff. u.v.a.

import
  "murus/host"
type
  ReaderWriter interface { // protocols for the reader writer problem

// Pre: The calling goroutine is neither reading or writing.
// The calling goroutine is reading; no goroutine is writing.
// If at the time of the call there was a writing goroutine,
// the goroutine has been delayed, until there was no writing goroutine.
  ReaderIn ()

// Pre: The calling goroutine is reading.
// The calling goroutine is neither reading or writing.
  ReaderOut ()

// Pre: The calling goroutine is neither reading or writing.
// The calling goroutine is writing;
// no other goroutine is writing and there are no reading goroutines.
// If at the time of the call there were reading goroutines or a writing one,
// the calling goroutine has been delayed, until there were no reading or writing goroutines.
  WriterIn ()

// Pre: The calling goroutine is writing.
// The calling goroutine is neither reading or writing.
  WriterOut ()

  Fin()
}

func New() ReaderWriter { return new_() }
func New1() ReaderWriter { return new1() }
func New2() ReaderWriter { return new2() }
func NewSemaphore() ReaderWriter { return newSem() }
func NewAddSemaphore (m uint) ReaderWriter { return newASem(m) }
func NewChannel() ReaderWriter { return newChan() }
func NewGuardedSelect() ReaderWriter { return newGSel() }
func NewCriticalSection() ReaderWriter { return newCS() }
func NewCriticalSection2() ReaderWriter { return newCS2() }
func NewCriticalSectionBounded (m uint) ReaderWriter { return newCSB(m) }
func NewCriticalSectionFair() ReaderWriter { return newCSFair() }
func NewCriticalResource() ReaderWriter { return newCR() }
func NewMonitor() ReaderWriter { return newMon() }
func NewConditionedMonitor() ReaderWriter { return newCondMon() }

// Returns a new ReaderWriter for message passing over the net.
// Arguments: h must be the server and p the used port.
// s determines whether the underlying process runs on the server (s == true) or a client.
func NewFarMonitor (h host.Host, p uint16, s bool) ReaderWriter { return newFMon(h,p,s) }
