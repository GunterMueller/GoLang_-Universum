package audio

// (c) Christian Maurer   v. 210515 - license see µU.go

import
  "µU/text"
var
  con = []string {"Claudio Abbado",
                  "Geza Anda",
                  "Ernest Ansermet",
                  "Daniel Barenboim",
                  "Leonard Bernstein",
                  "Karl Böhm",
                  "Colin Davis",
                  "Eugen Jochum",
                  "Herbert von Karajan",
                  "Otto Klemperer",
                  "Kurt Masur",
                  "Eugene Ormandy",
                  "Gennadi Roschdestwenski",
                  "Claudio Scimone"}
var
  conductor []text.Text

func init() {
  n := len(con)
  conductor = make([]text.Text, n)
  for i := 0; i < n; i++ {
    conductor[i] = text.New (len0)
    conductor[i].Defined (con[i])
  }
}
