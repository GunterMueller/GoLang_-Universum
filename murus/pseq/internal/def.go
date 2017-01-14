package internal

// (c) murus.org  v. 150122 - license see murus.go

type
  File interface {

  Fin()
  Name (n string)
  Rename (n string)
  Empty() bool
  Clr()
  Length() uint64
  Seek (p uint64)
  Position() uint64

  Read (bs []byte) (int, error)
  Write (bs []byte) (int, error)
}

func DirectLength (n string) uint64 { return directLength(n) }
func Erase (n string) { erase(n) }
