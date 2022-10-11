package masks

// (c) Christian Maurer   v. 220805 - license see µU.go

type
  Masks interface { // Persistent sets of constant texts with certain positions on the screen

  Name (n string)
  Empty() bool
  Clr()
  Write()
  Edit()
  Print()
}

func New() Masks { return new_() }
