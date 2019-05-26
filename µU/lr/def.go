package lr

// (c) Christian Maurer   v. 190329 - license see µU.go

type
  LeftRight interface { // protocols for the left right problem

// Pre: The calling goroutine is not a rightie.
// The calling goroutine is a leftie; no goroutine is a rightie.
// If at the time of the call there were active righties,
// the calling goroutine has been delayed, until there was no active rightie.
  LeftIn ()

// Pre: The calling goroutine is an active leftie.
// The calling goroutine is not any more active.
  LeftOut ()

// Pre: The calling goroutine is not a leftie.
// The calling goroutine is a rightie; no goroutine is a leftie.
// If at the time of the call there were active lefties, 
// the calling goroutine has been delayed, until there were no active leftie.
  RightIn ()

// Pre: The calling goroutine is an active rightie.
// The calling goroutine is not any more active.
  RightOut ()

  Fin()
}

func NewMutex() LeftRight { return newM() }
func NewSemaphore() LeftRight { return newS() }
func NewCriticalSection1() LeftRight { return newCS1() }
func NewCriticalSection2() LeftRight { return newCS2() }
func NewCriticalSectionBounded (l, r uint) LeftRight { return newCSB(l,r) }
func NewCriticalResource (l, r uint) LeftRight { return newCR(l,r) }
func NewMonitor1() LeftRight { return newM1() }
func NewMonitor2() LeftRight { return newM2() }
func NewMonitorBounded (l, r uint) LeftRight { return newMB(l, r) }
func NewConditionedMonitor1() LeftRight { return newCM1() }
func NewConditionedMonitor2() LeftRight { return newCM2() }
func NewConditionedMonitorBounded (l, r uint) LeftRight { return newCMB(l,r) }
func NewChannel() LeftRight { return newCh() }
func NewChannelBounded(l, r uint) LeftRight { return newChB(l, r) }
func NewGuardedSelect() LeftRight { return newGS() }

func NewFarMonitor (h string, p uint16, s bool) LeftRight { return newFMon(h,p,s) }
func NewFarMonitorBounded (l, r uint, h string, p uint16, s bool) LeftRight {
  return newFMonB(l,r,h,p,s)
}
