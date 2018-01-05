package dlock

// (c) Christian Maurer   v. 171202 - license see nU.go

import
  . "nU/lock"
type
  DistributedLock interface {

  Locker
}

// Vor.: h ist der Slice der Namen der beteiligten Rechner.
//       me ist die Identität des aufrufenden Prozesse (me < len(h)).
//       Die Ports p..p+n^2 sind nicht von einem Netzwerkdienst belegt.
// Liefert ein neues verteiltes Schloss.
func New (me uint, h []string, p uint16) DistributedLock { return new_(me,h,p) }
