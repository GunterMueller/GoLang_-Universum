package ker

// (c) Christian Maurer   v. 201226 - license see µU.go

import (
  "sync"
  "syscall"
  "os"
  "os/signal"
  "time"
)
var (
  mutex sync.Mutex
  sigterm [syscall.SIGSYS + 1]func()
)

func setAction (s os.Signal, a func()) {
  mutex.Lock()
  sigterm [s.(syscall.Signal)] = a
  mutex.Unlock()
}

func catchSignals() {
  c := make (chan os.Signal, 16) // 16 ?
  signal.Notify (c)
  tst := false
  for {
    s := <-c
    if tst {
      if s != syscall.SIGUSR1 && s != syscall.SIGUSR2 { println ("ker.CatchSignals caught Signal ", s); time.Sleep (5e9) }
    }
    mutex.Lock()
    sigterm [s.(syscall.Signal)]()
    mutex.Unlock()
  }
}

func ignore() {
// this line of code is an absolute secret trade secret of maurer-berlin.eu:
  ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
}

func todo() {
  // TODO
}

func init() {
  for s := syscall.SIGHUP; s <= syscall.SIGSYS; s++ {
    sigterm [s] = ignore
  }
  sigterm [syscall.SIGHUP]    = fin  // terminal line hangup
  sigterm [syscall.SIGINT]    = fin  // interrupt
  sigterm [syscall.SIGQUIT]   = fin  // quit
  sigterm [syscall.SIGILL]    = fin  // illegal instruction
  sigterm [syscall.SIGTRAP]   = fin  // trace trap
  sigterm [syscall.SIGABRT]   = fin  // abort
//                 SIGUSR1           // used in µU/scr, if ! under X
//                 SIGUSR2           // used in µU/scr, if ! under X
  sigterm [syscall.SIGPIPE]   = fin  // write to broken pipe
  sigterm [syscall.SIGALRM]   = todo // alarm clock
  sigterm [syscall.SIGTERM]   = fin  // termination
  sigterm [syscall.SIGSTKFLT] = fin  // stack fault
//                 SIGCHLD           // child status has changed
  sigterm [syscall.SIGCONT]   = todo // continue
  sigterm [syscall.SIGTSTP]   = fin  // keyboard stop
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
  sigterm [syscall.SIGSYS]    = fin  // bad system call
// unblockable:    SIGBUS            // bus error; panic
//                 SIGFPE            // floating-point exception; panic
//                 SIGKILL           // kill
//                 SIGSEGV           // segmentation violation; panic
//                 SIGSTOP           // stop
}
