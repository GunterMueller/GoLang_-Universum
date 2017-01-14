package comp

// (c) murus.org  v. 151117 - license see murus.org
//
// XXX Modify according to your needs !

import
  "os"
var (
//  hosts = []string { "jupiter", "saturn", "venus", "uranus", "mars", "apollo", "diana", "neptun" }
//  hosts = []string { "jupiter", "venus", "mars" }
  hosts = []string { "jupiter", "jupiter", "jupiter", "jupiter", "jupiter", "jupiter", "jupiter", "jupiter" }
  server = hosts[0]
  localHost, _ = os.Hostname()
)
