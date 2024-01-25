package phil

// (c) Christian Maurer   v. 171229 - license see nU.go

// >>> Solution with critical sections

import
  . "nU/cs"
type
  criticalSection struct {
                         CriticalSection
                         }

func newCS() Philos {
  nForks := make([]uint, 5)
  for i := uint(0); i < 5; i++ {
    nForks[i] = 2
  }
  c := func (i uint) bool {
         return nForks[i] == 2
       }
  f := func (i uint) uint {
         nForks[left(i)]--
         nForks[right(i)]--
         return uint(0)
       }
  l := func (i uint) {
         nForks[left(i)]++
         nForks[right(i)]++
       }
  return &criticalSection { New (5, c, f, l) }
}

func (x *criticalSection) Lock (i uint) {
  changeStatus (i, hungry)
  x.Enter (i)
  changeStatus (i, dining)
}

func (x *criticalSection) Unlock (i uint) {
  changeStatus (i, thinking)
  x.Leave (i)
}
