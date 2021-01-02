package ker

// (c) Christian Maurer   v. 201226 - license see µU.go

import
  "runtime"

func init() {
  switch runtime.GOARCH {
  case "386": // arm mips mipsle ?
    Panic ("µU does not support any longer 32-bit computers")
  case "amd64":
    // at the time being the only supported architecture
/*/
  case arm64: // s390x ppc64 ppc64le mips64 mips64le ?
    TODO
/*/
  default:
    Panic ("$GOARCH not yet supported by µU")
  }
  switch runtime.GOOS {
  case "linux":
    // at the time beeing the only supported system
//  case openbsd, freebsd, netbsd, solaris:
//    Panic ("please give me a note, if $GOOS is your operating system")
  case "windows":
    Panic ("windows is not yet supported by µU, but support is under work")
  default: // aix, android, darwin, plan9, dragonfly
    Panic ("$GOOS not yet supported by µU")
  }
}
