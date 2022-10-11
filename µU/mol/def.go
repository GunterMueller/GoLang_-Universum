package mol

// (c) Christian Maurer   v. 220831 - license see µU.go

import
  . "µU/obj"
const
  Suffix = ".s.dat"
type
  Molecule interface { // structs with atoms as components.
                       // Any molecule has at least one order.
                       // One of the orders is always the actual one.

//  DefineName (n string)
// Pre for Edit: DefineName has to be called immediately before.
  Editor
  Print()
  NumAtoms() uint
  Indexer
  Rotator

  Sub (Y any) bool
  Construct (n string)
}

// Returns a new empty moledule.
func New() Molecule { return new_() }

// Returns the molecule that was built by the call to Construct.
func Constructed (n string) Molecule { return constructed(n) }

