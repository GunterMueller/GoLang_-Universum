package env

// (c) Christian Maurer   v. 170923 - license see µu.go

import
  "os"
const (
  host_ = "HOST"
  user_ = "USER"
  home_ = "HOME"
)

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

func host() string {
  return os.Getenv (host_)
}

func home() string {
  return os.Getenv (home_)
}

func gosrc() string {
  return val("GOSRC")
}

func user() string {
  return os.Getenv (user_)
}

func par1() byte {
  if uint(len (os.Args)) > 1 {
    return os.Args[1][0]
  }
  return 0
}

func par2() byte {
  if uint(len (os.Args)) > 2 {
    return os.Args[2][0]
  }
  return 0
}

func nPars() uint {
  return uint(len (os.Args)) - 1
}

func par (i uint) string {
  if uint(len (os.Args)) > i {
    return os.Args[i]
  }
  return ""
}

func call() string {
  return os.Args[0]
}
