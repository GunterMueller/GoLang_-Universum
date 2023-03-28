package main

// (c) Christian Maurer   v. 230305 - license see µU.go

import (
  "µU/ker"
  "µU/env"
  "µU/kbd"
  "µU/files"
  "bahn/bahnhof"
  "bahn/netz"
)

func main() {
  if env.NArgs() == 0 {
    netz.MeinBahnhof = netz.Server
    netz.Aktivieren()
    kbd.Wait (false)
    return
  }
  n, m := env.N(1), netz.N - 1
  if n >= m { ker.Panic1 ("Das Argument des Aufrufs \"bahn\" muss kleiner sein als", m) }
  netz.MeinBahnhof = n
  files.Cds()
  bahnhof.New().Betreiben()
}
