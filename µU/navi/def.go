package navi

// (c) Christian Maurer   v. 120909 - license see µU.go

// Pre: /dev/input/navi is readable for world.
// If there is e.g. a rule in /etc/udev/rules.d with the line:
// SYSFS{idVendor}=="046d", SYSFS{idProduct}=="c626", MODE="444", SYMLINK+="input/navi"
// then a Space Navigator of 3dconnexion is initialized.

import
  "µU/spc"
type
  Command byte

func Exists () bool { return exists() }

func Channel () (chan Command) { return navipipe }

func Read () (spc.GridCoord, spc.GridCoord) { return read() }
