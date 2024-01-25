package phil

// (c) Christian Maurer   v. 171229 - license see nU.go

// >>> Solution with conditioned monitor

import
  "nU/cmon"
type
  conditionedMonitor struct {
                            cmon.Monitor
                            }

func newCM() Philos {
  nForks := make([]uint, 5)
  for i := uint(0); i < 5; i++ {
    nForks[i] = 2
  }
  c := func (i uint) bool {
         if i < 5 { // Lock
           return nForks[i] == 2
         }
         return true // Unlock
       }
  f := func (i uint) uint {
         if i < 5 {
           nForks[left(i)]--
           nForks[right(i)]--
           return i
         }
         i -= 5
         nForks[left(i)]++
         nForks[right(i)]++
         return i
       }
  return &conditionedMonitor { cmon.New (5, f, c) }
}

func (x *conditionedMonitor) Lock (i uint) {
  changeStatus (i, hungry)
  x.F (lock + i)
  changeStatus (i, dining)
}

func (x *conditionedMonitor) Unlock (i uint) {
  changeStatus (i, thinking)
  x.F (5 + i)
}
