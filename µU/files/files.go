package files

// (c) Christian Maurer   v. 210419 - license see µU.go

import (
  "os"
  "path"
  "sort"
  . "µU/obj"
  "µU/ker"
  "µU/str"
  "µU/env"
  "µU/n"
  "µU/files/pair"
)
const (
  RWX = 7 * 8 * 8
  worldRX = RWX + 7 * 8 + 7 // = rwxrwxrwx, which is changed to rwxr-xr-x
                            // by & with ~umask for the default umask = 022
  tmp = "/tmp"
)
var (
  pa = pair.New()
  seq = make([]pair.Pair, 0)
  actDir string
  actPath = actualPath()
)

func actualPath() string {
  actPath, e := os.Getwd()
  if e != nil { ker.Panic ("os.Getwd did not work") }
  _, actDir = path.Split (actPath)
  return actPath
}

func actualDir() string {
  _, actDir := path.Split (actualPath())
  return actDir
}

func typeString (t Type) string {
  switch t {
  case Fifo:
    return "Fifo"
  case Device:
    return "device"
  case Dir:
    return "dir"
  case File:
    return "file"
  case Link:
    return "link"
  case Socket:
    return "socket"
  }
  return "unknown"
}

func isPath (p string) bool {
  a := actualPath()
  r := os.Chdir (p)
  if r != nil {
    return false
  }
  if os.Chdir (a) != nil { ker.Oops() }
  return true
}

func isFile (s string) bool {
  fi, r := os.Stat (s)
  if r != nil {
    return false
  }
  return fi.Mode().IsRegular()
}

func isDir (s string) bool {
  fi, r := os.Stat (s)
  if r != nil {
    return false
  }
  return fi.Mode().IsDir()
}

func less (i, j int) bool {
  return str.EquivLess (seq[i].Name(), seq[j].Name())
}

func actualize (path string) {
  if path == actPath {
    return
  } else {
    actPath = path
  }
  seq = make([]pair.Pair, 0)
  f, e := os.Open (".")
  defer f.Close()
  if e != nil { ker.Shit() }
  fileinfos, ef := f.Readdir (-1)
  if ef != nil { ker.Shit() }
  var t Type
  n := len(fileinfos)
  for i := 0; i < n; i++ {
    f := fileinfos[i]
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
    if true { // t != Dir || n != "." && n != "..") {
      pp := pair.New()
      pp.Set (f.Name(), byte(t))
      seq = append (seq, pp)
    }
  }
  sort.Slice (seq, less)
}

func cd (path string) {
  str.OffSpc (&path)
  if path == "" {
    path = env.Home()
  }
  if path[0] == '~' {
    str.Rem (&path, 0, 2)
    cd ("")
  }
  if os.Chdir (path) != nil { ker.Panic ("files cd error; path == " + path) }
  actualize (path)
}

func cdp() {
  home, v := env.Home(), "/." + env.Call()
  p := home + v
  if ! isPath (p) {
    ins (home, v)
  }
  cd (p)
}

func cds() {
  cd (env.Gosrc() + "/" + env.Call())
}

func cd0() {
  home, v := env.Home(), ".µU/" + env.Call()
  p := home + v
  if ! isPath (p) {
    ins (home, v)
  }
  cd (p)
}

/*/
func Temp (filename *string) {
  path, name := path.Split (*filename)
  str.OffSpc (&name)
  n := uint(len (name))
  if n == 0 { return }
  if n + 11 > maxN {
    name = str.Part (name, 0, maxN - 11)
  }
  *filename = path + "." + tmp + ".µU" + "." + name
}
/*/

func sub (d string) {
  os.Mkdir (d, 7 * 8 * 8 + 5 * 8 + 5)
}

func ins (path, d string) bool {
  return os.Mkdir (path + "/" + d, worldRX) != nil
}

func del (path, s string) bool {
  return os.Remove (path + "/" + s) != nil
}

func move (f, d string) {
  os.Rename (f, d + "/" + f)
}

func num() uint {
  return uint(len(seq))
}

func names() []string {
  n := num()
  s := make([]string, n)
  for i := uint(0); i < n; i++ {
    s[i] = seq[i].Name()
  }
  return s
}

func num1 (t Type) uint {
  n := uint(0)
  for i := 0; i < len(seq); i++ {
    if seq[i].Typ() == t {
      n++
    }
  }
  return n
}

func names1 (t Type) []string {
  n := names()
  s := make([]string, 0)
  for i := uint(0); i < num(); i++ {
    if seq[i].Typ() == t {
      s = append (s, n[i])
    }
  }
  return s
}

func typ (n string) Type {
  for i := 0; i < len(seq); i++ {
    pa = seq[i]
    if pa.Name() == n {
      return Type(pa.Typ())
    }
  }
  return Unknown
}

func TmpDir() string {
  d := "µU" + "-" + env.User()
  ins (tmp, d)
  return tmp + "/" + d + "/"
}

func Tmp() string {
  return TmpDir() + n.StringFmt (uint(os.Getpid()), 5, true) + "."
}

var (
  seqPred []pair.Pair
  np uint
)

func numPred (p Pred) uint {
  seqPred = make([]pair.Pair, 0)
  for _, pa := range seq {
    if p (pa) {
      seqPred = append (seqPred, pa)
    }
  }
  np = uint(len(seqPred))
  return np
}

func namePred (p Pred, i uint) string {
  if i < np {
    return seqPred[i].Name()
  }
  return ""
}
