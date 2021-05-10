package piset

// (c) Christian Maurer   v. 210509 - license see µU.go

import
  . "µU/obj"
type
  PersistentIndexedSet interface { // persistent ordered sets of elements,
                                   // that have an index, by which they are ordered
  Collector
  Persistor

  Operate()

}

// Returns a new empty persistent indexed set for objects of the type of o.
func New (o Indexer) PersistentIndexedSet { return new_(o) }

func TeXstring() string { return tex }
