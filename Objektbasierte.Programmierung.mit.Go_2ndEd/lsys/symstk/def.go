package symstk

// (c) Christian Maurer   v. 230303 - license see µU.go.

// A stack of pairs (byte, uint); initially empty.

type
  Symbol = byte

// (s, i) is pushed onto the stack.
func Push (s Symbol, i uint) { push(s,i) }

// Returns true, iff x is empty.
func Empty() bool { return empty() }

// Pre: x is not empty.
// Returns the pair on top of x. That pair is now removed from x.
func Pop() (Symbol, uint) { return pop() }
