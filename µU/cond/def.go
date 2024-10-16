package cond

// (c) Christian Maurer    v. 241001 - license see  mU.go

type
  Condition interface {

// The calling process is blocked on x.
// the monitor is freigegeben.
  Wait()

// The first der auf x blocked process is unblocked,
// if there is such (wann er in Konkurrenz zu
// anderen Prozessen im Monitor weiterarbeiten kann,
// depends on the implemented signal-semantics);
// otherwise nothing had happened
  Signal()

// Pre: The function is called inside the implementation of a
//      monitor function of x (see Bem).
// Returns true, if (at least) one process is blocked on x.

  Awaited() bool
// Bem.: The value kann sich immediately after the call geändert
//       haben; it is consequently only verwertbar,
//       if the Unteilbarkeit of the call
//       vor späteren Bearbeitungen sichergestellt ist.
//       This is the case bei Einhaltung of the precondition.
}

// Returns a new condition.
func New() Condition { return new_() }
