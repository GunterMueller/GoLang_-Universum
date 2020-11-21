package stack

// (c) Christian Maurer  v. 201103 - license see µU.go

// Two stacks of N numbers of type float64.
const N = 9

// Returns true, iff stack is empty.
func Empty() bool { return empty() }

// The N numbers r[i] are pushed onto stack.
func Push (r ...float64) { push(r...) }

// Returns 0's, if stack is empty.
// returns otherwise the N numbers, that were on top of stack before,
// and these numbers are now removed from stack.
func Pop() []float64 { return pop() }

// Returns true, iff stack1 is empty.
func Empty1() bool { return empty1() }

// The N numbers r[i] are pushed onto stack1.
func Push1 (r ...float64) { push1(r...) }

// Returns 0's, if stack1 is empty.
// returns otherwise the N numbers, that were on top of stack1 before,
// and these numbers are now removed from stack1.
func Pop1() []float64 { return pop1() }
