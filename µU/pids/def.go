package pids

// (c) Christian Maurer   v. 201004 - license see µU.go

import
  . "µU/obj"
type
  PersistentIndexerSequence interface {

  Operate (l, c uint)
}
// Pre: a implements Indexer and Colourer.
// Returns a new persistent sequence of editor objects with name n.
func New (a Any, n string) PersistentIndexerSequence { return new_(a,n) }
