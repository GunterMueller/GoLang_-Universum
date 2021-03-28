package obj

// (c) Christian Maurer   v. 210316 - license see µU.go

type
  Persistor interface {
// An object "is defined with a name" means, that it is stored
// in a persistent file with that name as "handle" in the filesystem.

// Pre: n is a valid name in the filesystem and there exists no object
// of a type different from the type of x, but defined with name n.
// x is now defined with name n, i.e. it is the object, that is stored
// in a file with that name, if there exists such; otherwise it is empty.
  Name (n string)

// x is defined with that name.
// Another file with that name is now destroyed.
  Rename (n string)

// Pre: x is defined with a name.
// x is secured in the file system.
  Fin()
}

func IsPersistor (a Any) bool {
  if a == nil { return false }
  _, ok := a.(Persistor)
  return ok
}
