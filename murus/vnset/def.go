package vnset

// (c) Christian Maurer   v. 170418

import
  . "murus/obj"
type
  Predicate func (s VonNeumannSet) bool
type
  VonNeumannSet interface { // "x" means always the calling set;
                            // "in" means "is Element of".

  Object // Eq due to extensionality axiom, Less means "proper subset"
  Adder // Sub means diffence set.
  Stringer

// Returns true, iff x is a subset of y (= Less or Eq).
  Subset (y VonNeumannSet) bool

// Returns true, if s in y.
  Element (y VonNeumannSet) bool

// Returns {x}.
  Singleton() VonNeumannSet

// Returns {x, y}.
  Doubleton (y VonNeumannSet) VonNeumannSet

// Returns the number of elements of s.
  Num() uint

// Returns the union of x and y.
  Union (y VonNeumannSet) VonNeumannSet

// Returns the union of x, i.e. the set of all elements of elements of x.
  BigUnion() VonNeumannSet

// Returns the intersection of x and y.
  Intersection (y VonNeumannSet) VonNeumannSet

// Returns the intersection of x, i.e. the set of all a s.t. a in y for all y in x.
  BigIntersection() VonNeumannSet

// Returns the power set of x.
  Powerset() VonNeumannSet

// Returns the union of x and {x}.
  Succ() VonNeumannSet

// Returns the subset of x of all elements x with p(x) == true.
  Comprehension (p Predicate) VonNeumannSet

// Returns true, if x is transitive, i.e. iff x.Union().Subset(x).
  Transitive() bool

// Returns the set {{x}, {x,y}}.
  KuratowskiPair (y VonNeumannSet) VonNeumannSet

// Returns the ordered pair (x, y).
// >>> Ch. Maurer: Ein rekursiv definiertes geordnetes Paar.
//     Zeitschr. f. math. Logik und Grundlagen d. Math. 22 (1976) 211-214
  Pair (y VonNeumannSet) VonNeumannSet
}

// Returns a new empty mathematical set.
func EmptySet() VonNeumannSet { return emptySet() }

// Returns the ordinal n.
func Ordinal(n uint) VonNeumannSet { return ordinal(n) }

// Returns the set with the elements of x.
func SetOf (x ...VonNeumannSet) VonNeumannSet { return setOf(x...) }
