package internal

// (c) Christian Maurer   v. 210324 - license see µU.go

// The terminal is active.
func New() { new_() }

// Pre: The terminal is active.
// Returns the byte read from the terminal.
func Read() byte { return read() }

// The terminal is not any more active.
func Fin() { fin() }
