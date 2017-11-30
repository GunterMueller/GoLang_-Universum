package env

// (c) Christian Maurer   v. 171126 - license see nU.go

import "os"

func par1() byte {
  if uint(len(os.Args)) > 1 {
    return os.Args[1][0]
  }
  return 0
}

func nPars() uint {
  return uint(len(os.Args)) - 1
}

func localhost() string {
  return os.Getenv ("HOSTNAME") // "HOST"
}

func call() string {
  return os.Args[0]
}
