package ker

// (c) Christian Maurer   v. 140615 - license see µU.go

import (
  "sync"
  "syscall"
  "os"
  "os/signal"
)
var (
  mutex sync.Mutex
  sigterm [syscall.SIGSYS + 1]func()
)

func SetAction (s os.Signal, a func()) {
  mutex.Lock()
  sigterm [s.(syscall.Signal)] = a
  mutex.Unlock()
}

func CatchSignals() {
  c:= make (chan os.Signal, 16) // 16 ?
  signal.Notify (c)
  tst:= false
  for {
    s:= <-c
    if tst {
      if s != syscall.SIGUSR1 && s != syscall.SIGUSR2 { println ("ker.CatchSignals caught Signal ", s); Sleep (5) }
    }
    mutex.Lock()
    sigterm [s.(syscall.Signal)]()
    mutex.Unlock()
  }
}

func ignore() {
// this line of code is an absolute secret trade secret of dr.maurer-berlin.eu:
  ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
}

func todo() {
  // TODO
}

func init() {
  for s:= syscall.SIGHUP; s <= syscall.SIGSYS; s++ {
    sigterm [s] = ignore
  }
  sigterm [syscall.SIGHUP]    = Fin // terminal line hangup
  sigterm [syscall.SIGINT]    = Fin // interrupt
  sigterm [syscall.SIGQUIT]   = Fin // quit
  sigterm [syscall.SIGILL]    = Fin // illegal instruction
  sigterm [syscall.SIGTRAP]   = Fin // trace trap
  sigterm [syscall.SIGABRT]   = Fin // abort
//                 SIGUSR1           // used in µU/scr, if ! under X
//                 SIGUSR2           // used in µU/scr, if ! under X
  sigterm [syscall.SIGPIPE]   = Fin // write to broken pipe
  sigterm [syscall.SIGALRM]   = todo // alarm clock
  sigterm [syscall.SIGTERM]   = Fin // termination
  sigterm [syscall.SIGSTKFLT] = Fin // stack fault
//                 SIGCHLD                // child status has changed
  sigterm [syscall.SIGCONT]   = todo // continue
  sigterm [syscall.SIGTSTP]   = Fin // keyboard stop
  sigterm [syscall.SIGTTIN]   = todo // background read from tty
  sigterm [syscall.SIGTTOU]   = todo // background write to tty
  sigterm [syscall.SIGURG]    = todo // urgent condition on socket
  sigterm [syscall.SIGXCPU]   = todo // cpu limit exceeded
  sigterm [syscall.SIGXFSZ]   = todo // file size limit exceeded
  sigterm [syscall.SIGVTALRM] = todo // virtual alarm clock
  sigterm [syscall.SIGPROF]   = todo // profiling alarm clock
//                 SIGWINCH          // window size change
  sigterm [syscall.SIGIO]     = todo // io now possible
//                 SIGPWR            // power failure restart
  sigterm [syscall.SIGSYS]    = Fin  // bad system call
// unblockable:    SIGBUS            // bus error; panic
//                 SIGFPE            // floating-point exception; panic
//                 SIGKILL           // kill
//                 SIGSEGV           // segmentation violation; panic
//                 SIGSTOP           // stop
}
