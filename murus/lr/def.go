package lr

// (c) murus.org  v. 120910 - license see murus.go
//
//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 79 ff., 93 ff., 183

type
  LeftRight interface { // protocols for the left right problem

  LeftIn ()
  LeftOut ()

  RightIn ()
  RightOut ()
}
