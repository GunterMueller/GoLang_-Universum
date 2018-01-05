package env

// (c) Christian Maurer   v. 171227 - license see nU.go

import "os"

func arg1() byte {
  if uint(len(os.Args)) > 1 {
    return os.Args[1][0]
  }
  return 0
}

func arg (i uint) string {
  if uint(len(os.Args)) > i {
    return os.Args[i]
  }
  return ""
}

func localhost() string {
  return os.Getenv ("HOSTNAME") // "HOST"
}
