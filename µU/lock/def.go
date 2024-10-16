package lock

// (c) Christian Maurer   v. 241001 - license see µU.go

import
  . "sync"

// Protocols for access to a crittial section.
// The functions Lock and Unlock cannot be interrupted
// by calls of Lock or Unlock of other processes.

type
  Lock interface {

// Pre: The calling process is not in the critical section.
// It is the only one in the critical section.
// It might have been delayed, until this was possible.
  Lock()

// Pre: The calling process is in the critical section.
// It is not in the critical section.
  Unlock()
}

// Return new unlocked locks
// with an implementation revealed by their names.
func NewChannel() Locker { return newChan() }
func NewTAS() Locker { return newTAS() }
func NewXCHG() Locker { return newXCHG() }
func NewCAS() Locker { return newCAS() }
func NewDEC() Locker { return newDEC() }
func NewMutex() Locker { return newMutex() }
func NewUdding() Locker { return newUdding() }
func NewMorris() Locker { return newMorris() }
