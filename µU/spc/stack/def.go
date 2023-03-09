package stack

// (c) Christian Maurer  v. 220225 - license see µU.go

// 2 stacks of vectors

import
  "µU/vect"

// Returns true, iff stack is empty.
func Empty() bool { return empty() }

// v is pushed onto stack.
func Push (v vect.Vector) { push(v) }

// Returns an empty vectors, if stack is empty.
// returns otherwise the vector, that was on top of stack before,
// and this vectors is now removed from stack.
func Pop() vect.Vector { return pop() }

// Returns true, iff stack1 is empty.
func Empty1() bool { return empty1() }

// v is pushed onto stack1.
func Push1 (v vect.Vector) { push1(v) }

// Returns an empty vectors, if stack1 is empty.
// returns otherwise the vector, that was on top of stack1 before,
// and this vector is now removed from stack1.
func Pop1() vect.Vector { return pop() }
