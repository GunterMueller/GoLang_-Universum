package prog

// (c) Christian Maurer  v. 210401 - license see µU.go

type
  Program interface {

  Empty() bool
  GetLines()
  Parse() (string, uint)
  Write()
  Edit()
  Run()
}

func New() Program { return new_() }
