package rw

// (c) Christian Maurer   v. 230105 - license see µU.go

type
  ReaderWriter interface { // protocols for the readers/writers problem

// Pre: The calling goroutine is neither reading or writing.
// The calling goroutine is reading; no goroutine is writing.
// If at the time of the call there was a writing goroutine,
// the calling goroutine has been delayed, until there was no writing goroutine.
  ReaderIn()

// Pre: The calling goroutine is reading.
// The calling goroutine is neither reading or writing.
  ReaderOut()

// Pre: The calling goroutine is neither reading or writing.
// The calling goroutine is writing;
// no other goroutine is writing and there are no reading goroutines.
// If at the time of the call there were reading goroutines or a writing one,
// the calling goroutine has been delayed, until there were no reading or writing goroutines.
  WriterIn()

// Pre: The calling goroutine is writing.
// The calling goroutine is neither reading or writing.
  WriterOut()

  Fin()
}

// Parameter m means the maximally admissible number of concurrent readers

// implementations of the 1st and 2nd rw-problem with mutexes:
func New1() ReaderWriter { return new1() }
func New2() ReaderWriter { return new2() }

// implementation of the 1st and 2nd rw-problem with semaphores:
func NewSemaphore() ReaderWriter { return newS() }

// implementation of the bounded rw-problem with an additive semaphore:
func NewAddS (m uint) ReaderWriter { return newAddS(m) }

// implementation of the rw-problem with Googles solution:
func NewGo() ReaderWriter { return newG() }

// implementations of the 1st, 2nd and the bounded rw-problem with critical sections:
func NewCriticalSection1() ReaderWriter { return newCS1() }
func NewCriticalSection2() ReaderWriter { return newCS2() }
func NewCriticalSectionBounded (m uint) ReaderWriter { return newCSB(m) }

// fair implementation of the rw-problem with critical sections:
func NewCriticalSectionFair() ReaderWriter { return newCSF() }

// implementations of the bounded rw-problem with critical sections:
func NewCriticalResource (m uint) ReaderWriter { return newCR(m) }

// implementations of the 1st and the 2nd rw-problem with monitors:
func NewMonitor1() ReaderWriter { return newM1() }
func NewMonitor2() ReaderWriter { return newM2() }

// implementations of the 1st and the bounded rw-problem with conditioned monitors:
func NewConditionedMonitor() ReaderWriter { return newCM() }
func NewConditionedMonitorBounded (m uint) ReaderWriter { return newCMB(m) }

// implementation of the 1st rw-problem with message passing:
func NewChannel() ReaderWriter { return newCh() }

// implementation of the 1st, 2nd and the bounded rw-problem with critical sections:
func NewGuardedSelect() ReaderWriter { return newGS() }

// implementation of the rw-problem by Kand and Lee:
func NewKangLee() ReaderWriter { return newKL() }

// implementation of the 1st rw-problem with a far monitor (h = name of the server,
// p = used port; the underlying process runs for s == true on the server, otherwise as client):
func NewFarMonitor (h string, p uint16, s bool) ReaderWriter { return newFM(h,p,s) }

// implementation of the 1st rw-problem with a far monitor (see above):
func NewFarMonitorBounded (m uint, h string, p uint16, s bool) ReaderWriter {
                                                               return newFMB(m,h,p,s) }
