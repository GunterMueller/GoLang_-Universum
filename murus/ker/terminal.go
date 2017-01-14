package ker

// (c) murus.org  v. 140615 - license see murus.go

import (
  "unsafe"
  "syscall"
)
const (
  TC = 'T' * 0x100 + iota
  TCGETS
  TCSETS
  TCSETSW
)
const (
  KDGKBMODE = 0x4B44 + iota
  KDSKBMODE
)
const (
  K_RAW = iota // <linux/kd.h>
  K_XLATE
  K_MEDIUMRAW
)
type // <bits/termios.h>, <bits/types.h>
  termios struct {
    iflag, oflag,
    cflag, lflag uint32
            line byte
              cc [32]byte
  ispeed, ospeed uint32
                 }
var
  t termios

func ReadTerminal (b *byte) {
  syscall.Syscall (syscall.SYS_READ, 0, uintptr(unsafe.Pointer(b)), 1)
}

func TerminalFin() {
  syscall.Syscall (syscall.SYS_IOCTL, 0, KDSKBMODE, K_XLATE)
  syscall.Syscall (syscall.SYS_IOCTL, 0, TCSETS, uintptr(unsafe.Pointer(&t)))
}

func InitTerminal() {
  syscall.Syscall (syscall.SYS_IOCTL, 0, TCGETS, uintptr(unsafe.Pointer(&t)))
  var t1 termios
  syscall.Syscall (syscall.SYS_IOCTL, 0, TCGETS, uintptr(unsafe.Pointer(&t1)))
  t1.iflag = 0
  const (ISIG = 1; ICANON = 2; ECHO = 8; ECHONL = 64)
  t1.lflag &^= ( ISIG | ICANON | ECHO )
  const (VTIME = 5; VMIN = 6)
  t1.cc [VTIME] = 1
  t1.cc [VMIN] = 3 // or 18 ?
  syscall.Syscall (syscall.SYS_IOCTL, 0, TCSETSW, uintptr(unsafe.Pointer(&t1)))
  syscall.Syscall (syscall.SYS_IOCTL, 0, KDSKBMODE, K_MEDIUMRAW)
}
