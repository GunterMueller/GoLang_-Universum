package env

// (c) Christian Maurer   v. 171202 - license see µU.go

import
  "os"

func set (Variable string, content *string) {
  for i:= 0; i < len (Variable); i++ {
    switch Variable[i] { case ' ', '=': return }
  }
  err:= os.Setenv (Variable, *content) // int64
  if err != nil { panic ("no Variable") }
}

func val (Variable string) string {
  return os.Getenv (Variable)
}

func localhost() string {
  return os.Getenv ("HOSTNAME") // or "HOST" ?
}

func home() string {
  return os.Getenv ("HOME")
}

func gosrc() string {
  return val("GOSRC")
}

func user() string {
  return os.Getenv ("USER")
}

func arg1() byte {
  if uint(len (os.Args)) > 1 {
    return os.Args[1][0]
  }
  return 0
}

func arg2() byte {
  if uint(len (os.Args)) > 2 {
    return os.Args[2][0]
  }
  return 0
}

func nArgs() uint {
  return uint(len (os.Args)) - 1
}

func arg (i uint) string {
  if uint(len (os.Args)) > i {
    return os.Args[i]
  }
  return ""
}

func call() string {
  return os.Args[0]
}
