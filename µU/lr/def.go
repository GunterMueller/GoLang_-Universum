package lr

// (c) Christian Maurer   v. 230105 - license see µU.go

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

// implementations of the 1st lr-problem with mutexes and semaphores:
func NewMutex() LeftRight { return newM() }
func NewSemaphore() LeftRight { return newS() }

// implementations of the 1st, 2nd and the bounded lr-problem with cricital sections:
func NewCriticalSection1() LeftRight { return newCS1() }
func NewCriticalSection2() LeftRight { return newCS2() }
func NewCriticalSectionBounded (l, r uint) LeftRight { return newCSB(l,r) }

// implementation of the bounded lr-problem with critical resoures:
func NewCriticalResource (l, r uint) LeftRight { return newCR(l,r) }

// implementations of the 1st, 2nd and the bounded lr-problem with monitors:
func NewMonitor1() LeftRight { return newM1() }
func NewMonitor2() LeftRight { return newM2() }
func NewMonitorBounded (l, r uint) LeftRight { return newMB(l, r) }

// implementations of the 1st, 2nd and the bounded lr-problem with conditioned monitors:
func NewConditionedMonitor1() LeftRight { return newCM1() }
func NewConditionedMonitor2() LeftRight { return newCM2() }
func NewConditionedMonitorBounded (l, r uint) LeftRight { return newCMB(l,r) }

// implementations of the 1st and the bounded lr-problem with message passing:
func NewChannel() LeftRight { return newCh() }
func NewChannelBounded(l, r uint) LeftRight { return newChB(l, r) }

// implementation of the 1st lr-problem with guarded select:
func NewGuardedSelect() LeftRight { return newGS() }

// implementations of the 1st and the bounded lr-problem with far monitors:
func NewFarMonitor (h string, p uint16, s bool) LeftRight { return newFMon(h,p,s) }
func NewFarMonitorBounded (l, r uint, h string, p uint16, s bool) LeftRight {
                                                                  return newFMonB(l,r,h,p,s) }
