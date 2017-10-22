package smok

// (c) Christian Maurer   v. 170627 - license see µU.go

type
  Smokers interface {

// Vor.: u < 3.
// Die aufrufende Wirtin hat die zu u komplementären Utensilien
// verfügbar gemacht. Sie war ggf. solange blockiert, bis keiner raucht.
  Agent (u uint)

// Vor.: u < 3.
// Die zu u komplementären Utensilien sind nicht mehr verfügbar,
// sondern im exklusiven Besitz des aufrufenden Rauchers, der jetzt raucht.
// Er war ggf. solange blockiert, bis das möglich war.
  SmokerIn (u uint)

// Der aufrufende Raucher raucht nicht mehr.
  SmokerOut()
}
func WriteAgent (u uint) { writeAgent(u) }
func WriteSmoker (u uint) { writeSmoker(u) }

func NewNaive() Smokers { return new_() }
func NewParnas() Smokers { return newP() }
func NewCriticalSection() Smokers { return newCS() }
func NewMonitor() Smokers { return newM() }
func NewConditionedMonitor() Smokers { return newCM() }
