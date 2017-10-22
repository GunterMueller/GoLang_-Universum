package obj

// (c) Christian Maurer   v. 170424 - license see µU.go

type
  Iterator interface {

  Collector

// Returns the number of those elements in x, for which p returns true.
  NumPred (p Pred) uint

// Returns NumPred(p) == Num(), i.e. returns true, iff p returns true
// on all elements in x (particularly if x has no elements).
  All (p Pred) bool

// Returns true, iff there is an element in x, for which p returns true.
// In that case the actual element of x is for f/!f the last/first such
// element, otherwise the actual element of x is the same as before.
  ExPred (p Pred, f bool) bool

// Returns true, iff there is an element in x in direction f
// from the actual element of x, for which p returns true.
// In that case the actual element of x is for f/!f the
// next/previous such element, otherwise the actual element of x
// is the same as before.
  StepPred (p Pred, f bool) bool

// Pre: If x is ordered, o is strongly monotone with respect
//      to that order, i.e. x < y implies o(x) < o(y) 
//      (where < denotes the order of x).
// o was applied to all elements in x (in their order in x).
// The actual element of x is the same as before.
  Trav (op Op)

// op was applied to all elements in x, for which p returns true.
//  TravPred (p Pred, o Op) = Trav (PredOp2Op (p, o))

// o(-, true) was applied to all elements in x,
// for which p returns true, otherwise p(-, false).
//  TravCond (p Pred, o CondOp) = Trav (PredCondOp2Op (p, o))

// Pre: y is a collector of elements of the same type as x
//      (especially contains elements of the same type as a).
// y consists exactly of those elements in x before
// (in their order in x), for which p returns true.
// The actual element of x is undefined; x is unchanged.
  Filter (y Iterator, p Pred)

// Pre: See Filter.
// y contains exactly those elements in x (in their order in x),
// for which p returns true, and exactly those elements are
// removed from x. The actual elements of x and y are undefined.
  Cut (y Iterator, p Pred)

// In x all elements, for which p returns true, are removed.
// If the actual element of x was one of them, now it is undefined.
  ClrPred (p Pred)

// Pre: See Filter.
// If the actual element of x was undefined, x is not changed
// and y is empty. Otherwise, y contains exactly the elements
// before the former actual element in x and y contains the former
// actual element of x and exactly all elements behind it, both in
// their former order in x. In this case the actual element of x
// is undefined and the actual element of y is its first element.
  Split (y Iterator)

// Pre: See Filter.
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
  Join (y Iterator)

// If x is empty or x is ordered, nothing has changed.
// Otherwise x contains exactly the same elements as before
// in their former order, with the following exception:
// for f == true the former last element in x is now the first
// one and for f == false the former first element in x is now
// the last one. The actual element of x is the same as before.
//  Rotate () // ? TODO
}
