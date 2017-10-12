package ker

// (c) Christian Maurer   v. 140615 - license see µu.go

import (
  "syscall"
  "unsafe"
)
const
  maxConsole = 12 // 12 consoles is enough; 64 or more are confusing
const (
  _AUTO = iota // <linux/vt.h>
  _PROCESS
  _ACKACQ
)
const (
  _OPENQRY = 'V' * 0x100 + iota
  _GETMODE
  _SETMODE
  _GETSTATE
  _SENDSIG
  _RELDISP
  _ACTIVATE
  _WAITACTIVE
)
type (
  mode struct {
         mode,
        waitv byte
       relsig,
       acqsig,
        frsig int16
             }
  stat struct {
       active,
       signal,
        state uint16
             }
)

func ConsoleInit() {
  var m mode
  syscall.Syscall (syscall.SYS_IOCTL, 0, _GETMODE, uintptr(unsafe.Pointer(&m)))
  m.mode = _PROCESS
  m.waitv = 0
  m.relsig = int16(syscall.SIGUSR1)
  m.acqsig = int16(syscall.SIGUSR2)
  syscall.Syscall (syscall.SYS_IOCTL, 0, _SETMODE, uintptr(unsafe.Pointer(&m)))
}

func Console1 (forward bool) {
  const (
    xdm = 7
    msg = 10 // messages
  )
  var s stat
  syscall.Syscall (syscall.SYS_IOCTL, 0, _GETSTATE, uintptr(unsafe.Pointer(&s)))
  a:= s.active
  if forward {
    switch (a) { case xdm - 1:
      a = xdm + 1
    case msg - 1:
      a = msg + 1
    case maxConsole:
      a = 1
    default:
      a++
    }
  } else {
    switch (a) { case 1:
      a = maxConsole
    case xdm + 1:
      a = xdm - 1
    case msg + 1:
      a = msg - 1
    default:
      a--
    }
  }
  syscall.Syscall (syscall.SYS_IOCTL, 0, _ACTIVATE, uintptr(a))
  syscall.Syscall (syscall.SYS_IOCTL, 0, _WAITACTIVE, uintptr(a))
}

func Console (a uint8) {
  if a <= 0 || a > maxConsole { return }
  syscall.Syscall (syscall.SYS_IOCTL, 0, _ACTIVATE, uintptr(a))
  syscall.Syscall (syscall.SYS_IOCTL, 0, _WAITACTIVE, uintptr(a))
}

func ActualConsole() uint {
  var s stat
  syscall.Syscall (syscall.SYS_IOCTL, 0, _GETSTATE, uintptr(unsafe.Pointer(&s)))
  return uint(s.active)
}

func DeactivateConsole() {
  syscall.Syscall (syscall.SYS_IOCTL, 0, _RELDISP, _ACKACQ)
}

func ActivateConsole() {
  syscall.Syscall (syscall.SYS_IOCTL, 0, _RELDISP, _PROCESS)
}
