package smok

// (c) murus.org  v. 170627 - license see murus.go

type
  Smokers interface {

// Vor.: u < NUtensils.
// Die aufrufende Wirtin hat die zu u komplementären Utensilien
// verfügbar gemacht. Sie war ggf. solange blockiert, bis keiner raucht.
  Agent(u uint)

// Vor.: u < NUtensils.
// Die zu u komplementären Utensilien sind nicht mehr verfügbar,
// sondern im exklusiven Besitz des aufrufenden Rauchers, der jetzt raucht.
// Er war ggf. solange blockiert, bis sie verfügbar waren.
  SmokerIn(u uint)

// Der aufrufende Raucher raucht nicht mehr.
  SmokerOut()
}

func NewNaive() Smokers { return new_() }
func NewParnas() Smokers { return newP() }
func NewCriticalSection() Smokers { return newCS() }
func NewMonitor() Smokers { return newM() }
func NewConditionedMonitor() Smokers { return newCM() }
