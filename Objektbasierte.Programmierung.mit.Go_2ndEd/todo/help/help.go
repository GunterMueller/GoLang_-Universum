package help

// (c) Christian Maurer   v. 211215 - license see µU.go

import
  "µU/errh"
var
  txt []string = make ([]string, 1)

func init() {
  txt[0] = "Bedienungsanleitung siehe Buch"
}

func help() {
  errh.Help (txt)
}
