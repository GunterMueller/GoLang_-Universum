package comp

// (c) Christian Maurer   v. 180810 - license see µU.go

import
  "os"
var (
//  hosts = []string { "jupiter", "saturn", "mars", "uranus", "neptun", "venus" }
//  hosts = []string { "jupiter", "saturn", "mars" }
  hosts = []string { "jupiter", "jupiter", "jupiter", "jupiter", "jupiter", "jupiter", "jupiter", "jupiter" }
  server = hosts[0]
  localHost, _ = os.Hostname()
)
