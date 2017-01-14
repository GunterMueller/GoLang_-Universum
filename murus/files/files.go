package files

// (c) murus.org  v. 150826 - license see murus.go

// >>> alpha-version - still a lot TODO, also some bugs there

import (
  "os"; "path" // ; "exec"
  "murus/ker"; . "murus/obj"
  "murus/str"; "murus/env"
  "murus/nat"; "murus/seq"
  "murus/files/internal"
)
const (
  pack = "files"
  dot = "."
  RWX = 7 * 8 * 8
  worldRX = RWX + 7 * 8 + 7 // = rwxrwxrwx, which is changed to rwxr-xr-x
                            // by & with ~umask for the default umask = 022
  tmp = "/tmp"
  tempPraefix = dot + "tmp" + dot + ker.Murus + dot
)
var (
  pa = internal.New()
  sq = seq.New (pa)
  relsq = seq.New ("")
  sq1 [NTypes]seq.Sequence
  actualDirectory string
  initialized bool
  progvar string
)

func init() {
//  sq.Sort()
  for t := Unknown; t < NTypes; t++ {
    sq1[t] = seq.New ("")
    sq1[t].Sort() // ?
  }
}

func actualPath() string {
  p, err := os.Getwd()
  if err != nil { ker.Stop (pack, 1) }
  return p
}

func actualDir() string {
  _, actualDirectory = path.Split (actualPath())
  return actualDirectory
}

func ex (path string) bool {
  p := actualPath()
  if os.Chdir (path) != nil {
    return false
  }
  if os.Chdir (p) != nil { ker.Stop (pack, 2) }
  return true
}

func defined (name string) bool {
  if str.Empty (name) { return false }
  _, err := os.Stat (name)
  if err == nil { // file exists
    return true
  }
  file, err := os.Open (name)
  defer file.Close()
  return err == nil
}

func contained (name string) bool {
  fi, err := os.Stat (name)
  if err != nil {
    return false
  }
  return fi.Size() > 0
}

func initialize (b bool) { // TODO
  if b {
// println ("initialize (true)")
//    if ! initialized {
//      initialized = true
//    }
  } else {
    if initialized {
      return
    } else {
      initialized = true
// println ("initialize (false)")
    }
  }
  sq.Clr()
  for t := Unknown; t < NTypes; t++ { sq1 [t].Clr() }
  file, e := os.Open (dot)
  defer file.Close()
  if e != nil { ker.Stop (pack, 3) }
  fi, err := file.Readdir (-1)
  if err != nil { ker.Stop (pack, 4) }
  t := Unknown
  for i := 0; i < len (fi); i++ {
    f := fi [i]
/*/
    ms := f.Mode().String()
    if ms[0] == '-' {
      t = File
    } else if str.Contains (ms, 'p', &nn) {
      t = Fifo
    } else if str.Contains (ms, 'c', &nn) {
      t = Device
    } else if str.Contains (ms, 'D', &nn) {
      t = Device
    } else if str.Contains (ms, 'd', &nn) {
      t = Dir
    } else if str.Contains (ms, 'L', &nn) {
      t = Link
    } else if str.Contains (ms, 'S', &nn) {
      t = Socket
    } else {
      t = Unknown
    }
/*/
    switch f.Mode().String()[0] {
    case 'p':
      t = Fifo
    case 'c', 'D':
      t = Device
    case 'd':
      t = Dir
    case '-':
      t = File
    case 'L':
      t = Link
    case 'S':
      t = Socket
    default:
      t = Unknown
    }
    if t != Dir || (f.Name() != dot && f.Name() != "..") {
      pp := internal.New()
      pp.Set (f.Name(), byte(t))
      sq.Ins (pp)
      sq1 [t].Ins (f.Name())
    }
  }
  sq.Seek (0)
  for t := Unknown; t < NTypes; t++ {
    sq1[t].Seek (0)
  }
}

func cd (path string) {
  str.OffSpc (&path)
  if path == "" {
    path = env.Val ("HOME")
  }
  if os.Chdir (path) != nil {
    ker.Stop (pack + " cd error; path == " + path, 5)
    return
  }
  initialized = true
  initialize (true)
}

