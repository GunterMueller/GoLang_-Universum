package menue

// (c) Christian Maurer   v. 161216 - license see µU.go

import
  . "µU/obj"
type
  Menue interface { // Multiway trees of menues and statements.
                    // Each leaf contains a statement, that can be executed;
                    // the other nodes are menues, from which a node or leaf
                    // of the level below them can be selected,
                    // Nodes and leaves are identified by strings.

// If there is a level below x, nothing has happened.
// Otherwise, x is a leaf with statement s.
// While executing s, the name of x appears
// in the top line of the screen, iff t == true.
  Leaf (s Stmt, t bool)

// If x is a leaf, nothing has happened.
// Otherwise, y is inserted into the level below x.
  Ins (y Menue)

// If x is a leaf, the statement of x was executed and now
// the menue, from which x was selected, is again presented.
// Otherwise, a menue is presented, which allows to select
// a node or leaf from the level below x.
  Exec ()
}

// Returns an node with name s without a level below.
func New (s string) Menue { return new_(s) }
