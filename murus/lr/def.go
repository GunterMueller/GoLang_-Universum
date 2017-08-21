package lr

// (c) murus.org  v. 170411 - license see murus.go

//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 79 ff., 93 ff., 183

import
  "murus/host"
type
  LeftRight interface { // protocols for the left right problem
                        // Specs: You had better bye my book and read.
  LeftIn ()
  LeftOut ()
  RightIn ()
  RightOut ()

  Fin()
}

func New() LeftRight { return new_() }
func NewChannel() LeftRight { return newChan() }
func NewGuardedSelect() LeftRight { return newGSel() }
func NewMonitor() LeftRight { return newMon() }
func NewMonitorBounded (l, r uint) LeftRight { return newMB(l, r) }
func NewConditionedMonitor() LeftRight { return newCM() }
func NewConditionedMonitorBounded (l, r uint) LeftRight { return newCMB(l,r) }
func NewCriticalSection() LeftRight { return newCS() }
func NewCriticalSectionBounded (l, r uint) LeftRight { return newCSB(l,r) }
func NewCriticalResource() LeftRight { return newCR() }
// func NewSzymanski() LeftRight ( return newSz() }
func NewFarMonitor (h host.Host, p uint16, s bool) LeftRight { return newFMon(h,p,s) }
func NewFarMonitorBounded (l, r uint, h host.Host, p uint16, s bool) LeftRight {
  return newFMonB(l,r,h,p,s)
}