var
  gesetzt bool

func set (v string) {
  progvar = v
  gesetzt = true
}

func cd0() {
  pth := ""
  if gesetzt {
    pth = progvar
  } else {
    _, call := path.Split (env.Par(0))
    progvar = call
    pth = env.Val (progvar)
  }
  if ex (pth) {
    // pth = $progvar
  } else {
    Home := env.Val ("HOME")
    v := dot + progvar
    pth = Home + "/" + v // $HOME/.progvar
    if ! ex (pth) {
      Ins (Home, v)
    }
  }
  cd (pth)
}

/*
func CopyFilesOfMurus (prog string) {
  if Num() == 0 {
//    exec.Command ("cp", "$MURUS/." + prog + "/*", dot).Run() // TODO
  }
}
*/

func Temp (filename *string) {
  path, name := path.Split (*filename)
  str.OffSpc (&name)
  n := uint(len (name))
  if n == 0 { return }
  if n + 11 > maxN {
    name = str.Part (name, 0, maxN - 11)
  }
  *filename = path + tempPraefix + name
}

func ins (path, dir string) {
  if str.ProperLen (path) == 0 { return }
  err := os.Mkdir (path + "/" + dir, worldRX)
//  if err != os.EEXIST { ker.Stop (pack, 6) }
  if err != nil { /* then WHAT ? */ }
}

func del (path, dir string) {
  os.Remove (path + "/" + dir)
}

func empty() bool {
  if ! initialized {
//    println ("Empty")
    initialize (false)
  }
  return sq.Empty()
}

func num() uint {
  if ! initialized {
//    println ("Num")
    initialize (false)
  }
  return sq.Num()
}

func typ (name string) (Type, bool) {
  if ! initialized {
//    println ("Type")
    initialize (false)
  }
  for i := uint(0); i < sq.Num(); i++ {
    sq.Seek (i)
    pa = sq.Get().(internal.Pair)
    if pa.Name() == name {
      return Type(pa.Typ()), true
    }
  }
  return Unknown, false
}

func entry (i uint) (string, Type, uint64) {
  if ! initialized {
//    println ("Entry")
    initialize (false)
  }
  sq.Seek (i)
  pa = sq.Get().(internal.Pair)
  fi, err := os.Stat (pa.Name())
  e := uint64(0)
  if err != nil {
    e = uint64(fi.Size())
  }
  return pa.Name(), Type(pa.Typ()), e
}

func empty1 (typ Type) bool {
  if ! initialized {
//    println ("Empty1")
    initialize (false)
  }
  return sq1 [typ].Empty()
}

func num1 (typ Type) uint {
  if ! initialized {
//    println ("Num1")
    initialize (false)
  }
  return sq1 [typ].Num()
}

func contained1 (name string, typ Type) bool {
  if ! initialized {
//    println ("Contained1")
    initialize (false)
  }
  for t := Unknown; t < NTypes; t++ {
    if sq1 [t].Ex (name) {
      return true
    }
  }
  return false
}

func name1 (typ Type, i uint) (string, uint64) {
  if ! initialized {
//    println ("Name1")
    initialize (false)
  }
  sq1[typ].Seek (i)
  x := sq1[typ].Get()
  if x == nil {
    return "", 0
  }
  N := x.(string)
  fi, err := os.Stat (N)
  if err == nil {
    return N, uint64(fi.Size())
  }
  return N, 0
}

var np uint

func numPred (p Pred) uint {
  if ! initialized {
//    println ("NumPred")
    initialize (false)
  }
  sq1[File].Filter (relsq, p)
  relsq.Sort() // überflüssig ?
  np = relsq.Num()
  return np
}

func namePred (p Pred, i uint) string {
  if ! initialized {
//    println ("NamePred")
    initialize (false)
  }
  if i < np {
    relsq.Seek (i)
    return relsq.Get().(string)
  }
  return ""
}

func TmpDir() string {
  N := ker.Murus + "-" + env.User()
  Ins (tmp, N)
  return tmp + "/" + N + "/" // /tmp/ker.Murus-$USER/
}

func Tmp() string {
  return TmpDir() + nat.StringFmt (uint(os.Getpid()), 5, true) + dot
}
