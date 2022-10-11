package obj

// (c) Christian Maurer   v. 220702 - license see nU.go

// Collections of elements of type object or of variables of
// a concrete atomic type (bool, [u]int.., float.., string, ...).
// Every collection has either exactly one actual element
// or its actual element is undefined.
//
// In all specifications x denotes the calling collection.

// Constructors have to return a new collection for elements of the type of a,
// that does not contain any elements; so its actual object is undefined.

type
  Collector interface {

// Empty:   Returns true, iff x does not contain any element.
// Clr:     x is empty; its actual element is undefined.
//  Clearer // ! included to avoid clash in pseq

// Returns true, iff the actual element of x is undefined.
  Offc () bool

// Returns the nunber of elements in x.
  Num () uint

// Pre: a has the type of the elements in x. 
// If x does not carry any order:
//   If the actual element of x was undefined, a copy of a
//   is appended in x (i.e. it is now the last element in x),
//   otherwise x is inserted directly before the actual element.
// Otherwise, i.e. if x is ordered w.r.t. to an order
// relation r (a reflexive, transitive and antisymmetric
// relation "<=") or a strict order relation r (an irreflexive
// and transitive relation "<"):
//   x is inserted behind the last element b in x, for which
//   r(b,a) == true, i.e. that under r "is smaller" than a.
//   If an element b with b == a or b.Eq(a) resp. was already
//   contained in x, then, if r is a strict order, nothing
//   has changed; otherwise, i.e. if r is an order,
//   then a copy of a is contained once more in x.
//   So x is now still ordered w.r.t. r.
// In both cases all other elements and their order in x
// and the actual element in x are not influenced.
  Ins (a any)

// If f and if the actual element of x was defined, then
// the actual element is now the element behind the former actual
// element, if that was defined; otherwise it is undefined.
// If !f and if the actual element of x was defined and was not
// the first element in x, then the actual element of x is now
// the element before the former one; if it was undefined,
// then it is now the last element of x.
// In all other cases, nothing has happened.
  Step (f bool)

// If f is empty, the actual element is undefined; otherwise for
// f/!f the actual element of x now is the last/first element of x.
  Jump (f bool)

// Returns true, iff for f/!f the last/first element of x is its actual element.
  Eoc (f bool) bool

// Returns a copy of the actual element of x, if that is defined; nil otherwise.
  Get () any

// Pre: a has the type of the elements in x. 
// If x is not ordered:
//   If x was empty or if the actual element of x was undefined, a copy of a
//   is appended behind the end of x and is now the actual element of x.
//   Otherwise the actual element of x is replaced by a.
// Otherwise, i.e. if x is ordered:
//   If x was empty, a copy of a is now the only element in x.
//   Otherwise, the actual element in x is deleted and a is inserted into x
//   where the order of x is preserved.
  Put (a any)

// Returns nil, if the actual element of x is not undefined,
// otherwise, the actual element and that was removed from x,
// and the actual element is now the element after it,
// if the former actual element was not the last element of x.
// In that case the actual element of x is now undefined.
  Del () any

// Returns true, iff a is contained in x. In that case
// case the first such element is the actual element of x;
// otherwise, the actual element is the same as before.
  Ex (a any) bool

// Pre: If x is ordered, o is strongly monotone with respect
//      to that order, i.e. x < y implies o(x) < o(y) 
//      (where < denotes the order of x).
// o was applied to all elements in x (in their order in x).
// The actual element of x is the same as before.
  Trav (op Op)

// Pre: y is a collector of elements of the same type as x
//      (especially contains elements of the same type as a).
// If x == y or if x and y do not have the same type,
// nothing has changed. Otherwise:
// If x does not carry any order:
//   x consists of exactly all elements in x before (in their
//   order in x) and behind them all exactly all elements of y
//   before (in their order in y).
//   If the actual element of x was undefined, now the former
//   first element in y is the actual element of x, otherwise
//   the actual element of x is the same as before.
//   y is empty; so its actual element is undefined.
// Otherwise, i.e. if x is ordered w.r.t. to an order relation,
//   Pre: r is either an order (see collector.go) or
//        r is a strict order and x and y are strictly ordered
//        w.r.t. r (i.e. do not contain any two elements a and b
//        with a == b or a.Eq(b) resp.).
//   x consists exactly of all elements in x and y before.
//   If r is strict, then the elements, which are contained
//   in x as well as in y, are contained in x only once,
//   otherwise, i.e. if r is an order, in their multiplicity.
//   x is ordered w.r.t. r and y is empty.
//   The actual elements of x and y are undefined.
  Join (y Collector)
}
