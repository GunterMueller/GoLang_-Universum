package smok

// (c) Christian Maurer   v. 230105 - license see µU.go

type
  Smokers interface {

// Pre: u < 3.
// The calling agent has put the utensils complementary to u on the table
// (she was blocked, until no smoker smokes).
  Agent (u uint)

// Pre: u < 3.
// The utensils complementary to u do not lie on the table any more,
// but are in exclusive posssession of the calling smoker, who now is smoking
// (he was blocked, until that was possible).
  SmokerIn (u uint)

// The calling smoker does not smoke any more.
  SmokerOut()
}

func WriteAgent (u uint) { writeAgent(u) }
func WriteSmoker (u uint) { writeSmoker(u) }

// naive implementation with danger of deadlock:
func NewNaive() Smokers { return new_() }

// implementation with helper processes (due to Parnas):
func NewParnas() Smokers { return newP() }

// implementation with critical sections:
func NewCriticalSection() Smokers { return newCS() }

// implementation with a universal monitor:
func NewMonitor() Smokers { return newM() }

// implementation with a conditioned monitor:
func NewConditionedMonitor() Smokers { return newCM() }

// implementation with message passing:
func NewChannel() Smokers { return newCh() }
