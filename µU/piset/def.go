package piset

// (c) Christian Maurer   v. 210213 - license see µU.go

import
  . "µU/obj"
type
  PersistentIndexedSet interface { // persistent ordered sets of elements,
                                   // that have an index, by which they are ordered.
  Collector
  Persistor
}

// Pre: f is the function that returns the index for the objects in x.
//      (If f returns for every object that object itself,
//      the package just handles persistent ordered sets.)
// Returns a new empty persistent indexed set for objects of the type of o.
func New (o Object, f Func) PersistentIndexedSet { return new_(o,f) }
