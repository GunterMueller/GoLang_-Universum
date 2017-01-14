package cr

// (c) murus.org  v. 161216 - license see murus.go

/* Objects, that synchronize the access to critical sections
   (e.g. operating resources or shared data) for p process classes
   P(0), P(1), ..., P(p-1) (p > 1), r resources R(0), R(1), ... R(r-1)
   (r > 0) and p * r numbers m(c,n) (c < p, n < r), s.t. each resource
   can be accessed by processes of different classes only under mutual
   exclusion and the resource R(n) (n < r) can be accessed concurrently
   only by at most m(c,n) processes of the class P(c) (c < p).
   Examples:
   a) Readers/Writers problem:
      p = 2, P(0) = readers and P(1) = writers; r = 1 (R(0) = shared data),
      m(0,0) = MaxNat, m(1,0) = 1.
   b) Left/Right problem:
      p = 2, P(1) = lefties and P(2) = righties; r = 1 (R(0) = shared track),
      m(0,0) = m(1,0) = MaxNat.
   c) bounded Left/Right problem:
      same as above with bounds m(0,0), m(1,0) < MaxNat.
   d) Bowling problem: 
      p = number of participating clubs (P(c) = players of club c);
      r = number of available bowling alleys (R(a) = bowling alley a),
      m(c,a) = maximal number of players of club c on alley a.
*/
type
  CriticalResource interface {

// Pre: m[c][r] is defined for all c < number of classes
//      and for all r < number of resources of x.
// The resource r of x can be accessed by at most m[c][r] processes of the class c.
  Limit (m [][]uint)

// Pre: k < number of classes of x.
//      The calling process has no access to a resource of x.
// Returns the number of the resource, to which the calling process now has access.
// The calling process may have been blocked, until that was possible.
  Enter (k uint) uint

// Pre: k < number of classes of x.
//      The calling process has access to a resource of x.
// It now does not have the access any more.
  Leave (k uint)
}
// Returns a new critical resource with c classes and r resources.
func New (c, r uint) CriticalResource { return newCr(c,r) }
