package rw

// (c) Christian Maurer   v. 171019 - license see µU.go

import
  "µU/host"
type
  ReaderWriter interface { // protocols for the readers/writers problem

// Pre: The calling goroutine is neither reading or writing.
// The calling goroutine is reading; no goroutine is writing.
// If at the time of the call there was a writing goroutine,
// the calling goroutine has been delayed, until there was no writing goroutine.
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

func New1() ReaderWriter { return new1() }
func New2() ReaderWriter { return new2() }
func NewSemaphore() ReaderWriter { return newS() }
func NewAddSemaphore (m uint) ReaderWriter { return newAS(m) }
func NewGo() ReaderWriter { return newG() }
func NewCriticalSection1() ReaderWriter { return newCS1() }
func NewCriticalSection2() ReaderWriter { return newCS2() }
func NewCriticalSectionBounded (m uint) ReaderWriter { return newCSB(m) }
func NewCriticalSectionFair() ReaderWriter { return newCSF() }
func NewCriticalResource (r uint) ReaderWriter { return newCR(r) }
func NewMonitor1() ReaderWriter { return newM1() }
func NewMonitor2() ReaderWriter { return newM2() }
func NewConditionedMonitor() ReaderWriter { return newCM() }
func NewConditionedMonitorBounded (m uint) ReaderWriter { return newCMB(m) }
func NewChannel() ReaderWriter { return newCh() }
func NewGuardedSelect() ReaderWriter { return newGS() }
func NewKangLee() ReaderWriter { return newKL() }

// Returns a new ReaderWriter for message passing over the net.
// Arguments: h must be the server and p the used port.
// s determines whether the underlying process runs on the server (s == true) or a client.
func NewFarMonitor (h host.Host, p uint16, s bool) ReaderWriter { return newFM(h,p,s) }
// As above; additionally at most m readers can read concurrently.
func NewFarMonitorBounded (m uint, h host.Host, p uint16, s bool) ReaderWriter {
  return newFMB(m,h,p,s)
}
