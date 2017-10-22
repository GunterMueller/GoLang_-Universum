package obj

// (c) Christian Maurer   v. 150122 - license see µU.go

type
  Persistor interface {
// An object "is defined with a name" means, that it is stored
// in a persistent file with that name as "handle" in the filesystem.

// In all specifications x denotes the calling object.

// Pre: n is a valid name in the filesystem and there exists no object
// of a type different from the type of x, but defined with name n.
// x is now defined with name n, i.e. it is the object, that is stored
// in a file with that name, if there exists such; otherwise it is empty.
  Name (n string) // TODO error

// x is defined with that name.
// Another file with that name is now destroyed.
  Rename (n string) // TODO error

// Pre: x is defined with a name.
// TODO Spec
  Fin()
}
